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
	s, err := Summarize(insight, pop, nil)
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
	s, err := Summarize(insight, pop, nil)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range s.Composition {
		if c.Value == "rare" {
			t.Fatalf("a value covering fewer than %d group samples must not be reported: %+v", minCellSize, c)
		}
	}
}

func TestVisualizationIDMatchesSkillHash(t *testing.T) {
	if got := visualizationID("training_4821"); got != "training_o4j1YuKG_fKJhvklNXn00A" {
		t.Fatalf("got %s", got)
	}
}

func TestRankSamplesByAffinityWithinGroup(t *testing.T) {
	csv := []byte("sample_id,aggressor_affinity_score,is_low_perf_root_member\ntraining_1,0.2,True\ntraining_2,0.9,True\ntraining_3,0.99,False\ntraining_4,0.5,True\n")
	s, err := Summarize(csv, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(s.RankedIDs, ",") != "training_2,training_4,training_1" || s.RankedBy != "aggressor_affinity_score" {
		t.Fatalf("ranked %v by %s; a healthy neighbour must never be ranked", s.RankedIDs, s.RankedBy)
	}
}

func TestDuplicateInsightsUseTheirOwnMembersAndKeepPairsTogether(t *testing.T) {
	shared := "sample_id,metrics.loss,duplication_ids_x\ntraining_1,0,7\ntraining_2,0,9\nvalidation_5,0,7\ntraining_3,0,9\nunlabeled_4,0,4\n"
	members := clusterMembers([]byte(`{"samples_index":{"training":[1,3],"validation":[5]}}`))
	s, err := Summarize([]byte(shared), nil, members)
	if err != nil {
		t.Fatal(err)
	}
	if s.GroupSize != 3 || s.Split["unlabeled"] != 0 || !strings.Contains(s.GroupDefinition, "own members") {
		t.Fatalf("leakage must describe its own members, not the shared list: %+v", s)
	}
	if strings.Join(s.RankedIDs, ",") != "training_1,validation_5,training_3" {
		t.Fatalf("pairs must stay adjacent, training first: %v", s.RankedIDs)
	}
}
