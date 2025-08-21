package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/log"
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
			err = server.RunInstallCmd(cmd, flags)
			if err != nil {
				return mapInstallationErr(err)
			}
			if err := localLogin(flags.Port); err != nil {
				return err
			}

			// Track successful installation
			properties := map[string]interface{}{}
			
			if err := analytics.TrackServerInstall(properties); err != nil {
				// Log error but don't fail the installation
				log.Warnf("Failed to track installation event: %v", err)
			}

			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
