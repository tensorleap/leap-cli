package auth

import (
	"github.com/spf13/cobra"
)

var AuthCommand = &cobra.Command{
	Use:   "auth",
  Short: "auth commands",
	Long:  `auth commands`,
	Run: func(cmd *cobra.Command, args []string) {
    cmd.Usage();
	},
}

func init() {
  AuthCommand.AddCommand(AuthWhoAmICommand)
}
