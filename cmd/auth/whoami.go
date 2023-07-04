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
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.CheckLoggedIn(); err != nil {
				return err
			}
			fmt.Println("API Url: " + config.GetApiUrl())

			userData, _, err := ApiClient.WhoAmI(cmd.Context()).Execute()
			if err != nil {
				return err
			}
			fmt.Println("User email: " + userData.Local.Email)
			fmt.Println("Team name: " + userData.TeamName)
			return nil
		},
	}
	RootCommand.AddCommand(cmd)
}
