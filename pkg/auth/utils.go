package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/version"
)

var ErrNotLoggedIn = fmt.Errorf("not logged in")

func Login(env *Env) (err error) {

	setCurrentEnv(env)
	if env.Name != "" {
		setEnvAuth(env)
	}

	err = config.Save()
	if err != nil {
		return fmt.Errorf("Login failed: %v", err)
	}
	log.Println("Saved credentials to: ", viper.ConfigFileUsed())

	// Track successful login
	properties := map[string]interface{}{
		"api_url":  env.ApiUrl,
		"env_name": env.Name,
		"method":   "api_key",
		"success":  true,
		"is_local": IsLocalUrl(env.ApiUrl),
	}

	// Try to get user data to enrich the analytics event
	// Create a context with the API URL for the whoami call
	ctx := api.CreateAuthenticatedContext(context.Background(), env.ApiKey, env.ApiUrl)
	if userData, err := GetWhoami(ctx); err == nil {
		properties["email"] = userData.Local.Email
		properties["user_name"] = userData.Local.Name
		properties["team_name"] = userData.TeamName
		properties["team_id"] = userData.Local.TeamId
		properties["role"] = string(userData.Role)
		properties["user_id"] = userData.Cid
		if envInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute(); err == nil && envInfo.LicenseMetadata != nil {
			properties["license_type"] = string(envInfo.LicenseMetadata.Type)
			properties["license_activated"] = LicenseActivated(envInfo.LicenseMetadata)
		}
	} else {
		// Log warning but don't fail the tracking
		log.Warnf("Failed to get user data for analytics: %v", err)
	}

	analytics.SendEvent(analytics.EventAuthLoginSuccess, properties)

	return
}

func SelectAndSetEnv(envName string) error {
	auth, err := SelectEnv("Select an environment to use", envName)
	if err != nil {
		return err
	}
	setCurrentEnv(auth)
	err = config.Save()
	if err != nil {
		return fmt.Errorf("select failed: %v", err)
	}
	log.Println("Selected environment: ", auth.Name)
	return nil
}

// UseEnv makes name the current environment, as `leap auth select <name>` does, without prompting or printing
func UseEnv(name string) error {
	for _, env := range GetEnvs() {
		if env.Name == name {
			setCurrentEnv(&env)
			return config.Save()
		}
	}
	return fmt.Errorf("no 'leap auth' environment named %q", name)
}

func SelectEnv(msg string, envName string) (*Env, error) {
	currentAuth := GetCurrentEnv().Name
	envs := GetEnvs()
	if envName != "" {
		for _, env := range envs {
			if env.Name == envName {
				return &env, nil
			}
		}
		log.Warn("No found environment with name: ", envName)
	}

	options := make([]string, len(envs))
	defaultOption := 0
	for i, env := range envs {
		if env.Name == currentAuth {
			defaultOption = i
		}
		options[i] = env.Name
	}

	selectIndex := -1
	prompt := &survey.Select{
		Message: msg,
		Options: options,
		Default: defaultOption,
	}

	err := survey.AskOne(prompt, &selectIndex)
	if err != nil {
		return nil, err
	}
	return &envs[selectIndex], nil
}

func Logout(envName string) error {
	removeEnvAuth(envName)
	err := config.Save()
	if err != nil {
		return fmt.Errorf("Logout failed: %v", err)
	}
	log.Infof("Logged out")
	return nil
}

func IsLocalUrl(url string) bool {
	return !strings.Contains(url, ".tensorleap.ai")
}

func EnvNameFromUrl(url string) string {
	isLocal := IsLocalUrl(url)
	if isLocal {
		return "local"
	}
	trimUrl := strings.TrimPrefix(url, "https://api.")
	return strings.Split(trimUrl, ".")[0]
}

func PrintWhoami(ctx context.Context) error {
	apiUrl, _ := api.GetAuthFromContext(ctx)
	fmt.Println("API Url: " + apiUrl)

	userData, err := GetWhoami(ctx)
	if err != nil {
		return err
	}

	fmt.Println("User email: " + userData.Local.Email)
	fmt.Println("Team name: " + userData.TeamName)
	if envInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute(); err == nil {
		if line := LicenseLine(envInfo); line != "" {
			fmt.Println(line)
		}
	}
	return nil
}

// installedServerUrl is the address of the server `leap server install` set up on this machine, "" when there is none
var installedServerUrl = func() string {
	params, err := server.LoadInstallationParamsFromPrevious()
	if err != nil {
		return ""
	}
	return params.CalcUrl()
}

// KnownServerUrl is the server the user most likely signs in to, from what this machine knows without network:
// the configured API URL, else the local install's address; "" when neither is known
func KnownServerUrl() string {
	if url := GetCurrentEnv().ApiUrl; url != "" {
		return api.ChangeToUIUrl(url)
	}
	return installedServerUrl()
}

// NotSignedInError tells how to sign in, naming the server when the CLI knows it
func NotSignedInError() error {
	url := KnownServerUrl()
	if url == "" {
		url = "<url>"
	}
	return fmt.Errorf("Not signed in. Run: leap auth login %s", url)
}

// signInCheckTimeout bounds CheckSignIn, so a server that never answers doesn't hang the command
var signInCheckTimeout = 15 * time.Second

// CheckSignIn asks the server whose key env holds. Only a key the server rejects means "not signed in";
// a server that can't be reached is reported as such, since signing in again wouldn't help.
func CheckSignIn(ctx context.Context, env *Env) error {
	ctx, cancel := context.WithTimeout(env.AuthContext(ctx), signInCheckTimeout)
	defer cancel()
	_, resp, err := api.ApiClient.WhoAmI(ctx).Execute()
	switch {
	case err == nil:
		return nil
	// the API client turns a 401 into ErrAuth without a response
	case errors.Is(err, api.ErrAuth), resp != nil && resp.StatusCode == http.StatusForbidden:
		return NotSignedInError()
	case resp == nil:
		return fmt.Errorf("Can't reach the Tensorleap server at %s: %s", api.ChangeToUIUrl(env.ApiUrl), transportCause(err))
	}
	return fmt.Errorf("Could not check your sign-in at %s: %v", api.ChangeToUIUrl(env.ApiUrl), err)
}

// transportCause drops the request line Go puts in front of a network error (Get "<url>": ...)
func transportCause(err error) string {
	var urlErr *url.Error
	if errors.As(err, &urlErr) {
		err = urlErr.Err
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return "no answer in time"
	}
	return err.Error()
}

// RequireAuth checks if the user is authenticated and validates the credentials.
// If authentication is missing or invalid, it automatically triggers the interactive login flow.
// Returns the authenticated context and user data if successful.
func RequireAuth(ctx context.Context) (context.Context, *tensorleapapi.UserData, error) {
	env := GetCurrentEnv()

	if env.ApiUrl != "" && env.ApiKey != "" {
		authCtx := env.AuthContext(ctx)
		userData, err := GetWhoami(authCtx)
		if err == nil {
			version.CheckServerCompatibility(authCtx)
			return authCtx, userData, nil
		}
		log.Warnf("Current authentication is invalid or expired: %v", err)
	}

	return InteractiveLogin(ctx)
}

// InteractiveLogin prompts the user through the authentication flow interactively.
// It reuses the existing API URL if one is configured, otherwise asks for it.
// Returns the authenticated context and user data after successful login.
func InteractiveLogin(ctx context.Context) (context.Context, *tensorleapapi.UserData, error) {
	log.Info("Authentication required. Starting login flow...")

	env := GetCurrentEnv()
	var apiUrl string

	if env.ApiUrl != "" {
		log.Infof("Using configured API URL: %s", env.ApiUrl)
		apiUrl = env.ApiUrl
	} else {
		baseUrl, err := AskForUrl("")
		if err != nil {
			return ctx, nil, err
		}
		normalized, err := api.NormalizeAPIUrl(baseUrl)
		if err != nil {
			return ctx, nil, err
		}
		apiUrl = normalized
	}

	useLogin, err := AskIfUseLogin()
	if err != nil {
		return ctx, nil, err
	}

	var apiKey string
	if useLogin {
		useBrowser, err := AskIfOpenBrowser()
		if err != nil {
			return ctx, nil, err
		}
		if useBrowser {
			apiKey, err = LoginAndGetAuthTokenWithBrowser(ctx, apiUrl, true)
			if err != nil {
				return ctx, nil, err
			}
		} else {
			userName, password, err := AskForUserNameAndPassword("", "")
			if err != nil {
				return ctx, nil, err
			}
			apiKey, err = LoginAndGetAuthToken(apiUrl, userName, password)
			if err != nil {
				return ctx, nil, err
			}
		}
	} else {
		apiKey, err = AskForApiKey()
		if err != nil {
			return ctx, nil, err
		}
	}

	apiKey = strings.TrimSpace(apiKey)

	name := EnvNameFromUrl(apiUrl)
	if env.Name != "" {
		name = env.Name
	}

	newEnv := NewEnv(name, apiUrl, apiKey)
	authCtx := newEnv.AuthContext(ctx)

	userData, err := GetWhoami(authCtx)
	if err != nil {
		return ctx, nil, fmt.Errorf("login verification failed: %v", err)
	}

	version.CheckServerCompatibility(authCtx)

	fmt.Printf("Logged in as: %s (%s)\n", userData.Local.Email, userData.TeamName)

	if err := Login(newEnv); err != nil {
		return ctx, nil, err
	}

	return authCtx, userData, nil
}

func InitMaybeUnauthedContext(ctx context.Context, defaultUrl string) (context.Context, error) {
	baseUrl, _ := api.GetAuthFromContext(ctx)
	if baseUrl != "" {
		return ctx, nil
	}

	if defaultUrl == "" {
		log.Info("No URL found use `leap auth login` or enter the url")
		prompt := survey.Input{
			Message: "Server URL",
			Default: fmt.Sprintf("http://localhost:%v", server.DefaultHttpPort),
		}
		err := survey.AskOne(&prompt, &baseUrl)
		if err != nil {
			return ctx, err
		}
	} else {
		baseUrl = defaultUrl
	}
	baseUrl, err := api.NormalizeAPIUrl(baseUrl)
	if err != nil {
		return ctx, err
	}
	return api.CreateAuthenticatedContext(ctx, "", baseUrl), nil
}
