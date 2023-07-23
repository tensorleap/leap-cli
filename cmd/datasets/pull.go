package datasets

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/datasets"
	"github.com/tensorleap/cli-go/pkg/log"
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func NewPullCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pull [dataset-id]",
		Short: "Pull dataset into a new directory",
		Long:  `Pull dataset into a new directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedDataset *tensorleapapi.Dataset
			var err error
			ctx := cmd.Context()
			if len(args) == 0 {
				selectedDataset, err = datasets.AskForDataset(ctx)
			} else {
				datasetId := args[0]
				selectedDataset, err = datasets.GetDatasetById(ctx, datasetId)
			}
			if err != nil {
				return err
			}

			log.Infof("Pulling dataset %s", selectedDataset.GetName())

			datasetName := selectedDataset.GetName()
			_, dirExistsErr := os.Stat(datasetName)
			if dirExistsErr == nil {
				return fmt.Errorf("Can't pull '%s' dataset, directory named '%s' already exists on current directory", datasetName, datasetName)
			}
			latestVersion, err := datasets.GetLatestVersion(ctx, selectedDataset.GetCid())
			if err == nil {
				files, err := datasets.CloneDatasetVersionCode(ctx, latestVersion, datasetName)
				if err != nil {
					return err
				}
				datasetConfig := datasets.NewDatasetConfig(selectedDataset.GetCid(), latestVersion.GetCodeEntryFile(), files)
				err = datasets.SetDatasetConfig(datasetConfig, datasetName)
				if err != nil {
					return err
				}
			} else if err == datasets.EmptyDatasetVersionError {
				log.Warn("The selected dataset is empty, Create default template")
				err = datasets.CreateDatasetTemplate(selectedDataset.GetCid(), datasetName)
				if err != nil {
					return err
				}
			} else {
				return err
			}

			log.Infof("'%s' dataset pulled into '%s' directory. Run 'cd %s' to navigate into", datasetName, datasetName, datasetName)
			return nil
		},
	}
	return cmd
}

func init() {
	RootCommand.AddCommand(NewPullCmd())
}
