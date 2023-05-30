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
		Run: func(cmd *cobra.Command, args []string) {
			config.VerifyLoggedIn()
			params := *tensorleapapi.NewGetDatasetVersionsParams(datasetId)
			data, _, err := ApiClient.GetDatasetVersions(cmd.Context()).
				GetDatasetVersionsParams(params).
				Execute()
			cobra.CheckErr(err)
			for _, datasetVersion := range data.DatasetVersions {
				fmt.Println(datasetVersion.GetId())
			}
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId to get versions from")
	cmd.MarkFlagRequired("datasetId")
	RootCommand.AddCommand(cmd)
}
