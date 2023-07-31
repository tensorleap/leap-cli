package code

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
)

func NewDeleteCmd() *cobra.Command {
	var codeIntegrationId string
	var skipConfirm bool

	var cmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete code integration",
		Long:  `Delete code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var selected *code.CodeIntegration
			desc := code.CodeIntegrationEntityDesc
			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				return err
			}
			if len(codeIntegrationId) > 0 {
				selected, err = entity.GetEntityById(codeIntegrationId, codeIntegrations, desc)
			} else {
				selected, err = entity.SelectEntity(codeIntegrations, desc)
			}
			if err != nil {
				return err
			}
			if !skipConfirm {
				confirmed, err := entity.ConfirmDeletion(selected, desc)
				if !confirmed {
					return err
				}
			}

			err = code.DeleteCodeIntegration(ctx, selected)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "Code Integration Id to delete")
	cmd.Flags().BoolVarP(&skipConfirm, "skipConfirm", "y", false, "Skip deletion confirm")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewDeleteCmd())
}
