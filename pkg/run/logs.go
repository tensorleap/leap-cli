package run

import (
	"context"
	"math"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/api"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func GetRunLogs(ctx context.Context, jobId string) ([]tlApi.PodLogs, error) {

	job, _, err := api.ApiClient.GetJobLogs(ctx).GetJobLogsParams(tlApi.GetJobLogsParams{JobId: jobId}).Execute()
	if err != nil {
		return nil, err
	}

	return job.PodsLogs, nil
}

func GetTopLogs(logs []tlApi.PodLogs, prefix string, count int) string {
	logsStr := ""
	for _, log := range logs {
		if strings.HasPrefix(log.Name, prefix) {
			s := strings.Split(log.Logs, "\n")
			fromIndex := int(math.Max(0, float64(len(s)- count)))
			logsStr = strings.Join(s[fromIndex:], "\n")
			break
		}
	}
	return logsStr
}
