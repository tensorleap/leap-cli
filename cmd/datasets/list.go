package datasets

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/auth"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List available datasets",
		Long:  `List available datasets`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			data, _, err := ApiClient.GetDatasets(cmd.Context()).Execute()
			if err != nil {
				return err
			}
			for _, dataset := range data.Datasets {
				fmt.Println(dataset.GetCid(), "-", dataset.GetName())
			}

			return nil
		},
	})
}
