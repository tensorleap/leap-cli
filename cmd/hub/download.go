package hub

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/hub/storage"
)

func NewDownloadCmd() *cobra.Command {
	var version string

	cmd := &cobra.Command{
		Use:   "download [project name]",
		Short: "Download project from the hub",
		Long:  `Download project from the hub`,
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
			if err != nil {
				return err
			}
			projectNames := meta.GetProjectNames()
			projectName, err := hub.ValidateOrSelectProject(projectNameArg, projectNames)
			if err != nil {
				return err
			}

			projectVersions := meta.GetProjectVersions(projectName)

			projectVersion, err := hub.ValidateOrSelectLatestVersion(projectName, version, projectVersions)
			if err != nil {
				return err
			}

			destFileName := fmt.Sprintf("%s-v%d.tar.gz", projectName, projectVersion)

			err = hubStorage.DownloadProject(projectName, projectVersion, destFileName)
			return err
		},
	}

	cmd.Flags().StringVarP(&version, "version", "v", "", "Version of the project to download, by default latest")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewDownloadCmd())
}
