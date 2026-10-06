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
	"path/filepath"
	"strings"
	"testing"
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
		case strings.HasSuffix(p, "/blob/fix.csv"):
			_, _ = w.Write([]byte("sample_id\nunlabeled_7\n"))
		case strings.HasSuffix(p, "/blob/code.tar.gz"):
			_, _ = w.Write(codeTarGz(t, map[string]string{"leap_integration.py": "print('hi')\n"}))
		case strings.HasSuffix(p, "/assets/data.png"):
			_, _ = w.Write(pngBuf.Bytes())
		case strings.HasSuffix(p, "/image/vis/payload.json"):
			_, _ = w.Write([]byte(`{"data":{"blob":"x"}}`))
		case strings.HasSuffix(p, "/hbar/bars/payload.json"):
			_, _ = w.Write([]byte(`{"data":{"body":[1,2]}}`))
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
					{"cid": "c1", "index": 1, "status": "InReview", "csvUrl": u("samples.csv"), "clusterBlobUrl": u("cluster.json"), "fixingCsvUrl": u("fix.csv"),
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
			}}}})
			_, _ = w.Write(b)
		}
	}))
	t.Cleanup(ts.Close)
	return newServer(NewClient(ts.URL+"/api/v2", "k"))
}

func TestExportWritesEverythingAllowed(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true, SampleRows: true, Visuals: true, Code: true})
	dir := t.TempDir()
	_, out, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range []string{out.PopulationCsv, out.Manifest, filepath.Join(out.IntegrationDir, "leap_integration.py"),
		out.Insights[0].SamplesCsv, out.Insights[0].ClusterJson, out.Insights[0].FixingCsv} {
		if p == "" {
			t.Fatalf("missing path in %+v", out)
		}
		if _, err := os.Stat(p); err != nil {
			t.Fatalf("%s not written: %v", p, err)
		}
	}
	if out.Insights[0].Summary == nil || out.Insights[0].Summary.GroupSize != 2 || out.ClassLabels["classes"][0] != "cat" {
		t.Fatalf("summary/labels: %+v", out.Insights[0].Summary)
	}
	if !strings.HasPrefix(out.Insights[1].Dir, filepath.Join(dir, "insight_1_low_performance", "sub_2_low_performance")) {
		t.Fatalf("sub-insight dir: %s", out.Insights[1].Dir)
	}
	smp := out.Insights[0].Samples[0]
	if smp.ID != "training_1" || !smp.Rendered || len(smp.Files) != 3 {
		t.Fatalf("rendered sample files: %+v", smp)
	}
	for _, f := range smp.Files {
		if !strings.HasPrefix(f, filepath.Join(out.Insights[0].Dir, "samples", "training_1")) {
			t.Fatalf("file outside the sample dir: %s", f)
		}
	}
	if _, err := os.Stat(filepath.Join(dir, "escape")); err == nil {
		t.Fatal("a traversal path from the server must not be written")
	}
	var manifest ExportOut
	b, _ := os.ReadFile(out.Manifest)
	if err := json.Unmarshal(b, &manifest); err != nil || len(manifest.Insights) != 2 || manifest.Insights[0].Engine["n_samples"] != 3.0 {
		t.Fatalf("manifest: %v %+v", err, manifest.Insights)
	}
	if len(out.Skipped) != 0 || out.FilesWritten < 8 {
		t.Fatalf("skipped=%v files=%d", out.Skipped, out.FilesWritten)
	}
}

func TestExportExplainsWhatThePolicyWithheld(t *testing.T) {
	s := exportServer(t, AiAccess{Stats: true})
	_, out, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: projectHex, VersionID: versionHex, Dir: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(out.Skipped, "\n")
	for _, want := range []string{"population csv", "integration code", "insight sample lists"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("skipped should mention %q: %v", want, out.Skipped)
		}
	}
	if out.PopulationCsv != "" || out.IntegrationDir != "" || out.Insights[0].SamplesCsv != "" || out.Insights[0].Engine["n_samples"] != 3.0 {
		t.Fatalf("got %+v", out)
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
