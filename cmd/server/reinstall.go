package server

import (
	"os"

	"github.com/spf13/cobra"
	servercmd "github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewReinstallCmd() *cobra.Command {
	flags := &servercmd.ReinstallFlags{}
	licenseFlag := auth.NewLicenseFlag()
	var nonInteractive bool
	cmd := &cobra.Command{
		Use:   "reinstall",
		Short: "Reinstall tensorleap",
		Long:  "Reinstall tensorleap",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Set non-interactive mode via environment variable
			// This signals to helm-charts to use defaults and skip prompts
			if nonInteractive {
				_ = os.Setenv("TL_USE_DEFAULT_OPTION", "true")
			}

			logCurrentServerVersion(cmd.Context(), "Current server version")
			// before the first event, so a saved license gives the same events as --license-token
			licenseFlag.ResolveWithSaved()
			email := licenseTokenEmail(licenseFlag)

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

			sendServerEvent(analytics.EventServerReinstallStarted, startProperties, email)

			isReinstalled, err := initDataDir(cmd.Context(), "")
			if err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "init_data_dir",
				}
				sendServerEvent(analytics.EventServerReinstallFailed, failProperties, email)
				return err
			}
			_, err = runReinstall(cmd, flags, isReinstalled)
			ensureNotes()
			if err != nil {
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "run_reinstall_cmd",
				}
				sendServerEvent(analytics.EventServerReinstallFailed, failProperties, email)
				return mapInstallationErr(err)
			}

			logCurrentServerVersion(cmd.Context(), "Installed server version")

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

			sendServerEvent(analytics.EventServerReinstallSuccess, successProperties, email)
			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			err = applyLicense(cmd, licenseFlag)
			if err != nil {
				return err
			}

			return nil
		},
	}

	flags.SetFlags(cmd)
	licenseFlag.AddFlags(cmd)
	cmd.Flags().BoolVarP(&nonInteractive, "yes", "y", false, "Run in non-interactive mode (skip prompts)")
	return cmd
}
