package config

import (
	"fmt"
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
		configName := "config"
		configType := "yaml"
		configPath := path.Join(configDir, fmt.Sprintf("%s.%s", configName, configType))
		if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
			return err
		}
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			file, err := os.Create(configPath)
			if err != nil {
				return fmt.Errorf("error creating config the file: %v", err)
			}
			file.Close()
		}
		viper.AddConfigPath(configDir)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)

	}

	viper.AutomaticEnv()
	return viper.ReadInConfig()

}
