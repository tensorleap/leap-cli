package model

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	runPkg "github.com/tensorleap/leap-cli/pkg/run"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	tlApi "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"k8s.io/utils/strings/slices"
)

var MODEL_TYPES = []string{
	string(tlApi.IMPORTMODELTYPE_JSON_TF2),
	string(tlApi.IMPORTMODELTYPE_ONNX),
	string(tlApi.IMPORTMODELTYPE_PB_TF2),
	string(tlApi.IMPORTMODELTYPE_H5_TF2),
}

func GetDefaultMessageFromModelPath(filePath string) string {
	fileName := filepath.Base(filePath)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func InitMessage(msg *string, defaultMsg string) error {
	if len(*msg) > 0 {
		return nil
	}
	validate := func(val interface{}) error {
		str, ok := val.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", val)
		}
		str = strings.TrimSpace(str)
		if len(str) == 0 {
			return fmt.Errorf("message cannot be empty")
		}
		return nil
	}
	msgPrompt := &survey.Input{
		Message: "Enter model name",
		Help:    "This name will be displayed in the model version list",
		Default: defaultMsg,
	}
	err := survey.AskOne(msgPrompt, msg, survey.WithValidator(validate))
	if err != nil {
		return err
	}
	return nil
}

func SelectModelType(modelType *string, modelPath string) error {

	// Extract the file extension from the modelPath
	ext := strings.ToLower(filepath.Ext(modelPath))

	if len(*modelType) > 0 && slices.Contains(MODEL_TYPES, *modelType) {
		log.Warn(fmt.Sprintf("Model type %s not supported. Supported types are: %s", *modelType, MODEL_TYPES))
	}

	switch ext {
	case ".json":
		*modelType = string(tlApi.IMPORTMODELTYPE_JSON_TF2)
		return nil
	case ".h5":
		*modelType = string(tlApi.IMPORTMODELTYPE_H5_TF2)
		return nil
	case ".onnx":
		*modelType = string(tlApi.IMPORTMODELTYPE_ONNX)
		return nil
	case ".tar.gz":
		*modelType = string(tlApi.IMPORTMODELTYPE_PB_TF2)
		return nil
	}

	modelTypePrompt := &survey.Select{
		Message: "Select model type?",
		Options: MODEL_TYPES,
	}
	err := survey.AskOne(modelTypePrompt, modelType)
	if err != nil {
		return err
	}
	return nil
}

type VersionStatus string

const (
	VersionStatus_RUNNING  VersionStatus = "RUNNING"
	VersionStatus_FAILED   VersionStatus = "FAILED"
	VersionStatus_FINISHED VersionStatus = "FINISHED"
)

type VersionInfo struct {
	VersionId   string
	VersionName string
	HasModel    bool
	Status      VersionStatus
}

func AskUserForModelPathOrOverwrite(ctx context.Context, projectId string) (isOverwrite bool, overwriteVersionInfo *VersionInfo, modelPath string, err error) {
	overwriteVersionInfo, err = AskUserForNewVersionOrSelectExistingVersion(ctx, projectId)
	if err != nil {
		return false, nil, "", err
	}
	if overwriteVersionInfo != nil {
		if overwriteVersionInfo.HasModel {
			return true, overwriteVersionInfo, "", nil
		}
		log.Warn("The selected version does not have a model, you will be asked for a model path to import")
	}
	modelPath, err = AskUserForModelPath(ctx)
	if err != nil {
		return false, nil, "", err
	}
	return false, overwriteVersionInfo, modelPath, nil
}

func AskUserForNewVersionOrSelectExistingVersion(ctx context.Context, projectId string) (*VersionInfo, error) {

	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return nil, err
	}

	subTypes := []tensorleapapi.JobSubType{
		tensorleapapi.JOBSUBTYPE_IMPORT_MODEL,
		tensorleapapi.JOBSUBTYPE_CODE_PARSE,
	}

	runs, err := runPkg.GetRuns(ctx, subTypes, []tensorleapapi.JobStatus{}, projectId)
	if err != nil {
		return nil, err
	}
	runsStatusesPerVersionId := make(map[string][]tensorleapapi.RunProcess)
	for _, run := range runs {
		runsStatusesPerVersionId[run.GetVersionId()] = append(runsStatusesPerVersionId[run.GetVersionId()], run)
	}

	maxLengthOfVersionName := 0
	for _, version := range versions {
		if len(version.GetNotes()) > maxLengthOfVersionName {
			maxLengthOfVersionName = len(version.GetNotes())
		}
	}

	versionInfos := []VersionInfo{}
	options := []string{
		"Create new",
	}

	for _, version := range versions {
		status, hasModel := calcVersionStatus(&version, runsStatusesPerVersionId[version.GetCid()])

		updateAt := api.FormatDateToLocalTime(version.GetUpdatedAt())
		displayName := version.GetNotes()

		if len(displayName) < maxLengthOfVersionName {
			displayName = displayName + strings.Repeat(" ", maxLengthOfVersionName-len(displayName))
		}
		prefixIcon := ""

		switch status {
		case VersionStatus_FAILED:
			prefixIcon = text.FgRed.Sprint("✖")
		case VersionStatus_RUNNING:
			prefixIcon = text.FgYellow.Sprint("▶")
		case VersionStatus_FINISHED:
			prefixIcon = text.FgGreen.Sprint("✔")
		}
		displayName = fmt.Sprintf("%s %s %s", prefixIcon, displayName, updateAt)
		options = append(options, displayName)

		versionInfos = append(versionInfos, VersionInfo{
			VersionId:   version.GetCid(),
			VersionName: version.GetNotes(),
			HasModel:    hasModel,
			Status:      status,
		})
	}

	selectedIndex := 0
	hasOptions := len(options) > 1
	if hasOptions {
		prompt := &survey.Select{
			Message: "Create new, or overwrite existing model version",
			Options: options,
			Default: selectedIndex,
		}
		err = survey.AskOne(prompt, &selectedIndex)
		if err != nil {
			return nil, err
		}
	}
	isCreateNew := selectedIndex == 0
	if isCreateNew {
		return nil, nil
	}
	return &versionInfos[selectedIndex-1], nil
}

func AskUserForModelPath(ctx context.Context) (modelPath string, err error) {
	allowedExt := []string{".h5", ".onnx"}
	return local.SelectFile(allowedExt, "Select model file")
}

func calcVersionStatus(version *tensorleapapi.SlimVersion, jobs []tensorleapapi.RunProcess) (VersionStatus, bool) {
	hasModel := len(version.GetSessions()) > 0
	if jobs == nil {
		jobs = []tensorleapapi.RunProcess{}
	}

	lastJobByType := make(map[string]tensorleapapi.RunProcess)
	for _, job := range jobs {
		if last, ok := lastJobByType[job.GetJobType()]; !ok || job.GetCreatedAt() > last.GetCreatedAt() {
			lastJobByType[job.GetJobType()] = job
		}
	}
	for _, job := range lastJobByType {
		if api.IsJobRunning(job.GetStatus()) {
			return VersionStatus_RUNNING, hasModel
		}
	}
	if !hasModel {
		return VersionStatus_FAILED, hasModel
	}
	for _, job := range lastJobByType {
		if api.IsJobFailed(job.GetStatus()) {
			return VersionStatus_FAILED, hasModel
		}
	}

	return VersionStatus_FINISHED, hasModel
}
