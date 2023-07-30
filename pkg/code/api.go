package code

import (
	"context"
	"fmt"

	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type CodeIntegration = tensorleapapi.Dataset

var CodeIntegrationEntityDesc = entity.NewEntityDescriptor[CodeIntegration](
	"code integration",
	"code integrations",
	func(p *CodeIntegration) string { return p.GetName() },
	func(p *CodeIntegration) string { return p.GetCid() },
)

func GetCodeIntegrations(ctx context.Context) ([]CodeIntegration, error) {
	data, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return data.Datasets, nil
}

func AddCodeIntegration(ctx context.Context, name string) (*CodeIntegration, error) {
	log.Println("Creating code integration:", name)
	dataset, _, err := ApiClient.AddDataset(ctx).
		NewDatasetParams(*tensorleapapi.NewNewDatasetParams(*tensorleapapi.NewNullableString(&name))).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("failed to create a new code integration: %v", err)
	}
	codeIntegration := &dataset.Dataset
	entity.InfoCreation(codeIntegration, CodeIntegrationEntityDesc)
	return codeIntegration, nil
}

func GetLatestVersion(ctx context.Context, codeIntegrationId string) (*tensorleapapi.DatasetVersion, error) {
	version, _, err := ApiClient.GetLatestDatasetVersion(ctx).
		GetLatestDatasetVersionParams(*tensorleapapi.NewGetLatestDatasetVersionParams(codeIntegrationId)).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version for code integration id: %s", codeIntegrationId)
	}
	if isDatasetVersionEmpty(version.LatestVersion) {
		return version.LatestVersion, ErrEmptyCodeIntegrationVersion
	}
	return version.LatestVersion, nil
}
