package auth

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "logout",
    Short: "Remove api key from the machine",
    Long:  `Remove api key from the machine`,
    Run: func(cmd *cobra.Command, args []string) {
      viper.Set("auth.api_key", "")
      err := viper.WriteConfig()
      cobra.CheckErr(err)
    },
  })
}
