package code

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List available code integration",
		Long:  `List available code integration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			codeIntegrations, err := code.GetCodeIntegrations(cmd.Context())
			if err != nil {
				return err
			}
			entity.PrintList(codeIntegrations, code.CodeIntegrationEntityDesc)

			return nil
		},
	})
}
