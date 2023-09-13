package models

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewImportCmd() *cobra.Command {
	var projectId string
	var message string
	var modelType string
	var branchName string
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
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			modelPath := args[0]
			ctx := cmd.Context()

			err := model.SelectModelType(&modelType, modelPath)
			if err != nil {
				return err
			}
			err = model.InitMessage(&message)
			if err != nil {
				return err
			}

			if len(projectId) == 0 {
				projects, err := project.GetProjects(ctx)
				if err != nil {
					return err
				}
				selected, _, err := project.SelectOrCreateProject(ctx, projects, true)
				if err != nil {
					return err
				}
				projectId = selected.GetCid()
			}

			err = model.ImportModel(ctx, modelPath, projectId, message, modelType, branchName, codeIntegrationId, transformInput, !noWait)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&projectId, "projectId", "", "ProjectId is the id of the project the model will be imported to")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Version message")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&branchName, "branch", "", "Branch is the name of the branch [OPTIONAL]")
	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "This is a code integration id (Will use the last valid dataset version)")
	cmd.Flags().BoolVar(&transformInput, "transform-input", true, "Transform input in case of ONNX model")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewImportCmd())
}
