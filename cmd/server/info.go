package server

import (
	"github.com/spf13/cobra"
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
