package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewPushCmd() *cobra.Command {

	var secretId string
	var message string
	var modelType string
	var modelBranch string
	var codeBranch string
	var transformInput bool
	var force bool
	var noWait bool
	var pythonVersion string
	var leapMappingPath string

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
			if err != nil {
				return err
			}

			err = model.SelectModelType(&modelType, modelPath)
			if err != nil {
				return err
			}
			err = model.InitMessage(&message)
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
				currentProject, _, err = project.SelectOrCreateProject(ctx, projects, true)
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

			codeIntegration, _, err := code.GetAndUpdateCodeIntegrationIfNotExists(ctx, workspaceConfig)
			if err != nil {
				return err
			}

			codeIntegrationBranches := code.BranchesFromCodeIntegration(codeIntegration)
			codeBranch, err = code.SyncBranchFromFlagAndConfig(codeBranch, workspaceConfig, codeIntegrationBranches, codeIntegration.GetDefaultBranch())
			if err != nil {
				return err
			}

			secretId, err := secret.SyncSecretIdFromFlagAndConfig(ctx, secretId, workspaceConfig)
			if err != nil {
				return err
			}

			pythonVersion, err = code.SyncPythonVersionFromFlagAndConfig(ctx, pythonVersion, workspaceConfig)
			if err != nil {
				return err
			}

			close, tarGzFile, err := code.BundleCodeIntoTempFile(".", workspaceConfig, leapMappingPath)
			if err != nil {
				return err
			}
			defer close()

			pushed, currentVersion, err := code.PushCode(ctx, force, codeIntegration.Cid, tarGzFile, workspaceConfig.EntryFile, secretId, codeBranch, message, pythonVersion)
			if err != nil {
				return err
			}
			if pushed || code.IsCodeParsing(currentVersion) {

				if err != nil {
					return err
				}

				ok, codeIntegrationVersion, err := code.WaitForCodeIntegrationStatus(ctx, currentVersion.Cid)
				if err != nil {
					return err
				}
				if ok {
					log.Info("Code parsed successfully")
				} else {
					code.PrintCodeIntegrationVersionParserErr(codeIntegrationVersion)
					return fmt.Errorf("code parsing failed")
				}
			} else if code.IsCodeParseFailed(currentVersion) {
				code.PrintCodeIntegrationVersionParserErr(currentVersion)
				return fmt.Errorf("latest code parsing failed, add --force to push anyway")
			}

			err = model.ImportModel(ctx, modelPath, currentProject.GetCid(), message, modelType, modelBranch, codeIntegration.GetCid(), codeBranch, transformInput, !noWait)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&message, "message", "m", "", "Model version and code integration message")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&modelBranch, "model-branch", "", "Name of the model branch [OPTIONAL]")
	cmd.Flags().StringVar(&codeBranch, "code-branch", "", "Name of the code branch [OPTIONAL]")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret id")
	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force push code integration")
	cmd.Flags().BoolVar(&transformInput, "transform-input", false, "Transform input in case of ONNX model")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")
	cmd.Flags().StringVar(&leapMappingPath, "leap-mapping", "", "Path to external leap mapping file")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
