package mcp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUIBaseStripsTheCloudAPIHost(t *testing.T) {
	cases := map[string]string{
		"https://api.acme.tensorleap.ai/api/v2": "https://acme.tensorleap.ai",
		"http://localhost:4589/api/v2":          "http://localhost:4589",
		"https://tl.corp.internal/api/v2/":      "https://tl.corp.internal",
	}
	for in, want := range cases {
		if got := NewClient(in, "k").UIBase(); got != want {
			t.Errorf("UIBase(%s) = %s, want %s", in, got, want)
		}
	}
}

const engineInsight = `{"contractVersion":1,"deepLinkPath":"/p","predictionLabels":{"classes":["cat","dog"]},
 "visualizers":[{"name":"Image","type":"Image","argNames":["data"]}],
 "insights":[{"cid":"c1","index":1,"status":"InReview","insightType":{"id_":"i1","type":"low_performance","severity":2,"n_samples":500,
   "min_hash":[1,2],"display_filters":[{"metric":"x"}],"csv_path":"vis/a.csv","blob_path":"vis/b.json","top_panel_path":"vis/t.json",
   "metrics_info":[{"metric_name":"metrics.loss","metric_statistics":[{"name":"Cluster Average","value":0.9}]}],
   "is_train_aggressor":true,"overfitting_metrics":["metrics.loss"],"overfitting_evidence":[{"metric":"metrics.loss","contrast":2.4}],
   "mutual_info_elements":[{"features":[{"feature_name":"metadata.fog","feature_value":"yes","direction":"up","is_categorical":true}],"score":0.8}],
   "aggressor_fixing":{"num_of_samples_to_label":40,"num_of_samples_to_acquire":10,"csv_path":"vis/fix.csv"},
   "automatic_tests":[{"test_name":"Bad loss","filter":{"operator":"cluster","value":{"blob_paths":["organizations/x"]}},"metric_name":"metrics.loss","metric_value":0.5,"operator":"less_than"}]}}]}`

func TestInsightsCarryTheEnginePayloadLabelsAndVisualizers(t *testing.T) {
	s, _ := fakeServer(t, &AiAccess{Stats: true}, map[string]func(http.ResponseWriter){
		"analysis-export/exportAnalysis": func(w http.ResponseWriter) { _, _ = w.Write([]byte(engineInsight)) },
	})
	_, out, err := s.getInsights(context.Background(), nil, VersionIn{ProjectID: projectHex, VersionID: versionHex})
	if err != nil {
		t.Fatal(err)
	}
	if out.ClassLabels["classes"][1] != "dog" || len(out.Visualizers) != 1 || out.Visualizers[0].Name != "Image" {
		t.Fatalf("labels/visualizers missing: %+v", out)
	}
	eng := out.Insights[0].Engine
	if eng["is_train_aggressor"] != true || eng["metrics_info"] == nil || eng["overfitting_evidence"] == nil || eng["mutual_info_elements"] == nil {
		t.Fatalf("engine fields missing: %v", eng)
	}
	for _, k := range []string{"min_hash", "display_filters", "csv_path", "blob_path", "top_panel_path", "id_", "type"} {
		if _, ok := eng[k]; ok {
			t.Fatalf("internal key %q leaked", k)
		}
	}
	fix := eng["aggressor_fixing"].(map[string]any)
	if _, ok := fix["csv_path"]; ok || fix["num_of_samples_to_label"] != 40.0 {
		t.Fatalf("aggressor_fixing: %v", fix)
	}
	test := eng["automatic_tests"].([]any)[0].(map[string]any)
	if _, ok := test["filter"]; ok || test["metric_name"] != "metrics.loss" {
		t.Fatalf("automatic_tests should keep the condition and drop the filter: %v", test)
	}
}

func codeTarGz(t *testing.T, files map[string]string) []byte {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	for name, body := range files {
		if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o644, Size: int64(len(body)), Typeflag: tar.TypeReg}); err != nil {
			t.Fatal(err)
		}
		_, _ = tw.Write([]byte(body))
	}
	_ = tw.Close()
	_ = gz.Close()
	return buf.Bytes()
}

func codeServer(t *testing.T, code bool) *Server {
	t.Helper()
	archive := codeTarGz(t, map[string]string{
		"leap_integration.py": "def preprocess():\n    api_key = \"AKIAABCDEFGHIJKLMNOP\"\n    return []\n",
		"mnist/config.py":     "CLASSES = ['0', '1']\n",
		"../../etc/passwd":    "root:x:0:0\n",
	})
	var ts *httptest.Server
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/blob/code.tar.gz") {
			_, _ = w.Write(archive)
			return
		}
		switch strings.TrimPrefix(r.URL.Path, "/api/v2/") {
		case "analysis-export/listTargets":
			_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true,"code":` + map[bool]string{true: "true", false: "false"}[code] + `},"me":{}}`))
		case "analysis-export/exportAnalysis":
			_, _ = w.Write([]byte(`{"contractVersion":1,"deepLinkPath":"/p","insights":[],"integrationEntryFile":"leap_integration.py","integrationCodeUrl":"` + ts.URL + `/blob/code.tar.gz"}`))
		default:
			t.Fatalf("unexpected call %s", r.URL.Path)
		}
	}))
	t.Cleanup(ts.Close)
	return newServer(NewClient(ts.URL+"/api/v2", "k"))
}

func TestIntegrationCodeToolReadsTheEntryFileAndScrubs(t *testing.T) {
	s := codeServer(t, true)
	_, out, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex})
	if err != nil {
		t.Fatal(err)
	}
	if out.EntryFile != "leap_integration.py" || out.File != "leap_integration.py" || !strings.Contains(out.Content, "def preprocess") {
		t.Fatalf("entry file not returned: %+v", out)
	}
	if strings.Contains(out.Content, "AKIAABCDEFGHIJKLMNOP") {
		t.Fatal("secrets must be scrubbed from code")
	}
	names := []string{}
	for _, f := range out.Files {
		names = append(names, f.Path)
	}
	if strings.Join(names, ",") != "leap_integration.py,mnist/config.py" {
		t.Fatalf("file list (traversal member must be dropped): %v", names)
	}
	_, out, err = s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex, File: "mnist/config.py"})
	if err != nil || !strings.Contains(out.Content, "CLASSES") {
		t.Fatalf("named file: %v %+v", err, out)
	}
	_, out, _ = s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex, File: "nope.py"})
	if out.Reason != "file-not-found" {
		t.Fatalf("unknown file should explain, got %+v", out)
	}
}

func TestIntegrationCodeToolRespectsTheCodeClass(t *testing.T) {
	s := codeServer(t, false)
	_, _, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex})
	if err == nil || !strings.Contains(err.Error(), `"Integration code"`) {
		t.Fatalf("got %v", err)
	}
}
