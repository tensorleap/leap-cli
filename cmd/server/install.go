package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"

	leapdocker "github.com/tensorleap/leap-cli/pkg/docker"
)

func NewInstallCmd() *cobra.Command {
	flags := &server.InstallFlags{}

	cmd := &cobra.Command{
		Use:   "install",
		Short: server.InstallCmdDescription,
		Long:  server.InstallCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := initDataDir(cmd.Context(), flags.DataDir)
			if err != nil {
				return err
			}
			leapdocker.CheckNvidiaRuntime()
			err = server.RunInstallCmd(cmd, flags)
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
