package models

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "models",
  Short: "Manage tensorleap models",
	Long:  `Manage tensorleap models`,
}
