package hub

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/hub/storage"
)

func NewDeleteCmd() *cobra.Command {
	var version string
	var allVersions bool

	cmd := &cobra.Command{
		Use:   "delete [project name]",
		Short: "Delete project from the hub",
		Long:  `Delete project from the hub`,
		RunE: func(cmd *cobra.Command, args []string) error {
			projectNameArg := ""
			if len(args) > 0 {
				projectNameArg = args[0]
			}
			storageClient, err := storage.CreateStorageClient()
			if err != nil {
				return err
			}
			hubStorage := hub.NewHubStorageApi(hub.DEFAULT_NAMESPACE, storageClient)
			meta, err := hubStorage.GetMetaWrapper()
			hubApi := hub.NewHubApi(hubStorage)

			if err != nil {
				return err
			}
			projectNames := meta.GetProjectNames()
			projectName, err := hub.ValidateOrSelectProject(projectNameArg, projectNames)
			if err != nil {
				return err
			}

			if allVersions {
				err = hubApi.DeleteProject(projectName)
				if err != nil {
					return err
				}
				return nil
			}

			allVersions := meta.GetProjectVersions(projectName)
			selectedVersion, err := hub.ValidateOrSelectLatestVersion(projectName, version, allVersions)
			if err != nil {
				return err
			}

			return hubApi.DeleteProjectVersion(projectName, selectedVersion)
		},
	}
	cmd.Flags().StringVarP(&version, "version", "v", "", "Version of the project to delete, by default latest")
	cmd.Flags().BoolVarP(&allVersions, "all", "a", false, "Delete all versions of the project")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewDeleteCmd())
}
