package k3d

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	"github.com/tensorleap/cli-go/pkg/log"
)

type Registry = k3d.Registry

const REGISTRY_NAME = "k3d-tensorleap-registry"
const CONTAINER_NAME = "k3d-tensorleap-server-0"
const REGISTRY_DOMAIN = "k3d-tensorleap-registry:5000"

type RegistryTagListResponse struct {
	Name string
	Tags []string
}

func GetLocalRegistryPort(ctx context.Context) (string, error) {
	reg, err := client.RegistryGet(ctx, runtimes.SelectedRuntime, REGISTRY_NAME)
	if err != nil {
		log.SendCloudReport("error", "Failed getting local registry port", "Failed",
			&map[string]interface{}{"registryName": REGISTRY_NAME, "selectedRuntime": runtimes.SelectedRuntime, "error": err.Error()})
		return "", err
	}

	return GetRegistryPort(ctx, reg)
}

func GetRegistryPort(ctx context.Context, reg *Registry) (string, error) {
	registryNode, err := runtimes.SelectedRuntime.GetNode(ctx, &k3d.Node{Name: reg.Host})
	if err != nil {
		return "", err
	}

	regPort := registryNode.Ports["5000/tcp"][0].HostPort
	return regPort, nil
}

func CreateLocalRegistry(ctx context.Context, port uint, volumes []string) (*Registry, error) {
	if existingRegistry, _ := client.RegistryGet(ctx, runtimes.SelectedRuntime, REGISTRY_NAME); existingRegistry != nil {
		log.Println("Found existing registry!")
		log.SendCloudReport("info", "Found existing registry", "Running", &map[string]interface{}{"registryName": REGISTRY_NAME, "existingRegistry": existingRegistry})

		return existingRegistry, nil
	}

	reg := createRegistryConfig(port, volumes)
	_, err := client.RegistryRun(ctx, runtimes.SelectedRuntime, reg)
	if err != nil {
		log.SendCloudReport("error", "Failed running k3d registry", "Failed",
			&map[string]interface{}{"registryName": REGISTRY_NAME, "selectedRuntime": runtimes.SelectedRuntime, "port": port, "volumes": volumes, "error": err.Error()})
		return nil, err
	}

	log.SendCloudReport("info", "Successfully created k3d regisrty", "Running", &map[string]interface{}{"registryName": REGISTRY_NAME})
	return reg, nil
}

func createRegistryConfig(port uint, volumes []string) *Registry {
	exposePort, err := cliutil.ParsePortExposureSpec(strconv.FormatUint(uint64(port), 10), k3d.DefaultRegistryPort)
	if err != nil {
		log.SendCloudReport("error", "Failed creating k3d registry config", "Failed",
			&map[string]interface{}{"defaultRegistry": k3d.DefaultRegistryPort, "port": port, "error": err.Error()})
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

func isImageInRegistry(ctx context.Context, image string, regPort string) (bool, error) {
	imageParts := strings.SplitN(image, ":", 2)
	imageTag := imageParts[1]
	urlLength := strings.IndexRune(imageParts[0], '/')
	imageFullPath := imageParts[0][urlLength:]
	tagsListUrl := fmt.Sprintf("http://127.0.0.1:%s/v2%s/tags/list", regPort, imageFullPath)

	resp, err := http.Get(tagsListUrl)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return false, nil
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return false, fmt.Errorf("Local registry returned bad status code: %v", resp.StatusCode)
	}

	tagsListRaw, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	tagsList := RegistryTagListResponse{}
	if err = json.Unmarshal(tagsListRaw, &tagsList); err != nil {
		return false, err
	}

	for _, tag := range tagsList.Tags {
		if tag == imageTag {
			return true, nil
		}
	}

	return false, nil
}

func CacheImage(ctx context.Context, image string, regPort string) error {
	imageAlreadyInRegistry, err := isImageInRegistry(ctx, image, regPort)
	if err != nil {
		return err
	}
	if imageAlreadyInRegistry {
		log.Infof("Image already cached '%s'\n", image)
		return nil
	}

	dockerClient, err := docker.GetDockerClient()
	if err != nil {
		return err
	}

	resp, err := dockerClient.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		return fmt.Errorf("docker failed to pull the image '%s': %w", image, err)
	}
	defer resp.Close()

	log.Infof("Pulling image '%s'\n", image)

	// this prints out status of the pull, consider having that under some flags or doing fancy stuff to display it
	// _, err = io.Copy(os.Stdout, resp)
	_, err = io.Copy(io.Discard, resp)
	if err != nil {
		log.Warnf("Couldn't get docker output: %v", err)
	}

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

	log.Infof("Pushing image '%s'\n", targetImage)

	_, err = io.Copy(io.Discard, resp)
	if err != nil {
		log.Printf("Couldn't get docker output: %v", err)
	}

	log.Printf("Pushed image '%s'\n", targetImage)

	return nil
}

func CacheImagesInParallel(ctx context.Context, images []string, regPort string) {
	var wg sync.WaitGroup
	log.Info("Downloading docker images...")
	for _, img := range images {
		go func(img string) {
			wg.Add(1)
			defer wg.Done()
			if err := CacheImage(ctx, img, regPort); err != nil {
				log.SendCloudReport("error", "Failed caching image", "Failed", &map[string]interface{}{"image": img, "error": err.Error()})
				log.Fatalf("Failed to cache %s: %s", img, err)
			}
		}(img)
	}
	wg.Wait()
	log.SendCloudReport("info", "Successfully cached images in parallel", "Running", nil)
}

func CacheImageInTheBackground(ctx context.Context, image string) error {
	regPort, err := GetLocalRegistryPort(ctx)
	if err != nil {
		return err
	}
	imageAlreadyInRegistry, err := isImageInRegistry(ctx, image, regPort)
	if err != nil {
		return err
	}

	dockerClient, err := docker.GetDockerClient()
	if err != nil {
		log.SendCloudReport("error", "Failed fetching docker client", "Failed", &map[string]interface{}{"error": err.Error()})
		return err
	}

	urlLength := strings.IndexRune(image, '/')
	targetImage := REGISTRY_DOMAIN + image[urlLength:]

	shellScript := fmt.Sprintf("crictl pull %s", image)
	if !imageAlreadyInRegistry {
		shellScript = strings.Join([]string{
			shellScript,
			fmt.Sprintf("ctr image convert %s %s", image, targetImage),
			fmt.Sprintf("ctr image push --plain-http %s", targetImage),
		}, " && ")
	}
	exec, err := dockerClient.ContainerExecCreate(ctx, CONTAINER_NAME, dockerTypes.ExecConfig{
		Privileged: true,
		Detach:     true,
		Cmd:        []string{"sh", "-c", shellScript},
	})
	if err != nil {
		log.SendCloudReport("error", "Failed creating exec config for node", "Failed",
			&map[string]interface{}{"containerName": CONTAINER_NAME, "error": err.Error()})
		return fmt.Errorf("docker failed to create exec config for node '%s': %w", CONTAINER_NAME, err)
	}

	log.SendCloudReport("info", "Successfully cached images in background", "Running", nil)
	return dockerClient.ContainerExecStart(ctx, exec.ID, dockerTypes.ExecStartCheck{})
}
