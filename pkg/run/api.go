package run

import (
	"context"
	"fmt"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type RunEntity = tensorleapapi.RunProcess

var RunEntityDesc = entity.NewEntityDescriptor(
	"run",
	"runs",
	func(r *RunEntity) string {
		createdAt, err := ParseAndFormatDate(r.GetCreatedAt())

		if err != nil {
			return fmt.Sprintf("%s \t %s \t %s", r.GetJobType(), r.GetCreatedAt(), r.GetStatus())
		}
		return fmt.Sprintf("%s \t %s \t %s", r.GetJobType(), createdAt, r.GetStatus())
	},
	func(r *RunEntity) string { return r.GetJobId() },
)

func GetRuns(ctx context.Context, subTypes []tensorleapapi.JobSubType, statuses []tensorleapapi.JobStatus) ([]RunEntity, error) {
	req := api.ApiClient.GetTeamJobs(ctx).GetJobsFilterParams(tlApi.GetJobsFilterParams{
		SubTypes: subTypes,
		Status:   statuses,
	})
	resp, _, err := req.Execute()
	if err != nil {
		return nil, err
	}
	return resp.Processes, nil
}
