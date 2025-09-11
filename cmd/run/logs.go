package run

import (
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/log"
	runPkg "github.com/tensorleap/leap-cli/pkg/run"
)

func NewLogsCmd() *cobra.Command {
	var output string

	logsCmd := &cobra.Command{
		Use:   "logs [runId]",
		Short: "Show the logs for a run",
		Run: func(cmd *cobra.Command, args []string) {
			runId := args[0]
			if runId == "" {
				log.Error("Run id is required")
				return
			}
			log.Infof("Getting logs for %s", runId)
			logs, err := runPkg.GetRunLogs(cmd.Context(), runId)
			if err != nil {
				log.Error(err)
				return
			}
			if output != "" {
				log.Infof("Wrapping logs in tar file")
				savedFile, err := runPkg.WrapLogsInTarFile(logs, output, runId)
				if err != nil {
					log.Error(err)
					return
				}
				log.Printf("Logs wrapped in tar file: %s", savedFile)
				return

			}
			log.Infof("Printing logs for %s:", runId)
			runPkg.PrintLogs(logs)
		},
	}

	logsCmd.Flags().StringVar(&output, "output", "", "Download the logs to a file")
	return logsCmd
}

func init() {
	RootCommand.AddCommand(NewLogsCmd())
}
