package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewInstallCmd() *cobra.Command {
	flags := &server.InstallFlags{}
	licenseFlag := auth.NewLicenseFlag()

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

			analytics.SendEvent(analytics.EventServerInstallStarted, startProperties)

			_, err := initDataDir(cmd.Context(), flags.DataDir)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"data_dir": flags.DataDir,
					"error":    err.Error(),
					"stage":    "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
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
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
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
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
				return err
			}

			successProperties := map[string]interface{}{
				"data_dir": flags.DataDir,
				"port":     flags.Port,
			}

			analytics.SendEvent(analytics.EventServerInstallSuccess, successProperties)

			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			err = handleLicenseAfterInstall(cmd, licenseFlag)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags.SetFlags(cmd)
	licenseFlag.AddFlags(cmd)

	return cmd
}
