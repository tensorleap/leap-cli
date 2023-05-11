package datasets

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "datasets",
  Short: "Manage tensorleap datasets",
	Long:  `Manage tensorleap datasets`,
}
