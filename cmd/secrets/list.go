package secrets

import (
	"github.com/tensorleap/leap-cli/pkg/log"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
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
			data, _, err := ApiClient.GetSecretManagerList(cmd.Context()).Execute()
			if err != nil {
				return err
			}
			for _, secret := range data.Results {
				log.Println(secret.GetName(), "-", secret.GetCid())
			}

			return nil
		},
	}

	return cmd
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}
