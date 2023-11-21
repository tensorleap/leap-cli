package auth

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	API_URL  = "api_url"
	API_KEY  = "api_key"
	ENV_NAME = "name"

	ENVIRONMENTS_CONFIG_PATH = "envs"
	AUTH_CONFIG_PATH         = "auth"

	CURRENT_ENV = "current_env"
)

var (
	// just for backward compatibility, will use just current_env and envs in the future
	API_KEY_CONFIG_PATH = getAuthPath(API_KEY)
	API_URL_CONFIG_PATH = getAuthPath(API_URL)
)

func getAuthPath(key string) string {
	return fmt.Sprintf("%s.%s", AUTH_CONFIG_PATH, key)
}

func getEnvAuthPath(name, key string) string {
	return fmt.Sprintf("%s.%s.%s", ENVIRONMENTS_CONFIG_PATH, name, key)
}

func GetEnvAuth(name string) (*Env, error) {
	if !viper.IsSet(getEnvAuthPath(name, API_URL)) {
		return nil, ErrNotLoggedIn
	}
	return &Env{
		ApiUrl: viper.GetString(getEnvAuthPath(name, API_URL)),
		ApiKey: viper.GetString(getEnvAuthPath(name, API_KEY)),
	}, nil
}

func removeEnvAuth(name string) {
	currentEnv := viper.GetString(CURRENT_ENV)

	if len(name) == 0 {
		name = currentEnv
	}

	if len(name) > 0 {
		envs := viper.GetStringMap(ENVIRONMENTS_CONFIG_PATH)
		delete(envs, name)
		viper.Set(ENVIRONMENTS_CONFIG_PATH, envs)
	}

	isEqualOrNotSet := currentEnv == name || currentEnv == ""
	if isEqualOrNotSet {
		removeCurrentEnv()
	}
}

func getApiUrl() string {
	return viper.GetString(getAuthPath(API_URL))
}

func IsApiBaseUrlSet() bool {
	return len(getApiUrl()) > 0
}

func getApiKey() string {
	return viper.GetString(getAuthPath(API_KEY))
}

func GetEnvs() []Env {
	obj := viper.GetStringMap(ENVIRONMENTS_CONFIG_PATH)
	envs := make([]Env, 0, len(obj))
	for k, v := range obj {
		env := v.(map[string]interface{})
		envs = append(envs, Env{
			Name:   k,
			ApiUrl: env[API_URL].(string),
			ApiKey: env[API_KEY].(string),
		})

	}
	return envs
}

func CheckLoggedIn() error {
	if IsApiBaseUrlSet() {
		return nil
	}
	return ErrNotLoggedIn
}

func GetCurrentEnv() *Env {
	return &Env{
		Name:   viper.GetString(CURRENT_ENV),
		ApiUrl: getApiUrl(),
		ApiKey: getApiKey(),
	}
}

func setCurrentEnv(env *Env) {
	viper.Set(API_URL_CONFIG_PATH, env.ApiUrl)
	viper.Set(API_KEY_CONFIG_PATH, env.ApiKey)
	viper.Set(CURRENT_ENV, env.Name)
}

func removeCurrentEnv() {
	viper.Set(CURRENT_ENV, "")
	viper.Set(API_KEY_CONFIG_PATH, "")
	viper.Set(API_URL_CONFIG_PATH, "")
}

func setEnvAuth(auth *Env) {
	viper.Set(getEnvAuthPath(auth.Name, API_URL), auth.ApiUrl)
	viper.Set(getEnvAuthPath(auth.Name, API_KEY), auth.ApiKey)
}

func save() error {
	err := viper.SafeWriteConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileAlreadyExistsError)
		if ok {
			err = viper.WriteConfig()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
