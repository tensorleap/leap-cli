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
 "insights":[{"cid":"c1","index":1,"status":"InReview","csvUrl":"http://x/blob",
   "insightType":{"id_":"i1","type":"low_performance","severity":2,"n_samples":500,"csv_path":"vis/a.csv"},
   "engine":{"n_samples":500,"is_train_aggressor":true,"metrics_info":[{"metric_name":"metrics.loss"}],"overfitting_evidence":[{"metric_name":"metrics.loss","score":2.4}]},
   "digest":{"groupSize":120,"groupDefinition":"root members","csvRows":500,"split":{"training":120},"composition":[],
     "contrast":[{"metric":"metrics.loss","groupMedian":0.9,"groupMean":1,"baselineMedian":0.1,"baselineMean":0.2,"baseline":"all data"}],
     "rankedBy":"aggressor_affinity_score","rankedSampleIds":["training_1"]}}]}`

func TestInsightsReadTheServerDigestAndEnginePayload(t *testing.T) {
	s, _ := fakeServer(t, &AiAccess{Stats: true, SampleRows: true}, map[string]func(http.ResponseWriter){
		"analysis-export/exportAnalysis": func(w http.ResponseWriter) { _, _ = w.Write([]byte(engineInsight)) },
	})
	_, out, err := s.getInsights(context.Background(), nil, VersionIn{ProjectID: projectHex, VersionID: versionHex})
	if err != nil {
		t.Fatal(err)
	}
	if out.ClassLabels["classes"][1] != "dog" || len(out.Visualizers) != 1 || out.Visualizers[0].Name != "Image" {
		t.Fatalf("labels/visualizers missing: %+v", out)
	}
	ins := out.Insights[0]
	if ins.Engine["is_train_aggressor"] != true || ins.Engine["overfitting_evidence"] == nil || ins.Engine["csv_path"] != nil {
		t.Fatalf("the engine object is the server's stripped payload: %v", ins.Engine)
	}
	if ins.GroupSize == nil || *ins.GroupSize != 120 || ins.RankedBy != "aggressor_affinity_score" || *ins.Contrast[0].BaselineMedian != 0.1 {
		t.Fatalf("group summary must come from the server digest: %+v", ins)
	}
	if len(ins.TopSamples) != 1 || ins.TopSamples[0].ID != "training_1" {
		t.Fatalf("top samples come from the digest's ranking: %+v", ins.TopSamples)
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

func codeServer(t *testing.T, code *bool) *Server {
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
			_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true,"code":` + map[bool]string{true: "true", false: "false"}[*code] + `},"me":{}}`))
		case "analysis-export/exportAnalysis":
			url := ""
			if *code {
				url = ts.URL + "/blob/code.tar.gz"
			}
			_, _ = w.Write([]byte(`{"contractVersion":1,"deepLinkPath":"/p","insights":[],"integrationEntryFile":"leap_integration.py","integrationCodeUrl":"` + url + `"}`))
		default:
			t.Fatalf("unexpected call %s", r.URL.Path)
		}
	}))
	t.Cleanup(ts.Close)
	return newServer(NewClient(ts.URL+"/api/v2", "k"))
}

func TestIntegrationCodeToolReadsTheEntryFileAndScrubs(t *testing.T) {
	on := true
	s := codeServer(t, &on)
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
	off := false
	s := codeServer(t, &off)
	_, _, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex})
	if err == nil || !strings.Contains(err.Error(), `"Integration code"`) {
		t.Fatalf("got %v", err)
	}
}

func TestCachedCodeIsNotServedAfterAnAdminTurnsCodeOff(t *testing.T) {
	code := true
	s := codeServer(t, &code)
	if _, _, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex}); err != nil {
		t.Fatal(err)
	}
	code = false
	_, _, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex})
	if err == nil || !strings.Contains(err.Error(), `"Integration code"`) {
		t.Fatalf("the cached archive must not outlive the policy, got %v", err)
	}
}
