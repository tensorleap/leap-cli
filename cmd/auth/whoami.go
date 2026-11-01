package auth

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/log"
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
			ctx := env.AuthContext(cmd.Context())
			if _, err := auth.GetWhoami(ctx); err != nil {
				log.Warnf("Could not verify the saved sign-in for %s: %v", env.ApiUrl, err)
				return auth.NotSignedInError()
			}
			version.CheckServerCompatibility(ctx)
			return auth.PrintWhoami(ctx)
		},
	}
}
