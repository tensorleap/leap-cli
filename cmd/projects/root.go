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
		authCtx, _, err := auth.RequireAuth(context.Background())
		if err != nil {
			return err
		}
		cmd.SetContext(authCtx)
		return nil
	},
}
