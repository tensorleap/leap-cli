package projects

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/hub"
	"github.com/tensorleap/leap-cli/pkg/hub/storage"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewPublishCmd() *cobra.Command {
	var override bool
	var all bool
	var useSignedUrl bool

	cmd := &cobra.Command{
		Use:   "publish [project names]",
		Short: "publish project to the hub",
		Long:  `publish project to the hub`,
		RunE: func(cmd *cobra.Command, projectNames []string) error {
			ctx := cmd.Context()

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

			environmentInfo, _, err := api.ApiClient.GetEnvironmentInfo(ctx).Execute()
			if err != nil {
				return err
			}
			version := int(environmentInfo.SchemaVersion)

			publishProject := func(projectEntity *project.ProjectEntity) error {
				projectContext, err := project.BuildProjectContext(ctx, projectEntity, version)
				if err != nil {
					return fmt.Errorf("failed to extract project context: %v", err)
				}

				err = hub.HandleProjectExistsOnPublish(ctx, hubApi, meta, &projectContext.Meta, override)
				if err != nil {
					return err
				}

				if useSignedUrl {
					return StreamProjectToHubBySignedUrl(ctx, hubApi, projectContext)
				}
				return StreamProjectToHub(ctx, hubApi, projectContext)
			}

			projectsToPublish, err := project.SelectProjectsToPublish(ctx, projectNames, all)
			if err != nil {
				return err
			}

			if len(projectsToPublish) == 0 {
				log.Warnf("No projects to publish")
				return nil
			}

			for _, projectEntity := range projectsToPublish {
				err := publishProject(&projectEntity)
				if err == hub.ErrProjectExists {
					log.Warnf("Project %s already exists, use -o, --override to re-upload", projectEntity.Name)
					continue
				}
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&override, "override", "o", false, "Override existing project")
	cmd.Flags().BoolVarP(&all, "all", "a", false, "Publish all public projects")
	cmd.Flags().BoolVarP(&useSignedUrl, "use-signed-url", "s", false, "Use signed url to publish project, make sure you have enough permission to publish project")

	return cmd
}

func init() {
	if hub.IsHubEnabled() {
		RootCommand.AddCommand(NewPublishCmd())
	}
}

func StreamProjectToHub(ctx context.Context, hubApi *hub.HubApi, projectContext *hub.ProjectContext) error {
	res, err := project.ExportProject(ctx, projectContext.Meta.SourceProjectId)
	if err != nil {
		return fmt.Errorf("failed to export project: %v", err)
	}
	defer res.Body.Close()

	return hubApi.PublishProject(res.Body, projectContext)
}

func StreamProjectToHubBySignedUrl(ctx context.Context, hubApi *hub.HubApi, projectContext *hub.ProjectContext) error {
	tarSignedUrl, filesPath, err := hubApi.PublishProjectContentBySignedUrl(&projectContext.Meta)
	if err != nil {
		return fmt.Errorf("failed to get signed url: %v", err)
	}
	err = project.PublishProject(ctx, projectContext.Meta.SourceProjectId, tarSignedUrl)
	if err != nil {
		return fmt.Errorf("failed to publish project by signed url: %v", err)
	}

	return hubApi.PublishProjectMeta(filesPath, projectContext)
}
