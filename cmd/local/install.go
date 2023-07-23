package local

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/auth"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/log"
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
			log.SetCommandName("install")
			log.SendCloudReport("info", "Starting installation", "Starting",
				&map[string]interface{}{"params": map[string]interface{}{"port": port, "registryPort": registryPort,
					"useGpu": useGpu, "dataVolume": dataVolume, "args": args}})

			close, err := local.SetupInfra("install")
			if err != nil {
				return err
			}
			defer close()

			if err := local.InitDataVolumeDir(dataVolume); err != nil {
				log.SendCloudReport("error", "Failed initializing data volume directory", "Failed",
					&map[string]interface{}{"dataVolume": dataVolume, "error": err.Error()})
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

			if err := k3d.CacheImageInTheBackground(ctx, imageToCacheInTheBackground); err != nil {
				return err
			}

			if err := auth.Login("", "http://localhost:4589/api/v2"); err != nil {
				return err
			}

			log.SendCloudReport("info", "Successfully completed installation", "Success", nil)
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
