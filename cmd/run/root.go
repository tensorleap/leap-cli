package run

import (
	"github.com/spf13/cobra"
)

func NewRunCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "run",
		Aliases: []string{"runs", "job", "jobs"},
		Short:   "Manage runs",
		Long:    "Manage runs",
	}
}

var RootCommand = NewRunCmd()
