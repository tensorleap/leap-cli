package projects

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"project", "pro"},
	Short:   "Manage project",
	Long:    `Manage project`,
}
