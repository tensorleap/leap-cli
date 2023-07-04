package datasets

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	var datasetId string

	var cmd = &cobra.Command{
		Use:   "list-versions",
		Short: "List available dataset versions",
		Long:  `List available dataset versions`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.CheckLoggedIn(); err != nil {
				return err
			}
			params := *tensorleapapi.NewGetDatasetVersionsParams(datasetId)
			data, _, err := ApiClient.GetDatasetVersions(cmd.Context()).
				GetDatasetVersionsParams(params).
				Execute()
			if err != nil {
				return err
			}
			for _, datasetVersion := range data.DatasetVersions {
				fmt.Println(datasetVersion.GetCid())
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId to get versions from")
	cmd.MarkFlagRequired("datasetId")
	RootCommand.AddCommand(cmd)
}
