package root_cmd

import (
	"context"
	"fmt"

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
	batchSize           int
	overwriteVersionRef string
	updateParts         []string
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

  # Overwrite an existing version by id, run eval (auto-detects changes)
  leap push -o 6a16a0cf -e

  # Overwrite by name — picker opens if the name is ambiguous
  leap push -o my-model -e

  # Overwrite + force a specific artifact refresh (implies --eval)
  leap push -o my-model -u viz

  # Refresh multiple artifacts in one run
  leap push -o my-model -u metadata -u insights
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
	cmd.Flags().IntVar(&in.batchSize, "batch", 0, "Batch size for evaluation (only valid with --eval)")
	cmd.Flags().StringVarP(&in.overwriteVersionRef, "overwrite", "o", "", "Overwrite an existing version (id, or name — picker shown if name is ambiguous)")
	cmd.Flags().StringVar(&in.overwriteVersionRef, "overwrite-version", "", "")
	_ = cmd.Flags().MarkDeprecated("overwrite-version", "use --overwrite (-o) instead")
	cmd.Flags().StringSliceVarP(&in.updateParts, "update", "u", nil, "Artifact(s) to refresh on overwrite (repeatable; implies --eval). Values: metadata, metric, insights, visualization (viz)")
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

	if err := s.resolveOverwriteTarget(); err != nil {
		return err
	}
	if err := s.collectModelTypeAndName(); err != nil {
		return err
	}
	if err := s.syncBranchSecretAndPython(); err != nil {
		return err
	}

	evalBatchSize, updateActions, runUpdateEvaluate, err := s.resolveEvalPlan()
	if err != nil {
		return err
	}

	codeResp, codePushed, err := s.pushCodeAndWaitForParsing()
	if err != nil {
		return err
	}

	if err := s.applyModelToVersion(codeResp.VersionId); err != nil {
		return err
	}

	s.sendSuccessEvent(codeResp, codePushed)

	return s.triggerEvaluate(codeResp.VersionId, evalBatchSize, updateActions, runUpdateEvaluate)
}

func validatePushInputs(in *pushInputs) error {
	if in.batchSize > 0 && !in.runEval {
		return fmt.Errorf("--batch requires --eval")
	}
	if len(in.updateParts) > 0 && !in.runEval {
		in.runEval = true
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

func (s *pushState) resolveEvalPlan() (batchSize int, updateActions []tensorleapapi.UpdateAction, runUpdateEvaluate bool, err error) {
	in := s.inputs
	if len(in.updateParts) > 0 && !s.isOverwrite {
		err = fmt.Errorf("--update (-u) only applies when overwriting an existing version (use --overwrite or choose overwrite in the prompt)")
		return
	}
	if !in.runEval {
		return
	}

	if !s.isOverwrite {
		batchSize, err = s.askOrDefaultBatchSize()
		return
	}

	parsedActions, err := model.ParseUpdateActionsFromFlags(in.updateParts)
	if err != nil {
		return
	}
	if len(parsedActions) > 0 {
		plan := model.PlanFromUpdateActions(parsedActions)
		if plan.Kind == model.EvaluatePlanReset {
			batchSize, err = s.askOrDefaultBatchSize()
			return
		}
		updateActions = plan.UpdateActions
		runUpdateEvaluate = true
		return
	}

	hasEvalData, evalErr := model.HasEvaluatedAncestorOrSelf(s.ctx, s.projectId(), s.overwriteVersion.VersionId)
	if evalErr != nil {
		log.Warnf("failed to check evaluation data for override chain: %v", evalErr)
		hasEvalData = true
	}
	if !hasEvalData {
		log.Info("No evaluation data found in the override chain — running a fresh evaluate.")
		batchSize, err = s.askOrDefaultBatchSize()
		return
	}

	plan, planErr := model.AskForEvaluatePlan()
	if planErr != nil {
		err = fmt.Errorf("failed to get update plan: %w", planErr)
		return
	}
	if plan.Kind == model.EvaluatePlanReset {
		batchSize, err = s.askOrDefaultBatchSize()
		return
	}
	updateActions = plan.UpdateActions
	runUpdateEvaluate = true
	return
}

func (s *pushState) askOrDefaultBatchSize() (int, error) {
	if s.inputs.batchSize > 0 {
		return s.inputs.batchSize, nil
	}
	defaultBatchSize, err := model.GetLatestEvaluateBatchSize(s.ctx, s.projectId())
	if err != nil {
		log.Warnf("failed to get latest evaluate batch size: %v", err)
	}
	chosen, err := model.AskForBatchSize(defaultBatchSize)
	if err != nil {
		return 0, fmt.Errorf("failed to get batch size: %w", err)
	}
	return chosen, nil
}

func (s *pushState) pushCodeAndWaitForParsing() (*tensorleapapi.PushCodeSnapshotResponse, bool, error) {
	in := s.inputs
	closeBundle, tarGzFile, err := code.BundleCodeIntoTempFile(".", s.workspaceConfig)
	if err != nil {
		return nil, false, s.fail("bundle_code", err)
	}
	defer closeBundle()

	overwriteVersionId := ""
	if s.overwriteVersion != nil {
		overwriteVersionId = s.overwriteVersion.VersionId
	}
	pushed, codeResp, err := code.PushCode(s.ctx, tarGzFile, s.workspaceConfig.EntryFile, in.secretId, in.pythonVersion, in.modelVersionName, s.projectId(), in.branch, overwriteVersionId)
	if err != nil {
		s.tagCodeResp(codeResp)
		return codeResp, false, s.fail("push_code", err)
	}

	if pushed || !code.IsCodeEnded(&codeResp.CodeSnapshot) {
		ok, waitErr := code.WaitForCodeIntegrationStatus(s.ctx, s.projectId(), codeResp.CodeSnapshot.Cid)
		if waitErr != nil {
			s.tagCodeResp(codeResp)
			return codeResp, pushed, s.fail("wait_for_code_parsing", waitErr)
		}
		if !ok {
			s.tagCodeResp(codeResp)
			return codeResp, pushed, s.fail("code_parsing", fmt.Errorf("code parsing failed"))
		}
		log.Info("Code parsed successfully")
	} else if code.IsCodeParseFailed(&codeResp.CodeSnapshot) {
		s.tagCodeResp(codeResp)
		return codeResp, pushed, s.fail("previous_code_parsing_failed", fmt.Errorf("latest code parsing failed, add --force to push anyway"))
	}
	return codeResp, pushed, nil
}

func (s *pushState) tagCodeResp(codeResp *tensorleapapi.PushCodeSnapshotResponse) {
	if codeResp == nil {
		return
	}
	s.properties["code_snapshot_id"] = codeResp.CodeSnapshot.Cid
	s.properties["version_id"] = codeResp.VersionId
}

func (s *pushState) applyModelToVersion(versionId string) error {
	in := s.inputs
	if !s.isOverwrite {
		importModelInfo, err := model.PrepareImportModelFromFilePath(s.ctx, s.projectId(), in.modelPath, in.transformInput, in.modelType)
		if err != nil {
			return err
		}
		if _, err = model.ImportModel(s.ctx, s.projectId(), versionId, importModelInfo, !in.noWait); err != nil {
			s.properties["code_snapshot_id"] = versionId
			s.properties["version_id"] = versionId
			return s.fail("import_model", err)
		}
		return nil
	}

	var importModelInfo *tensorleapapi.ImportModelInfo
	if !s.overwriteVersion.HasModel && !s.overwriteVersion.HasUploadedModel {
		var err error
		importModelInfo, err = model.PrepareImportModelFromFilePath(s.ctx, s.projectId(), in.modelPath, in.transformInput, in.modelType)
		if err != nil {
			return err
		}
	}
	if _, err := model.OverrideModel(s.ctx, s.projectId(), versionId, !in.noWait, importModelInfo); err != nil {
		s.properties["version_id"] = versionId
		return s.fail("override_model", err)
	}
	return nil
}

func (s *pushState) sendSuccessEvent(codeResp *tensorleapapi.PushCodeSnapshotResponse, codePushed bool) {
	in := s.inputs
	s.properties["code_snapshot_id"] = codeResp.CodeSnapshot.Cid
	s.properties["version_id"] = codeResp.VersionId
	s.properties["project_id"] = s.projectId()
	s.properties["final_secret_id"] = in.secretId
	s.properties["final_python_version"] = in.pythonVersion
	s.properties["final_model_type"] = in.modelType
	s.properties["final_transform_input"] = in.transformInput
	s.properties["final_wait"] = !in.noWait
	s.properties["code_pushed"] = codePushed
	analytics.SendEvent(analytics.EventCliProjectsPushSuccess, s.properties)
}

func (s *pushState) triggerEvaluate(versionId string, batchSize int, updateActions []tensorleapapi.UpdateAction, runUpdateEvaluate bool) error {
	if !s.inputs.runEval {
		return nil
	}
	if runUpdateEvaluate {
		if err := model.RunUpdateEvaluateArtifact(s.ctx, s.projectId(), versionId, updateActions); err != nil {
			return fmt.Errorf("failed to run update evaluate: %w", err)
		}
		return nil
	}
	if err := model.RunEvaluate(s.ctx, s.projectId(), versionId, batchSize); err != nil {
		return fmt.Errorf("failed to run evaluation: %w", err)
	}
	return nil
}
