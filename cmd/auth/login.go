package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	. "github.com/tensorleap/cli-go/pkg/api"
)

var cmd = &cobra.Command{
  Use:   "login [api key] [api url]",
  Args: cobra.ExactArgs(2),
  Short: "Loging with Tensorleap API key",
  Long:  `Loging with Tensorleap API key`,
  Run: func(cmd *cobra.Command, args []string) {
    apiKey := args[0]
    baseUrl := args[1]
    ctx := CreateAuthenticatedContext(cmd.Context(), baseUrl, apiKey)

    userData, _, apiErr := ApiClient.WhoAmI(ctx).Execute()
    cobra.CheckErr(apiErr)
    fmt.Println("User email: " + userData.Local.Email)
    fmt.Println("Team name: " + userData.OrganizationName)

    viper.Set("auth.api_url", baseUrl)
    viper.Set("auth.api_key", apiKey)
    err := viper.SafeWriteConfig()
    if (err != nil) {
      _, ok := err.(viper.ConfigFileAlreadyExistsError)
      if (ok) {
        err = viper.WriteConfig()
      }
    }
    cobra.CheckErr(err)
    fmt.Println("Saved credentials to: ", viper.ConfigFileUsed())
  },
}

func init() {
  RootCommand.AddCommand(cmd)
}

