package model

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

const DEFAULT_BATCH_SIZE = 1

// ErrEvaluateCancelled is returned from the interactive concurrent-evaluate
// recovery flow when the user explicitly declines to terminate any running
// job (or aborts the prompt via Ctrl+C). It's a sentinel so RunEvaluate can
// short-circuit to a clean exit instead of treating user intent as a failure.
var ErrEvaluateCancelled = errors.New("evaluation cancelled by user")

func GetLatestEvaluateBatchSize(ctx context.Context, projectId string) (int, error) {
	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return 0, err
	}

	var latest *tensorleapapi.SlimVersion
	for _, v := range versions {
		if v.HasEvaluateParams() && (latest == nil || latest.GetCreatedAt().Before(v.GetCreatedAt())) {
			latest = &v
		}
	}
	if latest == nil {
		return 0, nil
	}
	return int(latest.EvaluateParams.GetBatchSize()), nil
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

// ParseUpdateActionsFromFlags maps CLI tokens to API update actions (repeatable -u / --update).
// The flag still operates on the engine-contract enum (metadata, metric_config,
// insights, visualizations) — the friendlier "what changed in your code" model
// only applies to the interactive prompt; flag callers are usually scripts that
// already know which artifact bucket they want to refresh.
func ParseUpdateActionsFromFlags(parts []string) ([]tensorleapapi.UpdateAction, error) {
	if len(parts) == 0 {
		return nil, nil
	}
	seen := make(map[tensorleapapi.UpdateAction]struct{})
	out := make([]tensorleapapi.UpdateAction, 0, len(parts))
	for _, raw := range parts {
		p := strings.TrimSpace(strings.ToLower(raw))
		if p == "" {
			continue
		}
		act, err := tensorleapapi.NewUpdateActionFromValue(p)
		if err != nil {
			return nil, fmt.Errorf("invalid --update value %q: %w (allowed: metadata, metric_config, insights, visualizations)", raw, err)
		}
		if _, ok := seen[*act]; ok {
			continue
		}
		seen[*act] = struct{}{}
		out = append(out, *act)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no valid --update values after parsing")
	}
	return out, nil
}

// EvaluatePlanKind tells the push flow whether the selected changes can be
// lifted onto the existing evaluation (Update) or need a full re-evaluate
// (Reset). The dialog computes this for us so the call site only has to
// dispatch.
type EvaluatePlanKind int

const (
	EvaluatePlanUpdate EvaluatePlanKind = iota
	EvaluatePlanReset
)

// EvaluatePlan is the result of the interactive "what changed in your code"
// dialog. When Kind == Reset, UpdateActions is empty and the caller should
// run RunEvaluate. Otherwise it's the deduped UpdateAction[] for
// RunUpdateEvaluateArtifact.
type EvaluatePlan struct {
	Kind          EvaluatePlanKind
	UpdateActions []tensorleapapi.UpdateAction
}

// changeOption mirrors the UI's ChangeOption — same 5 buckets, same hints,
// same cascade rules (a metadata change implies insights need to regen).
// Keeping the wording in sync with the dialog in web-ui means a user who
// learned the flow in one surface doesn't have to relearn it in the other.
type changeOption struct {
	key   string
	label string
	hint  string
}

var changeOptions = []changeOption{
	{key: "metadata", label: "Metadata", hint: "Added / updated / removed"},
	{key: "metric_direction", label: "Metric direction", hint: "Direction only — values unchanged"},
	{key: "metrics", label: "Metrics", hint: "Added / updated / removed"},
	{key: "insights", label: "Insights config", hint: ""},
	{key: "visualizations", label: "Visualizations", hint: "Added / updated / removed"},
}

// planUpdateEvaluate is the CLI mirror of the web-ui's planUpdateEvaluate.
// Keeping the two in sync (same cascade, same "metrics ⇒ reset" rule) is what
// makes the prompts behave identically across surfaces.
func planUpdateEvaluate(selected map[string]bool) EvaluatePlan {
	if selected["metrics"] {
		return EvaluatePlan{Kind: EvaluatePlanReset}
	}

	actions := make(map[tensorleapapi.UpdateAction]struct{})
	if selected["metadata"] {
		// Metadata cascades into insights — insights analysis reads from
		// metadata so it has to regenerate too.
		actions[tensorleapapi.UPDATEACTION_METADATA] = struct{}{}
		actions[tensorleapapi.UPDATEACTION_INSIGHTS] = struct{}{}
	}
	if selected["metric_direction"] {
		actions[tensorleapapi.UPDATEACTION_METRIC_CONFIG] = struct{}{}
	}
	if selected["insights"] {
		actions[tensorleapapi.UPDATEACTION_INSIGHTS] = struct{}{}
	}
	if selected["visualizations"] {
		actions[tensorleapapi.UPDATEACTION_VISUALIZATIONS] = struct{}{}
	}

	// Emit in canonical pipeline order so the printed action list reads
	// like the engine actually runs.
	ordered := []tensorleapapi.UpdateAction{
		tensorleapapi.UPDATEACTION_METADATA,
		tensorleapapi.UPDATEACTION_METRIC_CONFIG,
		tensorleapapi.UPDATEACTION_INSIGHTS,
		tensorleapapi.UPDATEACTION_VISUALIZATIONS,
	}
	out := make([]tensorleapapi.UpdateAction, 0, len(actions))
	for _, a := range ordered {
		if _, ok := actions[a]; ok {
			out = append(out, a)
		}
	}
	return EvaluatePlan{Kind: EvaluatePlanUpdate, UpdateActions: out}
}

// AskForEvaluatePlan shows the "what changed in your code" prompt, matching
// the web-ui's UpdateEvaluateDialog. Returning an EvaluatePlan lets the
// caller dispatch to RunEvaluate (Reset) or RunUpdateEvaluateArtifact
// (Update) without having to know how the dialog mapped the picks.
func AskForEvaluatePlan() (EvaluatePlan, error) {
	labels := make([]string, len(changeOptions))
	labelToKey := make(map[string]string, len(changeOptions))
	for i, opt := range changeOptions {
		label := opt.label
		if opt.hint != "" {
			label = fmt.Sprintf("%s — %s", opt.label, opt.hint)
		}
		labels[i] = label
		labelToKey[label] = opt.key
	}

	var selectedLabels []string
	prompt := &survey.MultiSelect{
		Message: "What changed in your code?",
		Options: labels,
	}
	if err := survey.AskOne(prompt, &selectedLabels, survey.WithValidator(survey.MinItems(1))); err != nil {
		return EvaluatePlan{}, err
	}

	selected := make(map[string]bool, len(selectedLabels))
	for _, l := range selectedLabels {
		selected[labelToKey[l]] = true
	}
	return planUpdateEvaluate(selected), nil
}

// hasOwnEvalData mirrors hasOwnEvalData in web-ui/VersionControlContext.
// A version has eval data iff any of these three resource fields is set —
// they're the three signals the engine writes when it lands evaluation
// results on a version.
func hasOwnEvalData(v *tensorleapapi.SlimVersion) bool {
	res := v.GetResources()
	if res.GetEsMetricsIndex() != "" {
		return true
	}
	if res.GetEsModelId() != "" {
		return true
	}
	if res.GetInferenceArtifactId() != "" {
		return true
	}
	return false
}

// HasEvaluatedAncestorOrSelf walks the override chain starting at versionId
// and returns true if any version in the chain (including the start) carries
// evaluation data. Mirrors the UI's hasEvaluatedAncestor logic so the two
// surfaces agree on when there's actually something to "update".
//
// When this returns false the caller should skip the dialog entirely —
// there's no source to lift evaluation results from, so the only sensible
// action is a fresh evaluate.
func HasEvaluatedAncestorOrSelf(ctx context.Context, projectId, versionId string) (bool, error) {
	versions, err := GetVersions(ctx, projectId)
	if err != nil {
		return false, fmt.Errorf("failed to load versions for eval-data check: %w", err)
	}
	byCid := make(map[string]*tensorleapapi.SlimVersion, len(versions))
	for i := range versions {
		byCid[versions[i].GetCid()] = &versions[i]
	}

	cur, ok := byCid[versionId]
	if !ok {
		// Couldn't resolve the starting version — fall back to the safe
		// answer (assume eval data exists so we still show the dialog
		// rather than silently bypassing the user's choice).
		return true, nil
	}

	seen := make(map[string]struct{})
	for cur != nil {
		if hasOwnEvalData(cur) {
			return true, nil
		}
		if _, dup := seen[cur.GetCid()]; dup {
			break
		}
		seen[cur.GetCid()] = struct{}{}
		nextId := cur.GetParentVersionId()
		if nextId == "" {
			break
		}
		cur = byCid[nextId]
	}
	return false, nil
}

func RunEvaluate(ctx context.Context, projectId, versionId string, batchSize int) error {
	existingVersionParams := tensorleapapi.NewEvaluateExistingVersionParams(
		versionId,
		projectId,
		float64(batchSize),
		0,
	)

	evaluateParams := tensorleapapi.EvaluateParams{
		EvaluateExistingVersionParams: existingVersionParams,
	}

	// callEvaluate is captured so it can be re-invoked on the retry path after
	// the user terminates one of their running evaluate jobs.
	callEvaluate := func() (*tensorleapapi.Job, error) {
		job, res, err := api.ApiClient.Evaluate(ctx).EvaluateParams(evaluateParams).Execute()
		if err = api.CheckRes(res, err); err != nil {
			return nil, err
		}
		return job, nil
	}

	log.Info("Starting evaluation...")
	job, err := callEvaluate()
	if err != nil {
		// Special-case the structured concurrent-evaluate-limit error so the
		// user gets a guided recovery prompt instead of a JSON dump.
		var limitErr *api.ConcurrentEvaluateLimitError
		if errors.As(err, &limitErr) {
			job, err = handleConcurrentEvaluateLimit(ctx, limitErr, callEvaluate)
			if err != nil {
				// User-initiated cancel (or Ctrl+C) is a deliberate choice,
				// not a failure: log a friendly note and return nil so the
				// parent command exits cleanly without a scary "Error:" line
				// or a usage dump.
				if errors.Is(err, ErrEvaluateCancelled) {
					log.Warn("Evaluation skipped — cancelled by user.")
					return nil
				}
				return err
			}
		} else {
			return fmt.Errorf("failed to start evaluation: %w", err)
		}
	}

	log.Infof("Evaluation started with job ID: %s", job.GetCid())
	return nil
}

// handleConcurrentEvaluateLimit interactively walks the user through
// terminating one of their running Evaluate/Continue-Evaluate jobs and
// retrying the original evaluate call. Returns the newly-started job on
// success, or an error if the user cancels or termination/retry fails.
//
// Mirrors the web-ui's ConcurrentEvaluateLimitDialog terminate-and-retry
// flow so users get the same recovery experience from either entry point.
func handleConcurrentEvaluateLimit(
	ctx context.Context,
	initial *api.ConcurrentEvaluateLimitError,
	callEvaluate func() (*tensorleapapi.Job, error),
) (*tensorleapapi.Job, error) {
	current := initial
	for {
		log.Warn(current.Error())
		if len(current.RunningJobs) == 0 {
			// Shouldn't happen — backend always sends at least one row when
			// the limit is hit — but guard anyway.
			return nil, errors.New(
				"concurrent evaluate jobs limit reached, but the server " +
					"reported no running jobs to terminate",
			)
		}

		options, byLabel := buildEvaluateJobOptions(current.RunningJobs)
		const cancelLabel = "Cancel — do not start this evaluation"
		options = append(options, cancelLabel)

		var selected string
		prompt := &survey.Select{
			Message: "Select a running Evaluate job to terminate, then this evaluation will be retried:",
			Options: options,
		}
		if err := survey.AskOne(prompt, &selected); err != nil {
			// Treat any prompt failure (most commonly Ctrl+C / InterruptErr)
			// as a deliberate cancellation so RunEvaluate exits cleanly.
			return nil, errors.Join(ErrEvaluateCancelled, err)
		}
		if selected == cancelLabel {
			return nil, ErrEvaluateCancelled
		}

		chosen := byLabel[selected]
		log.Infof("Terminating job %s ...", chosen.JobId)
		if err := terminateJob(ctx, chosen.JobId); err != nil {
			return nil, fmt.Errorf("failed to terminate job %s: %w", chosen.JobId, err)
		}

		log.Info("Retrying evaluation...")
		job, err := callEvaluate()
		if err == nil {
			return job, nil
		}

		// If we're still blocked (e.g., a different job started in the
		// meantime), loop with the fresh running-jobs list. Any other error
		// short-circuits with the original error message.
		var nextLimit *api.ConcurrentEvaluateLimitError
		if errors.As(err, &nextLimit) {
			log.Warn("Still at the concurrent-evaluate-jobs limit after termination.")
			current = nextLimit
			continue
		}
		return nil, fmt.Errorf("failed to start evaluation: %w", err)
	}
}

func buildEvaluateJobOptions(
	jobs []api.RunningEvaluateJobInfo,
) ([]string, map[string]api.RunningEvaluateJobInfo) {
	options := make([]string, 0, len(jobs))
	byLabel := make(map[string]api.RunningEvaluateJobInfo, len(jobs))
	for _, j := range jobs {
		startedLabel := j.StartedAt
		if formatted, err := api.ParseAndFormatDateToLocalTime(j.StartedAt); err == nil {
			startedLabel = formatted
		}
		// Prefer human-readable names that the backend resolves for us;
		// fall back to the raw ObjectIds when the server didn't (or couldn't)
		// provide a name. The disambiguating jobId stays at the end so two
		// identically-named jobs are still distinguishable.
		projectLabel := j.ProjectName
		if projectLabel == "" {
			projectLabel = j.ProjectId
		}
		versionLabel := j.VersionName
		if versionLabel == "" {
			versionLabel = j.VersionId
		}
		label := fmt.Sprintf(
			"%s  ·  %s  ·  started=%s  ·  jobId=%s",
			projectLabel, versionLabel, startedLabel, j.JobId,
		)
		options = append(options, label)
		byLabel[label] = j
	}
	return options, byLabel
}

func terminateJob(ctx context.Context, jobId string) error {
	params := tensorleapapi.NewTerminateJobParams(jobId)
	_, res, err := api.ApiClient.TerminateJob(ctx).TerminateJobParams(*params).Execute()
	return api.CheckRes(res, err)
}

func RunUpdateEvaluateArtifact(ctx context.Context, projectId, versionId string, updateActions []tensorleapapi.UpdateAction) error {
	params := tensorleapapi.NewUpdateEvaluateArtifactParams(versionId, projectId, updateActions)

	log.Info("Starting update evaluate artifact...")
	job, res, err := api.ApiClient.UpdateEvaluateArtifact(ctx).UpdateEvaluateArtifactParams(*params).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to start update evaluate artifact: %w", err)
	}

	log.Infof("Update evaluate artifact started with job ID: %s", job.GetCid())
	return nil
}
