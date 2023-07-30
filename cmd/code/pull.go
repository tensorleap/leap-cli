package code

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewPullCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pull [dataset-id]",
		Short: "Pull dataset into a new directory",
		Long:  `Pull dataset into a new directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedDataset *code.CodeIntegration
			var err error
			ctx := cmd.Context()
			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				return err
			}
			if len(args) == 0 {
				selectedDataset, err = entity.SelectEntity(codeIntegrations, code.CodeIntegrationEntityDesc)

			} else {
				datasetId := args[0]
				selectedDataset, err = entity.GetEntityById(datasetId, codeIntegrations, code.CodeIntegrationEntityDesc)
			}
			if err != nil {
				return err
			}

			log.Infof("Pulling dataset %s", selectedDataset.GetName())

			datasetName := selectedDataset.GetName()
			_, dirExistsErr := os.Stat(datasetName)
			if dirExistsErr == nil {
				return fmt.Errorf("can't pull '%s' dataset, directory named '%s' already exists on current directory", datasetName, datasetName)
			}
			latestVersion, err := code.GetLatestVersion(ctx, selectedDataset.GetCid())
			if err == nil {
				files, err := code.CloneCodeIntegrationVersion(ctx, latestVersion, datasetName)
				if err != nil {
					return err
				}
				datasetConfig := code.NewDatasetConfig(selectedDataset.GetCid(), latestVersion.GetCodeEntryFile(), files)
				err = code.SetCodeIntegrationConfig(datasetConfig, datasetName)
				if err != nil {
					return err
				}
			} else if err == code.ErrEmptyCodeIntegrationVersion {
				log.Warn("The selected dataset is empty, Create default template")
				err = code.CreateCodeTemplate(selectedDataset.GetCid(), datasetName)
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
