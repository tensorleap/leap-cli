package root_cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/analytics"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/secret"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
	"golang.org/x/term"
)

type pushInputs struct {
	secretId            string
	modelVersionName    string
	codeVersionMessage  string
	modelType           string
	modelPath           string
	branch              string
	pythonVersion       string
	runEval             bool
	transformInput      bool
	noWait              bool
	batch               string
	overwriteVersionRef string
	updateParts         []string
	// noVisualization opts out of the visualize_samples step in the engine
	// when evaluating. Only meaningful alongside --eval / -u, since those
	// are the paths that dispatch an evaluate job.
	noVisualization bool
	// yes acknowledges pre-push warnings (e.g. the latest version having
	// insight-settings overrides) without prompting.
	yes bool
}

type pushState struct {
	ctx              context.Context
	inputs           *pushInputs
	workspaceConfig  *workspace.WorkspaceConfig
	project          *tensorleapapi.Project
	isOverwrite      bool
	overwriteVersion *model.VersionInfo
	properties       map[string]interface{}
}

func NewPushCmd() *cobra.Command {
	in := &pushInputs{}

	cmd := &cobra.Command{
		Use:   "push",
		Short: "Push new or overwrite model version",
		Long: `Push new or overwrite model version into a project with its code integration.

Examples:
  # Push a new version (interactive — prompts for name and model file)
  leap push

  # Push and run evaluation after
  leap push -e

  # Push a new version non-interactively
  leap push -n my-model -m ./model.h5

  # Overwrite an existing version by id, then prompt for what changed
  leap push -o 6a16a0cf -e

  # Overwrite by name — picker opens if the name is ambiguous
  leap push -o my-model -e

  # Overwrite + name what changed (implies --eval; skips the prompt)
  leap push -o my-model -u viz
  leap push -o my-model -u metric_config
  leap push -o my-model -u metric          # triggers a full re-evaluation

  # Refresh multiple artifacts in one run
  leap push -o my-model -u visualization -u metric_config

  # Push + evaluate without running the visualizations step (faster — no dashboard samples)
  leap push -e --novis
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runPush(cmd.Context(), in)
		},
	}

	cmd.Flags().StringVarP(&in.modelVersionName, "name", "n", "", "Model version name")
	cmd.Flags().StringVar(&in.modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&in.branch, "branch", "", "Name of the branch [OPTIONAL]")
	cmd.Flags().StringVar(&in.secretId, "secretId", "", "Secret id")
	cmd.Flags().BoolVar(&in.transformInput, "transform-input", false, "Transpose the input data to channel-last format")
	cmd.Flags().BoolVar(&in.noWait, "no-wait", false, "Do not wait for push to complete")
	cmd.Flags().StringVarP(&in.modelPath, "model-path", "m", "", "Path to the model file")
	cmd.Flags().BoolVarP(&in.runEval, "eval", "e", false, "Run evaluation on the model after push completes")
	cmd.Flags().StringVarP(&in.batch, "batch", "b", "", "Batch size for evaluation: a number or 'latest' (requires --eval; if omitted, prompts with the latest as default)")
	cmd.Flags().StringVarP(&in.overwriteVersionRef, "overwrite", "o", "", "Overwrite an existing version (id, or name — picker shown if name is ambiguous)")
	cmd.Flags().StringVar(&in.overwriteVersionRef, "overwrite-version", "", "")
	_ = cmd.Flags().MarkDeprecated("overwrite-version", "use --overwrite (-o) instead")
	cmd.Flags().StringSliceVarP(&in.updateParts, "update", "u", nil, "What changed in the code on overwrite (repeatable; implies --eval; skips the prompt). Values: metadata, metric, metric_config, visualization (viz). metadata+metric trigger a full re-evaluation.")
	cmd.Flags().BoolVar(&in.noVisualization, "novis", false, "Skip the visualize_samples step on the subsequent evaluate (requires --eval / -u)")
	cmd.Flags().BoolVar(&in.yes, "yes", false, "Acknowledge pre-push warnings and proceed without prompting")
	return cmd
}

func runPush(cmdCtx context.Context, in *pushInputs) error {
	if err := validatePushInputs(in); err != nil {
		return err
	}

	ctx, _, err := auth.RequireAuth(cmdCtx)
	if err != nil {
		return err
	}

	s, err := newPushState(ctx, in)
	if err != nil {
		return err
	}

	analytics.SendEvent(analytics.EventCliProjectsPushStarted, s.properties)

	if err := s.confirmInsightsSettingsOverrides(); err != nil {
		return err
	}

	if err := s.resolveOverwriteTarget(); err != nil {
		return err
	}
	if err := s.collectModelTypeAndName(); err != nil {
		return err
	}
	if err := s.syncBranchSecretAndPython(); err != nil {
		return err
	}

	dispatch, err := s.resolveEvalPlan()
	if err != nil {
		return err
	}

	versionId, codeSnapshotCid, codePushed, err := s.runPushFlow()
	if err != nil {
		return err
	}

	s.sendSuccessEvent(versionId, codeSnapshotCid, codePushed)

	return s.triggerEvaluate(versionId, dispatch)
}

// runPushFlow pushes the version and returns its id and code-snapshot id.
// A new model goes through the combined push job (code parse + model import in
// one). Overwriting a version that already has a model goes through the
// override-push job (re-parse code + re-validate the existing model, no upload).
func (s *pushState) runPushFlow() (versionId, codeSnapshotCid string, codePushed bool, err error) {
	if s.needsNewModel() {
		return s.runCombinedPush()
	}
	return s.runOverridePush()
}

func (s *pushState) runOverridePush() (versionId, codeSnapshotCid string, codePushed bool, err error) {
	in := s.inputs

	closeBundle, tarGzFile, err := code.BundleCodeIntoTempFile(".", s.workspaceConfig)
	if err != nil {
		return "", "", false, s.fail("bundle_code", err)
	}
	defer closeBundle()

	fileStat, err := tarGzFile.Stat()
	if err != nil {
		return "", "", false, s.fail("bundle_code", fmt.Errorf("failed to get file stat: %w", err))
	}

	pushResp, err := code.PushOverride(
		s.ctx, tarGzFile, fileStat.Size(),
		s.workspaceConfig.EntryFile, in.secretId, in.pythonVersion,
		s.projectId(), in.branch, s.overwriteVersion.VersionId,
	)
	if err != nil {
		return "", "", false, s.fail("push_override", err)
	}

	s.properties["code_snapshot_id"] = pushResp.CodeSnapshot.Cid
	s.properties["version_id"] = pushResp.VersionId
	s.properties["job_id"] = pushResp.JobId

	if in.noWait {
		log.Info("Starting push job. JobId: ", pushResp.JobId)
		return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, nil
	}

	if err := model.WaitForPushJob(s.ctx, s.projectId(), pushResp.VersionId, pushResp.JobId); err != nil {
		return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, s.fail("push_job", err)
	}
	return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, nil
}

func (s *pushState) runCombinedPush() (versionId, codeSnapshotCid string, codePushed bool, err error) {
	in := s.inputs

	closeBundle, tarGzFile, err := code.BundleCodeIntoTempFile(".", s.workspaceConfig)
	if err != nil {
		return "", "", false, s.fail("bundle_code", err)
	}
	defer closeBundle()

	modelInfo, err := model.PrepareImportModelFromFilePath(s.ctx, s.projectId(), in.modelPath, in.transformInput, in.modelType)
	if err != nil {
		return "", "", false, s.fail("prepare_model", err)
	}

	fileStat, err := tarGzFile.Stat()
	if err != nil {
		return "", "", false, s.fail("bundle_code", fmt.Errorf("failed to get file stat: %w", err))
	}

	overwriteVersionId := ""
	if s.overwriteVersion != nil {
		overwriteVersionId = s.overwriteVersion.VersionId
	}

	pushResp, err := code.PushCodeAndModel(
		s.ctx, tarGzFile, fileStat.Size(),
		s.workspaceConfig.EntryFile, in.secretId, in.pythonVersion,
		in.modelVersionName, s.projectId(), in.branch, overwriteVersionId,
		*modelInfo,
	)
	if err != nil {
		return "", "", false, s.fail("push", err)
	}

	s.properties["code_snapshot_id"] = pushResp.CodeSnapshot.Cid
	s.properties["version_id"] = pushResp.VersionId
	s.properties["job_id"] = pushResp.JobId

	if in.noWait {
		log.Info("Starting push job. JobId: ", pushResp.JobId)
		return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, nil
	}

	if err := model.WaitForPushJob(s.ctx, s.projectId(), pushResp.VersionId, pushResp.JobId); err != nil {
		return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, s.fail("push_job", err)
	}
	return pushResp.VersionId, pushResp.CodeSnapshot.Cid, true, nil
}

func validatePushInputs(in *pushInputs) error {
	if in.batch != "" && !in.runEval {
		return fmt.Errorf("--batch requires --eval")
	}
	if len(in.updateParts) > 0 && !in.runEval {
		in.runEval = true
	}
	// --novis only makes sense when we're actually going to evaluate.
	// --update implies --eval (set above), so we just check the final
	// resolved runEval state.
	if in.noVisualization && !in.runEval {
		return fmt.Errorf("--novis requires --eval (-e) or --update (-u)")
	}
	return nil
}

func newPushState(ctx context.Context, in *pushInputs) (*pushState, error) {
	properties := map[string]interface{}{
		"secret_id":             in.secretId,
		"model_version_name":    in.modelVersionName,
		"code_version_message":  in.codeVersionMessage,
		"model_type":            in.modelType,
		"branch":                in.branch,
		"transform_input":       in.transformInput,
		"no_wait":               in.noWait,
		"python_version":        in.pythonVersion,
		"model_path":            in.modelPath,
		"overwrite_version_ref": in.overwriteVersionRef,
		"update_parts":          in.updateParts,
		"no_visualization":      in.noVisualization,
	}
	s := &pushState{ctx: ctx, inputs: in, properties: properties}

	workspaceConfig, err := workspace.GetWorkspaceConfig()
	if err != nil {
		return nil, s.fail("get_workspace_config", err)
	}
	s.workspaceConfig = workspaceConfig

	currentProject, err := project.SyncProjectIdToWorkspaceConfig(ctx, workspaceConfig)
	if err != nil {
		return nil, s.fail("sync_project", err)
	}
	s.project = currentProject

	return s, nil
}

func (s *pushState) fail(stage string, err error) error {
	s.properties["error"] = err.Error()
	s.properties["stage"] = stage
	analytics.SendEvent(analytics.EventCliProjectsPushFailed, s.properties)
	return err
}

func (s *pushState) projectId() string {
	return s.project.GetCid()
}

func (s *pushState) resolveOverwriteTarget() error {
	in := s.inputs
	if in.overwriteVersionRef != "" {
		info, err := model.ResolveVersionInfoFromRef(s.ctx, s.projectId(), in.overwriteVersionRef)
		if err != nil {
			return err
		}
		model.PrintResolvedOverwriteTarget(in.overwriteVersionRef, info)
		s.overwriteVersion = info
		s.isOverwrite = info.HasModel || info.HasUploadedModel
		if !s.isOverwrite && in.modelPath == "" {
			return fmt.Errorf("version %q has no model; set --model-path (-m)", in.overwriteVersionRef)
		}
	} else if in.modelPath == "" {
		isOverwrite, info, chosenPath, err := model.AskUserForModelPathOrOverwrite(s.ctx, s.projectId(), &in.modelVersionName, in.runEval)
		if err != nil {
			return err
		}
		s.isOverwrite = isOverwrite
		s.overwriteVersion = info
		in.modelPath = chosenPath
	}

	s.properties["is_overwrite"] = s.isOverwrite
	if s.overwriteVersion != nil {
		s.properties["version_id"] = s.overwriteVersion.VersionId
		s.properties["version_name"] = s.overwriteVersion.VersionName
		s.properties["version_has_model"] = s.overwriteVersion.HasModel
	}
	s.properties["model_path"] = in.modelPath
	return nil
}

// confirmInsightsSettingsOverrides warns when the project's latest version has
// insight-settings overrides — these are per-version and won't carry over to
// the version being pushed — and requires the user to acknowledge before the
// push proceeds (interactively, or via --yes).
func (s *pushState) confirmInsightsSettingsOverrides() error {
	latest, err := model.GetLatestProjectVersion(s.ctx, s.projectId())
	if err != nil {
		return s.fail("check_insights_overrides", err)
	}
	if latest == nil {
		return nil
	}

	overrides, err := model.GetInsightsSettingsOverrides(s.ctx, s.projectId(), latest.Cid)
	if err != nil {
		return s.fail("check_insights_overrides", err)
	}
	if !model.HasAnyInsightsOverride(overrides) {
		return nil
	}

	versionLabel := latest.Notes
	if versionLabel == "" {
		versionLabel = latest.Cid
	}

	log.Warnf("The latest version %q has customized insight settings which will not be applied to the version you're about to push.\nIf you wish to carry over these settings, please apply them to the code integration first and then re-run `leap push`.", versionLabel)
	if s.inputs.yes {
		log.Info("Proceeding — insight-settings override warning acknowledged via --yes.")
		return nil
	}
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return fmt.Errorf("the latest version has insight settings overrides that won't carry over to the new version; re-run with --yes to acknowledge and proceed")
	}

	fmt.Print("\nPress enter to continue with the push, or Ctrl+C to cancel: ")
	if _, err := bufio.NewReader(os.Stdin).ReadString('\n'); err != nil {
		return err
	}
	return nil
}

func (s *pushState) needsNewModel() bool {
	if !s.isOverwrite {
		return true
	}
	return !s.overwriteVersion.HasModel && !s.overwriteVersion.HasUploadedModel
}

func (s *pushState) collectModelTypeAndName() error {
	in := s.inputs
	if s.needsNewModel() {
		if err := model.SelectModelType(&in.modelType, in.modelPath); err != nil {
			return s.fail("select_model_type", err)
		}
	}
	if s.overwriteVersion == nil {
		defaultMessage := model.GetDefaultMessageFromModelPath(in.modelPath)
		if err := model.InitMessage(&in.modelVersionName, defaultMessage); err != nil {
			return s.fail("init_message", err)
		}
	} else {
		in.modelVersionName = s.overwriteVersion.VersionName
	}
	return nil
}

func (s *pushState) syncBranchSecretAndPython() error {
	in := s.inputs
	branch, err := code.SyncBranchFromFlagAndConfig(in.branch, s.workspaceConfig)
	if err != nil {
		return s.fail("sync_branch", err)
	}
	in.branch = branch

	secretId, err := secret.SyncSecretIdFromFlagAndConfig(s.ctx, in.secretId, s.workspaceConfig)
	if err != nil {
		return s.fail("sync_secret", err)
	}
	in.secretId = secretId

	pythonVersion, err := code.SyncPythonVersionFromFlagAndConfig(s.ctx, in.pythonVersion, s.workspaceConfig)
	if err != nil {
		return s.fail("sync_python_version", err)
	}
	in.pythonVersion = pythonVersion
	return nil
}

type evalDispatch struct {
	batchSize         int
	updateActions     []tensorleapapi.UpdateAction
	runUpdateEvaluate bool
	persistOnly       bool
	// noVisualization mirrors the CLI --novis flag; threaded through to
	// the evaluate API call so the engine marks visualize_samples SKIPPED.
	// Only consulted on the RunEvaluate path — update-evaluate-artifact
	// has no visualization step to skip.
	noVisualization bool
}

func (s *pushState) resolveEvalPlan() (evalDispatch, error) {
	in := s.inputs
	if len(in.updateParts) > 0 && !s.isOverwrite {
		return evalDispatch{}, fmt.Errorf("--update (-u) only applies when overwriting an existing version (use --overwrite or choose overwrite in the prompt)")
	}

	if !s.isOverwrite {
		// New version → a full evaluate. Offer to run it (default Yes).
		if !in.runEval {
			run, err := s.promptRunEvaluate(false)
			if err != nil {
				return evalDispatch{}, err
			}
			if !run {
				return evalDispatch{}, nil
			}
			in.runEval = true
		}
		batchSize, err := s.askOrDefaultBatchSize()
		return evalDispatch{batchSize: batchSize, noVisualization: in.noVisualization}, err
	}

	// Overwrite → ask what changed first; that decides whether the run would be
	// an update-evaluate or a full re-evaluate. Then offer to run it (default
	// Yes), labelling the prompt accordingly.
	plan, err := s.resolveOverwriteEvalPlan()
	if err != nil {
		return evalDispatch{}, err
	}

	if !in.runEval {
		run, err := s.promptRunEvaluate(plan.Kind == model.EvaluatePlanUpdate)
		if err != nil {
			return evalDispatch{}, err
		}
		if !run {
			return evalDispatch{updateActions: plan.UpdateActions, persistOnly: true}, nil
		}
		in.runEval = true
	}

	if plan.Kind == model.EvaluatePlanReset {
		batchSize, err := s.askOrDefaultBatchSize()
		return evalDispatch{
			batchSize:       batchSize,
			updateActions:   plan.UpdateActions,
			noVisualization: in.noVisualization,
		}, err
	}
	return evalDispatch{updateActions: plan.UpdateActions, runUpdateEvaluate: true}, nil
}

// promptRunEvaluate offers to run evaluation after the push (default Yes) when
// the user didn't pass --eval (or -u, which implies it). isUpdateEvaluate picks
// the label ("update-evaluate" vs "evaluate"). Returns false without prompting
// when stdin isn't a terminal, so scripted/piped pushes keep the explicit
// --eval requirement.
func (s *pushState) promptRunEvaluate(isUpdateEvaluate bool) (bool, error) {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return false, nil
	}
	return model.AskRunEvaluate(isUpdateEvaluate)
}

func (s *pushState) resolveOverwriteEvalPlan() (model.EvaluatePlan, error) {
	plan, err := s.collectUserUpdatePlan()
	if err != nil {
		return model.EvaluatePlan{}, err
	}
	if plan.Kind == model.EvaluatePlanUpdate && !s.canUpdateEvaluate() {
		log.Info("No evaluation data found in the override chain — running a fresh evaluate.")
		return model.EvaluatePlan{Kind: model.EvaluatePlanReset, UpdateActions: plan.UpdateActions}, nil
	}
	return plan, nil
}

func (s *pushState) collectUserUpdatePlan() (model.EvaluatePlan, error) {
	if len(s.inputs.updateParts) > 0 {
		parsed, err := model.ParseUpdateActionsFromFlags(s.inputs.updateParts)
		if err != nil {
			return model.EvaluatePlan{}, err
		}
		return model.PlanFromUpdateActions(parsed), nil
	}
	plan, err := model.AskForEvaluatePlan()
	if err != nil {
		return model.EvaluatePlan{}, fmt.Errorf("failed to get update plan: %w", err)
	}
	return plan, nil
}

func (s *pushState) canUpdateEvaluate() bool {
	hasEvalData, err := model.HasEvaluatedAncestorOrSelf(s.ctx, s.projectId(), s.overwriteVersion.VersionId)
	if err != nil {
		log.Warnf("failed to check evaluation data for override chain: %v", err)
		return true
	}
	return hasEvalData
}

func (s *pushState) askOrDefaultBatchSize() (int, error) {
	raw := strings.TrimSpace(s.inputs.batch)
	if raw != "" && !strings.EqualFold(raw, "latest") {
		n, err := strconv.Atoi(raw)
		if err != nil || n <= 0 {
			return 0, fmt.Errorf("--batch expects a positive integer or 'latest', got %q", s.inputs.batch)
		}
		return n, nil
	}
	latestBatchSize, err := model.GetLatestEvaluateBatchSize(s.ctx, s.projectId())
	if err != nil {
		log.Warnf("failed to get latest evaluate batch size: %v", err)
	}
	if strings.EqualFold(raw, "latest") {
		if latestBatchSize <= 0 {
			return 0, fmt.Errorf("--batch latest: no previous evaluation found in this project; pass an explicit batch size")
		}
		return latestBatchSize, nil
	}
	chosen, err := model.AskForBatchSize(latestBatchSize)
	if err != nil {
		return 0, fmt.Errorf("failed to get batch size: %w", err)
	}
	return chosen, nil
}

func (s *pushState) sendSuccessEvent(versionId, codeSnapshotCid string, codePushed bool) {
	in := s.inputs
	s.properties["code_snapshot_id"] = codeSnapshotCid
	s.properties["version_id"] = versionId
	s.properties["project_id"] = s.projectId()
	s.properties["final_secret_id"] = in.secretId
	s.properties["final_python_version"] = in.pythonVersion
	s.properties["final_model_type"] = in.modelType
	s.properties["final_transform_input"] = in.transformInput
	s.properties["final_wait"] = !in.noWait
	s.properties["code_pushed"] = codePushed
	analytics.SendEvent(analytics.EventCliProjectsPushSuccess, s.properties)
}

func (s *pushState) triggerEvaluate(versionId string, dispatch evalDispatch) error {
	if dispatch.persistOnly {
		if len(dispatch.updateActions) == 0 {
			return nil
		}
		return model.PersistUpdateActions(s.ctx, s.projectId(), versionId, dispatch.updateActions)
	}
	if !s.inputs.runEval {
		return nil
	}
	if dispatch.runUpdateEvaluate {
		if err := model.RunUpdateEvaluateArtifact(s.ctx, s.projectId(), versionId, dispatch.updateActions); err != nil {
			return fmt.Errorf("failed to run update evaluate: %w", err)
		}
		return nil
	}
	if err := model.RunEvaluate(s.ctx, s.projectId(), versionId, dispatch.batchSize, dispatch.noVisualization); err != nil {
		return fmt.Errorf("failed to run evaluation: %w", err)
	}
	return nil
}
