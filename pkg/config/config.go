package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const TL_CLI_CONFIG_FILE = "TL_CLI_CONFIG_FILE"

func getDefaultConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %v", err)
	}
	return filepath.Join(homeDir, ".config", "tensorleap/config.yaml"), nil
}

func validateConfigPath(cfgFile string) error {
	if cfgFile == "" {
		return fmt.Errorf("config file is required")
	}

	if !filepath.IsAbs(cfgFile) {
		return fmt.Errorf("cli config file must be an absolute path, current path: %s", cfgFile)
	}
	return nil
}

func InitConfig(cfgFile string) error {
	cfgFileFromEnv := os.Getenv(TL_CLI_CONFIG_FILE)

	if cfgFile == "" && cfgFileFromEnv != "" {
		cfgFile = cfgFileFromEnv
	} else if cfgFile == "" {
		var err error
		cfgFile, err = getDefaultConfigPath()
		if err != nil {
			return err
		}
	}
	err := validateConfigPath(cfgFile)
	if err != nil {
		return err
	}

	cfgFile = strings.TrimSuffix(cfgFile, filepath.Ext(cfgFile))

	configDir := filepath.Dir(cfgFile)
	configName := filepath.Base(cfgFile)
	configType := "yaml"
	configPath := filepath.Join(configDir, fmt.Sprintf("%s.%s", configName, configType))
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		return err
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		file, err := os.Create(configPath)
		if err != nil {
			return fmt.Errorf("error creating config the file: %v", err)
		}
		_ = file.Close()
	}

	viper.AddConfigPath(configDir)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func Save() error {
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
