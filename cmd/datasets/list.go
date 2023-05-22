package datasets

import (
	"fmt"

	"github.com/spf13/cobra"
	. "github.com/tensorleap/cli-go/pkg/api"
	"github.com/tensorleap/cli-go/pkg/config"
)

func init() {
  RootCommand.AddCommand(&cobra.Command{
    Use:   "list",
    Short: "List available datasets",
    Long:  `List available datasets`,
    Run: func(cmd *cobra.Command, args []string) {
      config.VerifyLoggedIn()
      data, _, err := ApiClient.GetDatasets(cmd.Context()).Execute()
      cobra.CheckErr(err)
      for _, dataset := range data.Datasets {
        fmt.Println(dataset.GetId(), "-",  dataset.GetName())
      }
    },
  })
}
