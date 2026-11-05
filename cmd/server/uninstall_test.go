package server

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	servercmd "github.com/tensorleap/helm-charts/cmd/server"
	"github.com/tensorleap/helm-charts/pkg/local"
)

// runUninstallSeeingDataDir runs `leap server uninstall` with the real work stubbed and returns the
// data dir the uninstall would act on
func runUninstallSeeingDataDir(t *testing.T, configDataDir string) string {
	t.Helper()
	prevConfig := viper.GetString(DATA_DIR_CONFIG_PATH)
	viper.Set(DATA_DIR_CONFIG_PATH, configDataDir)
	t.Cleanup(func() { viper.Set(DATA_DIR_CONFIG_PATH, prevConfig) })

	var seen string
	prevRun := runUninstall
	runUninstall = func(*cobra.Command, *servercmd.UninstallFlags) error {
		seen = local.GetServerDataDir()
		return nil
	}
	t.Cleanup(func() { runUninstall = prevRun })

	cmd := NewUninstallCmd()
	cmd.SetArgs([]string{})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("uninstall: %v", err)
	}
	return seen
}

// G88: with no data_dir in the config, uninstall used to act on the current folder
func TestUninstallWithoutConfiguredDataDirUsesTheDefault(t *testing.T) {
	t.Setenv(local.DATA_DIR_ENV_NAME, "")
	if got := runUninstallSeeingDataDir(t, ""); got != local.DEFAULT_DATA_DIR {
		t.Fatalf("uninstall acts on %q, want the default %q", got, local.DEFAULT_DATA_DIR)
	}
}

func TestUninstallUsesTheConfiguredDataDir(t *testing.T) {
	t.Setenv(local.DATA_DIR_ENV_NAME, "/somewhere/else")
	if got := runUninstallSeeingDataDir(t, "/data/tensorleap"); got != "/data/tensorleap" {
		t.Fatalf("uninstall acts on %q, want the configured /data/tensorleap", got)
	}
}

func TestUninstallWithoutConfiguredDataDirHonorsTheEnvVariable(t *testing.T) {
	t.Setenv(local.DATA_DIR_ENV_NAME, "/from/env")
	if got := runUninstallSeeingDataDir(t, ""); got != "/from/env" {
		t.Fatalf("uninstall acts on %q, want TL_DATA_DIR /from/env", got)
	}
}
