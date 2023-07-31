package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewDeleteCmd() *cobra.Command {
	var projectId string
	var skipConfirm bool

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete project",
		Long:  `Delete project`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var selected *project.ProjectEntity
			desc := project.ProjectEntityDesc

			projects, err := project.GetProjects(ctx)
			if err != nil {
				return err
			}
			if len(projectId) > 0 {
				selected, err = entity.GetEntityById(projectId, projects, desc)
			} else {
				selected, err = entity.SelectEntity(projects, desc)
			}
			if err != nil {
				return err
			}
			if !skipConfirm {
				confirmed, err := entity.ConfirmDeletion(selected, desc)
				if !confirmed {
					return err
				}
			}

			err = project.DeleteProject(ctx, selected)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "Project Id to delete")
	cmd.Flags().BoolVarP(&skipConfirm, "skipConfirm", "y", false, "Skip deletion confirm")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewDeleteCmd())
}
