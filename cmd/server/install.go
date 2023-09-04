package server

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewInstallCmd() *cobra.Command {
	flags := &server.InstallFlags{}

	cmd := &cobra.Command{
		Use:   "install",
		Short: server.InstallCmdDescription,
		Long:  server.InstallCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := server.RunInstallCmd(cmd, flags)
			if err != nil {
				return err
			}
			baseLink := fmt.Sprintf("http://127.0.0.1:%v", flags.Port)
			apiLink := fmt.Sprintf("%s/api/v2", baseLink)
			if err := auth.Login("", apiLink); err != nil {
				return err
			}
			return nil
		},
	}

	server.SetInstallCmdFlags(cmd, flags)

	return cmd
}
