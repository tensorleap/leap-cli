package modelversion

import (
	"context"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type SlimVersionEntity = tensorleapapi.SlimVersion

var SlimVersionEntityDesc = entity.NewEntityDescriptor[SlimVersionEntity](
	"model version",
	"model versions",
	func(p *SlimVersionEntity) string { return p.GetNotes() },
	func(p *SlimVersionEntity) string { return p.GetCid() },
)

func GetSlimVersions(ctx context.Context, projectId string) ([]SlimVersionEntity, error) {
	resp, _, err := api.ApiClient.GetProjectSlimVersions(ctx).GetProjectVersionsParams(*tensorleapapi.NewGetProjectVersionsParams(projectId)).Execute()
	if err != nil {
		return nil, err
	}
	return resp.Versions, nil
}
