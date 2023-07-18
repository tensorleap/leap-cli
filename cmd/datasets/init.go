package datasets

import (
	"errors"

	"github.com/spf13/cobra"
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
			var dataset *tensorleapapi.Dataset = nil
			var err error = nil
			if len(newDatasetName) > 0 {
				dataset, err = datasets.CreateNewDataset(cmd.Context(), newDatasetName)
			} else if len(datasetId) > 0 {
				dataset, err = datasets.GetDatasetById(cmd.Context(), datasetId)
			}
			if err != nil {
				return err
			}
			return datasets.CreateDatasetTemplate(dataset.GetCid(), "")
		},
	}

	cmd.Flags().StringVar(&datasetId, "datasetId", "", "DatasetId of existing dataset")
	cmd.Flags().StringVar(&newDatasetName, "new", "", "Name for new database")
	cmd.MarkFlagsMutuallyExclusive("new", "datasetId")
	RootCommand.AddCommand(cmd)
}
