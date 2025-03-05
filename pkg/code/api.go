package code

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/tensorleap/leap-cli/pkg/api"
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

func GetLatestVersion(ctx context.Context, codeIntegrationId string, branch string) (*tensorleapapi.DatasetVersion, error) {
	params := *tensorleapapi.NewGetLatestDatasetVersionParams(codeIntegrationId)
	if len(branch) == 0 {
		branch = "master"
	}
	params.SetBranch(branch)

	version, _, err := ApiClient.GetLatestDatasetVersion(ctx).
		GetLatestDatasetVersionParams(params).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to get latest version for code integration id: %s", codeIntegrationId)
	}
	if isDatasetVersionEmpty(version.LatestVersion) {
		return version.LatestVersion, ErrEmptyCodeIntegrationVersion
	}
	return version.LatestVersion, nil
}

func GetCodeIntegration(ctx context.Context, id string) (*CodeIntegrationVersion, error) {
	res, response, err := ApiClient.GetDatasetVersion(ctx).GetDatasetVersionParams(*tensorleapapi.NewGetDatasetVersionParams(id)).Execute()
	if err = api.CheckRes(response, err); err != nil {
		return nil, err
	}
	return &res.DatasetVersion, err
}

func AddCodeIntegrationVersion(ctx context.Context, tarGzFile io.Reader, fileSize int64, codeIntegrationId, entryFile, secretId, branch, message string) (*CodeIntegrationVersion, error) {

	getDatasetVersionUploadUrlParams := *tensorleapapi.NewGetDatasetVersionUploadUrlParams(
		codeIntegrationId,
	)

	data, _, err := ApiClient.GetDatasetVersionUploadUrl(ctx).
		GetDatasetVersionUploadUrlParams(getDatasetVersionUploadUrlParams).
		Execute()

	if err != nil {
		return nil, err
	}

	uploadUrl := data.GetUrl()
	if err := UploadFile(uploadUrl, tarGzFile, fileSize); err != nil {
		return nil, err
	}

	saveDatasetVersionParams := *tensorleapapi.NewSaveDatasetVersionParams(
		codeIntegrationId,
		uploadUrl,
		entryFile,
	)

	if len(branch) == 0 {
		branch = "master"
	}
	saveDatasetVersionParams.SetBranch(branch)

	if len(secretId) > 0 {
		saveDatasetVersionParams.SecretManagerId = &secretId
	}

	if len(pythonVersion) > 0 {
		saveDatasetVersionParams.GenericBaseImageType = &pythonVersion
	}

	if len(message) > 0 {
		saveDatasetVersionParams.SetNote(message)
	}

	log.Info("Creating new code integration version...")
	res, _, err := ApiClient.SaveDatasetVersion(ctx).
		SaveDatasetVersionParams(saveDatasetVersionParams).
		Execute()
	if err != nil {
		return nil, err
	}
	for _, latestVersionPerBranch := range res.Dataset.LatestVersions {
		if latestVersionPerBranch.Branch == branch {
			return latestVersionPerBranch.Latest, nil
		}
	}

	return nil, fmt.Errorf("failed to create a new code integration version")
}

func WaitForCodeIntegrationStatus(ctx context.Context, codeIntegrationId string) (ok bool, codeIntegrationVersion *CodeIntegrationVersion, err error) {
	message := "Waiting for code parser result..."
	sleepDuration := 3 * time.Second
	timeoutDuration := 10 * time.Minute
	condition := func() (bool, error) {
		codeIntegrationVersion, err = GetCodeIntegration(ctx, codeIntegrationId)
		if err != nil {
			return false, fmt.Errorf("failed to wait for the integration code status: %v", err)
		}

		switch codeIntegrationVersion.TestStatus {
		case tensorleapapi.TESTSTATUS_TEST_SUCCESS:
			ok = true
			return true, nil
		case tensorleapapi.TESTSTATUS_TEST_FAIL:
			ok = false
			return true, nil
		}

		return false, nil
	}

	err = WaitForCondition(ctx, message, condition, sleepDuration, timeoutDuration)

	if err == ErrorTimeout {
		return false, nil, fmt.Errorf("timeout occurred while waiting for the integration code status")
	}
	if err != nil {
		return false, nil, fmt.Errorf("failed to wait for the integration code status: %v", err)
	}

	return
}
