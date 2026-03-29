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

// RequireAuth checks if the user is authenticated and validates the credentials.
// If authentication is missing or invalid, it automatically triggers the interactive login flow.
// Returns the authenticated context and user data if successful.
func RequireAuth(ctx context.Context) (context.Context, *tensorleapapi.UserData, error) {
	env := GetCurrentEnv()

	if env.ApiUrl != "" && env.ApiKey != "" {
		authCtx := env.AuthContext(ctx)
		userData, err := GetWhoami(authCtx)
		if err == nil {
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
			apiKey, err = LoginAndGetAuthTokenWithBrowser(ctx, apiUrl)
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
