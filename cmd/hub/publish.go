package hub

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/hub/storage"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewPublishCmd() *cobra.Command {
	var override bool
	cmd := &cobra.Command{
		Use:   "publish <project tar file path>",
		Short: "Publish project to the hub",
		Long:  `Publish project to the hub`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing project tar file path")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			tarPath := args[0]
			storageClient, err := storage.CreateStorageClient()
			if err != nil {
				return err
			}
			hubStorage := hub.NewHubStorageApi(hub.DEFAULT_NAMESPACE, storageClient)
			hubApi := hub.NewHubApi(hubStorage)
			meta, err := hubStorage.GetMetaWrapper()
			if err != nil {
				return err
			}
			tarFile, err := os.Open(tarPath)
			if err != nil {
				return fmt.Errorf("failed to open tar file: %v", err)
			}
			defer tarFile.Close()
			log.Infof("Extracting project context from: %s", tarPath)
			projectContext, err := hub.ExtractProjectContextFromTar(tarFile)
			if err != nil {
				return fmt.Errorf("failed to extract project context: %v", err)
			}
			_, err = tarFile.Seek(0, 0)
			if err != nil {
				return fmt.Errorf("failed to seek tar file: %v", err)
			}

			err = hub.HandleProjectExistsOnPublish(ctx, hubApi, meta, &projectContext.Meta, override)
			if err != nil {
				return err
			}

			return hubApi.PublishProject(tarFile, projectContext)
		},
	}

	cmd.Flags().BoolVarP(&override, "override", "o", false, "Override existing project")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewPublishCmd())
}
