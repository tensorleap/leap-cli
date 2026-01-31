package projects

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

var RootCommand = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"project", "pro"},
	Short:   "Manage project",
	Long:    `Manage project`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Set up auth context (same as parent root command)
		env := auth.GetCurrentEnv()
		authCtx := env.AuthContext(context.Background())
		cmd.SetContext(authCtx)

		// Then validate authentication
		return auth.RequireAuthSimple(authCtx)
	},
}
