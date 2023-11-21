package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

func init() {
	cmd := &cobra.Command{
		Use:   "whoami",
		Short: "Get information about the authenticated user",
		Long:  `Get information about the authenticated user`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}

			return auth.PrintWhoami(cmd.Context())
		},
	}
	RootCommand.AddCommand(cmd)
}
