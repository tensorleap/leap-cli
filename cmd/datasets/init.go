package datasets

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/cli-go/pkg/config"
)

func init() {
	RootCommand.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "Create a .tensorleap.yaml file in the current directory",
		Long:  `Create a .tensorleap.yaml file in the current directory`,
		Run: func(cmd *cobra.Command, args []string) {
			config.CreateDatasetConfig()
		},
	})
}
