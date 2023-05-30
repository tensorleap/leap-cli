package auth

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tensorleap/cli-go/pkg/config"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "logout",
		Short: "Remove api key from the machine",
		Long:  `Remove api key from the machine`,
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set(config.API_KEY_CONFIG_PATH, "")
			err := viper.WriteConfig()
			cobra.CheckErr(err)
		},
	})
}
