package code

import (
	"context"
	"fmt"
	"io"

	. "github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

type CodeIntegration = tensorleapapi.Dataset
type CodeIntegrationVersion = tensorleapapi.DatasetVersion

var CodeIntegrationEntityDesc = entity.NewEntityDescriptor[CodeIntegration](
	"code integration",
	"code integrations",
	func(p *CodeIntegration) string { return p.GetName() },
	func(p *CodeIntegration) string { return p.GetCid() },
)

var CodeIntegrationVersionEntityDesc = entity.NewEntityDescriptor[CodeIntegrationVersion](
	"code integration version",
	"code integration versions",
	func(p *CodeIntegrationVersion) string { return p.GetNote() },
	func(p *CodeIntegrationVersion) string { return p.GetCid() },
)

func GetCodeIntegrations(ctx context.Context) ([]CodeIntegration, error) {
	data, _, err := ApiClient.GetDatasets(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return data.Datasets, nil
}

func GetCodeIntegrationVersions(ctx context.Context, codeIntegrationId string) ([]CodeIntegrationVersion, error) {
	params := *tensorleapapi.NewGetDatasetVersionsParams(codeIntegrationId)
	data, _, err := ApiClient.GetDatasetVersions(ctx).
		GetDatasetVersionsParams(params).
		Execute()
	if err != nil {
		return nil, err
	}
	return data.DatasetVersions, nil
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

func DeleteCodeIntegration(ctx context.Context, codeIntegration *CodeIntegration) error {
	_, _, err := ApiClient.TrashDataset(ctx).
		TrashDatasetParams(*tensorleapapi.NewTrashDatasetParams(codeIntegration.GetCid())).
		Execute()
	if err != nil {
		return err
	}
	entity.InfoDeletion(codeIntegration, CodeIntegrationEntityDesc)
	return nil
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

func AddCodeIntegrationVersion(ctx context.Context, tarGzFile io.Reader, codeIntegration *CodeIntegration, entryFile string, secretId string) (*CodeIntegrationVersion, error) {

	data, _, err := ApiClient.GetDatasetVersionUploadUrl(ctx).Execute()
	if err != nil {
		return nil, err
	}

	uploadUrl := data.GetUrl()
	if err := UploadFile(uploadUrl, tarGzFile); err != nil {
		return nil, err
	}

	saveDatasetVersionParams := *tensorleapapi.NewSaveDatasetVersionParams(
		codeIntegration.GetCid(),
		uploadUrl,
		entryFile,
	)

	if len(secretId) > 0 {
		saveDatasetVersionParams.SecretManagerId = &secretId
	}

	log.Info("Creating new dataset version...")
	res, _, err := ApiClient.SaveDatasetVersion(ctx).
		SaveDatasetVersionParams(saveDatasetVersionParams).
		Execute()
	if err != nil {
		return nil, err
	}

	log.Info("Done!")
	return res.Dataset.LatestVersion, nil

}
