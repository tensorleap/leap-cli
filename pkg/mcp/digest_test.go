package mcp

import (
	"fmt"
	"strings"
	"testing"
)

func buildCSV(rows int, row func(i int) string, header string) []byte {
	var b strings.Builder
	b.WriteString(header + "\n")
	for i := 0; i < rows; i++ {
		b.WriteString(row(i) + "\n")
	}
	return []byte(b.String())
}

func TestSummarizeUsesRootMembersAndBaseline(t *testing.T) {
	header := "sample_id,dataset_state,metrics.loss,metadata.weather,is_low_perf_root_member"
	population := buildCSV(1000, func(i int) string {
		weather := "day"
		if i%10 == 0 {
			weather = "night"
		}
		return fmt.Sprintf("training_%d,training,0.1,%s,", i, weather)
	}, header)
	insight := buildCSV(300, func(i int) string {
		root, weather, loss := "False", "day", "0.1"
		if i < 100 {
			root, weather, loss = "True", "night", "2.0"
		}
		return fmt.Sprintf("training_%d,training,%s,%s,%s", i, loss, weather, root)
	}, header)

	pop, err := NewPopulation(population)
	if err != nil {
		t.Fatal(err)
	}
	s, err := Summarize(insight, pop)
	if err != nil {
		t.Fatal(err)
	}
	if s.GroupSize != 100 || s.CsvRows != 300 {
		t.Fatalf("group should be the 100 root members out of 300 rows, got %d of %d", s.GroupSize, s.CsvRows)
	}
	if s.Split["training"] != 100 {
		t.Fatalf("split: %v", s.Split)
	}
	if len(s.Composition) == 0 || s.Composition[0].Value != "night" || s.Composition[0].InGroup != 1 || s.Composition[0].InBaseline != 0.1 {
		t.Fatalf("night should lead the composition at 100%% vs 10%%: %+v", s.Composition)
	}
	if len(s.Contrast) != 1 || s.Contrast[0].GroupMedian != 2 || s.Contrast[0].BaselineMedian != 0.1 {
		t.Fatalf("contrast: %+v", s.Contrast)
	}
}

func TestSummarizeSuppressesSmallCells(t *testing.T) {
	header := "sample_id,metadata.site,is_low_perf_root_member"
	population := buildCSV(1000, func(i int) string {
		site := "a"
		if i < 5 {
			site = "rare"
		}
		return fmt.Sprintf("training_%d,%s,", i, site)
	}, header)
	insight := buildCSV(20, func(i int) string {
		site := "a"
		if i < 5 {
			site = "rare"
		}
		return fmt.Sprintf("training_%d,%s,True", i, site)
	}, header)
	pop, _ := NewPopulation(population)
	s, err := Summarize(insight, pop)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range s.Composition {
		if c.Value == "rare" {
			t.Fatalf("a value covering fewer than %d group samples must not be reported: %+v", minCellSize, c)
		}
	}
}
