package mcp

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	// the bundle is unpacked here first, so an interrupted download never mixes with a previous export
	stagingDir     = ".tensorleap-export-partial"
	maxBundleBytes = 4 << 30
	maxResultNotes = 10
)

type ExportIn struct {
	ProjectID     string   `json:"projectId" jsonschema:"project id or name"`
	VersionID     string   `json:"versionId" jsonschema:"version id or name, or \"latest\""`
	Dir           string   `json:"dir" jsonschema:"absolute path of an empty directory, or of a previous export of the same version to refresh (created if missing); e.g. <cwd>/tensorleap-analysis/<version>"`
	TopK          int      `json:"topK,omitempty" jsonschema:"samples per insight (default 24, max 100; sub-insights get at most 6)"`
	HeatmapLabels []string `json:"heatmapLabels,omitempty" jsonschema:"heatmap labels to render as overlays (default: the first 3 per sample)"`
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

// manifest is the part of the server's manifest.json this tool reads; the file itself is kept whole
type manifest struct {
	ProjectID     string   `json:"projectId"`
	VersionID     string   `json:"versionId"`
	Version       string   `json:"version"`
	InsightsPanel string   `json:"insightsPanelLink"`
	FilesWritten  int      `json:"filesWritten"`
	BytesWritten  int64    `json:"bytesWritten"`
	Skipped       []string `json:"skipped"`
	Notes         []string `json:"notes"`
	Insights      []struct {
		Index       int    `json:"index"`
		Type        string `json:"type"`
		Name        string `json:"name"`
		ParentIndex int    `json:"parentIndex"`
		Dir         string `json:"dir"`
		Summary     *struct {
			GroupSize int `json:"groupSize"`
		} `json:"summary"`
		Samples []struct {
			Rendered bool `json:"rendered"`
		} `json:"samples"`
	} `json:"insights"`
}

func (s *Server) addExportTool(srv *sdk.Server) {
	closed := false
	sdk.AddTool(srv, &sdk.Tool{Name: "tl_export_analysis",
		Description: "Write a version's analysis to a local directory for a report: per-insight sample lists (CSV), cluster membership, the platform's fixing-samples list, the population CSV, the integration code, and each insight's top samples (payloads, images, rendered heatmap overlays). Returns a short summary and the path of manifest.json, which lists every file written (paths relative to dir). Use it instead of tl_view_samples when you need files or more than a handful of samples.",
		Annotations: &sdk.ToolAnnotations{Title: "Export analysis", ReadOnlyHint: false, IdempotentHint: true, OpenWorldHint: &closed}},
		s.exportAnalysis)
}

// previousExport accepts a missing or empty dir, a manifest.json written by this tool, or the pre-MCP
// skill's output (insights.json); anything else is a directory the user owns
func previousExport(dir string) (projectID, versionID string, err error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return "", "", nil
	}
	if err != nil {
		return "", "", err
	}
	if len(entries) == 0 || len(entries) == 1 && entries[0].Name() == stagingDir {
		return "", "", nil
	}
	var m manifest
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

// unpack writes a server bundle under dir; names come from the server, so they are treated as untrusted
func unpack(r io.Reader, dir string) error {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return fmt.Errorf("the export bundle is not a gzip stream: %w", err)
	}
	tr := tar.NewReader(gz)
	var total int64
	for {
		h, err := tr.Next()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return fmt.Errorf("the export was cut short (%v); run tl_export_analysis again", err)
		}
		name := path.Clean(h.Name)
		if h.Typeflag != tar.TypeReg || path.IsAbs(name) || name == "." || name == ".." || strings.HasPrefix(name, "../") ||
			strings.ContainsAny(name, "\\\x00") {
			continue
		}
		if total += h.Size; total > maxBundleBytes {
			return fmt.Errorf("the export is larger than %d GB; pass a smaller topK", maxBundleBytes>>30)
		}
		full := filepath.Join(dir, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			return err
		}
		f, err := os.OpenFile(full, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, io.LimitReader(tr, h.Size))
		if cerr := f.Close(); err == nil {
			err = cerr
		}
		if err != nil {
			return fmt.Errorf("the export was cut short (%v); run tl_export_analysis again", err)
		}
	}
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
	staging := filepath.Join(in.Dir, stagingDir)
	if err := os.RemoveAll(staging); err != nil {
		return nil, ExportResult{}, err
	}
	if err := os.MkdirAll(staging, 0o755); err != nil {
		return nil, ExportResult{}, err
	}
	defer os.RemoveAll(staging)

	body, err := s.client.PostStream(ctx, "analysis-export/exportBundle", map[string]any{
		"projectId": in.ProjectID, "versionId": in.VersionID, "topK": in.TopK, "heatmapLabels": in.HeatmapLabels})
	if err != nil {
		return nil, ExportResult{}, explain(err)
	}
	err = unpack(body, staging)
	body.Close()
	if err != nil {
		return nil, ExportResult{}, err
	}
	raw, err := os.ReadFile(filepath.Join(staging, "manifest.json"))
	var m manifest
	if err != nil || json.Unmarshal(raw, &m) != nil {
		return nil, ExportResult{}, errors.New("the export was cut short (no manifest.json); run tl_export_analysis again")
	}
	if prevProject != "" && (prevProject != m.ProjectID || prevVersion != m.VersionID) {
		return nil, ExportResult{}, fmt.Errorf("%s holds an export of another version; choose another directory", in.Dir)
	}

	if err := clearExport(in.Dir); err != nil {
		return nil, ExportResult{}, err
	}
	entries, err := os.ReadDir(staging)
	if err != nil {
		return nil, ExportResult{}, err
	}
	for _, e := range entries {
		if e.Name() == "manifest.json" {
			continue
		}
		if err := os.Rename(filepath.Join(staging, e.Name()), filepath.Join(in.Dir, e.Name())); err != nil {
			return nil, ExportResult{}, err
		}
	}
	// the server cannot know where the bundle lands; dir is the one absolute path in the manifest
	var whole map[string]any
	_ = json.Unmarshal(raw, &whole)
	whole["dir"] = in.Dir
	out, _ := json.MarshalIndent(whole, "", "  ")
	manifestPath := filepath.Join(in.Dir, "manifest.json")
	if err := os.WriteFile(manifestPath, out, 0o644); err != nil {
		return nil, ExportResult{}, err
	}

	res := ExportResult{Dir: in.Dir, Manifest: manifestPath, Version: m.Version, InsightsPanel: m.InsightsPanel, Insights: []ExportedBrief{},
		FilesWritten: m.FilesWritten, BytesWritten: m.BytesWritten, Skipped: m.Skipped, Notes: m.Notes}
	if len(res.Notes) > maxResultNotes {
		res.Notes = append(res.Notes[:maxResultNotes:maxResultNotes], fmt.Sprintf("%d more notes in manifest.json", len(m.Notes)-maxResultNotes))
	}
	for _, ins := range m.Insights {
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
