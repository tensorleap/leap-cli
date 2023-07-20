package datasets

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/datasets"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func init() {
	var datasetId string
	var newDatasetName string

	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create a .tensorleap.yaml file in the current directory",
		Long:  `Create a .tensorleap.yaml file in the current directory`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(datasetId) == 0 && len(newDatasetName) == 0 {
				return errors.New("Error: flag(s) \"datasetId\" or \"new\" must be set")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(newDatasetName) > 0 {
				dataset, err := datasets.CreateNewDataset(cmd.Context(), newDatasetName)
				if err != nil {
					return err
				}
				datasetId = dataset.GetCid()
			} else if len(datasetId) > 0 {
				data, _, err := ApiClient.GetDatasets(cmd.Context()).Execute()
				if err != nil {
					return err
				}
				found := false
				for _, dataset := range data.Datasets {
					if dataset.GetCid() == datasetId {
						fmt.Println("Found dataset:", dataset.GetName())
						found = true
						break
					}
				}
				if !found {
					return fmt.Errorf("Didn't find dataset with id: %v", datasetId)
				}
			}

			return datasets.CreateDatasetTemplate(datasetId, "")
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId of existing dataset")
	cmd.Flags().StringVar(&newDatasetName, "new", "", "Name for new database")
	cmd.MarkFlagsMutuallyExclusive("new", "datasetId")
	RootCommand.AddCommand(cmd)
}
