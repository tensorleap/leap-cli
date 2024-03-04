package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func NewClearCmd() *cobra.Command {
	var dryRun bool

	var cmd = &cobra.Command{
		Use:   "clear",
		Short: "Clear storage",
		Long:  `Clear storage`,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			params := tensorleapapi.NewClearStorageRequest()
			params.SetDryRun(dryRun)

			report, res, err := api.ApiClient.ClearStorage(ctx).ClearStorageRequest(*params).Execute()
			if err := api.CheckRes(res, err); err != nil {
				return err
			}

			prettyPrintClearReport(report)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&dryRun, "dryRun", "d", false, "Dry run")

	return cmd
}

func init() {
	RootCommand.AddCommand(NewClearCmd())
}

func prettyPrintClearReport(c *tensorleapapi.ClearStorageResponse) {
	if len(c.UnusedDirs) == 0 {
		fmt.Println("No unused directories found")
		return
	}
	fmt.Println("Remove unused directories:")
	for _, unused := range c.UnusedDirs {
		fmt.Printf("  %s:\n", unused.Name)
		for _, dir := range unused.Dirs {
			fmt.Printf("    - %s\n", dir)
		}
	}
}
