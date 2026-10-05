package mcp

import (
	"encoding/binary"
	"errors"
	"image"
	"image/color"
	"math"
)

type heatmap struct {
	w, h int
	data []float32
}

// parseHeatmap decodes the engine's imageHeatmap.HeatmapData flatbuffer
// (table { data:[float]; width:int; height:int; }), the same file web-ui renders.
func parseHeatmap(b []byte) (*heatmap, error) {
	bad := errors.New("not a heatmap flatbuffer")
	u32 := func(at int) (int, bool) {
		if at < 0 || at+4 > len(b) {
			return 0, false
		}
		return int(binary.LittleEndian.Uint32(b[at:])), true
	}
	table, ok := u32(0)
	if !ok {
		return nil, bad
	}
	soff, ok := u32(table)
	if !ok {
		return nil, bad
	}
	vtable := table - int(int32(soff))
	if vtable < 0 || vtable+4 > len(b) {
		return nil, bad
	}
	vtSize := int(binary.LittleEndian.Uint16(b[vtable:]))
	field := func(i int) int {
		at := vtable + 4 + 2*i
		if 4+2*i >= vtSize || at+2 > len(b) {
			return 0
		}
		return int(binary.LittleEndian.Uint16(b[at:]))
	}
	f0, f1, f2 := field(0), field(1), field(2)
	if f0 == 0 || f1 == 0 || f2 == 0 {
		return nil, bad
	}
	rel, _ := u32(table + f0)
	vec := table + f0 + rel
	n, ok := u32(vec)
	w, ok1 := u32(table + f1)
	h, ok2 := u32(table + f2)
	if !ok || !ok1 || !ok2 || w <= 0 || h <= 0 || w > 10000 || h > 10000 || n != w*h || vec+4+4*n > len(b) {
		return nil, bad
	}
	data := make([]float32, n)
	for i := range data {
		data[i] = math.Float32frombits(binary.LittleEndian.Uint32(b[vec+4+4*i:]))
	}
	return &heatmap{w: w, h: h, data: data}, nil
}

// overlay blends the heatmap over base the way web-ui shows it by default:
// range clipped to 5%-95%, turbo colormap, 50% opacity, bilinear resize.
func overlay(base image.Image, hm *heatmap) image.Image {
	lo, hi := float32(math.Inf(1)), float32(math.Inf(-1))
	for _, v := range hm.data {
		lo, hi = min(lo, v), max(hi, v)
	}
	span := hi - lo
	lo, hi = lo+span*0.05, lo+span*0.95
	if hi <= lo {
		hi = lo + 1
	}
	b := base.Bounds()
	dst := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	sx, sy := float64(hm.w)/float64(b.Dx()), float64(hm.h)/float64(b.Dy())
	for y := 0; y < b.Dy(); y++ {
		for x := 0; x < b.Dx(); x++ {
			v := hm.sample((float64(x)+0.5)*sx-0.5, (float64(y)+0.5)*sy-0.5)
			t := math.Max(0, math.Min(1, float64((v-lo)/(hi-lo))))
			hr, hg, hb := turbo(t)
			r, g, bl, _ := base.At(b.Min.X+x, b.Min.Y+y).RGBA()
			dst.Set(x, y, color.RGBA{blend(r, hr), blend(g, hg), blend(bl, hb), 255})
		}
	}
	return dst
}

func blend(base uint32, heat float64) uint8 {
	return uint8(math.Round(float64(base>>8)*0.5 + heat*0.5))
}

func (hm *heatmap) sample(x, y float64) float32 {
	x = math.Max(0, math.Min(float64(hm.w-1), x))
	y = math.Max(0, math.Min(float64(hm.h-1), y))
	x0, y0 := int(x), int(y)
	x1, y1 := min(x0+1, hm.w-1), min(y0+1, hm.h-1)
	fx, fy := float32(x-float64(x0)), float32(y-float64(y0))
	at := func(xx, yy int) float32 { return hm.data[yy*hm.w+xx] }
	top := at(x0, y0)*(1-fx) + at(x1, y0)*fx
	bottom := at(x0, y1)*(1-fx) + at(x1, y1)*fx
	return top*(1-fy) + bottom*fy
}

// turbo is d3.interpolateTurbo (Mikhailov's polynomial approximation), as used by web-ui.
func turbo(t float64) (float64, float64, float64) {
	c := func(v float64) float64 { return math.Max(0, math.Min(255, math.Round(v))) }
	return c(34.61 + t*(1172.33-t*(10793.56-t*(33300.12-t*(38394.49-t*14825.05))))),
		c(23.31 + t*(557.33+t*(1225.33-t*(3574.96-t*(1073.77+t*707.56))))),
		c(27.2 + t*(3211.1-t*(15327.97-t*(27814-t*(22569.18-t*6838.66)))))
}
