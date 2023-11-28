package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewImportCmd() *cobra.Command {
	var filePath string
	cmd := &cobra.Command{
		Use:   "import [project name]",
		Short: "Import project from hub",
		Long:  `Import project from hub`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			projectNameArg := ""
			if len(args) > 0 {
				projectNameArg = args[0]
			}

			if filePath != "" {
				return project.ImportProjectFromFile(ctx, filePath, projectNameArg)
			}
			publicHubUrl := hub.GetPublicUrl()
			meta, err := hub.CreateWrapperMetaFromUrl(publicHubUrl)
			if err != nil {
				return err
			}

			environmentInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute()
			if err != nil {
				return err
			}
			version := int(environmentInfo.SchemaVersion)

			projectNames := meta.GetProjectNamesByVersion(version)
			projectName, err := hub.ValidateOrSelectProject(projectNameArg, projectNames)
			if err != nil {
				return err
			}
			versionMeta, err := meta.GetVersionMeta(projectName, version)
			if err != nil {
				return err
			}
			importUrl := fmt.Sprintf("%s/%s", publicHubUrl, versionMeta.Path)
			bgImageUrl := fmt.Sprintf("%s/%s", publicHubUrl, versionMeta.BgImagePath)

			fmt.Printf("Importing project %s from %s\n", projectName, importUrl)

			projectMeta := &hub.ProjectMeta{
				Name:          projectName,
				SchemaVersion: version,
				BgImagePath:   bgImageUrl,
				Description:   versionMeta.Description,
				Tags:          versionMeta.Tags,
				Categories:    versionMeta.Categories,
			}

			return project.ImportProject(ctx, projectName, importUrl, projectMeta)
		},
	}
	cmd.Flags().StringVarP(&filePath, "file", "f", "", "Import project from file instead of hub")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewImportCmd())
}
