package server

import (
	"github.com/spf13/cobra"
	servercmd "github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewInstallCmd() *cobra.Command {
	flags := &servercmd.InstallFlags{}
	licenseFlag := auth.NewLicenseFlag()

	cmd := &cobra.Command{
		Use:   "install",
		Short: servercmd.InstallCmdDescription,
		Long:  servercmd.InstallCmdDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Track installation started
			startProperties := map[string]interface{}{
				"cli_version":     version.CliVersion,
				"data_dir":        flags.DataDir,
				"tag":             flags.Tag,
				"port":            flags.Port,
				"registry_port":   flags.RegistryPort,
				"domain":          flags.Domain,
				"cpu_limit":       flags.CpuLimit,
				"gpus":            flags.Gpus,
				"gpu_devices":     flags.GpuDevices,
				"disable_metrics": flags.DisableMetrics,
				"proxy_url":       flags.ProxyUrl,
				"local":           flags.Local,
				"airgap_install":  len(flags.AirGapInstallationFilePath) > 0,
			}
			isInternetAvailable := checkInternetAvailability(len(flags.AirGapInstallationFilePath) > 0)

			analytics.SendEvent(analytics.EventServerInstallStarted, startProperties)

			_, err := initDataDir(cmd.Context(), flags.DataDir)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
				return err
			}

			installResult, err := servercmd.RunInstallCmd(cmd, flags)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "run_install_cmd",
				}
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
				return mapInstallationErr(err)
			}

			if err := localLogin(installResult.ServerURL, flags.Port); err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"port":        flags.Port,
					"error":       err.Error(),
					"stage":       "local_login",
				}
				analytics.SendEvent(analytics.EventServerInstallFailed, failProperties)
				return err
			}

			successProperties := map[string]interface{}{
				"cli_version":     version.CliVersion,
				"data_dir":        flags.DataDir,
				"tag":             flags.Tag,
				"port":            flags.Port,
				"registry_port":   flags.RegistryPort,
				"domain":          flags.Domain,
				"cpu_limit":       flags.CpuLimit,
				"gpus":            flags.Gpus,
				"gpu_devices":     flags.GpuDevices,
				"disable_metrics": flags.DisableMetrics,
				"proxy_url":       flags.ProxyUrl,
				"local":           flags.Local,
				"airgap_install":  len(flags.AirGapInstallationFilePath) > 0,
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
