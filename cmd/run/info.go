package run

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/model"
)

func NewInfoCmd() *cobra.Command {
	var output string

	cmd := &cobra.Command{
		Use:     "info [runId]",
		Aliases: []string{"get", "status"},
		Args:    cobra.ExactArgs(1),
		Short:   "Show information about a run",
		Long: `Show information about a run, including its status, steps, the id of any
evaluation chained to it, and — for a failed run — the same errors the
interactive push report shows.

Output is JSON by default so it can be consumed by scripts and agents.

Examples:
  leap run info 665f1c0e9a1b2c3d4e5f6a7b
  leap run info 665f1c0e9a1b2c3d4e5f6a7b --output text`,
		RunE: func(cmd *cobra.Command, args []string) error {
			info, err := model.BuildRunInfo(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			switch output {
			case "json":
				encoded, err := json.MarshalIndent(info, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to encode run info: %w", err)
				}
				// Straight to stdout, not the logger: its INFO/timestamp prefix
				// would corrupt the payload for `| jq` and every other consumer.
				fmt.Fprintln(cmd.OutOrStdout(), string(encoded))
				return nil
			case "text":
				fmt.Fprintln(cmd.OutOrStdout(), renderRunInfoText(info))
				return nil
			default:
				return fmt.Errorf("invalid --output %q (allowed: json, text)", output)
			}
		},
	}

	cmd.Flags().StringVarP(&output, "output", "o", "json", "Output format: json or text")
	return cmd
}

func renderRunInfoText(info *model.RunInfo) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Run %s\n", info.JobId)
	fmt.Fprintf(&b, "  type:      %s\n", strings.TrimSpace(info.Type+" "+info.SubType))
	fmt.Fprintf(&b, "  status:    %s\n", info.Status)
	if info.ProjectId != "" {
		fmt.Fprintf(&b, "  project:   %s\n", info.ProjectId)
	}
	if info.VersionId != "" {
		fmt.Fprintf(&b, "  version:   %s\n", info.VersionId)
	}
	fmt.Fprintf(&b, "  created:   %s\n", info.CreatedAt)
	fmt.Fprintf(&b, "  updated:   %s\n", info.UpdatedAt)

	if info.ChainedEvaluate != nil {
		fmt.Fprintf(&b, "\nChained evaluate: %s\n", info.ChainedEvaluate.Status)
		if info.ChainedEvaluate.JobId != "" {
			fmt.Fprintf(&b, "  jobId: %s\n", info.ChainedEvaluate.JobId)
		}
		if info.ChainedEvaluate.Error != "" {
			fmt.Fprintf(&b, "  error: %s\n", info.ChainedEvaluate.Error)
		}
	}

	if len(info.Steps) > 0 {
		b.WriteString("\nSteps:\n")
		for _, step := range info.Steps {
			progress := ""
			if step.Total > 0 {
				progress = fmt.Sprintf(" (%.0f/%.0f)", step.Current, step.Total)
			}
			fmt.Fprintf(&b, "  %-10s %s%s\n", step.Status, step.Name, progress)
		}
	}

	if info.ErrorReport != nil {
		// Same pages the interactive push viewer renders, flattened — so the
		// text output and the TUI show identical content.
		b.WriteString("\n")
		b.WriteString(info.ErrorReport.ToPushReportPages().RenderFullContent())
	}

	for _, hint := range info.Hints {
		fmt.Fprintf(&b, "\n%s", hint)
	}

	return b.String()
}

func init() {
	RootCommand.AddCommand(NewInfoCmd())
}
