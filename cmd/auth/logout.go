package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func NewLogoutCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Remove api key from the machine",
		Long:  `Remove api key from the machine`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return auth.Logout(name)
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name of the environment to logout from, by default logout from the current environment")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewLogoutCmd())
}
