package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewListCmd() *cobra.Command {

	var cmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List projects",
		Long:    `List projects`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projects, err := project.GetProjects(cmd.Context())
			if err != nil {
				return err
			}

			entity.PrintList(projects, project.ProjectEntityDesc)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}
