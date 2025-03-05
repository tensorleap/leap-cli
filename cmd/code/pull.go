package code

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

const MAPPING_FILE_NAME = "leap_mapping.yaml"

func NewPullCmd() *cobra.Command {
	var mappingOnly bool

	var cmd = &cobra.Command{
		Use:   "pull [dataset-id] [branch]",
		Short: "Pull dataset into a new directory",
		Long:  `Pull dataset into a new directory`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var selectedDataset *code.CodeIntegration
			var err error
			var selectedBranch string
			ctx := cmd.Context()
			codeIntegrations, err := code.GetCodeIntegrations(ctx)
			if err != nil {
				return err
			}
			if len(args) == 0 {
				selectedDataset, err = entity.SelectEntity(codeIntegrations, code.CodeIntegrationEntityDesc)
				if err != nil {
					return err
				}
				branches := code.BranchesFromCodeIntegration(selectedDataset)
				selectedBranch, err = code.SelectBranch(branches, selectedDataset.DefaultBranch)
			} else {
				datasetId := args[0]
				selectedDataset, err = entity.GetEntityById(datasetId, codeIntegrations, code.CodeIntegrationEntityDesc)
				if args[1] != "" {
					selectedBranch = args[1]
				}
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
			latestVersion, err := code.GetLatestVersion(ctx, selectedDataset.GetCid(), selectedBranch)
			var secretId string
			if err == nil {
				secretId = latestVersion.Metadata.GetSecretManagerId()
				var specificFileName = ""
				if mappingOnly {
					specificFileName = MAPPING_FILE_NAME
				}
				files, err := code.CloneCodeIntegrationVersion(ctx, latestVersion, datasetName, specificFileName)
				if err != nil {
					return err
				}

				if !mappingOnly {
					workspaceConfig := workspace.NewWorkspaceConfig(selectedDataset.GetCid(), "", latestVersion.GetCodeEntryFile(), secretId, latestVersion.Branch, latestVersion.GetGenericBaseImageType(), files)
					err = workspace.SetWorkspaceConfig(workspaceConfig, datasetName)
					if err != nil {
						return err
					}
				}
			} else if err == code.ErrEmptyCodeIntegrationVersion {
				log.Warn("The selected dataset is empty, Create default template")
				err = workspace.CreateCodeTemplate(selectedDataset.GetCid(), "", secretId, datasetName, selectedBranch, latestVersion.GetGenericBaseImageType())
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

	cmd.Flags().BoolVarP(&mappingOnly, "mapping-only", "f", false, "Pull mapping only")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPullCmd())
}
