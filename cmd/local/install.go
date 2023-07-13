package local

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
)

func NewInstallCmd() *cobra.Command {
	var port uint
	var registryPort uint
	var useGpu bool
	var dataVolume string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs tensorleap on the local machine, running in a docker container",
		Long:  `Installs tensorleap on the local machine, running in a docker container`,
		RunE: func(cmd *cobra.Command, args []string) error {

			close, err := local.SetupInfra("install")
			if err != nil {
				return err
			}
			defer close()

			if err := local.InitDataVolumeDir(dataVolume); err != nil {
				return err
			}

			ctx := cmd.Context()

			registryVolumes := []string{
				fmt.Sprintf("%v:%v", path.Join(local.STANDALONE_DIR, "registry"), "/var/lib/registry"),
			}
			registry, err := k3d.CreateLocalRegistry(ctx, registryPort, registryVolumes)
			if err != nil {
				return err
			}

			registryPortStr, err := k3d.GetRegistryPort(ctx, registry)
			if err != nil {
				return err
			}

			imagesToCache, imageToCacheInTheBackground, err := local.GetLatestImages(useGpu)
			if err != nil {
				return err
			}
			k3d.CacheImagesInParallel(ctx, imagesToCache, registryPortStr)

			if err := k3d.CreateCluster(
				ctx,
				port,
				[]string{fmt.Sprintf("%v:%v", local.STANDALONE_DIR, local.STANDALONE_DIR), dataVolume},
				useGpu,
			); err != nil {
				return err
			}

			dataContainerPath := strings.Split(dataVolume, ":")[1]
			if err := local.InstallHelm(useGpu, dataContainerPath); err != nil {
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

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInstallCmd())
}
