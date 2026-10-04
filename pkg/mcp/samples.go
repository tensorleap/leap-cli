package mcp

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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
	maxViewSamples  = 6
	thumbSize       = 640
	minViewSize     = 224
	defaultHeatmaps = 3
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
	Labels     []string `json:"heatmapLabels,omitempty" jsonschema:"heatmap labels (e.g. class names) to overlay; default: the first 3"`
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
			case dataType == "image_heatmap" && path.Base(f.Path) == "payload.json":
				content, err := s.heatmapOverlays(ctx, id, visualizer, f.URL, files, in)
				if err != nil {
					return nil, ViewOut{}, err
				}
				vs.Visualizers = appendOnce(vs.Visualizers, visualizer)
				result.Content = append(result.Content, content...)
			case dataType == "image_heatmap":
			case dataType == "video_heatmap":
				if !containsNote(result.Content, id, visualizer) {
					vs.Visualizers = appendOnce(vs.Visualizers, visualizer)
					result.Content = append(result.Content, &sdk.TextContent{Text: fmt.Sprintf(
						"sample %s · %s (%s): video heatmaps are not available through the API yet; open the sample in the UI to see it", id, visualizer, dataType)})
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

type heatmapItem struct {
	Label, Blob, HeatmapBlob string
}

func (s *Server) heatmapOverlays(ctx context.Context, id, visualizer, payloadURL string, files []assetFile, in ViewIn) ([]sdk.Content, error) {
	raw, err := s.client.Download(ctx, payloadURL)
	if err != nil {
		return nil, err
	}
	var payload any
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, fmt.Errorf("sample %s: unreadable heatmap payload: %w", id, err)
	}
	var items []heatmapItem
	collectHeatmaps(payload, "", &items)
	var chosen, skipped []heatmapItem
	for _, it := range items {
		wanted := len(in.Labels) == 0 && len(chosen) < defaultHeatmaps
		for _, l := range in.Labels {
			wanted = wanted || l == it.Label
		}
		if wanted {
			chosen = append(chosen, it)
		} else {
			skipped = append(skipped, it)
		}
	}
	var out []sdk.Content
	bases := map[string]image.Image{}
	for _, it := range chosen {
		baseURL, heatURL := urlFor(files, it.Blob), urlFor(files, it.HeatmapBlob)
		if baseURL == "" {
			continue
		}
		base, ok := bases[baseURL]
		if !ok {
			rawImg, err := s.client.Download(ctx, baseURL)
			if err != nil {
				return nil, err
			}
			if base, _, err = image.Decode(bytes.NewReader(rawImg)); err != nil {
				continue
			}
			base = viewSize(base, in.Full)
			bases[baseURL] = base
		}
		kind := "image_heatmap"
		if it.Label != "" {
			kind += fmt.Sprintf(", label %q", it.Label)
		}
		caption := fmt.Sprintf("sample %s · %s (%s): heatmap over the image, turbo colormap (red = strongest attention, blue = weakest)", id, visualizer, kind)
		img := base
		if heatURL == "" {
			caption = fmt.Sprintf("sample %s · %s (%s): base image only; this server does not return heatmap data, open the sample in the UI to see the overlay", id, visualizer, kind)
		} else {
			rawHeat, err := s.client.Download(ctx, heatURL)
			if err != nil {
				return nil, err
			}
			hm, err := parseHeatmap(rawHeat)
			if err != nil {
				return nil, fmt.Errorf("sample %s (%s): %w", id, kind, err)
			}
			img = overlay(base, hm)
		}
		var buf bytes.Buffer
		if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: 85}); err != nil {
			return nil, err
		}
		out = append(out, &sdk.TextContent{Text: caption}, &sdk.ImageContent{Data: buf.Bytes(), MIMEType: "image/jpeg"})
	}
	if len(skipped) > 0 {
		labels := make([]string, 0, len(skipped))
		for _, it := range skipped {
			labels = append(labels, it.Label)
		}
		out = append(out, &sdk.TextContent{Text: fmt.Sprintf("sample %s · %s: %d more heatmap label(s) not shown (%s); pass heatmapLabels to choose", id, visualizer, len(skipped), truncate(strings.Join(labels, ", "), 300))})
	}
	return out, nil
}

// collectHeatmaps finds every {blob, heatmap_blob} pair, whatever the payload nesting,
// labelled by the closest enclosing "label"
func collectHeatmaps(v any, label string, out *[]heatmapItem) {
	switch t := v.(type) {
	case []any:
		for _, x := range t {
			collectHeatmaps(x, label, out)
		}
	case map[string]any:
		if l, ok := t["label"].(string); ok && l != "" {
			label = l
		}
		if hb, ok := t["heatmap_blob"].(string); ok {
			blob, _ := t["blob"].(string)
			*out = append(*out, heatmapItem{Label: label, Blob: blob, HeatmapBlob: hb})
			return
		}
		for _, x := range t {
			collectHeatmaps(x, label, out)
		}
	}
}

func urlFor(files []assetFile, rel string) string {
	if rel == "" {
		return ""
	}
	for _, f := range files {
		if f.Path == rel || strings.HasSuffix(f.Path, "/"+rel) {
			return f.URL
		}
	}
	return ""
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
	small := max(b.Dx(), b.Dy()) < minViewSize
	if !small && (full || (b.Dx() <= thumbSize && b.Dy() <= thumbSize)) {
		mime := "image/jpeg"
		if format == "png" {
			mime = "image/png"
		}
		return raw, mime, nil
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, viewSize(img, full), &jpeg.Options{Quality: 85}); err != nil {
		return nil, "", err
	}
	return buf.Bytes(), "image/jpeg", nil
}

func viewSize(img image.Image, full bool) image.Image {
	b := img.Bounds()
	if longest := max(b.Dx(), b.Dy()); longest > 0 && longest < minViewSize {
		return upscale(img, (minViewSize+longest-1)/longest)
	}
	if !full && (b.Dx() > thumbSize || b.Dy() > thumbSize) {
		return downscale(img, thumbSize)
	}
	return img
}

// upscale enlarges tiny samples (e.g. 28px MNIST digits) so a vision model can make them out;
// nearest neighbour keeps the original pixels visible
func upscale(src image.Image, k int) image.Image {
	b := src.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx()*k, b.Dy()*k))
	for y := 0; y < b.Dy()*k; y++ {
		for x := 0; x < b.Dx()*k; x++ {
			dst.Set(x, y, src.At(b.Min.X+x/k, b.Min.Y+y/k))
		}
	}
	return dst
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
