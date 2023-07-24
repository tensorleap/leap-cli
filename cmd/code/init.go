package code

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/code"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	var codeIntegrationId string
	var newCodeIntegrationName string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create a .tensorleap.yaml file in the current directory",
		Long:  `Create a .tensorleap.yaml file in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(codeIntegrationId) == 0 && len(newCodeIntegrationName) == 0 {
				return errors.New("Error: flag(s) \"codeId\" or \"new\" must be set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			var codeIntegration *tensorleapapi.Dataset = nil
			var err error = nil
			if len(newCodeIntegrationName) > 0 {
				codeIntegration, err = code.CreateNewCodeIntegration(cmd.Context(), newCodeIntegrationName)
			} else if len(codeIntegrationId) > 0 {
				codeIntegration, err = code.GetCodeIntegrationById(cmd.Context(), codeIntegrationId)
			}
			if err != nil {
				return err
			}
			return code.CreateCodeTemplate(codeIntegration.GetCid(), "")
		},
	}

	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "Code integration id of existing dataset")
	cmd.Flags().StringVar(&newCodeIntegrationName, "new", "", "Name for new database")
	cmd.MarkFlagsMutuallyExclusive("new", "codeId")
	RootCommand.AddCommand(cmd)
}
