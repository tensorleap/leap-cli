package auth

import (
	"github.com/spf13/cobra"
)

var AuthCommand = &cobra.Command{
	Use:   "auth",
  Short: "auth commands",
	Long:  `auth commands`,
}

func init() {
  AuthCommand.AddCommand(AuthWhoAmICommand)
}
