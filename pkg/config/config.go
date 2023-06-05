package config

import (
	"errors"
	"os"
	"path"

	"github.com/spf13/viper"
)

const API_URL_CONFIG_PATH = "auth.api_url"
const API_KEY_CONFIG_PATH = "auth.api_key"

var NotLoggedInError = errors.New("Not logged in")

func GetApiUrl() string {
	return viper.GetString("auth.api_url")
}

func IsApiBaseUrlSet() bool {
	return len(GetApiUrl()) > 0
}

func GetApiKey() string {
	return viper.GetString("auth.api_key")
}

func CheckLoggedIn() error {
	if IsApiBaseUrlSet() {
		return nil
	}
	return NotLoggedInError
}

func InitConfig(cfgFile string) error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		configDir := path.Join(homeDir, ".config", "tensorleap")
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

		os.MkdirAll(configDir, os.ModePerm)
	}

	viper.AutomaticEnv()
	viper.ReadInConfig()

	return nil
}
