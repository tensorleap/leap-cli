package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"github.com/tensorleap/helm-charts/pkg/server"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
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
		properties["user_email"] = userData.Local.Email
		properties["user_name"] = userData.Local.Name
		properties["team_name"] = userData.TeamName
		properties["team_id"] = userData.Local.TeamId
		properties["user_role"] = string(userData.Role)
		properties["user_id"] = userData.Cid
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
	return nil
}

// ErrAuthRequired is returned when the user needs to authenticate
var ErrAuthRequired = fmt.Errorf("authentication required. Please run: leap auth login")

// ErrAuthInvalid is returned when the stored credentials are invalid
var ErrAuthInvalid = fmt.Errorf("authentication invalid or expired. Please run: leap auth login")

// RequireAuth checks if the user is authenticated and validates the credentials.
// It first checks if URL and API key are configured, then validates them by calling the API.
// Returns the authenticated context and user data if successful, or an error with a helpful message.
func RequireAuth(ctx context.Context) (context.Context, *tensorleapapi.UserData, error) {
	env := GetCurrentEnv()

	// Check if URL is configured
	if env.ApiUrl == "" {
		log.Warn("No API URL configured")
		return ctx, nil, ErrAuthRequired
	}

	// Check if API key is configured
	if env.ApiKey == "" {
		log.Warn("No API key configured")
		return ctx, nil, ErrAuthRequired
	}

	// Create authenticated context
	authCtx := env.AuthContext(ctx)

	// Validate credentials by calling WhoAmI
	userData, err := GetWhoami(authCtx)
	if err != nil {
		log.Warnf("Failed to validate authentication: %v", err)
		return ctx, nil, ErrAuthInvalid
	}

	return authCtx, userData, nil
}

// RequireAuthSimple is a simpler version of RequireAuth that only returns an error.
// Use this when you don't need the user data and the context is already set up.
func RequireAuthSimple(ctx context.Context) error {
	_, _, err := RequireAuth(ctx)
	return err
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
