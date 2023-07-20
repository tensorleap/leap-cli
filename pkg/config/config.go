package config

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

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
