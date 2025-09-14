package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
)

func NewReinstallCmd() *cobra.Command {
	flags := &server.ReinstallFlags{}

	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: "Reinstall tensorleap",
		Long:  "Reinstall tensorleap",
		RunE: func(cmd *cobra.Command, args []string) error {
			startProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
			}
			isInternetAvailable := checkInternetAvailability(len(flags.AirGapInstallationFilePath) > 0)

			analytics.SendEvent(analytics.EventServerReinstallStarted, startProperties)

			isReinstalled, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
				return err
			}
			err = server.RunReinstallCmd(cmd, flags, isReinstalled)
			if err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "run_reinstall_cmd",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
				return mapInstallationErr(err)
			}
			if err := localLogin(flags.Port); err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"port":     flags.Port,
					"error":    err.Error(),
					"stage":    "local_login",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
				return err
			}

			successProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
				"port":     flags.Port,
			}

			analytics.SendEvent(analytics.EventServerReinstallSuccess, successProperties)
			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
