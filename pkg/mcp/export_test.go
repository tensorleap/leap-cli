package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const exportCsv = "sample_id,is_low_perf_root_member,metrics.loss,metadata.fog\ntraining_1,true,2.5,yes\ntraining_2,true,1.5,yes\ntraining_3,false,0.1,no\n"
const populationCsv = "sample_id,metrics.loss,metadata.fog\ntraining_1,2.5,yes\ntraining_2,1.5,yes\ntraining_3,0.1,no\ntraining_4,0.1,no\n"

func exportServer(t *testing.T, access AiAccess) *Server {
	t.Helper()
	img := image.NewGray(image.Rect(0, 0, 8, 8))
	img.SetGray(1, 1, color.Gray{255})
	var pngBuf bytes.Buffer
	_ = png.Encode(&pngBuf, img)
	hashed := visualizationID("training_1")
	var ts *httptest.Server
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path
		switch {
		case strings.HasSuffix(p, "/blob/samples.csv"):
			_, _ = w.Write([]byte(exportCsv))
		case strings.HasSuffix(p, "/blob/population.csv"):
			_, _ = w.Write([]byte(populationCsv))
		case strings.HasSuffix(p, "/blob/cluster.json"):
			_, _ = w.Write([]byte(`{"samples_index":{"training_1":0,"training_2":1}}`))
		case strings.HasSuffix(p, "/blob/top.json"):
			_, _ = w.Write([]byte(`{"title":"t","summary":{"csv_path":"vis/x.csv"},"correlated_metadata":[{"population_a":{"filters":[{"value":{"blob_paths":["organizations/x"]}}],"n":3}}]}`))
		case strings.HasSuffix(p, "/blob/fix.csv"):
			_, _ = w.Write([]byte("sample_id\nunlabeled_7\n"))
		case strings.HasSuffix(p, "/blob/code.tar.gz"):
			_, _ = w.Write(codeTarGz(t, map[string]string{"leap_integration.py": "api_key = \"AKIAABCDEFGHIJKLMNOP\"\n"}))
		case strings.HasSuffix(p, "/assets/data.png"):
			_, _ = w.Write(pngBuf.Bytes())
		case strings.HasSuffix(p, "/image/vis/payload.json"):
			_, _ = w.Write([]byte(`{"data":{"blob":"x"}}`))
		case strings.HasSuffix(p, "/hbar/bars/payload.json"):
			_, _ = w.Write([]byte(`{"data":{"body":[1,2]}}`))
		case strings.HasSuffix(p, "/s/escaped.txt"):
			_, _ = w.Write([]byte("escaped"))
		}
		switch strings.TrimPrefix(p, "/api/v2/") {
		case "analysis-export/listTargets":
			b, _ := json.Marshal(map[string]any{"contractVersion": 1, "aiAccess": access, "me": map[string]any{"role": "user"}})
			_, _ = w.Write(b)
		case "analysis-export/exportAnalysis":
			u := func(name string) string {
				if !access.SampleRows && name != "code.tar.gz" || !access.Code && name == "code.tar.gz" {
					return ""
				}
				return ts.URL + "/blob/" + name
			}
			resp := map[string]any{"contractVersion": 1, "deepLinkPath": "/p", "version": map[string]any{"name": "v1"},
				"populationCsvUrl": u("population.csv"), "integrationCodeUrl": u("code.tar.gz"), "integrationEntryFile": "leap_integration.py",
				"predictionLabels": map[string][]string{"classes": {"cat", "dog"}},
				"insights": []map[string]any{
					{"cid": "c1", "index": 1, "status": "InReview", "csvUrl": u("samples.csv"), "clusterBlobUrl": u("cluster.json"), "fixingCsvUrl": u("fix.csv"), "topPanelUrl": u("top.json"),
						"insightType": map[string]any{"id_": "i1", "type": "low_performance", "severity": 2, "n_samples": 3, "aggressor_fixing": map[string]any{"num_of_samples_to_label": 1, "csv_path": "vis/fix.csv"}}},
					{"cid": "c2", "index": 2, "status": "InReview", "csvUrl": u("samples.csv"),
						"insightType": map[string]any{"id_": "i2", "parent_id": "i1", "type": "low_performance", "severity": 1, "n_samples": 2}},
				}}
			b, _ := json.Marshal(resp)
			_, _ = w.Write(b)
		case "analysis-export/getSampleAssets":
			if !access.Visuals {
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"error":"off","code":"AI_ACCESS_DISABLED"}`))
				return
			}
			prefix := "org/projects/p/vis/a/sample_visualizers/" + hashed + "/"
			b, _ := json.Marshal(map[string]any{"samples": []map[string]any{{"sampleId": hashed, "files": []map[string]string{
				{"path": prefix + "image/vis/assets/data.png", "url": ts.URL + "/s/image/vis/assets/data.png"},
				{"path": prefix + "image/vis/payload.json", "url": ts.URL + "/s/image/vis/payload.json"},
				{"path": prefix + "hbar/bars/payload.json", "url": ts.URL + "/s/hbar/bars/payload.json"},
				{"path": "org/projects/p/vis/a/../../../escape/payload.json", "url": ts.URL + "/s/hbar/bars/payload.json"},
				{"path": prefix + "../../../../integration/leap_integration.py", "url": ts.URL + "/s/escaped.txt"},
				{"path": prefix + "../../population.csv", "url": ts.URL + "/s/escaped.txt"},
			}}}})
			_, _ = w.Write(b)
		}
	}))
	t.Cleanup(ts.Close)
	return newServer(NewClient(ts.URL+"/api/v2", "k"))
}

func readManifest(t *testing.T, res ExportResult) ExportOut {
	t.Helper()
	var m ExportOut
	b, err := os.ReadFile(res.Manifest)
	if err != nil || json.Unmarshal(b, &m) != nil {
		t.Fatalf("manifest unreadable: %v", err)
	}
	return m
}

func TestExportWritesEverythingAllowed(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true, SampleRows: true, Visuals: true, Code: true})
	dir := t.TempDir()
	_, res, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	out := readManifest(t, res)
	for _, p := range []string{out.PopulationCsv, path.Join(out.IntegrationDir, "leap_integration.py"),
		out.Insights[0].SamplesCsv, out.Insights[0].ClusterJson, out.Insights[0].FixingCsv, out.Insights[0].TopPanelJson} {
		if p == "" || filepath.IsAbs(p) {
			t.Fatalf("manifest paths must be relative and present: %q in %+v", p, out)
		}
		if _, err := os.Stat(filepath.Join(dir, p)); err != nil {
			t.Fatalf("%s not written: %v", p, err)
		}
	}
	if out.Insights[0].Summary == nil || out.Insights[0].Summary.GroupSize != 2 || out.ClassLabels["classes"][0] != "cat" {
		t.Fatalf("summary/labels: %+v", out.Insights[0].Summary)
	}
	if out.Insights[1].Dir != "insight_1_low_performance/sub_2_low_performance" {
		t.Fatalf("sub-insight dir: %s", out.Insights[1].Dir)
	}
	smp := out.Insights[0].Samples[0]
	if smp.ID != "training_1" || smp.Rank != 1 || !smp.Rendered || len(smp.Files) != 3 {
		t.Fatalf("rendered sample files: %+v", smp)
	}
	if second := out.Insights[0].Samples[1]; second.ID != "training_2" || second.Rank != 2 || second.Rendered {
		t.Fatalf("samples keep their rank order whether or not they are rendered: %+v", out.Insights[0].Samples)
	}
	for _, f := range smp.Files {
		if !strings.HasPrefix(f, "insight_1_low_performance/samples/training_1/") {
			t.Fatalf("file outside the sample dir: %s", f)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "escape")); err == nil {
		t.Fatal("a traversal path from the server must not be written")
	}
	panel, _ := os.ReadFile(filepath.Join(dir, out.Insights[0].TopPanelJson))
	if strings.Contains(string(panel), "organizations/") || strings.Contains(string(panel), "csv_path") || !strings.Contains(string(panel), `"title"`) {
		t.Fatalf("top_panel.json must keep its content and drop storage internals: %s", panel)
	}
	if len(out.Insights) != 2 || out.Insights[0].Engine["n_samples"] != 3.0 {
		t.Fatalf("manifest insights: %+v", out.Insights)
	}
	if len(res.Skipped) != 0 || res.FilesWritten < 8 || len(res.Insights) != 2 || res.Insights[0].Rendered != 1 || res.Insights[0].GroupSize != 2 {
		t.Fatalf("result: %+v", res)
	}
	if b, _ := json.Marshal(res); strings.Contains(string(b), "payload.json") {
		t.Fatalf("the tool result must stay short; per-file paths belong in manifest.json: %s", b)
	}
}

func TestExportExplainsWhatThePolicyWithheld(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true})
	_, res, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(res.Skipped, "\n")
	for _, want := range []string{"population csv", "integration code", "insight sample lists"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("skipped should mention %q: %v", want, res.Skipped)
		}
	}
	out := readManifest(t, res)
	if out.PopulationCsv != "" || out.IntegrationDir != "" || out.Insights[0].SamplesCsv != "" || out.Insights[0].Engine["n_samples"] != 3.0 {
		t.Fatalf("got %+v", out)
	}
}

func TestExportDecidesFromTheCurrentPolicyNotTheCache(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true, SampleRows: true, Visuals: true, Code: true})
	s.policies.entries[projectHex] = policyEntry{access: &AiAccess{Stats: true, SampleRows: true}, at: time.Now()}
	_, res, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Skipped) != 0 || res.Insights[0].Rendered != 1 {
		t.Fatalf("an admin just turned visuals and code on; the export must see it: %+v", res)
	}
}

func TestReExportReplacesThePreviousOneAndGuardsOtherDirs(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true, SampleRows: true, Visuals: true, Code: true})
	dir := t.TempDir()
	stale := filepath.Join(dir, "insight_9_duplication", "samples.csv")
	_ = os.MkdirAll(filepath.Dir(stale), 0o755)
	_ = os.WriteFile(stale, []byte("old"), 0o644)
	_ = os.WriteFile(filepath.Join(dir, "insights.json"), []byte("{}"), 0o644)
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: dir}); err != nil {
		t.Fatalf("the pre-MCP skill's output dir must be accepted: %v", err)
	}
	if _, err := os.Stat(stale); err == nil {
		t.Fatal("stale insight dirs from an earlier export must be removed")
	}
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: dir}); err != nil {
		t.Fatalf("re-export of the same version: %v", err)
	}
	other := strings.Replace(versionHex, "1", "2", 1)
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: other, Dir: dir}); err == nil || !strings.Contains(err.Error(), "another version") {
		t.Fatalf("another version's export must not be overwritten, got %v", err)
	}
	web := t.TempDir()
	_ = os.WriteFile(filepath.Join(web, "manifest.json"), []byte(`{"name":"my pwa"}`), 0o644)
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: web}); err == nil || !strings.Contains(err.Error(), "not empty") {
		t.Fatalf("a foreign manifest.json must not pass as a previous export, got %v", err)
	}
}

func TestExportRequiresAnAbsoluteDirAndStatistics(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true})
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: "relative/dir"}); err == nil || !strings.Contains(err.Error(), "absolute") {
		t.Fatalf("got %v", err)
	}
	s = exportServer(t, AiAccess{})
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: t.TempDir()}); err == nil || !strings.Contains(err.Error(), `"Statistics and insights"`) {
		t.Fatalf("got %v", err)
	}
}

func TestExportKeepsServerPathsInsideTheSampleDirAndScrubsText(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true, SampleRows: true, Visuals: true, Code: true})
	dir := t.TempDir()
	_, res, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	out := readManifest(t, res)
	for _, p := range []string{filepath.Join(dir, "integration", "leap_integration.py"), filepath.Join(dir, "population.csv")} {
		b, _ := os.ReadFile(p)
		if strings.Contains(string(b), "AKIAABCDEFGHIJKLMNOP") || strings.Contains(string(b), "escaped") {
			t.Fatalf("%s was overwritten or unscrubbed: %q", p, b)
		}
	}
	for _, f := range out.Insights[0].Samples[0].Files {
		if strings.Contains(f, "escape") {
			t.Fatalf("escaping asset path written: %s", f)
		}
	}
	busy := t.TempDir()
	_ = os.WriteFile(filepath.Join(busy, "notes.txt"), []byte("mine"), 0o644)
	if _, _, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: busy}); err == nil || !strings.Contains(err.Error(), "not empty") {
		t.Fatalf("a non-empty non-export dir must be refused, got %v", err)
	}
}

func TestCodeWithheldByTheServerIsReportedAsThePolicy(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimPrefix(r.URL.Path, "/api/v2/") {
		case "analysis-export/listTargets":
			_, _ = w.Write([]byte(`{"contractVersion":1,"aiAccess":{"stats":true,"code":true},"me":{"role":"admin"}}`))
		case "analysis-export/exportAnalysis":
			_, _ = w.Write([]byte(`{"contractVersion":1,"deepLinkPath":"/p","insights":[],"integrationEntryFile":"leap_integration.py"}`))
		}
	}))
	defer ts.Close()
	s := newServer(NewClient(ts.URL+"/api/v2", "k"))
	_, _, err := s.getIntegrationCode(context.Background(), nil, CodeIn{ProjectID: projectHex, VersionID: versionHex})
	if err == nil || !strings.Contains(err.Error(), `"Integration code"`) || strings.Contains(err.Error(), "repository") {
		t.Fatalf("got %v", err)
	}
}

func TestCodeArchiveSkipsDotDotAndBoundsTheTotal(t *testing.T) {
	arc, err := readCodeArchive(codeTarGz(t, map[string]string{"..": "x", "a.py": "ok", `..\\b.py`: "x"}))
	if err != nil || len(arc.files) != 1 || arc.files["a.py"] == nil {
		t.Fatalf("got %v %v", err, arc)
	}
	big := codeTarGz(t, map[string]string{"zeros.bin": strings.Repeat("\x00", maxCodeArchive)})
	if _, err := readCodeArchive(big); err == nil || !strings.Contains(err.Error(), "limit") {
		t.Fatalf("expanded archive over the limit must be refused, got %v", err)
	}
}
