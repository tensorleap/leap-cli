package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewInitCmd() *cobra.Command {
	var projectId string
	var secretId string
	var pythonVersion string
	var branch string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create initial project environment files in the current directory",
		Long:  `Create initial project environment files in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			selectedProject, _, err := project.GetProjectFromProjectId(ctx, projectId, true)
			if err != nil {
				return err
			}

			pythonVersion, err := code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, nil)
			if err != nil {
				return err
			}

			err = workspace.CreateCodeTemplate(
				selectedProject.GetCid(),
				secretId,
				".",
				pythonVersion,
				branch,
			)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret manager id for new code integration")
	cmd.Flags().StringVar(&pythonVersion, "pythonVersion", "", "Python version for the code integration")
	cmd.Flags().StringVar(&branch, "branch", "", "Branch name")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewInitCmd())
}
