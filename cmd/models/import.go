package models

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewImportCmd() *cobra.Command {
	var projectId string
	var message string
	var modelType string
	var modelBranch string
	var codeIntegrationBranch string
	var transformInput bool
	var codeIntegrationId string
	var noWait bool

	var cmd = &cobra.Command{
		Use:   "import <modelPath>",
		Short: "Import a model into tensorleap",
		Long:  `Import a model into tensorleap`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf("missing model path argument")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Define base properties for all analytics events
			properties := map[string]interface{}{
				"project_id":              projectId,
				"message":                 message,
				"model_type":              modelType,
				"model_branch":            modelBranch,
				"code_integration_branch": codeIntegrationBranch,
				"transform_input":         transformInput,
				"code_integration_id":     codeIntegrationId,
				"no_wait":                 noWait,
			}

			// Track models import started
			if err := analytics.SendEvent(analytics.EventCliModelsImportStarted, properties); err != nil {
				log.Warnf("Failed to track models import start event: %v", err)
			}

			if err := auth.CheckLoggedIn(); err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "auth_check"
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}
			modelPath := args[0]
			ctx := cmd.Context()

			err := model.SelectModelType(&modelType, modelPath)
			if err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "select_model_type"
				properties["model_path"] = modelPath
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}
			err = model.InitMessage(&message)
			if err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "init_message"
				properties["model_path"] = modelPath
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}

			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if !os.IsNotExist(err) && err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "get_workspace_config"
				properties["model_path"] = modelPath
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}

			if workspaceConfig != nil {
				if len(projectId) == 0 && len(workspaceConfig.ProjectId) > 0 {
					log.Infof("Using projectId('%s') from workspace config", workspaceConfig.ProjectId)
					projectId = workspaceConfig.ProjectId
				}

				if len(codeIntegrationId) == 0 {
					codeIntegrationId = workspaceConfig.CodeIntegrationId
				}

			}

			currentProject, _, err := project.GetProjectFromProjectId(ctx, projectId, true)
			if err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "get_project"
				properties["model_path"] = modelPath
				properties["final_project_id"] = projectId
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}

			err = model.ImportModel(ctx, modelPath, currentProject.Cid, message, modelType, modelBranch, codeIntegrationId, codeIntegrationBranch, transformInput, !noWait)
			if err != nil {
				// Track models import failed
				properties["error"] = err.Error()
				properties["stage"] = "import_model"
				properties["model_path"] = modelPath
				properties["final_project_id"] = currentProject.Cid
				properties["final_model_type"] = modelType
				properties["final_model_branch"] = modelBranch
				properties["final_code_integration_id"] = codeIntegrationId
				properties["final_code_integration_branch"] = codeIntegrationBranch
				properties["final_transform_input"] = transformInput
				properties["final_wait"] = !noWait
				if err := analytics.SendEvent(analytics.EventCliModelsImportFailed, properties); err != nil {
					log.Warnf("Failed to track models import failure event: %v", err)
				}
				return err
			}
			// Track models import success
			properties["model_path"] = modelPath
			properties["final_project_id"] = currentProject.Cid
			properties["final_model_type"] = modelType
			properties["final_model_branch"] = modelBranch
			properties["final_code_integration_id"] = codeIntegrationId
			properties["final_code_integration_branch"] = codeIntegrationBranch
			properties["final_transform_input"] = transformInput
			properties["final_wait"] = !noWait
			if err := analytics.SendEvent(analytics.EventCliModelsImportSuccess, properties); err != nil {
				log.Warnf("Failed to track models import success event: %v", err)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project the model will be imported to")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Version message")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&modelBranch, "model-branch", "", "Name of the model branch [OPTIONAL]")
	cmd.Flags().StringVar(&codeIntegrationBranch, "code-branch", "", "Name of the code integration branch [OPTIONAL]")
	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "This is a code integration id (Will use the last valid dataset version)")
	cmd.Flags().BoolVar(&transformInput, "transform-input", false, "Transpose the input data to channel-last format")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewImportCmd())
}
