package mcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/png"
	"path"
	"sort"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	maxViewSamples = 6
	thumbSize      = 640
)

func visualizationID(sampleID string) string {
	state, index, ok := strings.Cut(sampleID, "_")
	if !ok {
		return sampleID
	}
	sum := sha256.Sum256([]byte(index))
	return state + "_" + strings.TrimRight(base64.URLEncoding.EncodeToString(sum[:16]), "=")
}

type assetFile struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

func (s *Server) sampleAssets(ctx context.Context, projectID, versionID string, sampleIDs []string) (map[string][]assetFile, error) {
	byHash := map[string]string{}
	hashed := make([]string, 0, len(sampleIDs))
	for _, id := range sampleIDs {
		h := visualizationID(id)
		byHash[h] = id
		hashed = append(hashed, h)
	}
	var resp struct {
		Samples []struct {
			SampleID string      `json:"sampleId"`
			Files    []assetFile `json:"files"`
		} `json:"samples"`
	}
	if err := s.client.Post(ctx, "analysis-export/getSampleAssets", map[string]any{"projectId": projectID, "versionId": versionID, "sampleIds": hashed}, &resp); err != nil {
		return nil, explain(err)
	}
	out := map[string][]assetFile{}
	for _, smp := range resp.Samples {
		if len(smp.Files) > 0 {
			out[byHash[smp.SampleID]] = smp.Files
		}
	}
	return out, nil
}

type SampleRef struct {
	ID       string `json:"id"`
	Rendered bool   `json:"rendered" jsonschema:"true when tl_view_samples can show its visualizations"`
}

func renderedFirst(ids []string, rendered map[string][]assetFile, n int) []SampleRef {
	refs := make([]SampleRef, 0, len(ids))
	for _, id := range ids {
		refs = append(refs, SampleRef{ID: id, Rendered: len(rendered[id]) > 0})
	}
	sort.SliceStable(refs, func(i, j int) bool { return refs[i].Rendered && !refs[j].Rendered })
	if len(refs) > n {
		refs = refs[:n]
	}
	return refs
}

type ViewIn struct {
	ProjectID  string   `json:"projectId"`
	VersionID  string   `json:"versionId"`
	SampleIDs  []string `json:"sampleIds" jsonschema:"sample ids such as training_123 (from tl_get_insights topSamples), at most 6"`
	Visualizer string   `json:"visualizer,omitempty" jsonschema:"optional: only this visualizer (name as listed in the response)"`
	Full       bool     `json:"fullResolution,omitempty" jsonschema:"original resolution instead of 640px thumbnails; use only when fine detail matters"`
}

type ViewedSample struct {
	ID          string   `json:"id"`
	Visualizers []string `json:"visualizers"`
	Rendered    bool     `json:"rendered"`
}

type ViewOut struct {
	Samples  []ViewedSample `json:"samples"`
	Note     string         `json:"note,omitempty"`
	NextStep string         `json:"nextStep,omitempty"`
}

func (s *Server) viewSamples(ctx context.Context, _ *sdk.CallToolRequest, in ViewIn) (*sdk.CallToolResult, ViewOut, error) {
	if len(in.SampleIDs) == 0 {
		return nil, ViewOut{}, errors.New("sampleIds is required")
	}
	if len(in.SampleIDs) > maxViewSamples {
		return nil, ViewOut{}, fmt.Errorf("at most %d samples per call; page through the rest", maxViewSamples)
	}
	assets, err := s.sampleAssets(ctx, in.ProjectID, in.VersionID, in.SampleIDs)
	if err != nil {
		return nil, ViewOut{}, err
	}
	result := &sdk.CallToolResult{}
	out := ViewOut{Samples: []ViewedSample{}}
	missing := 0
	for _, id := range in.SampleIDs {
		vs := ViewedSample{ID: id, Visualizers: []string{}}
		files := assets[id]
		vs.Rendered = len(files) > 0
		if !vs.Rendered {
			missing++
		}
		for _, f := range files {
			dataType, visualizer := assetKind(f.Path)
			if in.Visualizer != "" && visualizer != in.Visualizer {
				continue
			}
			ext := strings.ToLower(path.Ext(f.Path))
			switch {
			case dataType == "image_heatmap" || dataType == "video_heatmap":
				if !containsNote(result.Content, id, visualizer) {
					vs.Visualizers = appendOnce(vs.Visualizers, visualizer)
					result.Content = append(result.Content, &sdk.TextContent{Text: fmt.Sprintf(
						"sample %s · %s (%s): the heatmap overlay is not available through the API yet; open the sample in the UI to see it", id, visualizer, dataType)})
				}
			case strings.Contains(f.Path, "/assets/") && (ext == ".jpg" || ext == ".jpeg" || ext == ".png"):
				raw, err := s.client.Download(ctx, f.URL)
				if err != nil {
					return nil, ViewOut{}, err
				}
				img, mime, err := prepareImage(raw, in.Full)
				if err != nil {
					continue
				}
				vs.Visualizers = appendOnce(vs.Visualizers, visualizer)
				result.Content = append(result.Content,
					&sdk.TextContent{Text: fmt.Sprintf("sample %s · %s (%s)", id, visualizer, dataType)},
					&sdk.ImageContent{Data: img, MIMEType: mime})
			case path.Base(f.Path) == "payload.json" && !strings.HasPrefix(dataType, "image"):
				raw, err := s.client.Download(ctx, f.URL)
				if err != nil {
					return nil, ViewOut{}, err
				}
				vs.Visualizers = appendOnce(vs.Visualizers, visualizer)
				result.Content = append(result.Content, &sdk.TextContent{
					Text: fmt.Sprintf("sample %s · %s (%s) payload: %s", id, visualizer, dataType, truncate(Scrub(string(raw)), 1500))})
			}
		}
		out.Samples = append(out.Samples, vs)
	}
	if missing > 0 {
		out.Note = fmt.Sprintf("%d sample(s) have no rendered visualizations yet", missing)
		out.NextStep = "pick samples marked rendered in tl_get_insights, or render more from the Population Exploration dashlet (Visualize)"
	}
	if len(result.Content) == 0 {
		result.Content = append(result.Content, &sdk.TextContent{Text: "no visualizations to show"})
	}
	return result, out, nil
}

func assetKind(p string) (string, string) {
	parts := strings.Split(p, "/")
	for i := len(parts) - 1; i >= 2; i-- {
		if parts[i] == "assets" || parts[i] == "payload.json" {
			return parts[i-2], parts[i-1]
		}
	}
	return "", ""
}

func appendOnce(l []string, s string) []string {
	for _, x := range l {
		if x == s {
			return l
		}
	}
	return append(l, s)
}

func prepareImage(raw []byte, full bool) ([]byte, string, error) {
	img, format, err := image.Decode(bytes.NewReader(raw))
	if err != nil {
		return nil, "", err
	}
	b := img.Bounds()
	if full || (b.Dx() <= thumbSize && b.Dy() <= thumbSize) {
		mime := "image/jpeg"
		if format == "png" {
			mime = "image/png"
		}
		return raw, mime, nil
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, downscale(img, thumbSize), &jpeg.Options{Quality: 85}); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), "image/jpeg", nil
}

func downscale(src image.Image, max int) image.Image {
	b := src.Bounds()
	scale := float64(max) / float64(b.Dx())
	if s := float64(max) / float64(b.Dy()); s < scale {
		scale = s
	}
	w, h := int(float64(b.Dx())*scale), int(float64(b.Dy())*scale)
	dst := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		y0, y1 := b.Min.Y+int(float64(y)/scale), b.Min.Y+int(float64(y+1)/scale)
		for x := 0; x < w; x++ {
			x0, x1 := b.Min.X+int(float64(x)/scale), b.Min.X+int(float64(x+1)/scale)
			var r, g, bl, a, n uint32
			for sy := y0; sy < y1 && sy < b.Max.Y; sy++ {
				for sx := x0; sx < x1 && sx < b.Max.X; sx++ {
					cr, cg, cb, ca := src.At(sx, sy).RGBA()
					r, g, bl, a, n = r+cr, g+cg, bl+cb, a+ca, n+1
				}
			}
			if n > 0 {
				dst.Set(x, y, color.RGBA64{uint16(r / n), uint16(g / n), uint16(bl / n), uint16(a / n)})
			}
		}
	}
	return dst
}

func containsNote(content []sdk.Content, id, visualizer string) bool {
	prefix := "sample " + id + " · " + visualizer + " "
	for _, c := range content {
		if t, ok := c.(*sdk.TextContent); ok && strings.HasPrefix(t.Text, prefix) {
			return true
		}
	}
	return false
}
