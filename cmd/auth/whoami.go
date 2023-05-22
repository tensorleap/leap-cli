package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	. "github.com/tensorleap/cli-go/pkg/api"
)

func init() {
  cmd := &cobra.Command{
    Use:   "whoami",
    Short: "Get information about the authenticated user",
    Long:  `Get information about the authenticated user`,
    Run: func(cmd *cobra.Command, args []string) {
      apiUrl := viper.GetString("auth.api_url")
      if (len(apiUrl) == 0) {
        cobra.CheckErr("Not logged in!")
      }
      fmt.Println("API Url: " + apiUrl)
      userData, _, err := ApiClient.WhoAmI(cmd.Context()).Execute()
      cobra.CheckErr(err)
      fmt.Println("User email: " + userData.Local.Email)
      fmt.Println("Team name: " + userData.OrganizationName)
    },
  }
  RootCommand.AddCommand(cmd)
}
