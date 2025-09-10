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
			// Track installation started
			startProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
			}
			isInternetAvailable := checkInternetAvailability(len(flags.AirGapInstallationFilePath) > 0)

			if err := analytics.SendEvent(analytics.EventServerInstallStarted, startProperties); err != nil {
				log.Warnf("Failed to track installation start event: %v", err)
			}


			_, err := initDataDir(cmd.Context(), flags.DataDir)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "init_data_dir",
				}
				if err := analytics.SendEvent(analytics.EventServerInstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track installation failure event: %v", err)
				}
				return err
			}

			err = server.RunInstallCmd(cmd, flags)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "run_install_cmd",
				}
				if err := analytics.SendEvent(analytics.EventServerInstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track installation failure event: %v", err)
				}
				return mapInstallationErr(err)
			}

			if err := localLogin(flags.Port); err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"port":     flags.Port,
					"error":    err.Error(),
					"stage":    "local_login",
				}
				if err := analytics.SendEvent(analytics.EventServerInstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track installation failure event: %v", err)
				}
				return err
			}

			successProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
				"port":     flags.Port,
			}

			if err := analytics.SendEvent(analytics.EventServerInstallSuccess, successProperties); err != nil {
				// Log error but don't fail the installation
				log.Warnf("Failed to track installation success event: %v", err)
			}

			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}
			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
