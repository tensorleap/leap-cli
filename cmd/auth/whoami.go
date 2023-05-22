package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
)

func init() {
  cmd := &cobra.Command{
    Use:   "whoami",
    Short: "Get information about the authenticated user",
    Long:  `Get information about the authenticated user`,
    Run: func(cmd *cobra.Command, args []string) {
      config.VerifyLoggedIn()
      fmt.Println("API Url: " + config.GetApiUrl())

      userData, _, err := ApiClient.WhoAmI(cmd.Context()).Execute()
      cobra.CheckErr(err)
      fmt.Println("User email: " + userData.Local.Email)
      fmt.Println("Team name: " + userData.OrganizationName)
    },
  }
  RootCommand.AddCommand(cmd)
}
