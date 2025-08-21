package auth

import (
	"context"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/config"
	"github.com/tensorleap/leap-cli/pkg/log"
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
		"api_url": env.ApiUrl,
		"env_name": env.Name,
		"method": "api_key",
		"success": true,
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
	
	if err := analytics.SendEvent(analytics.EventAuthLoginSuccess, properties); err != nil {
		// Log error but don't fail the login
		log.Warnf("Failed to track login event: %v", err)
	}
	
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

const API_PATH = "/api/v2"
const API_CLOUD_SUBDOMAIN = "://api."

func isCloudUrl(url string) bool {
	return strings.Contains(url, ".tensorleap.ai")
}

func NormalizeAPIUrl(url string) (string, error) {
	url = strings.TrimSuffix(url, "/")
	hasProtocol := strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
	if !hasProtocol {
		return "", fmt.Errorf("URL must start with http:// or https://")
	}
	isCloud := isCloudUrl(url)
	hasApiSubdomain := strings.Contains(url, API_CLOUD_SUBDOMAIN)
	if isCloud && !hasApiSubdomain {
		url = strings.Replace(url, "://", API_CLOUD_SUBDOMAIN, 1)
	}

	if strings.HasSuffix(url, API_PATH) {
		return url, nil
	}
	return fmt.Sprintf("%s%s", url, API_PATH), nil
}

func ChangeToUIUrl(apiUrl string) string {
	url := strings.Replace(apiUrl, API_PATH, "", 1)
	return strings.Replace(url, API_CLOUD_SUBDOMAIN, "://", 1)
}
