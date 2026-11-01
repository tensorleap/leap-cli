package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func init() {
	RootCommand.AddCommand(NewWhoamiCmd())
}

// NewWhoamiCmd only reports who is signed in; it never starts a login, so scripts and AI assistants can call it safely
func NewWhoamiCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "whoami",
		Short: "Get information about the authenticated user",
		Long:  `Get information about the authenticated user. When nobody is signed in it says how to sign in and exits with an error, without asking anything.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			env := auth.GetCurrentEnv()
			if env.ApiUrl == "" || env.ApiKey == "" {
				return auth.NotSignedInError()
			}
			if err := auth.CheckSignIn(cmd.Context(), env); err != nil {
				return err
			}
			ctx := env.AuthContext(cmd.Context())
			version.CheckServerCompatibility(ctx)
			return auth.PrintWhoami(ctx)
		},
	}
}
