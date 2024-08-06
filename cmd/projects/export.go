package projects

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewExportCmd() *cobra.Command {
	var outputDir string
	var exportOptions = project.ExportProjectParams{}

	cmd := &cobra.Command{
		Use:   "export [project name]",
		Short: "Export project to file",
		Long:  `Export project to file`,
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx := cmd.Context()
			projectNameArg := ""
			if len(args) > 0 {
				projectNameArg = args[0]
			}
			projects, err := project.GetProjects(ctx)

			if err != nil {
				return err
			}

			projectEntity, err := entity.SelectEntityByNameOrAsk(projectNameArg, projects, project.ProjectEntityDesc)
			if err != nil {
				return err
			}

			return project.ExportProjectIntoFile(ctx, projectEntity, outputDir, exportOptions)
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output-dir", "o", "", "Output directory")
	cmd.Flags().BoolVar(&exportOptions.NoCache, "no-cache", false, "Do not use previously cached export")
	cmd.Flags().BoolVar(&exportOptions.ExcludeCalculatedFiles, "exclude-calc-files", false, "Do not include calculated files in the export (will decrease export size)")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewExportCmd())
}
