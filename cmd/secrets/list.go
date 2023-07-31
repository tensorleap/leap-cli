package secrets

import (
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
)

type SecretEntity = tensorleapapi.SecretManager

var SecretEntityDescriptor = entity.NewEntityDescriptor[SecretEntity](
	"secret",
	"secrets",
	func(e *SecretEntity) string { return e.GetName() },
	func(e *SecretEntity) string { return e.GetCid() },
)

func NewListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List secrets",
		Long:  `List secrets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			data, _, err := api.ApiClient.GetSecretManagerList(cmd.Context()).Execute()
			if err != nil {
				return err
			}
			entity.PrintList(data.Results, SecretEntityDescriptor)

			return nil
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}
