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
	var newDatasetName string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create a .tensorleap.yaml file in the current directory",
		Long:  `Create a .tensorleap.yaml file in the current directory`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(datasetId) == 0 && len(newDatasetName) == 0 {
				msg := "Error: flag(s) \"datasetId\" or \"new\" must be set"
				fmt.Println(msg)
				cmd.Usage()
				cobra.CheckErr(msg)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			if len(newDatasetName) > 0 {
				fmt.Println("Creating dataset:", newDatasetName)
				data, _, err := ApiClient.AddDataset(cmd.Context()).
					NewDatasetParams(*tensorleapapi.NewNewDatasetParams(
						*tensorleapapi.NewNullableString(&newDatasetName))).
					Execute()
				cobra.CheckErr(err)
				datasetId = data.Dataset.GetId()
			} else if len(datasetId) > 0 {
				data, _, err := ApiClient.GetDatasets(cmd.Context()).Execute()
				cobra.CheckErr(err)
				found := false
				for _, dataset := range data.Datasets {
					if dataset.GetId() == datasetId {
						fmt.Println("Found dataset:", dataset.GetName())
						found = true
						break
					}
				}
				if !found {
					cobra.CheckErr("Didn't find dataset with id: " + datasetId)
				}
			}
			config.CreateDatasetConfig(datasetId)
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId of existing dataset")
	cmd.Flags().StringVar(&newDatasetName, "new", "", "Name for new database")
	cmd.MarkFlagsMutuallyExclusive("new", "datasetId")
	RootCommand.AddCommand(cmd)
}
