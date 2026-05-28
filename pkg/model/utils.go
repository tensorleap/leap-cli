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
	"k8s.io/utils/strings/slices"
)

var MODEL_TYPES = []string{
	string(tensorleapapi.IMPORTMODELTYPE_JSON_TF2),
	string(tensorleapapi.IMPORTMODELTYPE_ONNX),
	string(tensorleapapi.IMPORTMODELTYPE_PB_TF2),
	string(tensorleapapi.IMPORTMODELTYPE_H5_TF2),
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

	ext := strings.ToLower(filepath.Ext(modelPath))

	if len(*modelType) > 0 && slices.Contains(MODEL_TYPES, *modelType) {
		log.Warn(fmt.Sprintf("Model type %s not supported. Supported types are: %s", *modelType, MODEL_TYPES))
	}

	switch ext {
	case ".json":
		*modelType = string(tensorleapapi.IMPORTMODELTYPE_JSON_TF2)
		return nil
	case ".h5":
		*modelType = string(tensorleapapi.IMPORTMODELTYPE_H5_TF2)
		return nil
	case ".onnx":
		*modelType = string(tensorleapapi.IMPORTMODELTYPE_ONNX)
		return nil
	case ".tar.gz":
		*modelType = string(tensorleapapi.IMPORTMODELTYPE_PB_TF2)
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
	VersionId        string
	VersionName      string
	HasModel         bool
	HasUploadedModel bool
	Status           VersionStatus
}

func AskUserForModelPathOrOverwrite(ctx context.Context, projectId string, currentName *string, runEval bool) (isOverwrite bool, overwriteVersionInfo *VersionInfo, modelPath string, err error) {
	overwriteVersionInfo, err = AskUserForNewVersionOrSelectExistingVersion(ctx, projectId, runEval)
	if err != nil {
		return false, nil, "", err
	}
	if overwriteVersionInfo != nil {
		if overwriteVersionInfo.HasModel || overwriteVersionInfo.HasUploadedModel {
			return true, overwriteVersionInfo, "", nil
		}
		log.Warn("The selected version does not have a model, you will be asked for a model path to import")
		modelPath, err = AskUserForModelPath(ctx)
		if err != nil {
			return false, nil, "", err
		}
		return false, overwriteVersionInfo, modelPath, nil
	}
	if err = InitMessage(currentName, ""); err != nil {
		return false, nil, "", err
	}
	modelPath, err = AskUserForModelPath(ctx)
	if err != nil {
		return false, nil, "", err
	}
	return false, nil, modelPath, nil
}

func AskUserForNewVersionOrSelectExistingVersion(ctx context.Context, projectId string, runEval bool) (*VersionInfo, error) {

	versions, err := GetSlimActiveVersions(ctx, projectId)
	if err != nil {
		return nil, err
	}

	runsStatusesPerVersionId, err := GetRunsStatusesPerVersionId(ctx, projectId)
	if err != nil {
		return nil, err
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
		status, hasModel, hasUploadedModel := CalcVersionStatus(&version, runsStatusesPerVersionId[version.GetCid()])
		displayName := FormatVersionDisplayName(&version, status, maxLengthOfVersionName)
		options = append(options, displayName)

		versionInfos = append(versionInfos, VersionInfo{
			VersionId:        version.GetCid(),
			VersionName:      version.GetNotes(),
			HasModel:         hasModel,
			HasUploadedModel: hasUploadedModel,
			Status:           status,
		})
	}

	selectedIndex := 0
	hasOptions := len(options) > 1
	if hasOptions {

		// Only surface the eval hint when the caller didn't pass --eval:
		// with --eval the diff is auto-detected and patched / re-evaluated
		// for them; without --eval we want users to know that overwriting
		// alone just replaces the version and they can drive evaluation
		// from the UI afterwards.
		if !runEval {
			fmt.Println(text.FgYellow.Sprint("\n\n NOTE: Overwriting replaces the version. Use the update-evaluate dialog in the UI to re-evaluate (or re-run with --eval). \n\n"))
		}

		prompt := &survey.Select{
			Message: "Push as new, or overwrite existing",
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

func versionInfoFromSlim(version *tensorleapapi.SlimVersion, runsStatusesPerVersionId map[string][]tensorleapapi.RunProcess) *VersionInfo {
	status, hasModel, hasUploadedModel := CalcVersionStatus(version, runsStatusesPerVersionId[version.GetCid()])
	return &VersionInfo{
		VersionId:        version.GetCid(),
		VersionName:      version.GetNotes(),
		HasModel:         hasModel,
		HasUploadedModel: hasUploadedModel,
		Status:           status,
	}
}

func ResolveVersionInfoFromRef(ctx context.Context, projectId, versionRef string) (*VersionInfo, error) {
	versionRef = strings.TrimSpace(versionRef)
	if versionRef == "" {
		return nil, fmt.Errorf("version reference is empty")
	}
	allVersions, err := GetVersions(ctx, projectId)
	if err != nil {
		return nil, err
	}
	runsStatusesPerVersionId, err := GetRunsStatusesPerVersionId(ctx, projectId)
	if err != nil {
		return nil, err
	}
	for i := range allVersions {
		v := &allVersions[i]
		if v.GetCid() == versionRef {
			return versionInfoFromSlim(v, runsStatusesPerVersionId), nil
		}
	}
	var activeNameMatches []*tensorleapapi.SlimVersion
	for i := range allVersions {
		v := &allVersions[i]
		if v.GetIsActive() && v.GetNotes() == versionRef {
			activeNameMatches = append(activeNameMatches, v)
		}
	}
	switch len(activeNameMatches) {
	case 0:
		return nil, fmt.Errorf("no active version found for %q (use the version id from the UI or API)", versionRef)
	case 1:
		return versionInfoFromSlim(activeNameMatches[0], runsStatusesPerVersionId), nil
	default:
		return pickAmbiguousVersion(versionRef, activeNameMatches, runsStatusesPerVersionId)
	}
}

func PrintResolvedOverwriteTarget(originalRef string, info *VersionInfo) {
	if info == nil {
		return
	}
	id := info.VersionId
	if len(id) > 8 {
		id = id[:8]
	}
	displayName := info.VersionName
	if displayName == "" {
		displayName = "(unnamed)"
	}
	if originalRef == info.VersionName || originalRef == info.VersionId {
		log.Infof("Overwriting %s (id %s)", displayName, id)
	} else {
		log.Infof("Overwriting %s (id %s, matched %q)", displayName, id, originalRef)
	}
}

func pickAmbiguousVersion(
	versionRef string,
	nameMatches []*tensorleapapi.SlimVersion,
	runsStatusesPerVersionId map[string][]tensorleapapi.RunProcess,
) (*VersionInfo, error) {
	maxLen := 0
	for _, v := range nameMatches {
		if len(v.GetNotes()) > maxLen {
			maxLen = len(v.GetNotes())
		}
	}
	options := make([]string, len(nameMatches))
	infos := make([]VersionInfo, len(nameMatches))
	for i, v := range nameMatches {
		status, hasModel, hasUploadedModel := CalcVersionStatus(v, runsStatusesPerVersionId[v.GetCid()])
		options[i] = FormatVersionDisplayName(v, status, maxLen)
		infos[i] = VersionInfo{
			VersionId:        v.GetCid(),
			VersionName:      v.GetNotes(),
			HasModel:         hasModel,
			HasUploadedModel: hasUploadedModel,
			Status:           status,
		}
	}

	idx := 0
	prompt := &survey.Select{
		Message: fmt.Sprintf("Multiple versions named %q — pick one", versionRef),
		Options: options,
		Default: idx,
	}
	if err := survey.AskOne(prompt, &idx); err != nil {
		// Non-TTY (script) or user-cancel — surface the candidate ids
		// so the caller can re-run with --overwrite <id>.
		ids := make([]string, 0, len(nameMatches))
		for _, v := range nameMatches {
			ids = append(ids, v.GetCid())
		}
		const maxList = 12
		if len(ids) > maxList {
			ids = append(ids[:maxList], "...")
		}
		return nil, fmt.Errorf("version name %q is not unique (%d matches); use --overwrite with a version id. Candidates: %s",
			versionRef, len(nameMatches), strings.Join(ids, ", "))
	}
	return &infos[idx], nil
}

func AskUserForModelPath(ctx context.Context) (modelPath string, err error) {
	allowedExt := []string{".h5", ".onnx"}
	ignoredDirs := []string{".venv"}
	return local.SelectFile(allowedExt, "Select model file", ignoredDirs)
}

func FormatVersionDisplayName(version *tensorleapapi.SlimVersion, status VersionStatus, maxLengthOfVersionName int) string {
	displayName := version.GetNotes()
	updateAt := api.FormatDateToLocalTime(version.GetUpdatedAt())

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

	return fmt.Sprintf("%s %s %s", prefixIcon, displayName, updateAt)
}

func GetRunsStatusesPerVersionId(ctx context.Context, projectId string) (map[string][]tensorleapapi.RunProcess, error) {
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

	return runsStatusesPerVersionId, nil
}

func CalcVersionStatus(version *tensorleapapi.SlimVersion, jobs []tensorleapapi.RunProcess) (VersionStatus, bool, bool) {
	hasModel := version.HasModelId()
	hasUploadedModel := version.HasUploadedModel()
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
			return VersionStatus_RUNNING, hasModel, hasUploadedModel
		}
	}
	if !hasModel {
		return VersionStatus_FAILED, hasModel, hasUploadedModel
	}
	for _, job := range lastJobByType {
		if api.IsJobFailed(job.GetStatus()) {
			return VersionStatus_FAILED, hasModel, hasUploadedModel
		}
	}

	return VersionStatus_FINISHED, hasModel, hasUploadedModel
}
