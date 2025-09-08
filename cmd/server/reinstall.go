package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/log"
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
			if err := analytics.SendEvent(analytics.EventServerReinstallStarted, startProperties); err != nil {
				log.Warnf("Failed to track reinstallation start event: %v", err)
			}

			isReinstalled, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "init_data_dir",
				}
				if err := analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track reinstallation failure event: %v", err)
				}
				return err
			}
			err = server.RunReinstallCmd(cmd, flags, isReinstalled)
			if err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "run_reinstall_cmd",
				}
				if err := analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track reinstallation failure event: %v", err)
				}
				return mapInstallationErr(err)
			}
			if err := localLogin(flags.Port); err != nil {
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"port":     flags.Port,
					"error":    err.Error(),
					"stage":    "local_login",
				}
				if err := analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties); err != nil {
					log.Warnf("Failed to track reinstallation failure event: %v", err)
				}
				return err
			}

			successProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
				"port":     flags.Port,
			}

			if err := analytics.SendEvent(analytics.EventServerReinstallSuccess, successProperties); err != nil {
				log.Warnf("Failed to track reinstallation success event: %v", err)
			}

			return nil
		},
	}

	flags.SetFlags(cmd)

	return cmd
}
