package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewInitCmd() *cobra.Command {
	var projectId string
	var secretId string
	var pythonVersion string
	var branch string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create initial project environment files (DEPRECATED)",
		Long:  `Create initial project environment files in the current directory (DEPRECATED)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Show deprecation warning
			log.Warn("⚠️  The command 'leap projects init' is deprecated and will be removed in a future version.")
			log.Warn("Please use 'leap init' instead.")
			fmt.Println()

			return fmt.Errorf("please use 'leap init' instead of 'leap projects init'")
		},
	}

	// Add the same flags as the new init command (matching the old projects init interface)
	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret manager id for new code integration")
	cmd.Flags().StringVar(&pythonVersion, "pythonVersion", "", "Python version for the code integration")
	cmd.Flags().StringVar(&branch, "branch", "", "Branch name")

	// Mark command as deprecated
	cmd.Deprecated = "use 'leap init' instead"

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInitCmd())
}
