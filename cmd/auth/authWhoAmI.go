package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/api"
)

var apiClient = api.Client

var AuthWhoAmICommand = &cobra.Command{
	Use:   "whoami",
  Short: "Get information about the authenticated user",
	Long:  `Get information about the authenticated user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Install command")
    userData, resp, err := apiClient.WhoAmI(cmd.Context()).Execute()
		fmt.Println(userData)
		fmt.Println(resp)
		fmt.Println(err)
	},
}
