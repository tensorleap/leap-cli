package server

import (
	"fmt"
	"time"

	"context"

	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/helm-charts/pkg/server/manifest"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/cli"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/version"
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

	// TODO: Skip login
	return nil
	// baseLink := fmt.Sprintf("http://127.0.0.1:%v", port)
	// apiLink := fmt.Sprintf("%s/api/v2", baseLink)
	// envName := auth.EnvNameFromUrl(apiLink)
	// authData := auth.NewEnv(envName, apiLink, "")
	// if err := auth.Login(authData); err != nil {
	// 	return err
	// }
	// return nil
}

func getConfigureDataDir() string {
	dataDir := viper.GetString(DATA_DIR_CONFIG_PATH)
	return dataDir
}

func initDataDir(ctx context.Context, flag string) (bool, error) {
	configureDir := getConfigureDataDir()
	previousDir := configureDir
	if previousDir == "" {
		previousDir = local.DEFAULT_DATA_DIR
	}
	err := local.SetDataDir(previousDir, flag)
	if err != nil {
		return false, err
	}
	currentDir := local.GetServerDataDir()

	isTransfer, err := server.TransferData(ctx)
	if err != nil {
		return false, err
	}
	if isTransfer || configureDir != currentDir {
		log.Infof("Saving data-dir (%s) to config", currentDir)
		viper.Set(DATA_DIR_CONFIG_PATH, currentDir)
		return isTransfer, config.Save()
	} else {
		log.Infof("Using data-dir: %s", currentDir)
	}
	return false, nil
}

func recommendCliUpgradeMessage() {
	currentVersion := version.CliVersion
	latestVersion, err := version.GetLatestVersion()
	if err != nil {
		log.Warnf("Failed to get latest version: %v", err)
		return
	}
	if currentVersion != latestVersion {
		log.Warnf("Your CLI version is %s, but the latest version is %s. Please upgrade your CLI. \n leap cli upgrade -s | bash", currentVersion, latestVersion)
	} else {
		log.Infof("Your CLI version is %s, which is the latest version.", currentVersion)
	}
}

func checkInternetAvailability(isAirGapInstallation bool) bool {
	internetShouldBeAvailable := !isAirGapInstallation
	if internetShouldBeAvailable {
		internetShouldBeAvailable = hasInternet()
		if !internetShouldBeAvailable {
			log.Warnf("No internet connection found, it may cause some issues during installation")
		}
	}
	return internetShouldBeAvailable
}

func hasInternet() bool {
	client := api.NewDefaultClient()
	client.Timeout = 3 * time.Second
	resp, err := client.Get("https://www.google.com/generate_204")
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == 204
}
