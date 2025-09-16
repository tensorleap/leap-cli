package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/local"
)

func NewUninstallCmd() *cobra.Command {
	var purge bool
	var cleanup bool
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove local Tensorleap installation",
		Long:  `Remove local Tensorleap installation`,
		RunE: func(cmd *cobra.Command, args []string) error {
			previousDir := getConfigureDataDir()
			err := local.SetDataDir(previousDir, previousDir)
			if err != nil {
				return err
			}

			return server.RunUninstallCmd(cmd, purge, cleanup)
		},
	}

	cmd.Flags().BoolVar(&purge, "purge", false, "Remove all data and cached files")
	cmd.Flags().BoolVar(&cleanup, "cleanup", false, "Remove cached data")

	return cmd
}
