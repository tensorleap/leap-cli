package mcp

import (
	"context"
	"encoding/json"
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
