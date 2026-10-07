package mcp

import (
	"context"
	"sort"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type VersionIn struct {
	ProjectID string `json:"projectId" jsonschema:"project id or name from tl_list_projects"`
	VersionID string `json:"versionId" jsonschema:"evaluated version id or name from tl_list_versions, or \"latest\""`
}

type LatentSpace struct {
	Name  string `json:"name"`
	Means string `json:"means"`
}

type Remedy struct {
	SamplesToLabel   int `json:"samplesToLabel"`
	SamplesToAcquire int `json:"samplesToAcquire"`
}

type Insight struct {
	Index        int            `json:"index"`
	Type         string         `json:"type"`
	Name         string         `json:"name"`
	Severity     int            `json:"severity"`
	Status       string         `json:"status"`
	Description  string         `json:"description,omitempty"`
	ParentIndex  int            `json:"parentIndex,omitempty"`
	LatentSpace  *LatentSpace   `json:"latentSpace,omitempty"`
	GroupSize    *int           `json:"groupSize,omitempty"`
	GroupMeaning string         `json:"groupMeaning"`
	ClusterSize  int            `json:"clusterSize" jsonschema:"platform n_samples; for Failure Mode it includes healthy neighbours, do not quote as the failing group"`
	Split        map[string]int `json:"split,omitempty"`
	Composition  []Composition  `json:"composition,omitempty"`
	Contrast     []Contrast     `json:"contrast,omitempty"`
	Remedy       *Remedy        `json:"remedy,omitempty"`
	TopSamples   []SampleRef    `json:"topSamples,omitempty" jsonschema:"most representative failing samples (ranked by affinity or loss), rendered ones first; pass to tl_view_samples"`
	RankedBy     string         `json:"rankedBy,omitempty" jsonschema:"what topSamples are ordered by"`
	Engine       map[string]any `json:"engine,omitempty" jsonschema:"the platform's own analysis of this insight: metrics_info, mutual_info_elements, is_train_aggressor, overfitting_metrics/evidence, cluster_extended_stats, severity_metrics, automatic_tests (see the glossary)"`
	HasTests     bool           `json:"hasSuggestedTests"`
	Link         string         `json:"link"`
	CreateTest   string         `json:"createTestLink,omitempty"`
	Warning      string         `json:"warning,omitempty"`
}

type InsightsOut struct {
	Server      string              `json:"server"`
	Version     string              `json:"version"`
	Note        string              `json:"note,omitempty"`
	ClassLabels map[string][]string `json:"classLabels,omitempty" jsonschema:"per prediction type, the class name at each index; fields ending in _prd_idx hold an index into that list (with one prediction type every such field uses it)"`
	Visualizers []Visualizer        `json:"visualizers,omitempty" jsonschema:"the visualizers the integration declared; tl_get_integration_code explains what each renders"`
	Insights    []Insight           `json:"insights"`
	Link        string              `json:"insightsPanelLink"`
	Reason      string              `json:"reason,omitempty"`
	NextStep    string              `json:"nextStep,omitempty"`
}

func (s *Server) getInsights(ctx context.Context, _ *sdk.CallToolRequest, in VersionIn) (*sdk.CallToolResult, InsightsOut, error) {
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, InsightsOut{}, err
	}
	access, err := s.policy(ctx, in.ProjectID)
	if err != nil {
		return nil, InsightsOut{}, err
	}
	e, err := s.export(ctx, in.ProjectID, in.VersionID, true)
	if err != nil {
		return nil, InsightsOut{}, err
	}
	out := InsightsOut{Server: s.client.UIBase(), Version: e.Version.Name, Insights: []Insight{}, Link: s.client.UIBase() + e.DeepLinkPath,
		ClassLabels: e.PredictionLabels, Visualizers: e.Visualizers}
	if !access.SampleRows {
		out.Note = "per-sample data is turned off for this project, so Failure Mode sizes, composition, metric contrast and representative samples are unavailable"
	}
	if len(e.Insights) == 0 {
		out.Reason, out.NextStep = "no-insights", "check tl_list_versions: the version must be evaluated; then generate insights from the Insights panel"
		return nil, out, nil
	}
	byID := map[string]int{}
	for _, raw := range e.Insights {
		if id, ok := raw.InsightType["id_"].(string); ok {
			byID[id] = int(raw.Index)
		}
	}
	ranked := map[int][]string{}
	var candidates []string
	for _, raw := range e.Insights {
		t := raw.InsightType
		typ := str(t["type"])
		ins := Insight{Index: int(raw.Index), Type: typ, Name: displayNames[typ], Severity: num(t["severity"]), Status: raw.Status,
			Description: raw.Description, ClusterSize: num(t["n_samples"]), GroupMeaning: "platform sample count", HasTests: len(list(t["automatic_tests"])) > 0,
			Engine: raw.Engine}
		if typ == "low_performance" {
			// n_samples counts healthy latent neighbours too; only the sample list gives the failing count
			ins.GroupMeaning = "unavailable without per-sample data"
		} else {
			n := ins.ClusterSize
			ins.GroupSize = &n
		}
		if ins.Name == "" {
			ins.Name = typ
		}
		if pid := str(t["parent_id"]); pid != "" {
			ins.ParentIndex = byID[pid]
		}
		if ls := str(t["latent_space"]); ls != "" {
			meaning := latentMeanings[ls]
			if meaning == "" {
				meaning = "grouped in the project's '" + ls + "' representation"
			}
			ins.LatentSpace = &LatentSpace{Name: ls, Means: meaning}
		}
		if fix, ok := t["aggressor_fixing"].(map[string]any); ok {
			ins.Remedy = &Remedy{SamplesToLabel: num(fix["num_of_samples_to_label"]), SamplesToAcquire: num(fix["num_of_samples_to_acquire"])}
		}
		link := raw.AnalyzeLinkPath
		if link == "" {
			link = e.DeepLinkPath
		}
		ins.Link = s.client.UIBase() + link
		if ins.HasTests {
			ins.CreateTest = createTestLink(ins.Link, raw.Cid)
		}
		if d := raw.Digest; d != nil {
			size := d.GroupSize
			ins.GroupSize, ins.GroupMeaning = &size, d.GroupDefinition
			ins.Split, ins.Composition, ins.Contrast, ins.RankedBy = d.Split, d.Composition, d.Contrast, d.RankedBy
			ids := d.RankedSampleIDs
			if len(ids) > 24 {
				ids = ids[:24]
			}
			ranked[len(out.Insights)] = ids
			candidates = append(candidates, ids...)
		} else if raw.DigestError != "" {
			ins.Warning = raw.DigestError
		} else if raw.CsvURL != "" {
			ins.Warning = olderServer
		}
		out.Insights = append(out.Insights, ins)
	}
	if len(candidates) > 0 {
		// with visualizations turned off the ids are still worth returning, just not viewable
		var rendered map[string][]assetFile
		if access.Visuals {
			if rendered, err = s.sampleAssets(ctx, in.ProjectID, in.VersionID, unique(candidates)); err != nil {
				rendered = nil
			}
		} else if out.Note == "" {
			out.Note = "sample visualizations are turned off for this project, so topSamples cannot be viewed here"
		}
		for i, ids := range ranked {
			out.Insights[i].TopSamples = renderedFirst(ids, rendered, maxViewSamples)
		}
	}
	sort.SliceStable(out.Insights, func(i, j int) bool { return out.Insights[i].Severity > out.Insights[j].Severity })
	return nil, out, nil
}

func createTestLink(link, cid string) string {
	sep := "?"
	if strings.Contains(link, "?") {
		sep = "&"
	}
	return link + sep + "addTestFromInsight=" + cid
}
