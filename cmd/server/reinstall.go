package server

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewReinstallCmd() *cobra.Command {
	flags := &server.ReinstallFlags{}
	licenseFlag := auth.NewLicenseFlag()

	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: "Reinstall tensorleap",
		Long:  "Reinstall tensorleap",
		RunE: func(cmd *cobra.Command, args []string) error {
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

			analytics.SendEvent(analytics.EventServerReinstallStarted, startProperties)

			isReinstalled, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "init_data_dir",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
				return err
			}
			err = server.RunReinstallCmd(cmd, flags, isReinstalled)
			if err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "run_reinstall_cmd",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
				return mapInstallationErr(err)
			}
			if err := localLogin(flags.Port); err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"port":        flags.Port,
					"error":       err.Error(),
					"stage":       "local_login",
				}
				analytics.SendEvent(analytics.EventServerReinstallFailed, failProperties)
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

			analytics.SendEvent(analytics.EventServerReinstallSuccess, successProperties)
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
