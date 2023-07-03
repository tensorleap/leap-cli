package local

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/tensorleap/cli-go/pkg/helm"
	"github.com/tensorleap/cli-go/pkg/k3d"
)

const VAR_DIR = "/var/lib/tensorleap/standalone"

var port uint
var registryPort uint
var useGpu bool

func init() {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs tensorleap on the local machine, running in a docker container",
		Long:  `Installs tensorleap on the local machine, running in a docker container`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := initVarDir(); err != nil {
				return err
			}

			ctx := cmd.Context()

			registryVolumes := []string{
				fmt.Sprintf("%v:%v", path.Join(VAR_DIR, "registry"), "/var/lib/registry"),
			}
			registry, err := k3d.CreateLocalRegistry(ctx, registryPort, registryVolumes)
			if err != nil {
				return err
			}

			imagesToCache, err := getLatestImages(useGpu)
			if err != nil {
				return err
			}

			var wg sync.WaitGroup
			for _, img := range imagesToCache {
				go func(img string) {
					wg.Add(1)
					defer wg.Done()
					if err := k3d.CacheImage(ctx, img, registry); err != nil {
						log.Fatalf("Failed to cache %s: %s", img, err)
					}
				}(img)
			}
			wg.Wait()

			if err := k3d.CreateCluster(
				ctx,
				port,
				[]string{fmt.Sprintf("%v:%v", VAR_DIR, VAR_DIR)},
				useGpu,
			); err != nil {
				return err
			}

			if err := helm.InstallLatestTensorleapChartVersion(
				ctx,
				"k3d-tensorleap",
				"tensorleap",
				helm.CreateTensorleapChartValues(useGpu),
			); err != nil {
				return err
			}

			return nil

		},
	}

	cmd.Flags().UintVarP(&port, "port", "p", 4589, "Port to be used for tensorleap installation")
	cmd.Flags().UintVar(&registryPort, "registry-port", 5699, "Port to be used for docker registry")
	cmd.Flags().BoolVar(&useGpu, "gpu", false, "Enable GPU usage for training and evaluating")

	RootCommand.AddCommand(cmd)
}

func initVarDir() error {
	_, err := os.Stat(VAR_DIR)
	if os.IsNotExist(err) {
		log.Printf("Creating directory: %s (you may be asked to enter the root user password)", VAR_DIR)
		mkdirCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo mkdir -p %s", VAR_DIR))
		if err := mkdirCmd.Run(); err != nil {
			return err
		}

		log.Print("Setting directory permissions")
		chmodCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo chmod -R 777 %s", VAR_DIR))
		if err := chmodCmd.Run(); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	for _, dir := range []string{"storage", "registry"} {
		fullPath := path.Join(VAR_DIR, dir)
		_, err := os.Stat(fullPath)
		if os.IsNotExist(err) {
			log.Printf("Creating directory: %s", fullPath)
			if err := os.MkdirAll(fullPath, 777); err != nil {
				return err
			}
		} else if err != nil {
			return err
		}
	}

	return nil
}

func getLatestImages(useGpu bool) ([]string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/tensorleap/helm-charts/master/images.txt")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Getting latest chart images returned bad status code: %v", resp.StatusCode)
	}

	tensorleapImages, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	k3sVersion := k3d.K3sVersion
	if useGpu {
		k3sVersion = k3d.K3sGpuVersion
	}

	resp, err = http.Get(fmt.Sprintf("https://github.com/k3s-io/k3s/releases/download/%s/k3s-images.txt", strings.Replace(k3sVersion, "-", "+", 1)))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Getting latest k3s images returned bad status code: %v", resp.StatusCode)
	}

	k3sImages, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	allImages := strings.Split(string(tensorleapImages), "\n")
	allImages = append(allImages, strings.Split(string(k3sImages), "\n")...)

	ret := []string{}
	for _, img := range allImages {
		if len(img) > 0 && !strings.Contains(img, "engine") {
			ret = append(ret, img)
		}
	}

	return ret, nil
}
