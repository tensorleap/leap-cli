package auth

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tensorleap/cli-go/pkg/log"
)

func Login(apiKey, baseUrl string) (err error) {
	viper.Set(API_URL_CONFIG_PATH, baseUrl)
	viper.Set(API_KEY_CONFIG_PATH, apiKey)

	err = viper.SafeWriteConfig()
	if err != nil {
		_, ok := err.(viper.ConfigFileAlreadyExistsError)
		if ok {
			err = viper.WriteConfig()
			if err != nil {
				return fmt.Errorf("Login failed: %v", err)
			}
		}
	}
	log.Println("Saved credentials to: ", viper.ConfigFileUsed())
	return
}

func Logout() error {
	viper.Set(API_KEY_CONFIG_PATH, "")
	return viper.WriteConfig()
}
