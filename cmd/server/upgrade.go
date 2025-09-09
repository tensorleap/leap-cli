package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/log"
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
			if err := analytics.SendEvent(analytics.EventServerUpgradeStarted, startProperties); err != nil {
				log.Warnf("Failed to track upgrade start event: %v", err)
			}

			_, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"tag":   flags.Tag,
					"error": err.Error(),
					"stage": "init_data_dir",
				}
				if err := analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties); err != nil {
					log.Warnf("Failed to track upgrade failure event: %v", err)
				}
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
				if err := analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties); err != nil {
					log.Warnf("Failed to track upgrade failure event: %v", err)
				}
				return mapInstallationErr(err)
			}

			successProperties := map[string]interface{}{
				"tag": flags.Tag,
			}

			if err := analytics.SendEvent(analytics.EventServerUpgradeSuccess, successProperties); err != nil {
				// Log error but don't fail the upgrade
				log.Warnf("Failed to track upgrade success event: %v", err)
			}
			
			recommendCliUpgradeMessage()

			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
