package model

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

const DEFAULT_BATCH_SIZE = 1

func GetLatestEvaluateBatchSize(ctx context.Context, projectId string) (int, error) {
	sessionRuns, err := GetSessionRunsEvaluate(ctx, projectId, []string{})
	if err != nil {
		return 0, err
	}

	var latestSessionRun *tensorleapapi.SessionRunData
	for _, sessionRun := range sessionRuns.GetEvaluateSessionRuns() {
		if sessionRun.HasEvaluateParams() && (latestSessionRun == nil || latestSessionRun.GetCreatedAt().Before(sessionRun.GetCreatedAt())) {
			latestSessionRun = &sessionRun
		}
	}
	if latestSessionRun == nil {
		return 0, nil
	}
	return int(latestSessionRun.EvaluateParams.GetBatchSize()), nil
}

func AskForBatchSize(defaultBatchSize int) (int, error) {

	if defaultBatchSize <= 0 {
		defaultBatchSize = DEFAULT_BATCH_SIZE
	}

	var batchSizeStr string
	prompt := &survey.Input{
		Message: "Enter batch size for evaluation:",
		Default: fmt.Sprintf("%v", defaultBatchSize),
	}
	err := survey.AskOne(prompt, &batchSizeStr, survey.WithValidator(func(val interface{}) error {
		batchSizeStr := val.(string)
		batchSize, err := strconv.Atoi(batchSizeStr)
		if err != nil {
			return fmt.Errorf("invalid batch size: %s", batchSizeStr)
		}
		if batchSize <= 0 {
			return fmt.Errorf("batch size must be greater than 0")
		}
		return nil
	}))
	if err != nil {
		return 0, err
	}

	batchSize, err := strconv.Atoi(batchSizeStr)
	if err != nil {
		return 0, fmt.Errorf("invalid batch size: %s", batchSizeStr)
	}

	return batchSize, nil
}

func GenerateShortHash() string {
	bytes := make([]byte, 2)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

var evalNameSuffixRegex = regexp.MustCompile(`-(\d+)$`)

// extractNumberSuffix extracts the trailing number suffix from a name like "eval_3".
// Returns -1 if no suffix is found.
func extractNumberSuffix(name string) int {
	matches := evalNameSuffixRegex.FindStringSubmatch(name)
	if len(matches) < 2 {
		return -1
	}
	number, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1
	}
	return number
}

func GenerateEvalName(versionName string, sessionRuns []tensorleapapi.SessionRunData) string {
	if len(sessionRuns) == 0 {
		return versionName
	}

	lastNumber := 1
	for _, sessionRun := range sessionRuns {
		number := extractNumberSuffix(sessionRun.GetName())
		if number > lastNumber {
			lastNumber = number
		}
	}
	return fmt.Sprintf("%s-%v", versionName, lastNumber+1)
}

func WaitForSessionAfterVersionPush(ctx context.Context, projectId, versionId string) (string, []string, error) {

	sessionIds := []string{}
	params := tensorleapapi.SessionVersionIdRequestParams{
		ProjectId: projectId,
		VersionId: versionId,
	}

	lastSessionId := ""

	waitInterval := 10 * time.Second
	waitTimeout := 10 * time.Minute

	err := api.WaitForCondition(ctx, "Waiting for version to be ready...", func() (bool, error) {

		result, res, err := api.ApiClient.GetSessionsByVersionId(ctx).SessionVersionIdRequestParams(params).Execute()
		if err = api.CheckRes(res, err); err != nil {
			return false, fmt.Errorf("failed to get last session id: %w", err)
		}
		var lastSession *tensorleapapi.Session
		for _, session := range result.GetSessions() {
			if lastSession == nil || lastSession.GetCreatedAt().Before(session.GetCreatedAt()) {
				lastSession = &session
			}
		}
		if lastSession == nil {
			return false, nil
		}
		for _, session := range result.GetSessions() {
			sessionIds = append(sessionIds, session.GetCid())
		}
		lastSessionId = lastSession.GetCid()

		return true, nil
	}, waitInterval, waitTimeout)

	if err != nil {
		return "", nil, fmt.Errorf("failed to wait for session after version push: %w", err)
	}

	return lastSessionId, sessionIds, nil
}

func RunEvaluate(ctx context.Context, projectId, versionId string, sessionId string, batchSize int, evalName string) error {
	existingSessionParams := tensorleapapi.NewEvaluateExistingSessionParams(
		versionId,
		projectId,
		float64(batchSize),
		evalName,
		"Evaluation triggered from CLI",
		0,
		sessionId,
	)

	evaluateParams := tensorleapapi.EvaluateParams{
		EvaluateExistingSessionParams: existingSessionParams,
	}

	log.Info("Starting evaluation...")
	job, res, err := api.ApiClient.Evaluate(ctx).EvaluateParams(evaluateParams).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to start evaluation: %w", err)
	}

	log.Infof("Evaluation started with job ID: %s", job.GetCid())
	return nil
}
