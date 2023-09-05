package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
)

func NewUpgradeCmd() *cobra.Command {
	flags := &server.UpgradeFlags{}

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: server.UpgradeCmdDescription,
		Long:  server.UpgradeCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.RunUpgradeCmd(cmd, flags)
			if err != nil {
				return mapInstallationErr(err)
			}

			return nil
		},
	}

	server.SetUpgradeCmdFlags(cmd, flags)

	return cmd
}
