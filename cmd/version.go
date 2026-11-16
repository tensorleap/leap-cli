package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/version"
)

func newVersionCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the CLI version, the server's skills level and the skills this CLI ships with",
		Long: `Show the CLI version, the signed-in server's skills level and the skills this CLI ships with.

The Tensorleap skills read this to install the skills version this CLI ships with
(skillsRef) and to tell you when the server is too old for them.
serverLevel is null when no server answers (not signed in, or unreachable).

Examples:
  leap version
  leap version -o json`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			info := version.GetSkillsInfo(cmd.Context())
			switch output {
			case "json":
				encoded, err := json.Marshal(info)
				if err != nil {
					return fmt.Errorf("failed to encode version info: %w", err)
				}
				fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
			case "text":
				server := "unknown (not signed in or unreachable)"
				if info.ServerLevel != nil {
					server = fmt.Sprint(*info.ServerLevel)
				}
				fmt.Fprintf(cmd.OutOrStdout(), "leap %s\nServer skills level: %s\nSkills: %s\n",
					info.Cli, server, info.SkillsRef)
			default:
				return fmt.Errorf("invalid --output %q (allowed: json, text)", output)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "text", "Output format: json or text")
	return cmd
}
