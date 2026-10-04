package mcp

import (
	"bytes"
	"encoding/csv"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

const (
	maxCategorical = 12
	minCellSize    = 10
	maxComposition = 8
)

var skipColumns = map[string]bool{
	"sample_id": true, "epoch": true, "batch": true, "dataset_name": true,
	"dataset_version_id": true, "inserted_at": true, "is_low_perf_root_member": true, "dataset_state": true,
}

type table struct {
	header []string
	rows   [][]string
}

func parseCSV(b []byte) (*table, error) {
	r := csv.NewReader(bytes.NewReader(b))
	r.FieldsPerRecord = -1
	header, err := r.Read()
	if err != nil {
		return nil, err
	}
	t := &table{header: header}
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		t.rows = append(t.rows, rec)
	}
	return t, nil
}

func (t *table) col(name string) int {
	for i, h := range t.header {
		if h == name {
			return i
		}
	}
	return -1
}

func analysedColumn(name string) bool {
	return !skipColumns[name] && !strings.HasPrefix(name, "metadata_is_none.") &&
		(strings.HasPrefix(name, "metrics.") || strings.HasPrefix(name, "metadata."))
}

type columnStats struct {
	numeric bool
	values  []float64
	counts  map[string]int
	n       int
}

func (c *columnStats) mean() float64 {
	s := 0.0
	for _, v := range c.values {
		s += v
	}
	return s / float64(len(c.values))
}

func (c *columnStats) median() float64 {
	v := append([]float64(nil), c.values...)
	sort.Float64s(v)
	m := len(v) / 2
	if len(v)%2 == 1 {
		return v[m]
	}
	return (v[m-1] + v[m]) / 2
}

func (c *columnStats) categorical() bool {
	return !c.numeric || len(c.counts) <= maxCategorical
}

func statsFor(t *table, rows [][]string) map[string]*columnStats {
	out := map[string]*columnStats{}
	for i, name := range t.header {
		if !analysedColumn(name) {
			continue
		}
		cs := &columnStats{numeric: true, counts: map[string]int{}}
		for _, r := range rows {
			if i >= len(r) || r[i] == "" {
				continue
			}
			cs.n++
			cs.counts[r[i]]++
			if f, err := strconv.ParseFloat(r[i], 64); err == nil && !math.IsNaN(f) {
				cs.values = append(cs.values, f)
			} else {
				cs.numeric = false
			}
		}
		if cs.n > 0 {
			out[name] = cs
		}
	}
	return out
}

type Population struct {
	Total int
	stats map[string]*columnStats
}

func NewPopulation(csvBytes []byte) (*Population, error) {
	t, err := parseCSV(csvBytes)
	if err != nil {
		return nil, err
	}
	return &Population{Total: len(t.rows), stats: statsFor(t, t.rows)}, nil
}

type Composition struct {
	Field      string  `json:"field"`
	Value      string  `json:"value,omitempty"`
	InGroup    float64 `json:"inGroup"`
	InBaseline float64 `json:"inBaseline"`
	Ratio      float64 `json:"ratio"`
	Kind       string  `json:"kind" jsonschema:"share = fraction of samples with this value; mean = average of a numeric field"`
	Semantics  string  `json:"semantics"`
}

type Contrast struct {
	Metric         string  `json:"metric"`
	GroupMedian    float64 `json:"groupMedian"`
	BaselineMedian float64 `json:"baselineMedian"`
	GroupMean      float64 `json:"groupMean"`
	BaselineMean   float64 `json:"baselineMean"`
	Baseline       string  `json:"baseline"`
}

type GroupSummary struct {
	GroupSize       int            `json:"groupSize"`
	GroupDefinition string         `json:"groupDefinition"`
	CsvRows         int            `json:"csvRows"`
	Split           map[string]int `json:"split"`
	Composition     []Composition  `json:"composition"`
	Contrast        []Contrast     `json:"contrast"`
	MemberSampleIDs []string       `json:"-"`
}

func Summarize(csvBytes []byte, pop *Population) (*GroupSummary, error) {
	t, err := parseCSV(csvBytes)
	if err != nil {
		return nil, err
	}
	group := t.rows
	definition := "every row of the insight's sample list"
	if rc := t.col("is_low_perf_root_member"); rc >= 0 {
		var root [][]string
		for _, r := range t.rows {
			if rc < len(r) && strings.EqualFold(r[rc], "true") {
				root = append(root, r)
			}
		}
		if len(root) > 0 {
			group = root
			definition = "samples that actually underperform (root members); the platform's n_samples also counts healthy latent neighbours"
		}
	}
	s := &GroupSummary{GroupSize: len(group), GroupDefinition: definition, CsvRows: len(t.rows), Split: map[string]int{}}
	stateCol, idCol := t.col("dataset_state"), t.col("sample_id")
	for _, r := range group {
		state := ""
		if stateCol >= 0 && stateCol < len(r) {
			state = r[stateCol]
		}
		if state == "" && idCol >= 0 && idCol < len(r) {
			if i := strings.LastIndex(r[idCol], "_"); i > 0 {
				state = r[idCol][:i]
			}
		}
		if state == "" {
			state = "unknown"
		}
		s.Split[state]++
		if idCol >= 0 && idCol < len(r) {
			s.MemberSampleIDs = append(s.MemberSampleIDs, r[idCol])
		}
	}
	gstats := statsFor(t, group)
	names := make([]string, 0, len(gstats))
	for n := range gstats {
		names = append(names, n)
	}
	sort.Strings(names)
	for _, name := range names {
		g := gstats[name]
		var base *columnStats
		if pop != nil {
			base = pop.stats[name]
		}
		if strings.HasPrefix(name, "metrics.") {
			if !g.numeric || len(g.values) == 0 {
				continue
			}
			c := Contrast{Metric: name, GroupMedian: round4(g.median()), GroupMean: round4(g.mean()), Baseline: "all data"}
			if base != nil && base.numeric && len(base.values) > 0 {
				c.BaselineMedian, c.BaselineMean = round4(base.median()), round4(base.mean())
			}
			s.Contrast = append(s.Contrast, c)
			continue
		}
		if base == nil {
			continue
		}
		if base.categorical() {
			for value, count := range g.counts {
				if count < minCellSize {
					continue
				}
				in := float64(count) / float64(len(group))
				out := float64(base.counts[value]) / float64(pop.Total)
				if out == 0 {
					continue
				}
				if ratio := in / out; ratio >= 1.25 {
					s.Composition = append(s.Composition, Composition{Field: name, Value: value, InGroup: round4(in), InBaseline: round4(out),
						Ratio: round4(ratio), Kind: "share", Semantics: "over-represented in the group, not necessarily its defining trait"})
				}
			}
		} else if g.numeric && base.numeric && len(base.values) > 0 && len(g.values) > 0 {
			gm, bm := g.mean(), base.mean()
			if bm != 0 {
				if ratio := gm / bm; ratio >= 1.25 || ratio <= 0.8 {
					s.Composition = append(s.Composition, Composition{Field: name, InGroup: round4(gm), InBaseline: round4(bm),
						Ratio: round4(ratio), Kind: "mean", Semantics: "group average differs from all data"})
				}
			}
		}
	}
	weight := func(c Composition) float64 {
		w := math.Abs(math.Log(c.Ratio))
		if c.Kind == "share" {
			w *= c.InGroup
		}
		return w
	}
	sort.Slice(s.Composition, func(i, j int) bool { return weight(s.Composition[i]) > weight(s.Composition[j]) })
	if len(s.Composition) > maxComposition {
		s.Composition = s.Composition[:maxComposition]
	}
	return s, nil
}

func round4(f float64) float64 {
	return math.Round(f*10000) / 10000
}
