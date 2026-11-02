package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/local"
)

func NewUninstallCmd() *cobra.Command {
	flags := &server.UninstallFlags{}
	cmd := &cobra.Command{
		Use:   "uninstall",
		Short: "Remove local Tensorleap installation",
		Long:  `Remove local Tensorleap installation`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Same order as install: the config's data_dir, else TL_DATA_DIR, else the default location.
			// Without the default, an empty config resolved to the current folder, and --purge then
			// re-permissioned it and deleted folders inside it
			err := local.SetDataDir(configuredOrDefaultDataDir(), getConfigureDataDir())
			if err != nil {
				return err
			}

			return runUninstall(cmd, flags)
		},
	}

	flags.AddToCommand(cmd)

	return cmd
}
