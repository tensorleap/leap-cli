package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
)

func NewUpgradeCmd() *cobra.Command {
	flags := &server.UpgradeFlags{}

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: server.UpgradeCmdDescription,
		Long:  server.UpgradeCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			startProperties := map[string]interface{}{
				"tag": flags.Tag,
			}
			isInternetAvailable := checkInternetAvailability(len(flags.AirGapInstallationFilePath) > 0)

			analytics.SendEvent(analytics.EventServerUpgradeStarted, startProperties)

			_, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"tag":   flags.Tag,
					"error": err.Error(),
					"stage": "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties)
				return err
			}
			err = server.RunUpgradeCmd(cmd, flags)
			if err != nil {
				// Track upgrade failed
				failProperties := map[string]interface{}{
					"tag":   flags.Tag,
					"error": err.Error(),
					"stage": "run_upgrade_cmd",
				}
				analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties)
				return mapInstallationErr(err)
			}

			successProperties := map[string]interface{}{
				"tag": flags.Tag,
			}

			analytics.SendEvent(analytics.EventServerUpgradeSuccess, successProperties)
			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
