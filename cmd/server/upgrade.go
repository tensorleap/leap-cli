package server

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewUpgradeCmd() *cobra.Command {
	flags := &server.UpgradeFlags{}
	var nonInteractive bool

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: server.UpgradeCmdDescription,
		Long:  server.UpgradeCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set non-interactive mode via environment variable
			// This signals to helm-charts to use defaults and skip prompts
			if nonInteractive {
				os.Setenv("TL_USE_DEFAULT_OPTION", "true")
			}

			startProperties := map[string]interface{}{
				"cli_version":    version.CliVersion,
				"tag":            flags.Tag,
				"local":          flags.Local,
				"airgap_install": len(flags.AirGapInstallationFilePath) > 0,
			}
			isInternetAvailable := checkInternetAvailability(len(flags.AirGapInstallationFilePath) > 0)

			analytics.SendEvent(analytics.EventServerUpgradeStarted, startProperties)

			_, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"tag":         flags.Tag,
					"error":       err.Error(),
					"stage":       "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties)
				return err
			}
			_, err = server.RunUpgradeCmd(cmd, flags)
			if err != nil {
				// Track upgrade failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"tag":         flags.Tag,
					"error":       err.Error(),
					"stage":       "run_upgrade_cmd",
				}
				analytics.SendEvent(analytics.EventServerUpgradeFailed, failProperties)
				return mapInstallationErr(err)
			}

			successProperties := map[string]interface{}{
				"cli_version":    version.CliVersion,
				"tag":            flags.Tag,
				"local":          flags.Local,
				"airgap_install": len(flags.AirGapInstallationFilePath) > 0,
			}

			analytics.SendEvent(analytics.EventServerUpgradeSuccess, successProperties)
			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			return nil
		},
	}

	flags.SetFlags(cmd)
	cmd.Flags().BoolVarP(&nonInteractive, "yes", "y", false, "Run in non-interactive mode (skip prompts)")

	return cmd
}
