package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/pkg/k3d"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/log"
	"gopkg.in/yaml.v3"
)

func NewInfoCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Server installation information",
		Long:  "Server installation information",
		RunE: func(cmd *cobra.Command, args []string) error {
			installationParams, err := server.LoadInstallationParamsFromPrevious()
			if err == server.ErrNoInstallationParams {
				log.Info("No installation information found")
				return nil
			}
			if err != nil {
				return err
			}

			log.Info("Installation information:")

			// Check if cluster exists before trying to get helm chart version
			cluster, err := k3d.GetCluster(cmd.Context())
			if err != nil {
				return err
			}
			if cluster == nil {
				log.Info("Server is not running (cluster not found)")
				installationParamsYaml, err := yaml.Marshal(installationParams)
				if err != nil {
					return err
				}
				log.Info(string(installationParamsYaml))
				return nil
			}

			currentVersion, err := server.GetCurrentInsalledHelmChartVersion(cmd.Context())
			if err != nil {
				return err
			}
			installationParams.Version = currentVersion
			installationParamsYaml, err := yaml.Marshal(installationParams)
			if err != nil {
				return err
			}
			log.Info(string(installationParamsYaml))
			return nil
		},
	}

	return cmd
}
