package server

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/helm-charts/pkg/server/manifest"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/cli"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
)

const DATA_DIR_CONFIG_PATH = "data_dir"

var ErrUnsupportedManifestVersion = fmt.Errorf(
	`unsupported manifest version. Please update your CLI.
Run the following command to upgrade:
%s`, cli.UpgradeCmd)

var ErrCliUpgradeRequired = fmt.Errorf(
	`please update your CLI to install latest version.
Run the following command to upgrade:
%s`, cli.UpgradeCmd)

var ErrOldManifest = fmt.Errorf(
	`you can't install this old manifest with this CLI. Try to use an newer version`,
)

func mapInstallationErr(err error) error {
	if err == manifest.ErrUnsupportedManifestVersion {
		return ErrUnsupportedManifestVersion
	}
	if err == server.ErrCliUpgradeRequired {
		return ErrCliUpgradeRequired
	}
	if err == server.ErrOldManifest {
		return ErrOldManifest
	}
	return err
}

func localLogin(port uint) error {
	baseLink := fmt.Sprintf("http://127.0.0.1:%v", port)
	apiLink := fmt.Sprintf("%s/api/v2", baseLink)
	envName := auth.EnvNameFromUrl(apiLink)
	authData := auth.NewEnv(envName, apiLink, "")
	if err := auth.Login(authData); err != nil {
		return err
	}
	return nil
}

func initDataDir(flag string) (func() error, error) {
	previousDir := viper.GetString(DATA_DIR_CONFIG_PATH)
	if previousDir == "" {
		previousDir = local.DEFAULT_DATA_DIR
		viper.Set(DATA_DIR_CONFIG_PATH, previousDir)
		err := config.Save()
		if err != nil {
			return nil, fmt.Errorf("failed to save data-dir to config: %v", err)
		}
	}
	err := local.SetDataDir(previousDir, flag)
	currentDir := local.GetServerDataDir()

	save := func() error {
		isDirNotChanged := previousDir == currentDir
		if isDirNotChanged {
			return nil
		}
		previousStatus, err := local.CheckDirectoryStatus(previousDir)
		if err != nil {
			return err
		}
		isStorageMigrated := !previousStatus.Exists
		if isStorageMigrated {
			log.Infof("Saving data-dir (%s) to config", currentDir)
			viper.Set(DATA_DIR_CONFIG_PATH, currentDir)
			return config.Save()
		}
		return nil
	}

	return save, err
}
