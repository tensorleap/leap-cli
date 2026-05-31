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

var updateActionAliases = map[string]tensorleapapi.UpdateAction{
	"metadata":         tensorleapapi.UPDATEACTION_UPDATE_METADATA,
	"metric":           tensorleapapi.UPDATEACTION_UPDATE_METRIC,
	"metric_direction": tensorleapapi.UPDATEACTION_UPDATE_METRIC_DIRECTION,
	"metric-direction": tensorleapapi.UPDATEACTION_UPDATE_METRIC_DIRECTION,
	"direction":        tensorleapapi.UPDATEACTION_UPDATE_METRIC_DIRECTION,
	"insights":         tensorleapapi.UPDATEACTION_UPDATE_INSIGHTS,
	"visualization":    tensorleapapi.UPDATEACTION_UPDATE_VISUALIZATION,
	"viz":              tensorleapapi.UPDATEACTION_UPDATE_VISUALIZATION,
}

const updateActionAllowedHint = "metadata, metric, metric_direction, insights, visualization, viz"

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
		var act tensorleapapi.UpdateAction
		if alias, ok := updateActionAliases[p]; ok {
			act = alias
		} else if full, ferr := tensorleapapi.NewUpdateActionFromValue(p); ferr == nil {
			act = *full
		} else {
			return nil, fmt.Errorf("invalid --update value %q (allowed: %s)", raw, updateActionAllowedHint)
		}
		if _, ok := seen[act]; ok {
			continue
		}
		seen[act] = struct{}{}
		out = append(out, act)
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no valid --update values after parsing")
	}
	return out, nil
}

// EvaluatePlanKind selects between Update and Reset.
type EvaluatePlanKind int

const (
	EvaluatePlanUpdate EvaluatePlanKind = iota
	EvaluatePlanReset
)

// EvaluatePlan is the resolved set of update actions.
type EvaluatePlan struct {
	Kind          EvaluatePlanKind
	UpdateActions []tensorleapapi.UpdateAction
}

type ChangeKey int

const (
	ChangeMetadata ChangeKey = iota
	ChangeMetric
	ChangeMetricDirection
	ChangeVisualization
	ChangeInsights
)

type changeOption struct {
	key    ChangeKey
	label  string
	hint   string
	action tensorleapapi.UpdateAction
}

var changeOptions = []changeOption{
	{
		key:    ChangeMetadata,
		label:  "Added or edited a metadata",
		hint:   "triggers full re-evaluation",
		action: tensorleapapi.UPDATEACTION_UPDATE_METADATA,
	},
	{
		key:    ChangeMetric,
		label:  "Added or edited a metric",
		hint:   "triggers full re-evaluation",
		action: tensorleapapi.UPDATEACTION_UPDATE_METRIC,
	},
	{
		key:    ChangeMetricDirection,
		label:  "Edited metric direction or insight-config",
		hint:   "cheap update — keeps existing evaluation",
		action: tensorleapapi.UPDATEACTION_UPDATE_METRIC_DIRECTION,
	},
	{
		key:    ChangeVisualization,
		label:  "Added or edited a visualization",
		hint:   "regenerate all visualizations",
		action: tensorleapapi.UPDATEACTION_UPDATE_VISUALIZATION,
	},
	{
		key:    ChangeInsights,
		label:  "Reinforce insights",
		hint:   "regenerate insights from scratch",
		action: tensorleapapi.UPDATEACTION_UPDATE_INSIGHTS,
	},
}

func triggersFullReeval(a tensorleapapi.UpdateAction) bool {
	return a == tensorleapapi.UPDATEACTION_UPDATE_METADATA ||
		a == tensorleapapi.UPDATEACTION_UPDATE_METRIC
}

func PlanFromUpdateActions(actions []tensorleapapi.UpdateAction) EvaluatePlan {
	for _, a := range actions {
		if triggersFullReeval(a) {
			return EvaluatePlan{Kind: EvaluatePlanReset, UpdateActions: actions}
		}
	}
	return EvaluatePlan{Kind: EvaluatePlanUpdate, UpdateActions: actions}
}

func planUpdateEvaluate(selected map[ChangeKey]bool) EvaluatePlan {
	actions := make([]tensorleapapi.UpdateAction, 0, len(changeOptions))
	for _, opt := range changeOptions {
		if selected[opt.key] {
			actions = append(actions, opt.action)
		}
	}
	return PlanFromUpdateActions(actions)
}

func init() {
	// Push the inline `[Use arrows … type to filter]` hint onto its own
	// line below the question; the upstream template glues it to the
	// message with two leading spaces, crowding longer prompts.
	survey.MultiSelectQuestionTemplate = `
{{- define "option"}}
    {{- if eq .SelectedIndex .CurrentIndex }}{{color .Config.Icons.SelectFocus.Format }}{{ .Config.Icons.SelectFocus.Text }}{{color "reset"}}{{else}} {{end}}
    {{- if index .Checked .CurrentOpt.Index }}{{color .Config.Icons.MarkedOption.Format }} {{ .Config.Icons.MarkedOption.Text }} {{else}}{{color .Config.Icons.UnmarkedOption.Format }} {{ .Config.Icons.UnmarkedOption.Text }} {{end}}
    {{- color "reset"}}
    {{- " "}}{{- .CurrentOpt.Value}}{{ if ne ($.GetDescription .CurrentOpt) "" }} - {{color "cyan"}}{{ $.GetDescription .CurrentOpt }}{{color "reset"}}{{end}}
{{end}}
{{- if .ShowHelp }}{{- color .Config.Icons.Help.Format }}{{ .Config.Icons.Help.Text }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color .Config.Icons.Question.Format }}{{ .Config.Icons.Question.Text }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else }}
	{{ "\n  " }}{{- color "cyan"}}[Use arrows to move, {{color "cyan+b"}}space{{color "cyan"}} to select,{{- if not .Config.RemoveSelectAll }} <right> to all,{{end}}{{- if not .Config.RemoveSelectNone }} <left> to none,{{end}} type to filter{{- if and .Help (not .ShowHelp)}}, {{ .Config.HelpInput }} for more help{{end}}]{{color "reset"}}
  {{- "\n"}}
  {{- range $ix, $option := .PageEntries}}
    {{- template "option" $.IterateOption $ix $option}}
  {{- end}}
{{- end}}`
}

func AskForEvaluatePlan() (EvaluatePlan, error) {
	labels := make([]string, len(changeOptions))
	labelToKey := make(map[string]ChangeKey, len(changeOptions))
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
		Message: "What did you change in your code?",
		Options: labels,
	}
	if err := survey.AskOne(
		prompt, &selectedLabels,
		survey.WithRemoveSelectAll(),
		survey.WithRemoveSelectNone(),
		survey.WithValidator(survey.Required),
	); err != nil {
		return EvaluatePlan{}, err
	}

	selected := make(map[ChangeKey]bool, len(selectedLabels))
	for _, l := range selectedLabels {
		selected[labelToKey[l]] = true
	}
	return planUpdateEvaluate(selected), nil
}

func FormatEvaluatePlan(plan EvaluatePlan) []string {
	if plan.Kind == EvaluatePlanReset {
		return []string{"Re-evaluate (full)"}
	}
	out := make([]string, 0, len(plan.UpdateActions))
	for _, a := range plan.UpdateActions {
		switch a {
		case tensorleapapi.UPDATEACTION_UPDATE_METRIC_DIRECTION:
			out = append(out, "Update metric direction")
		case tensorleapapi.UPDATEACTION_UPDATE_INSIGHTS:
			out = append(out, "Regenerate insights")
		case tensorleapapi.UPDATEACTION_UPDATE_VISUALIZATION:
			out = append(out, "Regenerate visualizations")
		default:
			out = append(out, string(a))
		}
	}
	return out
}

func PrintEvaluatePlan(plan EvaluatePlan) {
	log.Info("This will:")
	for _, item := range FormatEvaluatePlan(plan) {
		log.Infof("  • %s", item)
	}
}

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

// HasEvaluatedAncestorOrSelf reports whether versionId or any version in its override chain has evaluation data.
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

	log.Infof("Starting evaluation (batch size %d)...", batchSize)
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

	PrintEvaluatePlan(EvaluatePlan{
		Kind:          EvaluatePlanUpdate,
		UpdateActions: updateActions,
	})

	log.Info("Starting update evaluate...")
	job, res, err := api.ApiClient.UpdateEvaluateArtifact(ctx).UpdateEvaluateArtifactParams(*params).Execute()
	if err = api.CheckRes(res, err); err != nil {
		return fmt.Errorf("failed to start update evaluate: %w", err)
	}

	log.Infof("Update evaluate started with job ID: %s", job.GetCid())
	return nil
}
