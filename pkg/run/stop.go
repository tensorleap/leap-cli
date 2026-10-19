package run

import (
	"context"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func StopRun(ctx context.Context, jobId string) error {
	params := tensorleapapi.NewTerminateJobParams(jobId)
	_, res, err := api.ApiClient.TerminateJob(ctx).TerminateJobParams(*params).Execute()
	return api.CheckRes(res, err)
}
