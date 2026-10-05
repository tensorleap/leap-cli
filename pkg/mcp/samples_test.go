package mcp

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func TestPrepareImageMakesThumbnails(t *testing.T) {
	src := image.NewRGBA(image.Rect(0, 0, 1920, 1080))
	for x := 0; x < 1920; x++ {
		src.Set(x, 500, color.White)
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, src); err != nil {
		t.Fatal(err)
	}
	out, mime, err := prepareImage(buf.Bytes(), false)
	if err != nil {
		t.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(out))
	if err != nil {
		t.Fatal(err)
	}
	if b := img.Bounds(); b.Dx() != 640 || b.Dy() != 360 || mime != "image/jpeg" {
		t.Fatalf("thumbnail %dx%d %s, want 640x360 jpeg", b.Dx(), b.Dy(), mime)
	}
	full, mime, _ := prepareImage(buf.Bytes(), true)
	if !bytes.Equal(full, buf.Bytes()) || mime != "image/png" {
		t.Fatal("full resolution must return the original bytes")
	}
}

func TestAssetKind(t *testing.T) {
	dt, vis := assetKind("organizations/t/projects/p/vis/a/sample_visualizers/validation_X/image_heatmap/default_image_visualizer_heatmap/assets/data.jpg")
	if dt != "image_heatmap" || vis != "default_image_visualizer_heatmap" {
		t.Fatalf("got %s %s", dt, vis)
	}
}

func TestPrepareImageEnlargesTinySamples(t *testing.T) {
	src := image.NewGray(image.Rect(0, 0, 28, 28))
	src.SetGray(0, 0, color.Gray{255})
	var buf bytes.Buffer
	if err := png.Encode(&buf, src); err != nil {
		t.Fatal(err)
	}
	out, _, err := prepareImage(buf.Bytes(), true)
	if err != nil {
		t.Fatal(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(out))
	if b := img.Bounds(); b.Dx() != 224 || b.Dy() != 224 {
		t.Fatalf("got %dx%d, want 224x224 (8x nearest neighbour)", b.Dx(), b.Dy())
	}
	if r, _, _, _ := img.At(7, 7).RGBA(); r>>8 < 200 {
		t.Fatal("the bright source pixel must cover the first 8x8 block")
	}
}
