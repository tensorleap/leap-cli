package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
)

var AuthWhoAmICommand = &cobra.Command{
	Use:   "whoami",
  Short: "Get information about the authenticated user",
	Long:  `Get information about the authenticated user`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WhoAmI command")
    userData, resp, err := ApiClient.WhoAmI(cmd.Context()).Execute()
		fmt.Println(userData)
		fmt.Println(resp)
		fmt.Println(err)
	},
}
