package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

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
	Dir           string   `json:"dir" jsonschema:"absolute path of the directory to write into (created if missing)"`
	TopK          int      `json:"topK,omitempty" jsonschema:"samples per insight (default 24, max 100; sub-insights get at most 6)"`
	HeatmapLabels []string `json:"heatmapLabels,omitempty" jsonschema:"heatmap labels to render as overlays (default: the first 3 per sample)"`
}

type ExportedSample struct {
	ID       string   `json:"id"`
	Rendered bool     `json:"rendered"`
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

type ExportOut struct {
	Dir            string              `json:"dir"`
	Manifest       string              `json:"manifest"`
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
	Skipped        []string            `json:"skipped,omitempty" jsonschema:"what could not be exported and why"`
	Notes          []string            `json:"notes,omitempty"`
}

type exportWriter struct {
	dir   string
	mu    sync.Mutex
	files int
	bytes int64
}

// write keeps every file under dir; names come from the server, so they are treated as untrusted
func (w *exportWriter) write(rel string, b []byte) (string, error) {
	rel = filepath.FromSlash(path.Clean("/" + filepath.ToSlash(rel)))[1:]
	full := filepath.Join(w.dir, rel)
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
	return full, nil
}

func (s *Server) exportAnalysis(ctx context.Context, _ *sdk.CallToolRequest, in ExportIn) (*sdk.CallToolResult, ExportOut, error) {
	if !filepath.IsAbs(in.Dir) {
		return nil, ExportOut{}, errors.New("dir must be an absolute path")
	}
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, ExportOut{}, err
	}
	access, err := s.allowed(ctx, in.ProjectID, statsClass)
	if err != nil {
		return nil, ExportOut{}, err
	}
	e, err := s.export(ctx, in.ProjectID, in.VersionID)
	if err != nil {
		return nil, ExportOut{}, err
	}
	topK := in.TopK
	if topK <= 0 {
		topK = defaultExportTopK
	}
	if topK > maxExportTopK {
		topK = maxExportTopK
	}
	if err := os.MkdirAll(in.Dir, 0o755); err != nil {
		return nil, ExportOut{}, err
	}
	if entries, err := os.ReadDir(in.Dir); err == nil && len(entries) > 0 {
		if _, err := os.Stat(filepath.Join(in.Dir, "manifest.json")); err != nil {
			return nil, ExportOut{}, fmt.Errorf("%s is not empty and is not a previous export; choose an empty directory", in.Dir)
		}
	}
	w := &exportWriter{dir: in.Dir}
	out := ExportOut{Dir: in.Dir, ProjectID: in.ProjectID, VersionID: in.VersionID, Version: e.Version.Name, InsightsPanel: s.client.UIBase() + e.DeepLinkPath,
		ClassLabels: e.PredictionLabels, Visualizers: e.Visualizers, Insights: []ExportedInsight{}}
	var mu sync.Mutex
	note := func(format string, a ...any) {
		mu.Lock()
		out.Notes = append(out.Notes, fmt.Sprintf(format, a...))
		mu.Unlock()
	}
	skip := func(what string, class aiClass) {
		out.Skipped = append(out.Skipped, fmt.Sprintf("%s: %s", what, class.refusal(access.admin)))
	}
	fetch := func(url string) ([]byte, error) {
		b, err := s.client.Download(ctx, url)
		if err != nil {
			return nil, err
		}
		return decompress(b)
	}

	pop, popErr := s.population(ctx, in.VersionID, e.PopulationCsvURL)
	switch {
	case e.PopulationCsvURL == "" && !access.SampleRows:
		skip("population csv", sampleRowsClass)
	case e.PopulationCsvURL != "":
		if b, err := fetch(e.PopulationCsvURL); err != nil {
			note("population csv: %v", err)
		} else if p, err := w.write("population.csv", b); err != nil {
			return nil, ExportOut{}, err
		} else {
			out.PopulationCsv = p
		}
	}
	if popErr != nil && access.SampleRows {
		note("no all-data baseline for composition: %v", popErr)
	}

	switch {
	case !access.Code || e.IntegrationCodeURL == "" && e.IntegrationEntryFile != "":
		skip("integration code", codeClass)
	case e.IntegrationCodeURL == "":
		out.Skipped = append(out.Skipped, "integration code: this version has no code snapshot on the server")
	default:
		arc, err := s.integrationCode(ctx, in.ProjectID, in.VersionID, access.admin)
		if err != nil {
			note("integration code: %v", err)
		} else {
			for name, b := range arc.files {
				if _, err := w.write(path.Join("integration", name), []byte(Scrub(string(b)))); err != nil {
					return nil, ExportOut{}, err
				}
			}
			out.IntegrationDir, out.EntryFile = filepath.Join(in.Dir, "integration"), arc.entry
		}
	}

	byID := map[string]*exportedInsight{}
	for i := range e.Insights {
		if id, ok := e.Insights[i].InsightType["id_"].(string); ok {
			byID[id] = &e.Insights[i]
		}
	}
	dirOf := func(raw *exportedInsight) string {
		own := fmt.Sprintf("insight_%d_%s", int(raw.Index), str(raw.InsightType["type"]))
		if parent, ok := byID[str(raw.InsightType["parent_id"])]; ok && parent != raw {
			return path.Join(fmt.Sprintf("insight_%d_%s", int(parent.Index), str(parent.InsightType["type"])), "sub_"+strings.TrimPrefix(own, "insight_"))
		}
		return own
	}
	type sampleJob struct {
		insight *ExportedInsight
		pos     int
	}
	var jobs []sampleJob
	var candidates []string
	ranked := map[int][]string{}
	for i := range e.Insights {
		raw := &e.Insights[i]
		t := raw.InsightType
		typ := str(t["type"])
		ins := ExportedInsight{Index: int(raw.Index), Type: typ, Name: displayNames[typ], Status: raw.Status, Description: raw.Description,
			Dir: filepath.Join(in.Dir, filepath.FromSlash(dirOf(raw))), Engine: enginePayload(t), HasTests: len(list(t["automatic_tests"])) > 0, Samples: []ExportedSample{}}
		if ins.Name == "" {
			ins.Name = typ
		}
		if parent, ok := byID[str(t["parent_id"])]; ok {
			ins.ParentIndex = int(parent.Index)
		}
		link := raw.AnalyzeLinkPath
		if link == "" {
			link = e.DeepLinkPath
		}
		ins.Link = s.client.UIBase() + link
		if ins.HasTests {
			ins.CreateTest = createTestLink(ins.Link, raw.Cid)
		}
		rel := dirOf(raw)
		if raw.CsvURL == "" && !access.SampleRows && len(out.Insights) == 0 {
			skip("insight sample lists, cluster membership and the fixing-samples list", sampleRowsClass)
		}
		for name, url := range map[string]string{"samples.csv": raw.CsvURL, "cluster.json": raw.ClusterBlobURL, "fixing_samples.csv": raw.FixingCsvURL, "top_panel.json": raw.TopPanelURL} {
			if url == "" {
				continue
			}
			b, err := fetch(url)
			if err != nil {
				note("insight %d %s: %v", ins.Index, name, err)
				continue
			}
			p, err := w.write(path.Join(rel, name), b)
			if err != nil {
				return nil, ExportOut{}, err
			}
			switch name {
			case "samples.csv":
				ins.SamplesCsv = p
				if sum, err := Summarize(b, pop); err != nil {
					note("insight %d sample list unreadable: %v", ins.Index, err)
				} else {
					ins.Summary = sum
					ranked[len(out.Insights)] = sum.RankedIDs
					candidates = append(candidates, sum.RankedIDs...)
				}
			case "cluster.json":
				ins.ClusterJson = p
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
	for i := range out.Insights {
		ins := &out.Insights[i]
		k := topK
		if ins.ParentIndex != 0 && k > subInsightTopK {
			k = subInsightTopK
		}
		for _, ref := range renderedFirst(ranked[i], rendered, k) {
			ins.Samples = append(ins.Samples, ExportedSample{ID: ref.ID, Rendered: ref.Rendered})
			if ref.Rendered {
				jobs = append(jobs, sampleJob{insight: ins, pos: len(ins.Samples) - 1})
			}
		}
	}

	sem := make(chan struct{}, exportWorkers)
	var wg sync.WaitGroup
	var firstErr error
	for _, job := range jobs {
		wg.Add(1)
		sem <- struct{}{}
		go func(job sampleJob) {
			defer wg.Done()
			defer func() { <-sem }()
			smp := &job.insight.Samples[job.pos]
			files, err := s.writeSample(ctx, w, job.insight.Dir, smp.ID, rendered[smp.ID], in.HeatmapLabels)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err
				}
				mu.Unlock()
				note("sample %s: %v", smp.ID, err)
			}
			smp.Files = files
			if len(files) > 0 {
				smp.Dir = filepath.Join(job.insight.Dir, "samples", smp.ID)
			}
		}(job)
	}
	wg.Wait()
	if firstErr != nil && w.files == 0 {
		return nil, ExportOut{}, firstErr
	}

	out.FilesWritten, out.BytesWritten = w.files, w.bytes
	out.Manifest = filepath.Join(in.Dir, "manifest.json")
	manifest, _ := json.MarshalIndent(out, "", "  ")
	if _, err := w.write("manifest.json", manifest); err != nil {
		return nil, ExportOut{}, err
	}
	return nil, out, nil
}

// writeSample stores a sample's payloads and assets under samples/<id>/<dataType>/<visualizer>/...
// and renders heatmap overlays beside them
func (s *Server) writeSample(ctx context.Context, w *exportWriter, insightDir, id string, files []assetFile, labels []string) ([]string, error) {
	hashed := "/" + visualizationID(id) + "/"
	base := filepath.Join("samples", id)
	rel, _ := filepath.Rel(w.dir, insightDir)
	base = filepath.Join(rel, base)
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
				if p, err := w.write(filepath.Join(base, dataType, visualizer, name), o.jpeg); err == nil {
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
		if p, err := w.write(filepath.Join(base, filepath.FromSlash(suffix)), b); err == nil {
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
