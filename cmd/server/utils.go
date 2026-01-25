package server

import (
	"fmt"
	"time"

	"context"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/helm-charts/pkg/server/manifest"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/cli"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
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

type postInstallUserState struct {
	HasAnyUsers bool
	IsLoggedIn  bool
}

func fetchAuthStatus(ctx context.Context) (postInstallUserState, error) {
	res, _, err := api.ApiClient.GetAuthStatus(ctx).Execute()
	if err != nil {
		return postInstallUserState{}, err
	}

	switch res.Status {
	case tensorleapapi.AUTHSTATUS_LOGGED_IN:
		return postInstallUserState{HasAnyUsers: true, IsLoggedIn: true}, nil
	case tensorleapapi.AUTHSTATUS_NO_USERS:
		return postInstallUserState{HasAnyUsers: false, IsLoggedIn: false}, nil
	case tensorleapapi.AUTHSTATUS_NOT_LOGGED_IN:
		return postInstallUserState{HasAnyUsers: true, IsLoggedIn: false}, nil
	default:
		return postInstallUserState{}, fmt.Errorf("unknown auth status: %s", res.Status)
	}
}

func localLogin(baseUrl string, fallbackPort uint) error {
	apiLink := ""
	uiBaseUrl := ""

	if baseUrl != "" {
		normalizedApi, err := api.NormalizeAPIUrl(baseUrl)
		if err != nil {
			log.Warnf("Failed to normalize installation URL %q: %v", baseUrl, err)
		} else {
			apiLink = normalizedApi
			uiBaseUrl = api.ChangeToUIUrl(apiLink)
		}
	}

	if apiLink == "" {
		baseLink := fmt.Sprintf("http://127.0.0.1:%v", fallbackPort)
		apiLink = fmt.Sprintf("%s/api/v2", baseLink)
		uiBaseUrl = api.ChangeToUIUrl(apiLink)
	}

	// Check if we already have a stored token for this URL
	existingApiKey := ""
	currentEnv := auth.GetCurrentEnv()
	if currentEnv != nil && currentEnv.ApiKey != "" && currentEnv.ApiUrl == apiLink {
		existingApiKey = currentEnv.ApiKey
	} else {
		// Try to find token by env name (e.g. "local")
		envName := auth.EnvNameFromUrl(apiLink)
		if envAuth, err := auth.GetEnvAuth(envName); err == nil && envAuth.ApiKey != "" {
			existingApiKey = envAuth.ApiKey
		}
	}

	var ctx context.Context
	if existingApiKey != "" {
		ctx = api.CreateAuthenticatedContext(context.Background(), existingApiKey, apiLink)
	} else {
		ctx = api.CreateContextWithBaseURL(context.Background(), apiLink)
	}

	userState, err := fetchAuthStatus(ctx)
	if err != nil {
		log.Warnf("Failed to get auth status, falling back to login: %v", err)
		userState = postInstallUserState{HasAnyUsers: true, IsLoggedIn: false}
	}

	switch {
	case userState.IsLoggedIn:
		log.Infof("User already logged in, opening Tensorleap at %s", uiBaseUrl)
		return local.OpenLink(uiBaseUrl)
	default:
		ctx := context.Background()
		apiKey, err := auth.LoginAndGetAuthTokenWithBrowser(ctx, apiLink)
		if err != nil {
			return fmt.Errorf("automatic login failed: %w", err)
		}
		env := auth.NewEnv(auth.EnvNameFromUrl(apiLink), apiLink, apiKey)
		if err := env.PrintWhoami(ctx); err != nil {
			return fmt.Errorf("failed whoami after login: %w", err)
		}
		return auth.Login(env)
	}
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
	_ = resp.Body.Close()
	return resp.StatusCode == 204
}

func handleLicenseAfterInstall(cmd *cobra.Command, licenseFlag *auth.LicenseFlag) error {
	if !licenseFlag.HasLicense() {
		return nil
	}
	installationParams, err := server.LoadInstallationParamsFromPrevious()
	if err != nil {
		return err
	}
	baseUrl := installationParams.CalcUrl()
	var ctx context.Context
	ctx, err = auth.InitMaybeUnauthedContext(cmd.Context(), baseUrl)
	if err != nil {
		return err
	}
	var licenseToken string
	licenseToken, err = licenseFlag.PrepareLicenseToken()
	if err != nil {
		return err
	}
	return auth.SetLicense(ctx, licenseToken)
}
