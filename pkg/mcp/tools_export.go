package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	defaultExportTopK = 24
	subInsightTopK    = 6
	maxExportTopK     = 100
	exportWorkers     = 8
)

type ExportIn struct {
	ProjectID     string   `json:"projectId" jsonschema:"project id or name"`
	VersionID     string   `json:"versionId" jsonschema:"version id or name, or \"latest\""`
	Dir           string   `json:"dir" jsonschema:"absolute path of an empty directory, or of a previous export of the same version to refresh (created if missing); e.g. <cwd>/tensorleap-analysis/<version>"`
	TopK          int      `json:"topK,omitempty" jsonschema:"samples per insight (default 24, max 100; sub-insights get at most 6)"`
	HeatmapLabels []string `json:"heatmapLabels,omitempty" jsonschema:"heatmap labels to render as overlays (default: the first 3 per sample)"`
}

type ExportedSample struct {
	ID       string   `json:"id"`
	Rank     int      `json:"rank" jsonschema:"1 = most representative, in the insight's ranking (summary.rankedBy)"`
	Rendered bool     `json:"rendered" jsonschema:"true when files holds its visualizations"`
	Dir      string   `json:"dir,omitempty"`
	Files    []string `json:"files,omitempty"`
}

type ExportedInsight struct {
	Index        int              `json:"index"`
	Type         string           `json:"type"`
	Name         string           `json:"name"`
	Status       string           `json:"status"`
	Description  string           `json:"description,omitempty"`
	ParentIndex  int              `json:"parentIndex,omitempty"`
	Dir          string           `json:"dir"`
	SamplesCsv   string           `json:"samplesCsv,omitempty"`
	ClusterJson  string           `json:"clusterJson,omitempty"`
	FixingCsv    string           `json:"fixingCsv,omitempty" jsonschema:"the samples the platform selected for labeling"`
	TopPanelJson string           `json:"topPanelJson,omitempty"`
	Summary      *GroupSummary    `json:"summary,omitempty"`
	Engine       map[string]any   `json:"engine,omitempty"`
	Link         string           `json:"link"`
	CreateTest   string           `json:"createTestLink,omitempty"`
	HasTests     bool             `json:"hasSuggestedTests"`
	Samples      []ExportedSample `json:"samples"`
}

// ExportOut is manifest.json; every path in it except dir is relative to dir
type ExportOut struct {
	Dir            string              `json:"dir"`
	ProjectID      string              `json:"projectId"`
	VersionID      string              `json:"versionId"`
	Version        string              `json:"version"`
	InsightsPanel  string              `json:"insightsPanelLink"`
	ClassLabels    map[string][]string `json:"classLabels,omitempty"`
	Visualizers    []Visualizer        `json:"visualizers,omitempty"`
	PopulationCsv  string              `json:"populationCsv,omitempty"`
	IntegrationDir string              `json:"integrationDir,omitempty"`
	EntryFile      string              `json:"entryFile,omitempty"`
	Insights       []ExportedInsight   `json:"insights"`
	FilesWritten   int                 `json:"filesWritten"`
	BytesWritten   int64               `json:"bytesWritten"`
	Skipped        []string            `json:"skipped,omitempty"`
	Notes          []string            `json:"notes,omitempty"`
}

type ExportedBrief struct {
	Index       int    `json:"index"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	ParentIndex int    `json:"parentIndex,omitempty"`
	Dir         string `json:"dir"`
	GroupSize   int    `json:"groupSize,omitempty"`
	Exported    int    `json:"samplesExported"`
	Rendered    int    `json:"samplesRendered"`
}

// ExportResult is what the assistant sees; the full detail stays in manifest.json
type ExportResult struct {
	Dir           string          `json:"dir"`
	Manifest      string          `json:"manifest"`
	Version       string          `json:"version"`
	InsightsPanel string          `json:"insightsPanelLink"`
	Insights      []ExportedBrief `json:"insights"`
	FilesWritten  int             `json:"filesWritten"`
	BytesWritten  int64           `json:"bytesWritten"`
	Skipped       []string        `json:"skipped,omitempty" jsonschema:"what could not be exported and why"`
	Notes         []string        `json:"notes,omitempty"`
}

const maxResultNotes = 10

type exportWriter struct {
	dir   string
	mu    sync.Mutex
	files int
	bytes int64
}

// write keeps every file under dir and returns its slash-separated path relative to dir; names come from the server
func (w *exportWriter) write(rel string, b []byte) (string, error) {
	rel = path.Clean("/" + filepath.ToSlash(rel))[1:]
	full := filepath.Join(w.dir, filepath.FromSlash(rel))
	if r, err := filepath.Rel(w.dir, full); err != nil || r == "." || strings.HasPrefix(r, "..") {
		return "", fmt.Errorf("refusing to write outside %s: %q", w.dir, rel)
	}
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(full, b, 0o644); err != nil {
		return "", err
	}
	w.mu.Lock()
	w.files++
	w.bytes += int64(len(b))
	w.mu.Unlock()
	return rel, nil
}

// previousExport accepts a missing or empty dir, a manifest.json written by this tool, or the pre-MCP skill's output
// (insights.json); anything else is a directory the user owns
func previousExport(dir string) (projectID, versionID string, err error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) || err == nil && len(entries) == 0 {
		return "", "", nil
	}
	if err != nil {
		return "", "", err
	}
	var m ExportOut
	if b, err := os.ReadFile(filepath.Join(dir, "manifest.json")); err == nil && json.Unmarshal(b, &m) == nil && m.ProjectID != "" && m.VersionID != "" {
		return m.ProjectID, m.VersionID, nil
	}
	if _, err := os.Stat(filepath.Join(dir, "insights.json")); err == nil {
		return "", "", nil
	}
	return "", "", fmt.Errorf("%s is not empty and is not a previous export; choose an empty directory (e.g. a new subdirectory)", dir)
}

// clearExport removes what a previous export wrote so a re-export never mixes old and new samples
func clearExport(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		n := e.Name()
		if strings.HasPrefix(n, "insight_") || n == "integration" || n == "population.csv" || n == "manifest.json" {
			if err := os.RemoveAll(filepath.Join(dir, n)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Server) exportAnalysis(ctx context.Context, _ *sdk.CallToolRequest, in ExportIn) (*sdk.CallToolResult, ExportResult, error) {
	if !filepath.IsAbs(in.Dir) {
		return nil, ExportResult{}, errors.New("dir must be an absolute path")
	}
	in.Dir = filepath.Clean(in.Dir)
	prevProject, prevVersion, err := previousExport(in.Dir)
	if err != nil {
		return nil, ExportResult{}, err
	}
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, ExportResult{}, err
	}
	if prevProject != "" && (prevProject != in.ProjectID || prevVersion != in.VersionID) {
		return nil, ExportResult{}, fmt.Errorf("%s holds an export of another version; choose another directory", in.Dir)
	}
	s.fresh(in.ProjectID)
	access, err := s.allowed(ctx, in.ProjectID, statsClass)
	if err != nil {
		return nil, ExportResult{}, err
	}
	e, err := s.export(ctx, in.ProjectID, in.VersionID)
	if err != nil {
		return nil, ExportResult{}, err
	}
	topK := in.TopK
	if topK <= 0 {
		topK = defaultExportTopK
	}
	if topK > maxExportTopK {
		topK = maxExportTopK
	}
	if err := os.MkdirAll(in.Dir, 0o755); err != nil {
		return nil, ExportResult{}, err
	}
	if err := clearExport(in.Dir); err != nil {
		return nil, ExportResult{}, err
	}
	w := &exportWriter{dir: in.Dir}
	out := ExportOut{Dir: in.Dir, ProjectID: in.ProjectID, VersionID: in.VersionID, Version: e.Version.Name, InsightsPanel: s.client.UIBase() + e.DeepLinkPath,
		ClassLabels: e.PredictionLabels, Visualizers: e.Visualizers, Insights: []ExportedInsight{}, Notes: []string{"export did not finish; run tl_export_analysis again"}}
	// a stub first, so an interrupted export still marks the directory as ours
	stub, _ := json.Marshal(out)
	if _, err := w.write("manifest.json", stub); err != nil {
		return nil, ExportResult{}, err
	}
	out.Notes = nil
	var mu sync.Mutex
	note := func(format string, a ...any) {
		mu.Lock()
		out.Notes = append(out.Notes, fmt.Sprintf(format, a...))
		mu.Unlock()
	}
	skip := func(what string, class aiClass) {
		out.Skipped = append(out.Skipped, fmt.Sprintf("%s: %s", what, class.refusal(access.admin)))
	}

	popBytes, pop, popErr := s.populationWithBytes(ctx, in.VersionID, e.PopulationCsvURL, true)
	switch {
	case e.PopulationCsvURL == "" && !access.SampleRows:
		skip("population csv", sampleRowsClass)
	case popBytes != nil:
		p, err := w.write("population.csv", popBytes)
		if err != nil {
			return nil, ExportResult{}, err
		}
		out.PopulationCsv = p
	}
	if popErr != nil && access.SampleRows {
		note("no all-data baseline for composition: %v", popErr)
	}

	switch {
	case e.IntegrationCodeURL == "" && e.IntegrationEntryFile != "":
		skip("integration code", codeClass)
	case e.IntegrationCodeURL == "":
		out.Skipped = append(out.Skipped, "integration code: this version has no code snapshot on the server")
	default:
		arc, err := s.integrationCode(ctx, e, in.VersionID, access.admin)
		if err != nil {
			note("integration code: %v", err)
			break
		}
		for name, b := range arc.files {
			if utf8.Valid(b) {
				b = []byte(Scrub(string(b)))
			}
			if _, err := w.write(path.Join("integration", name), b); err != nil {
				note("integration %s: %v", name, err)
			}
		}
		out.IntegrationDir, out.EntryFile = "integration", arc.entry
	}

	byID := map[string]*exportedInsight{}
	for i := range e.Insights {
		if id, ok := e.Insights[i].InsightType["id_"].(string); ok {
			byID[id] = &e.Insights[i]
		}
	}
	dirName := func(raw *exportedInsight) string {
		return fmt.Sprintf("insight_%d_%s", int(raw.Index), safeName(str(raw.InsightType["type"])))
	}
	dirOf := func(raw *exportedInsight) string {
		if parent, ok := byID[str(raw.InsightType["parent_id"])]; ok && parent != raw {
			return path.Join(dirName(parent), "sub_"+strings.TrimPrefix(dirName(raw), "insight_"))
		}
		return dirName(raw)
	}
	var candidates []string
	ranked := map[int][]string{}
	for i := range e.Insights {
		raw := &e.Insights[i]
		t := raw.InsightType
		typ := str(t["type"])
		ins := ExportedInsight{Index: int(raw.Index), Type: typ, Name: displayNames[typ], Status: raw.Status, Description: raw.Description,
			Dir: dirOf(raw), Engine: enginePayload(t), HasTests: len(list(t["automatic_tests"])) > 0, Samples: []ExportedSample{}}
		if ins.Name == "" {
			ins.Name = typ
		}
		if parent, ok := byID[str(t["parent_id"])]; ok {
			ins.ParentIndex = int(parent.Index)
		}
		k := topK
		if ins.ParentIndex != 0 && k > subInsightTopK {
			k = subInsightTopK
		}
		link := raw.AnalyzeLinkPath
		if link == "" {
			link = e.DeepLinkPath
		}
		ins.Link = s.client.UIBase() + link
		if ins.HasTests {
			ins.CreateTest = createTestLink(ins.Link, raw.Cid)
		}
		if raw.CsvURL == "" && !access.SampleRows && len(out.Insights) == 0 {
			skip("insight sample lists, cluster membership and the fixing-samples list", sampleRowsClass)
		}
		var members map[string]bool
		// cluster.json first: it narrows a sample list that duplication-type insights share
		for _, f := range []struct{ name, url string }{{"cluster.json", raw.ClusterBlobURL}, {"samples.csv", raw.CsvURL}, {"fixing_samples.csv", raw.FixingCsvURL}, {"top_panel.json", raw.TopPanelURL}} {
			if f.url == "" {
				continue
			}
			b, err := s.download(ctx, f.url)
			if err != nil {
				note("insight %d %s: %v", ins.Index, f.name, err)
				continue
			}
			if f.name == "top_panel.json" {
				var v any
				if json.Unmarshal(b, &v) == nil {
					b, _ = json.MarshalIndent(stripInternal(v), "", "  ")
				}
			}
			p, err := w.write(path.Join(ins.Dir, f.name), b)
			if err != nil {
				return nil, ExportResult{}, err
			}
			switch f.name {
			case "cluster.json":
				ins.ClusterJson = p
				if typ != "low_performance" {
					members = clusterMembers(b)
				}
			case "samples.csv":
				ins.SamplesCsv = p
				sum, err := Summarize(b, pop, members)
				if err != nil {
					note("insight %d sample list unreadable: %v", ins.Index, err)
					break
				}
				ins.Summary = sum
				ids := sum.RankedIDs
				if len(ids) > k {
					ids = ids[:k]
				}
				ranked[len(out.Insights)] = ids
				candidates = append(candidates, ids...)
			case "fixing_samples.csv":
				ins.FixingCsv = p
			case "top_panel.json":
				ins.TopPanelJson = p
			}
		}
		out.Insights = append(out.Insights, ins)
	}

	rendered := map[string][]assetFile{}
	if len(candidates) > 0 {
		if !access.Visuals {
			skip("sample visualizations", visualsClass)
		} else if rendered, err = s.sampleAssets(ctx, in.ProjectID, in.VersionID, unique(candidates)); err != nil {
			note("sample visualizations: %v", err)
			rendered = map[string][]assetFile{}
		}
	}
	type sampleJob struct {
		insight *ExportedInsight
		pos     int
	}
	var jobs []sampleJob
	for i := range out.Insights {
		ins := &out.Insights[i]
		for r, id := range ranked[i] {
			ins.Samples = append(ins.Samples, ExportedSample{ID: id, Rank: r + 1})
			if len(rendered[id]) > 0 {
				jobs = append(jobs, sampleJob{insight: ins, pos: len(ins.Samples) - 1})
			}
		}
	}

	sem := make(chan struct{}, exportWorkers)
	var wg sync.WaitGroup
	for _, job := range jobs {
		wg.Add(1)
		sem <- struct{}{}
		go func(job sampleJob) {
			defer wg.Done()
			defer func() { <-sem }()
			smp := &job.insight.Samples[job.pos]
			dir := path.Join(job.insight.Dir, "samples", safeName(smp.ID))
			files, err := s.writeSample(ctx, w, dir, smp.ID, rendered[smp.ID], in.HeatmapLabels)
			if err != nil {
				note("sample %s: %v", smp.ID, err)
			}
			smp.Files, smp.Rendered = files, len(files) > 0
			if smp.Rendered {
				smp.Dir = dir
			}
		}(job)
	}
	wg.Wait()
	if err := ctx.Err(); err != nil {
		return nil, ExportResult{}, err
	}

	out.FilesWritten, out.BytesWritten = w.files, w.bytes
	manifest, _ := json.MarshalIndent(out, "", "  ")
	if _, err := w.write("manifest.json", manifest); err != nil {
		return nil, ExportResult{}, err
	}
	res := ExportResult{Dir: in.Dir, Manifest: filepath.Join(in.Dir, "manifest.json"), Version: out.Version, InsightsPanel: out.InsightsPanel,
		Insights: []ExportedBrief{}, FilesWritten: out.FilesWritten, BytesWritten: out.BytesWritten, Skipped: out.Skipped, Notes: out.Notes}
	if len(res.Notes) > maxResultNotes {
		res.Notes = append(res.Notes[:maxResultNotes:maxResultNotes], fmt.Sprintf("%d more notes in manifest.json", len(out.Notes)-maxResultNotes))
	}
	for _, ins := range out.Insights {
		b := ExportedBrief{Index: ins.Index, Type: ins.Type, Name: ins.Name, ParentIndex: ins.ParentIndex, Dir: ins.Dir, Exported: len(ins.Samples)}
		if ins.Summary != nil {
			b.GroupSize = ins.Summary.GroupSize
		}
		for _, smp := range ins.Samples {
			if smp.Rendered {
				b.Rendered++
			}
		}
		res.Insights = append(res.Insights, b)
	}
	return nil, res, nil
}

// writeSample stores a sample's payloads and assets under base/<dataType>/<visualizer>/...
// and renders heatmap overlays beside them
func (s *Server) writeSample(ctx context.Context, w *exportWriter, base, id string, files []assetFile, labels []string) ([]string, error) {
	hashed := "/" + visualizationID(id) + "/"
	var written []string
	var firstErr error
	for _, f := range files {
		i := strings.Index(f.Path, hashed)
		if i < 0 {
			continue
		}
		suffix := path.Clean(f.Path[i+len(hashed):])
		if suffix == ".." || strings.HasPrefix(suffix, "../") || path.IsAbs(suffix) || strings.ContainsRune(suffix, '\\') {
			continue
		}
		dataType, visualizer := assetKind(f.Path)
		if dataType == "image_heatmap" && path.Base(f.Path) == "payload.json" {
			overlays, _, err := s.renderOverlays(ctx, f.URL, files, labels, defaultHeatmaps, true)
			if err != nil && firstErr == nil {
				firstErr = err
			}
			for _, o := range overlays {
				name := "overlay.jpg"
				if o.label != "" {
					name = "overlay_" + safeName(o.label) + ".jpg"
				}
				if p, err := w.write(path.Join(base, dataType, visualizer, name), o.jpeg); err == nil {
					written = append(written, p)
				} else if firstErr == nil {
					firstErr = err
				}
			}
		}
		b, err := s.client.Download(ctx, f.URL)
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		if path.Base(suffix) == "payload.json" {
			b = []byte(Scrub(string(b)))
		}
		if p, err := w.write(path.Join(base, suffix), b); err == nil {
			written = append(written, p)
		} else if firstErr == nil {
			firstErr = err
		}
	}
	sort.Strings(written)
	return written, firstErr
}

func safeName(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r == '/' || r == '\\' || r == ':' || r == 0 || r == '.' && b.Len() == 0 {
			r = '_'
		}
		b.WriteRune(r)
	}
	r := []rune(b.String())
	if len(r) > 80 {
		r = r[:80]
	}
	return string(r)
}
