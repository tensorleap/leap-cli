package k3d

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/docker/docker/api/types"
	dockerTypes "github.com/docker/docker/api/types"
	cliutil "github.com/k3d-io/k3d/v5/cmd/util"
	"github.com/k3d-io/k3d/v5/pkg/client"
	"github.com/k3d-io/k3d/v5/pkg/runtimes"
	docker "github.com/k3d-io/k3d/v5/pkg/runtimes/docker"
	k3d "github.com/k3d-io/k3d/v5/pkg/types"
)

type Registry = k3d.Registry

const REGISTRY_NAME = "k3d-tensorleap-registry"
const CONTAINER_NAME = "k3d-tensorleap-server-0"
const REGISTRY_DOMAIN = "k3d-tensorleap-registry:5000"

func GetRegistry(ctx context.Context) (*Registry, error) {
	registry, err := client.RegistryGet(ctx, runtimes.SelectedRuntime, REGISTRY_NAME)
	if err != nil {
		return nil, err
	}
	return registry, nil
}

func CreateLocalRegistry(ctx context.Context, port uint, volumes []string) (*Registry, error) {
	if existingRegistry, _ := client.RegistryGet(ctx, runtimes.SelectedRuntime, REGISTRY_NAME); existingRegistry != nil {
		log.Println("Found existing registry!")
		return existingRegistry, nil
	}

	reg := createRegistryConfig(port, volumes)
	_, err := client.RegistryRun(ctx, runtimes.SelectedRuntime, reg)
	if err != nil {
		return nil, err
	}
	return reg, nil
}

func createRegistryConfig(port uint, volumes []string) *Registry {
	exposePort, err := cliutil.ParsePortExposureSpec(strconv.FormatUint(uint64(port), 10), k3d.DefaultRegistryPort)
	if err != nil {
		log.Fatalln(err)
	}

	reg := &Registry{
		Host:         REGISTRY_NAME,
		Image:        fmt.Sprintf("%s:%s", k3d.DefaultRegistryImageRepo, k3d.DefaultRegistryImageTag),
		ExposureOpts: *exposePort,
		Network:      k3d.DefaultRuntimeNetwork,
		Volumes:      volumes,
	}

	return reg
}

func CacheImage(ctx context.Context, image string, reg *Registry) error {
	dockerClient, err := docker.GetDockerClient()
	if err != nil {
		return err
	}

	resp, err := dockerClient.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("docker failed to pull the image '%s': %w", image, err)
	}
	defer resp.Close()

	log.Printf("Pulling image '%s'\n", image)

	// this prints out status of the pull, consider having that under some flags or doing fancy stuff to display it
	// _, err = io.Copy(os.Stdout, resp)
	_, err = io.Copy(io.Discard, resp)
	if err != nil {
		log.Printf("Couldn't get docker output: %v", err)
	}

	registryNode, err := runtimes.SelectedRuntime.GetNode(ctx, &k3d.Node{Name: reg.Host})
	if err != nil {
		return err
	}
	regPort := registryNode.Ports["5000/tcp"][0].HostPort
	targetImage := fmt.Sprintf(
		"127.0.0.1:%s%s",
		regPort,
		strings.TrimLeftFunc(image, func(r rune) bool {
			return r != '/'
		}),
	)

	if err = dockerClient.ImageTag(ctx, image, targetImage); err != nil {
		return err
	}

	resp, err = dockerClient.ImagePush(ctx, targetImage, types.ImagePushOptions{
		RegistryAuth: "empty auth",
	})
	if err != nil {
		return fmt.Errorf("docker failed to push the image '%s': %w", targetImage, err)
	}
	defer resp.Close()

	log.Printf("Pushing image '%s'\n", targetImage)

	_, err = io.Copy(io.Discard, resp)
	if err != nil {
		log.Printf("Couldn't get docker output: %v", err)
	}

	log.Printf("Pushed image '%s'\n", targetImage)

	return nil
}

func CacheImagesInParallel(ctx context.Context, images []string, registry *k3d.Registry) {
	var wg sync.WaitGroup
	for _, img := range images {
		go func(img string) {
			wg.Add(1)
			defer wg.Done()
			if err := CacheImage(ctx, img, registry); err != nil {
				log.Fatalf("Failed to cache %s: %s", img, err)
			}
		}(img)
	}
	wg.Wait()
}

func CacheImageInTheBackground(ctx context.Context, image string) error {
	dockerClient, err := docker.GetDockerClient()
	if err != nil {
		return err
	}

	urlLength := strings.IndexRune(image, '/')
	targetImage := REGISTRY_DOMAIN + image[urlLength:]

	shellScript := strings.Join([]string{
		fmt.Sprintf("crictl pull %s", image),
		fmt.Sprintf("ctr image convert %s %s", image, targetImage),
		fmt.Sprintf("ctr image push --plain-http %s", targetImage),
	}, " && ")
	exec, err := dockerClient.ContainerExecCreate(ctx, CONTAINER_NAME, dockerTypes.ExecConfig{
		Privileged: true,
		Detach:     true,
		Cmd:        []string{"sh", "-c", shellScript},
	})
	if err != nil {
		return fmt.Errorf("docker failed to create exec config for node '%s': %w", CONTAINER_NAME, err)
	}

	return dockerClient.ContainerExecStart(ctx, exec.ID, dockerTypes.ExecStartCheck{})
}
