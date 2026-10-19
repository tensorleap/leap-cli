package run

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/model"
	runPkg "github.com/tensorleap/leap-cli/pkg/run"
)

func NewStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "stop <runId> [runId...]",
		Aliases: []string{"terminate"},
		Args:    cobra.MinimumNArgs(1),
		Short:   "Stop (terminate) one or more runs",
		Long: `Terminate runs by id: queued runs are cancelled, running ones are killed.
A run that already ended is reported and skipped, so repeating the command is safe.

Examples:
  leap run stop 665f1c0e9a1b2c3d4e5f6a7b
  leap job stop 665f1c0e9a1b2c3d4e5f6a7b 665f1c0e9a1b2c3d4e5f6a7c`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var errs []error
			for _, runId := range args {
				job, err := model.GetJobById(cmd.Context(), runId)
				if err != nil {
					errs = append(errs, err)
					continue
				}
				status := job.GetStatus()
				if api.IsJobFailed(status) || api.IsJobFinished(status) {
					fmt.Fprintf(cmd.OutOrStdout(), "run %s is already %s; nothing to stop\n", runId, status)
					continue
				}
				if err := runPkg.StopRun(cmd.Context(), runId); err != nil {
					errs = append(errs, fmt.Errorf("failed to stop run %s: %w", runId, err))
					continue
				}
				fmt.Fprintf(cmd.OutOrStdout(), "stopping run %s (%s -> TERMINATED)\n", runId, status)
			}
			return errors.Join(errs...)
		},
	}
}

func init() {
	RootCommand.AddCommand(NewStopCmd())
}
