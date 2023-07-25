package code

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func init() {
	var datasetId string

	var cmd = &cobra.Command{
		Use:   "list-versions",
		Short: "List available dataset versions",
		Long:  `List available dataset versions`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.CheckLoggedIn(); err != nil {
				return err
			}
			params := *tensorleapapi.NewGetDatasetVersionsParams(datasetId)
			data, _, err := ApiClient.GetDatasetVersions(cmd.Context()).
				GetDatasetVersionsParams(params).
				Execute()
			if err != nil {
				return err
			}
			for _, codeIntegrationVersion := range data.DatasetVersions {
				fmt.Println(codeIntegrationVersion.GetCid())
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId to get versions from")
	_ = cmd.MarkFlagRequired("datasetId")
	RootCommand.AddCommand(cmd)
}
