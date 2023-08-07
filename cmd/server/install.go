package server

import (
	"fmt"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/k3d"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/server"
)

func NewInstallCmd() *cobra.Command {
	var port uint
	var registryPort uint
	var useGpu bool
	var useCpu bool
	var datasetDirectory string

	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs tensorleap on the local machine, running in a docker container",
		Long:  `Installs tensorleap on the local machine, running in a docker container`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetCommandName("install")

			err := k3d.CheckDockerRequirements()
			if err != nil {
				log.SendCloudReport("error", "Docker requirements not met", "Failed",
					&map[string]interface{}{"error": err.Error()})
				return err
			}

			close, err := local.SetupInfra("install")
			if err != nil {
				return err
			}
			defer close()

			if err := server.InitUseGPU(&useGpu, useCpu); err != nil {
				log.SendCloudReport("error", "Failed to initializing with gpu", "Failed",
					&map[string]interface{}{"useGpu": useCpu, "error": err.Error()})
				return err
			}

			if err := server.InitDatasetDirectory(&datasetDirectory); err != nil {
				log.SendCloudReport("error", "Failed initializing data volume directory", "Failed",
					&map[string]interface{}{"datasetDirectory": datasetDirectory, "error": err.Error()})
				return err
			}

			log.SendCloudReport("info", "Starting installation", "Starting",
				&map[string]interface{}{"params": map[string]interface{}{"port": port, "registryPort": registryPort,
					"useGpu": useGpu, "datasetDirectory": datasetDirectory, "args": args}})

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
				[]string{fmt.Sprintf("%v:%v", local.STANDALONE_DIR, local.STANDALONE_DIR), datasetDirectory},
				useGpu,
			); err != nil {
				return err
			}

			dataContainerPath := strings.Split(datasetDirectory, ":")[1]
			if err := server.InstallHelm(useGpu, dataContainerPath); err != nil {
				return err
			}

			if err := k3d.CacheImageInTheBackground(ctx, imageToCacheInTheBackground); err != nil {
				return err
			}
			baseLink := fmt.Sprintf("http://127.0.0.1:%v", port)
			apiLink := fmt.Sprintf("%s/api/v2", baseLink)
			if err := auth.Login("", apiLink); err != nil {
				return err
			}

			log.SendCloudReport("info", "Successfully completed installation", "Success", nil)
			log.Info("Successfully completed installation")
			_ = local.OpenLink(baseLink)
			log.Infof("You can now access Tensorleap at %s", baseLink)
			return nil
		},
	}

	cmd.Flags().UintVarP(&port, "port", "p", 4589, "Port to be used for tensorleap installation")
	cmd.Flags().UintVar(&registryPort, "registry-port", 5699, "Port to be used for docker registry")
	cmd.Flags().BoolVar(&useGpu, "gpu", false, "Enable GPU usage for training and evaluating")
	cmd.Flags().BoolVar(&useCpu, "cpu", false, "Use CPU for training and evaluating")
	cmd.Flags().StringVar(&datasetDirectory, "dataset-dir", "", "Dataset directory maps the user's local directory to the container's directory, enabling access to code integration for training and evaluation")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInstallCmd())
}
