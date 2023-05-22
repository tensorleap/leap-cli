package config

import (
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const API_URL_CONFIG_PATH = "auth.api_url"
const API_KEY_CONFIG_PATH = "auth.api_key"

func GetApiUrl() string {
  return viper.GetString("auth.api_url")
}

func IsApiBaseUrlSet() bool {
  return len(GetApiUrl()) > 0;
}

func GetApiKey() string {
  return viper.GetString("auth.api_key")
}

func VerifyLoggedIn() {
  if (!IsApiBaseUrlSet()) {
    cobra.CheckErr("Not logged in!")
  }
}

func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		homeDir, err := os.UserHomeDir()
		cobra.CheckErr(err)

    configDir := path.Join(homeDir, ".config", "tensorleap")
		viper.AddConfigPath(configDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")

    os.MkdirAll(configDir, os.ModePerm)
	}

	viper.AutomaticEnv()
  viper.ReadInConfig()
}

