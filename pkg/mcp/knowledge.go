package mcp

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const glossaryURI = "tensorleap://glossary"

const glossary = `# Tensorleap glossary

Tensorleap evaluates a trained model on a dataset and analyses where and why it fails. It records per-sample
metrics and metadata, and builds latent spaces from the model's internal activations; insights are groups of
samples found in those latent spaces.

## Insight types (names as shown in the Insights panel)
- Failure Mode (low_performance): a group of samples where the model underperforms. Its sample list also holds
  healthy latent neighbours; the group is the root members only (groupSize). Never quote clusterSize /
  n_samples as the number of failing samples. Check contrast before calling it a failure: a group whose loss or
  error is within ~1.2x of all data is not a real failure mode, whatever the insight's title says.
- Out of Distribution: samples in one subset unlike anything in the rest of the data.
- Duplication: near-identical samples within a subset.
- Data Leakage: near-identical samples on both sides of a split; evaluation metrics may be optimistic.
- Domain Gap: performance differs between two values of a metadata field.
- Mislabeled: samples whose ground truth looks wrong.
Sub-insights refine a parent insight (parentIndex). Statuses: InReview, Approved, Archived.

## Reading an insight
- composition: metadata values over-represented in the group compared with all data. Over-represented is not
  the same as defining; a value at 30% of the group can still be the minority.
- contrast: the group's metric median/mean against all data.
- remedy: the platform's selection of samples to label or acquire for a Failure Mode.
- topSamples: the most representative failing samples (by affinity score, else by loss), not a random draw.

## The engine object on each insight (the platform's own analysis)
- metrics_info: per metric, the cluster's median/average against the data outside the cluster.
- severity_metrics: the metrics that made the insight severe, with their values.
- mutual_info_elements: the features that separate the group from the rest, with the value inside and outside
  the cluster and a score; direction says whether the feature is higher ("up") or lower in the group.
- is_train_aggressor (Failure Mode): true when the failing group is dominated by training samples. Then more
  data will not help: review the top samples' labels, then look at model-side fixes (targeted loss term,
  loss weighting or oversampling of these samples).
- overfitting_metrics / overfitting_evidence (Failure Mode): the metrics on which the cluster's training samples
  do far better than its failing samples. overfitting_evidence has one row per flagged metric: score (a
  MAD-normalised gap between the two medians, weighted by support), the threshold it exceeded (2.0),
  cluster_median (the failing group), extended_train_median (training samples in the wider cluster), n_cluster
  and n_extended_train (samples behind each median), support (how train-heavy the wider cluster is relative
  to the dataset) and direction. When flagged, recommend rebalancing the dataset, e.g. moving samples from
  test to train for this population.
- aggressor_fixing (Failure Mode): num_of_samples_to_label chosen by similarity search near the cluster and
  num_of_samples_to_acquire; when present it is the platform's suggestion and must be surfaced with its
  numbers (tl_export_analysis writes the selected list).
- cluster_extended_stats: per metric, the median over the failing group (cluster_value) next to the median
  over the wider cluster, the group plus its latent neighbourhood (extended_value), with the metric's
  direction so you can tell which side is worse.
- direction (in the fields above): the metric's preferred direction as the integration declared it:
  Downward means lower is better (a loss), Upward means higher is better (an accuracy).
- automatic_tests: a regression test the platform suggests (metric, threshold, operator); create it through
  createTestLink.
- Type-specific: subset (Out of Distribution, Duplication, Mislabeled), first_subset/second_subset (Data
  Leakage), metadata_name/domain_a/domain_b/domain_gap_score (Domain Gap).

## Latent spaces (what "similar" means for a group)
- classification-semantic: similar in the features that drive the model's class decision; clusters here are
  about how the model reasons (confusions, label boundaries).
- image-non-semantic: visually similar (low-level appearance) regardless of class; clusters here are about
  appearance (corruptions, lighting, capture conditions).
- foreground: similar main subject, background discounted.
- balanced: a mix of semantic and visual similarity.
Other names are project-specific representations.

## Splits
training, validation, test, unlabeled, additional (field dataset_state.keyword).

## What this server cannot do
It does not change anything in Tensorleap. Creating tests, approving or archiving insights and running
evaluations happen through the links it returns or through the leap CLI (leap push --eval, or leap push -o <version> --eval to re-evaluate an existing version;
leap run list/logs only inspect jobs).
It does not read model weights. Integration code is available through tl_get_integration_code when the
project allows it.`

func registerKnowledge(srv *sdk.Server) {
	srv.AddResource(&sdk.Resource{URI: glossaryURI, Name: "glossary", Title: "Tensorleap glossary",
		Description: "Current insight names, how to read insights, latent-space meanings and what this server cannot do.",
		MIMEType:    "text/markdown"},
		func(context.Context, *sdk.ReadResourceRequest) (*sdk.ReadResourceResult, error) {
			return &sdk.ReadResourceResult{Contents: []*sdk.ResourceContents{{URI: glossaryURI, MIMEType: "text/markdown", Text: glossary}}}, nil
		})

	project := &sdk.PromptArgument{Name: "project", Description: "project name or id (optional; asked for if missing)"}
	version := &sdk.PromptArgument{Name: "version", Description: "model version name or id (optional; asked for if missing)"}
	for _, p := range prompts {
		p := p
		srv.AddPrompt(&sdk.Prompt{Name: p.name, Title: p.title, Description: p.description, Arguments: append([]*sdk.PromptArgument{project, version}, p.args...)},
			func(_ context.Context, req *sdk.GetPromptRequest) (*sdk.GetPromptResult, error) {
				return &sdk.GetPromptResult{Description: p.description, Messages: []*sdk.PromptMessage{{
					Role: "user", Content: &sdk.TextContent{Text: renderPrompt(p.body, req.Params.Arguments)}}}}, nil
			})
	}
}

type promptDef struct {
	name, title, description, body string
	args                           []*sdk.PromptArgument
}

func renderPrompt(body string, args map[string]string) string {
	target := []string{}
	for _, k := range []string{"project", "version", "baseline", "budget", "job"} {
		if v := strings.TrimSpace(args[k]); v != "" {
			target = append(target, fmt.Sprintf("%s: %s", k, v))
		}
	}
	head := "Use the Tensorleap MCP tools. Start with tl_status."
	if len(target) > 0 {
		head += " Target — " + strings.Join(target, ", ") + "."
	} else {
		head += " If the project or version is unclear, list them and ask me which one."
	}
	return head + "\n\n" + body
}

var prompts = []promptDef{
	{name: "analyze_version", title: "Analyze a model version",
		description: "Explain where and why an evaluated model version fails, with evidence and concrete next steps.",
		body: `1. Pick the version with tl_list_versions (it must be evaluated) and call tl_get_insights. Call
   tl_get_integration_code (entry file) to learn what each visualizer renders and how metadata is derived.
2. For each insight, worst first: say what fails in plain ML terms, how big the failing group is (groupSize), how it
   splits across training/validation, which metadata is over-represented (composition, with the all-data share
   beside it) and how its metrics compare with all data (contrast). Name the latent space and what "similar" means there.
3. Look at the evidence yourself: call tl_view_samples on 3-6 rendered topSamples (for more, or when the user wants
   files, use tl_export_analysis) and say what they visibly share,
   including patterns no metadata field captures (suggest such a field). Treat what you see as the most
   representative samples, not a random draw.
4. Use tl_query to confirm or refute a hypothesis over the whole population (e.g. the metric by the over-represented field).
5. End with 1-3 actions per insight an ML engineer can do this week, citing the platform's remedy counts when present,
   and the insight's link (and createTestLink when offered).
Keep platform evidence and your own observations clearly apart.`},
	{name: "compare_versions", title: "Compare two model versions",
		description: "Where a candidate version is better or worse than a baseline, by slice.",
		args:        []*sdk.PromptArgument{{Name: "baseline", Description: "baseline version name or id"}},
		body: `1. Identify the candidate (version) and the baseline; both must be evaluated (tl_list_versions).
2. Use tl_describe_fields, then tl_query with both versionIds: the main metrics overall, then by split and by the
   metadata fields that matter (per class, per site, per condition). Report n for every slice.
3. Call tl_get_insights for both versions and match failure modes by what they describe (composition and latent space):
   which persist, which are new in the candidate, which are gone.
4. Conclude: is the candidate better everywhere, or does it regress on a slice even if the overall metric improved?`},
	{name: "what_to_label_next", title: "What to label next",
		description: "Which data to label or collect to fix the model's failure modes.",
		args:        []*sdk.PromptArgument{{Name: "budget", Description: "labeling budget in samples (optional)"}},
		body: `1. Call tl_get_insights for the version. For each Failure Mode: is the group mostly training data (a model-side
   problem more data will not fix) or not (a data gap)? Use the split and contrast.
2. Where the platform offers a remedy, quote its numbers (samples to label / to acquire). To hand the list over,
   call tl_export_analysis: it writes fixing_samples.csv per Failure Mode; the insight link opens the same selection in the UI.
3. Describe the data to collect in domain terms from the over-represented metadata and what tl_view_samples shows.
4. If a budget is given, split it across insights by severity and group size, and say why.`},
	{name: "debug_failed_job", title: "Debug a failed job",
		description: "Why a Push or Evaluate failed and how to fix it.",
		args:        []*sdk.PromptArgument{{Name: "job", Description: "job id (optional; the latest failed job is used)"}},
		body: `1. Find the job with tl_list_jobs (status failed) unless a job id is given. Do not start the job again yet.
2. Call tl_get_job_logs and read errorLines first, then the log tails. Find the earliest real error.
3. Classify it: customer integration code (preprocess / encoders / metrics), dependencies, data not reachable from the
   server (data volume, credentials), out of memory, or a platform problem.
4. Give the fix. After the user fixes it, the job is re-run with the leap CLI (e.g. leap push -o <version> --eval);
   then use tl_wait_for_job instead of polling.`},
	{name: "summarize_for_stakeholders", title: "Summary for stakeholders",
		description: "A short, non-technical summary of a version's findings to forward to a manager.",
		body: `1. Call tl_get_insights and tl_query for the headline metrics by split.
2. Write at most 10 lines: overall quality, the 2-3 most important problems in plain language with their size,
   what the team will do next, and what decision (if any) is needed. No internal ids, field names or jargon.
3. Add the insights panel link for readers who want detail.`},
}
