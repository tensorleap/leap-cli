package server

import (
	"os"

	"github.com/spf13/cobra"
	servercmd "github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/hint"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func NewInstallCmd() *cobra.Command {
	flags := &servercmd.InstallFlags{}
	licenseFlag := auth.NewLicenseFlag()
	var nonInteractive bool
	cmd := &cobra.Command{
		Use:   "install",
		Short: servercmd.InstallCmdDescription,
		Long:  servercmd.InstallCmdDescription,
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

			sendStartedEvent(analytics.EventServerInstallStarted, startProperties, email)

			_, err := initDataDir(cmd.Context(), flags.DataDir)
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "init_data_dir",
				}
				sendServerEvent(analytics.EventServerInstallFailed, failProperties, email)
				return err
			}

			res, err := runInstall(cmd, flags)
			// also after a failed attempt: the install skill records what went wrong there too
			ensureNotes()
			if err != nil {
				// Track installation failed
				failProperties := map[string]interface{}{
					"cli_version": version.CliVersion,
					"data_dir":    flags.DataDir,
					"error":       err.Error(),
					"stage":       "run_install_cmd",
				}
				sendServerEvent(analytics.EventServerInstallFailed, failProperties, email)
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

			sendServerEvent(analytics.EventServerInstallSuccess, successProperties, email)

			if isInternetAvailable {
				recommendCliUpgradeMessage()
			}

			err = applyLicense(cmd, licenseFlag)
			if err != nil {
				return err
			}

			hint.Show(cmd.OutOrStdout(), "Next", hint.Integrate, hint.Registered(serverURL(res)))
			return nil
		},
	}

	flags.SetFlags(cmd)
	licenseFlag.AddFlags(cmd)
	cmd.Flags().BoolVarP(&nonInteractive, "yes", "y", false, "Run in non-interactive mode (skip prompts)")
	return cmd
}
