package server

import (
	"fmt"

	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/helm-charts/pkg/server/manifest"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/cli"
)

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
	if err := auth.Login("", apiLink); err != nil {
		return err
	}
	return nil
}
