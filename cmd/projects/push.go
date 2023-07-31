package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {

	var secretId string
	var modelName string
	var versionName string
	var modelType string
	var branchName string

	var cmd = &cobra.Command{
		Use:   "push <modelPath>",
		Short: "Push new version into a project with its model and code integration",
		Long:  `Push new version into a project with its model and code integration`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing model path argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			modelPath := args[0]
			projects, err := project.GetProjects(ctx)

			model.SelectModelType(&modelType, modelPath)

			if err != nil {
				return err
			}
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err != nil {
				return err
			}
			var currentProject *project.ProjectEntity

			configContainProjectId := len(workspaceConfig.ProjectId) > 0
			if configContainProjectId {
				currentProject, _ = entity.GetEntityById(workspaceConfig.ProjectId, projects, project.ProjectEntityDesc)
			}
			if currentProject == nil {
				if configContainProjectId {
					log.Infof("Not found project ID %s", workspaceConfig.ProjectId)
				} else {
					log.Info("Project ID not found on local config file")
				}
				currentProject, err = project.SelectOrCreateProject(ctx, projects, true)
				if err != nil {
					return err
				}
				log.Info("Updating projectId")
				workspaceConfig.ProjectId = currentProject.GetCid()
				err = workspace.SetWorkspaceConfig(workspaceConfig, "")
				if err != nil {
					return err
				}
			}

			codeIntegration, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				return err
			}
			if len(secretId) == 0 {
				secretId = workspaceConfig.SecretManagerId
			}

			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig)
			if err != nil {
				return err
			}
			defer close()

			_, err = code.AddCodeIntegrationVersion(ctx, tarGzFile, codeIntegration, workspaceConfig.EntryFile, secretId)
			if err != nil {
				return err
			}

			err = model.ImportModel(ctx, modelPath, currentProject.GetCid(), modelName, versionName, modelType, branchName, codeIntegration.GetCid())
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&modelName, "name", "", "Name of the model that will be created")
	_ = cmd.MarkFlagRequired("name")

	cmd.Flags().StringVar(&versionName, "version", "", "Version is the name of the version that will be created for the model")
	_ = cmd.MarkFlagRequired("version")

	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")

	cmd.Flags().StringVar(&branchName, "branch", "", "Branch is the name of the branch [OPTIONAL]")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
