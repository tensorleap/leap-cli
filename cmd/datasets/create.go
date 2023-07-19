package datasets

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/datasets"
	"github.com/tensorleap/cli-go/pkg/log"
)

func NewCreateCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create [dataset-name]",
		Short: "Create a new dataset",
		Long:  `Create a new dataset`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var newDatasetName string
			ctx := cmd.Context()
			if len(args) > 0 {
				newDatasetName = args[0]
			} else {
				newDatasetName = datasets.AskForDatasetName()
			}
			dataset, err := datasets.CreateNewDataset(ctx, newDatasetName)
			if err != nil {
				return err
			}
			log.Infof("Creating dataset %s", dataset.GetName())

			datasetName := dataset.GetName()
			_, dirExistsErr := os.Stat(datasetName)
			if dirExistsErr == nil {
				return fmt.Errorf("Can't pull '%s' dataset, directory named '%s' already exists on current directory", datasetName, datasetName)
			}

			err = datasets.CreateDatasetTemplate(dataset.GetCid(), datasetName)
			if err != nil {
				return err
			}

			log.Infof("'%s' dataset created into '%s' directory. Run 'cd %s' to navigate into", datasetName, datasetName, datasetName)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewCreateCmd())
}
