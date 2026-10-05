package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func catalogServer(t *testing.T) *Server {
	t.Helper()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		switch strings.TrimPrefix(r.URL.Path, "/api/v2/") {
		case "analysis-export/listTargets":
			if body["projectId"] == nil {
				_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true},"projects":[{"cid":"` + projectHex + `","name":"MNIST"}]}`))
				return
			}
			_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true},"versions":[
				{"cid":"aaaaaaaaaaaaaaaaaaaaaaaa","name":"baseline","createdAt":"2026-09-01T00:00:00Z","evaluated":true},
				{"cid":"` + versionHex + `","name":"augmented","createdAt":"2026-09-08T00:00:00Z","evaluated":true},
				{"cid":"bbbbbbbbbbbbbbbbbbbbbbbb","name":"draft","createdAt":"2026-09-09T00:00:00Z","evaluated":false}]}`))
		case "versions/getProjectSlimVersions":
			_, _ = w.Write([]byte(`{"versions":[{"cid":"` + versionHex + `","resources":{"inference_artifact_id":"ia","es_metrics_index":"idx"}}]}`))
		case "dashboards/getDashletFields":
			_, _ = w.Write([]byte(`{"aggregatableFields":["metrics.loss","metadata.label","dataset_state.keyword"],"numericFields":["metrics.loss"]}`))
		default:
			t.Fatalf("unexpected call %s", r.URL.Path)
		}
	}))
	t.Cleanup(ts.Close)
	return &Server{client: NewClient(ts.URL+"/api/v2", "k"), pops: map[string]*Population{}, fields: map[string]map[string]bool{}, policies: policyCache{entries: map[string]policyEntry{}}}
}

func TestResolveAcceptsNamesAndLatest(t *testing.T) {
	s := catalogServer(t)
	p, v := "mnist", "latest"
	if err := s.resolve(context.Background(), &p, &v); err != nil || p != projectHex || v != versionHex {
		t.Fatalf("got %s %s %v; latest must be the newest evaluated version", p, v, err)
	}
	p, v = "mnist", "Baseline"
	if err := s.resolve(context.Background(), &p, &v); err != nil || v != "aaaaaaaaaaaaaaaaaaaaaaaa" {
		t.Fatalf("version by name: %s %v", v, err)
	}
	p = "cifar"
	if err := s.resolve(context.Background(), &p); err == nil || !strings.Contains(err.Error(), "your projects: MNIST") {
		t.Fatalf("unknown project should list the real ones: %v", err)
	}
}

func TestQueryRejectsUnknownFieldsWithASuggestionAndAcceptsMean(t *testing.T) {
	s := catalogServer(t)
	in := QueryIn{ProjectID: "mnist", VersionIDs: []string{"latest"}, Measures: []Measure{{"metrics.los", "mean"}}}
	_, _, err := s.query(context.Background(), nil, in)
	if err == nil || !strings.Contains(err.Error(), `unknown field "metrics.los"`) || !strings.Contains(err.Error(), "did you mean metrics.loss") {
		t.Fatalf("got %v", err)
	}
}

func listServer(t *testing.T, sampleRows bool, honourValues bool) (*Server, *map[string]any) {
	t.Helper()
	var sent map[string]any
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/api/v2/")
		switch path {
		case "analysis-export/listTargets":
			_, _ = fmt.Fprintf(w, `{"contractVersion":1,"aiAccess":{"stats":true,"sampleRows":%v,"visuals":true},"me":{}}`, sampleRows)
		case "versions/getProjectSlimVersions":
			_, _ = w.Write([]byte(`{"versions":[{"cid":"` + versionHex + `","resources":{"inference_artifact_id":"ia","es_metrics_index":"idx"}}]}`))
		case "dashboards/getDashletFields":
			_, _ = w.Write([]byte(`{"aggregatableFields":["metrics.loss","metadata.label","metrics.pred"],"numericFields":["metrics.loss","metrics.pred"]}`))
		case "analysis-export/getSampleAssets":
			_, _ = w.Write([]byte(`{"samples":[{"sampleId":"` + visualizationID("training_4821") + `","files":[{"path":"p/image/assets/data.jpg","url":"u"}]}]}`))
		case "sample-collection/getVersionSampleOrder":
			_ = json.NewDecoder(r.Body).Decode(&sent)
			label := "2"
			if !honourValues {
				label = "5"
			}
			_, _ = w.Write([]byte(`{"total":36,"rows":[{"state":"training","index":4821,"metadata.label":"` + label + `","metrics.pred":7,"metrics.loss":6.4}]}`))
		default:
			t.Fatalf("unexpected call %s", path)
		}
	}))
	t.Cleanup(ts.Close)
	return &Server{client: NewClient(ts.URL+"/api/v2", "k"), pops: map[string]*Population{}, fields: map[string]map[string]bool{}, policies: policyCache{entries: map[string]policyEntry{}}}, &sent
}

func TestListSamplesFindsANamedCase(t *testing.T) {
	s, sent := listServer(t, true, true)
	in := ListSamplesIn{ProjectID: projectHex, VersionID: versionHex, SortBy: "metrics.loss",
		Filters: []Filter{{Field: "metadata.label", Operator: "equal", Value: "2"}, {Field: "metrics.pred", Operator: "equal", Value: 7}}}
	_, out, err := s.listSamples(context.Background(), nil, in)
	if err != nil || out.Matching != 36 || len(out.Samples) != 1 || out.Samples[0].ID != "training_4821" || out.Samples[0].Values["metrics.loss"] != 6.4 || out.Samples[0].Rendered == nil || !*out.Samples[0].Rendered {
		t.Fatalf("got %v %+v", err, out)
	}
	if f := (*sent)["filters"].([]any)[0].(map[string]any); f["values"].([]any)[0] != "2" {
		t.Fatalf("include filter not sent: %v", (*sent)["filters"])
	}
	if (*sent)["sort"].(map[string]any)["dir"] != "desc" {
		t.Fatal("highest loss first by default")
	}
}

func TestListSamplesRefusesToPassOffUnfilteredResults(t *testing.T) {
	s, _ := listServer(t, true, false)
	in := ListSamplesIn{ProjectID: projectHex, VersionID: versionHex, Filters: []Filter{{Field: "metadata.label", Operator: "equal", Value: "2"}}}
	if _, _, err := s.listSamples(context.Background(), nil, in); err == nil || !strings.Contains(err.Error(), "can't list samples by value") {
		t.Fatalf("an old server ignoring the filter must be reported, got %v", err)
	}
}

func TestListSamplesNeedsPerSampleData(t *testing.T) {
	s, _ := listServer(t, false, true)
	_, _, err := s.listSamples(context.Background(), nil, ListSamplesIn{ProjectID: projectHex, VersionID: versionHex})
	if err == nil || !strings.Contains(err.Error(), `"Per-sample data"`) {
		t.Fatalf("got %v", err)
	}
}
