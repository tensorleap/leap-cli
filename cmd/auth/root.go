package auth

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "auth",
  Short: "auth commands",
	Long:  `auth commands`,
}

