package server

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/helm"
	"github.com/tensorleap/cli-go/pkg/k3d"
	"github.com/tensorleap/cli-go/pkg/local"
	"github.com/tensorleap/cli-go/pkg/log"
	"github.com/tensorleap/cli-go/pkg/server"
)

func NewUpgradeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: "Upgrade an existing local tensorleap installation to the latest version",
		Long:  `Upgrade an existing local tensorleap installation to the latest version`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.SetCommandName("upgrade")
			log.SendCloudReport("info", "Starting upgrade", "Starting", &map[string]interface{}{"args": args})

			if err := server.ValidateStandaloneDir(); err != nil {
				return err
			}
			ctx := cmd.Context()

			close, err := local.SetupInfra("upgrade")
			if err != nil {
				return err
			}
			defer close()

			cluster, err := server.GetTensorleapCluster(ctx)
			if err != nil {
				return err
			}
			isGpuCluster := k3d.IsGpuCluster(cluster)

			helmConfig, err := helm.CreateHelmConfig(server.KUBE_CONTEXT, server.KUBE_NAMESPACE)
			if err != nil {
				return err
			}

			isHelmReleaseExisted, err := helm.IsHelmReleaseExists(helmConfig)
			if err != nil {
				return err
			} else if !isHelmReleaseExisted {
				return errors.New("Not found helm release, Please make sure to install before upgrade")
			}

			imagesToCache, imageToCacheInTheBackground, err := local.GetLatestImages(isGpuCluster)
			if err != nil {
				return err
			}

			registryPort, err := k3d.GetLocalRegistryPort(ctx)
			if err != nil {
				return err
			}
			k3d.CacheImagesInParallel(ctx, imagesToCache, registryPort)

			if err := helm.UpgradeTensorleapChartVersion(helmConfig); err != nil {
				return err
			}

			if err = k3d.CacheImageInTheBackground(ctx, imageToCacheInTheBackground); err != nil {
				return err
			}

			log.SendCloudReport("info", "Successfully completed upgrade", "Success", nil)
			return nil
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewUpgradeCmd())
}
