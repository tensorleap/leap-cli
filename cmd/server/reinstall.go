package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
)

func NewReinstallCmd() *cobra.Command {
	flags := &server.ReinstallFlags{}

	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: "Reinstall tensorleap",
		Long:  "Reinstall tensorleap",
		RunE: func(cmd *cobra.Command, args []string) error {
			isReinstalled, err := initDataDir(cmd.Context(), "")
			if err != nil {
				return err
			}
			err = server.RunReinstallCmd(cmd, flags, isReinstalled)
			if err != nil {
				return mapInstallationErr(err)
			}
			if err := localLogin(flags.Port); err != nil {
				return err
			}
			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
