package code

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

func NewListVersionsCmd() *cobra.Command {
	var codeIntegrationId string

	var cmd = &cobra.Command{
		Use:   "list-versions",
		Short: "List code integration versions",
		Long:  `List code integration versions`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			if len(codeIntegrationId) == 0 {
				if workspace.IsWorkspaceDir() {
					config, err := workspace.GetWorkspaceConfig()
					if err != nil {
						return err
					}
					codeIntegrationId = config.CodeIntegrationId
				}
				if len(codeIntegrationId) == 0 {
					codeIntegrations, err := code.GetCodeIntegrations(cmd.Context())
					if err != nil {
						return err
					}
					selected, err := entity.SelectEntity(codeIntegrations, code.CodeIntegrationEntityDesc)
					if err != nil {
						return err
					}
					if err != nil {
						return err
					}
					codeIntegrationId = selected.GetCid()
				}
			}

			codeIntegrationVersions, err := code.GetCodeIntegrationVersions(cmd.Context(), codeIntegrationId)
			if err != nil {
				return err
			}
			entity.PrintList(codeIntegrationVersions, code.CodeIntegrationVersionEntityDesc)
			return nil
		},
	}

	cmd.Flags().StringVar(&codeIntegrationId, "codeId", "", "Code integration id to get versions from")
	return cmd
}

func init() {
	RootCommand.AddCommand(NewListVersionsCmd())
}
