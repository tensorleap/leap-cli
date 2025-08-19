package auth

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
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
			// Track auth failure
			errorMessage, errorType := analytics.ExtractErrorInfo(err)
			analytics.GetClient().TrackCLIAuthFailed("copy_paste", "saas", baseUrl, errorMessage, errorType)
			return err
		}

		fmt.Println("User email: " + userData.Local.Email)
		fmt.Println("Team name: " + userData.TeamName)
		
		// Track successful auth
		analytics.GetClient().TrackCLIAuth("copy_paste", "saas", baseUrl)
		
		// Update user email for analytics
		analytics.GetClient().SetUserEmail(userData.Local.Email)
		
		return auth.Login(apiKey, baseUrl)
	},
}

func init() {
	RootCommand.AddCommand(cmd)
}
