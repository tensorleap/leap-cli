package mcp

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func fakeServer(t *testing.T, aiAccess *AiAccess, routes map[string]func(w http.ResponseWriter)) (*Server, *int) {
	t.Helper()
	calls := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		path := strings.TrimPrefix(r.URL.Path, "/api/v2/")
		if path == "analysis-export/listTargets" {
			body := map[string]any{"contractVersion": 1, "me": map[string]string{"email": "a@b.c", "role": "user"}}
			if aiAccess != nil {
				body["aiAccess"] = aiAccess
			}
			_ = json.NewEncoder(w).Encode(body)
			return
		}
		if h, ok := routes[path]; ok {
			h(w)
			return
		}
		t.Fatalf("unexpected call to %s", path)
	}))
	t.Cleanup(ts.Close)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "key"), pops: map[string]*Population{}, policies: policyCache{entries: map[string]policyEntry{}}}
	return s, &calls
}

func TestLegacyServerRefusedUnlessAllowed(t *testing.T) {
	s, _ := fakeServer(t, nil, map[string]func(http.ResponseWriter){
		"sessionmetrics/getTableChart": func(w http.ResponseWriter) { _, _ = w.Write([]byte(`{"charts":[]}`)) },
	})
	in := QueryIn{ProjectID: "p", VersionIDs: []string{"v"}, Measures: []Measure{{"metrics.loss", "Average"}}}
	if _, _, err := s.query(context.Background(), nil, in); err == nil || !strings.Contains(err.Error(), "predates AI access controls") {
		t.Fatalf("expected the legacy-server refusal, got %v", err)
	}
	s.allowLegacy = true
	s.policies.entries = map[string]policyEntry{}
	if _, out, err := s.query(context.Background(), nil, in); err != nil || out.Reason != "no-rows" {
		t.Fatalf("allowed legacy server should answer: %v %+v", err, out)
	}
}

func TestStatsOffBlocksQueriesWithoutCallingTheServer(t *testing.T) {
	s, calls := fakeServer(t, &AiAccess{}, nil)
	in := QueryIn{ProjectID: "p", VersionIDs: []string{"v"}, Measures: []Measure{{"metrics.loss", "Average"}}}
	_, _, err := s.query(context.Background(), nil, in)
	if err == nil || !strings.Contains(err.Error(), "turned off for this project") {
		t.Fatalf("expected a policy refusal, got %v", err)
	}
	if *calls != 1 {
		t.Fatalf("only the policy lookup may reach the server, got %d calls", *calls)
	}
	if _, _, err := s.describeFields(context.Background(), nil, VersionIn{ProjectID: "p", VersionID: "v"}); err == nil {
		t.Fatal("describe_fields must be refused too")
	}
}

func TestServerRefusalMessageIsSurfaced(t *testing.T) {
	s, _ := fakeServer(t, &AiAccess{Stats: true}, map[string]func(http.ResponseWriter){
		"analysis-export/exportAnalysis": func(w http.ResponseWriter) {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write([]byte(`{"error":"AI access to statistics and insights is turned off for this project.","code":"AI_ACCESS_DISABLED"}`))
		},
	})
	_, _, err := s.getInsights(context.Background(), nil, VersionIn{ProjectID: "p", VersionID: "v"})
	if err == nil || err.Error() != "AI access to statistics and insights is turned off for this project." {
		t.Fatalf("got %v", err)
	}
}

func TestInsightsWithoutSampleRowsExplainGroupSize(t *testing.T) {
	s, _ := fakeServer(t, &AiAccess{Stats: true}, map[string]func(http.ResponseWriter){
		"analysis-export/exportAnalysis": func(w http.ResponseWriter) {
			_, _ = w.Write([]byte(`{"contractVersion":1,"deepLinkPath":"/p","insights":[{"cid":"c","index":1,"status":"InReview",
				"insightType":{"type":"low_performance","severity":2,"n_samples":500}}]}`))
		},
	})
	_, out, err := s.getInsights(context.Background(), nil, VersionIn{ProjectID: "p", VersionID: "v"})
	if err != nil {
		t.Fatal(err)
	}
	ins := out.Insights[0]
	if ins.Name != "Failure Mode" || ins.ClusterSize != 500 || !strings.Contains(ins.Warning, "per-sample data is off") {
		t.Fatalf("got %+v", ins)
	}
}

func TestQueryWithoutGroupByCoversWholePopulation(t *testing.T) {
	var sent map[string]any
	s, _ := fakeServer(t, &AiAccess{Stats: true}, nil)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "listTargets") {
			_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true},"me":{}}`))
			return
		}
		_ = json.NewDecoder(r.Body).Decode(&sent)
		_, _ = w.Write([]byte(`{"charts":[{"data":{"data":[{"data":{"model.extId.keyword":"m","metrics.loss":0.5,"sample_id":70000}}]}}]}`))
	}))
	defer ts.Close()
	s.client = NewClient(ts.URL+"/api/v2", "key")
	_, out, err := s.query(context.Background(), nil, QueryIn{ProjectID: "p", VersionIDs: []string{"v"}, Measures: []Measure{{"metrics.loss", "Average"}}})
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Rows) != 1 || out.Rows[0].N == nil || *out.Rows[0].N != 70000 || out.Rows[0].Group != nil {
		t.Fatalf("want one whole-population row with n=70000 and no group, got %+v", out.Rows)
	}
	if b := sent["buckets"].([]any); len(b) != 1 || b[0].(map[string]any)["field"] != "model.extId.keyword" {
		t.Fatalf("request buckets: %v", sent["buckets"])
	}
}
