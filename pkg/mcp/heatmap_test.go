package mcp

import (
	"encoding/binary"
	"encoding/json"
	"image"
	"image/color"
	"math"
	"testing"
)

func heatmapBuffer(w, h int, data []float32) []byte {
	b := make([]byte, 36+4*len(data))
	le := binary.LittleEndian
	le.PutUint32(b[0:], 16)
	for i, v := range []uint16{10, 16, 4, 8, 12} {
		le.PutUint16(b[4+2*i:], v)
	}
	le.PutUint32(b[16:], 12)
	le.PutUint32(b[20:], 12)
	le.PutUint32(b[24:], uint32(w))
	le.PutUint32(b[28:], uint32(h))
	le.PutUint32(b[32:], uint32(len(data)))
	for i, v := range data {
		le.PutUint32(b[36+4*i:], math.Float32bits(v))
	}
	return b
}

func TestParseHeatmap(t *testing.T) {
	hm, err := parseHeatmap(heatmapBuffer(2, 1, []float32{0, 1}))
	if err != nil || hm.w != 2 || hm.h != 1 || hm.data[1] != 1 {
		t.Fatalf("got %+v %v", hm, err)
	}
	if _, err := parseHeatmap(heatmapBuffer(3, 1, []float32{0, 1})); err == nil {
		t.Fatal("size mismatch must be rejected")
	}
	if _, err := parseHeatmap(heatmapBuffer(2, 1, []float32{0, 1})[:30]); err == nil {
		t.Fatal("truncated buffer must be rejected")
	}
}

func TestTurboMatchesD3(t *testing.T) {
	r, g, b := turbo(0.2)
	if r != 47 || g != 157 || b != 245 {
		t.Fatalf("turbo(0.2) = %v %v %v, d3 gives rgb(47, 157, 245)", r, g, b)
	}
	if r, g, b := turbo(0.85); r != 219 || g != 61 || b != 17 {
		t.Fatalf("turbo(0.85) = %v %v %v, d3 gives rgb(219, 61, 17)", r, g, b)
	}
}

func TestOverlayBlendsHalfAndKeepsSize(t *testing.T) {
	base := image.NewGray(image.Rect(0, 0, 20, 10))
	for i := range base.Pix {
		base.Pix[i] = 100
	}
	out := overlay(base, &heatmap{w: 2, h: 1, data: []float32{0, 1}})
	if out.Bounds() != base.Bounds() {
		t.Fatal("overlay must keep the image size")
	}
	r, g, b := turbo(0)
	half := func(v float64) uint8 { return uint8(math.Round(50 + v/2)) }
	want := color.RGBA{half(r), half(g), half(b), 255}
	if got := out.At(0, 5).(color.RGBA); got != want {
		t.Fatalf("cold edge: got %v want %v", got, want)
	}
}

func TestURLForMatchesProjectRelativeBlob(t *testing.T) {
	files := []assetFile{{Path: "org/projects/p/vis/v/image_heatmap/assets/h.bin", URL: "u"}}
	if urlFor(files, "vis/v/image_heatmap/assets/h.bin") != "u" || urlFor(files, "assets/x.bin") != "" {
		t.Fatal("blob lookup")
	}
}

func TestCollectHeatmapsFromEnginePayload(t *testing.T) {
	var payload any
	_ = json.Unmarshal([]byte(`{"connection_name":"v","data":{"data":[
		{"label":"cat","data":{"blob":"a.jpg","heatmap_blob":"h1.bin","type":"image_heatmap"}},
		{"label":"","data":{"blob":"a.jpg","heatmap_blob":"h2.bin"}}]}}`), &payload)
	var items []heatmapItem
	collectHeatmaps(payload, "", &items)
	if len(items) != 2 || items[0] != (heatmapItem{"cat", "a.jpg", "h1.bin"}) || items[1].HeatmapBlob != "h2.bin" {
		t.Fatalf("got %+v", items)
	}
}
