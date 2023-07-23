package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/auth"
)

var cmd = &cobra.Command{
	Use:   "login [api key] [api url]",
	Args:  cobra.ExactArgs(2),
	Short: "Logging in with Tensorleap API key",
	Long:  `Logging in with Tensorleap API key`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, baseUrl := args[0], args[1]
		ctx := CreateAuthenticatedContext(cmd.Context(), apiKey, baseUrl)

		userData, _, err := ApiClient.WhoAmI(ctx).Execute()
		if err != nil {
			return err
		}

		fmt.Println("User email: " + userData.Local.Email)
		fmt.Println("Team name: " + userData.TeamName)
		auth.Login(apiKey, baseUrl)

		return nil
	},
}

func init() {
	RootCommand.AddCommand(cmd)
}
