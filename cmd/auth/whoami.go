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
			ctx, _, err := auth.RequireAuth(cmd.Context())
			if err != nil {
				return err
			}

			return auth.PrintWhoami(ctx)
		},
	}
	RootCommand.AddCommand(cmd)
}
