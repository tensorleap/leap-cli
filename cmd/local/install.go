package local

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/tensorleap/cli-go/pkg/helm"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
)

var port uint
var registryPort uint
var useGpu bool
var dataVolume string

func init() {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs tensorleap on the local machine, running in a docker container",
		Long:  `Installs tensorleap on the local machine, running in a docker container`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := local.InitVarDir(); err != nil {
				return err
			}

			if err := local.InitDataVolumeDir(dataVolume); err != nil {
				return err
			}

			ctx := cmd.Context()

			registryVolumes := []string{
				fmt.Sprintf("%v:%v", path.Join(local.VAR_DIR, "registry"), "/var/lib/registry"),
			}
			registry, err := k3d.CreateLocalRegistry(ctx, registryPort, registryVolumes)
			if err != nil {
				return err
			}

			imagesToCache, imageToCacheInTheBackground, err := local.GetLatestImages(useGpu)
			if err != nil {
				return err
			}
			k3d.CacheImagesInParallel(ctx, imagesToCache, registry)

			if err := k3d.CreateCluster(
				ctx,
				port,
				[]string{fmt.Sprintf("%v:%v", local.VAR_DIR, local.VAR_DIR), dataVolume},
				useGpu,
			); err != nil {
				return err
			}

			dataContainerPath := strings.Split(dataVolume, ":")[1]
			if err := helm.InstallLatestTensorleapChartVersion(
				ctx,
				local.KUBE_CONTEXT,
				local.KUBE_NAMESPACE,
				helm.CreateTensorleapChartValues(useGpu, dataContainerPath),
			); err != nil {
				return err
			}

			k3d.CacheImageInTheBackground(ctx, imageToCacheInTheBackground)

			return nil

		},
	}

	cmd.Flags().UintVarP(&port, "port", "p", 4589, "Port to be used for tensorleap installation")
	cmd.Flags().UintVar(&registryPort, "registry-port", 5699, "Port to be used for docker registry")
	cmd.Flags().BoolVar(&useGpu, "gpu", false, "Enable GPU usage for training and evaluating")
	cmd.Flags().StringVar(&dataVolume, "data-volume", local.GetDefaultDataVolume(), "Data Volume maps the user's local directory to the container's directory, enabling access to datasets for training and evaluation")

	RootCommand.AddCommand(cmd)
}
