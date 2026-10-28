package mcp

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func writeReportFiles(t *testing.T, dir string, files map[string]string) {
	t.Helper()
	for name, content := range files {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
}

func TestPublishSendsTheReportAndReturnsItsLink(t *testing.T) {
	dir := t.TempDir()
	writeReportFiles(t, dir, map[string]string{
		"manifest.json": `{"projectId":"p1","versionId":"v1","version":"model"}`,
		"report.html":   `<!doctype html><html><body><img src="data:image/png;base64,AA=="></body></html>`,
		"report.md":     "# report",
		"report.json":   `{"project":"mnist","version":"regenrate","executive_summary":"The <a href=\"#i\">test split</a> is a copy &amp;\n 100% means nothing."}`,
	})
	f := &fakeTensorleap{}
	ts := f.start(t)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "k")}
	_, res, err := s.publishReport(context.Background(), nil, PublishIn{Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	if res.ReportID != "r1" || !res.Replaced || res.Title != "mnist / regenrate" || !strings.HasSuffix(res.Link, "/project/p1/dashboard?report=r1") || strings.Contains(res.Link, "/api/v2") {
		t.Fatalf("result: %+v", res)
	}
	p := f.published
	if p["projectId"] != "p1" || p["versionId"] != "v1" || p["title"] != "mnist / regenrate" || p["summary"] != "The test split is a copy & 100% means nothing." || p["markdown"] != "# report" {
		t.Fatalf("published: %+v", p)
	}
	if h, _ := p["html"].(string); !strings.Contains(h, "data:image") {
		t.Fatalf("html must be sent as is, got %q", h)
	}
	if j, _ := p["json"].(string); !strings.Contains(j, "regenrate") {
		t.Fatalf("report.json must be sent along, got %q", j)
	}

	writeReportFiles(t, dir, map[string]string{"report.json": "not json"})
	if _, _, err = s.publishReport(context.Background(), nil, PublishIn{Dir: dir}); err != nil || f.published["title"] != "model" || f.published["summary"] != "" {
		t.Fatalf("without a readable report.json the version name is the title: %v %+v", err, f.published)
	}
}

func TestPublishRefusesWhatIsNotAFinishedReport(t *testing.T) {
	s := &Server{client: NewClient("http://x/api/v2", "k")}
	publish := func(dir string) error {
		_, _, err := s.publishReport(context.Background(), nil, PublishIn{Dir: dir})
		return err
	}
	if err := publish("relative"); err == nil {
		t.Fatal("a relative dir must be refused")
	}
	dir := t.TempDir()
	if err := publish(dir); err == nil || !strings.Contains(err.Error(), "manifest.json") {
		t.Fatalf("a directory without an export must be refused, got %v", err)
	}
	writeReportFiles(t, dir, map[string]string{"manifest.json": `{"projectId":"p1","versionId":"v1"}`})
	if err := publish(dir); err == nil || !strings.Contains(err.Error(), "report.html") {
		t.Fatalf("an export without a report must be refused, got %v", err)
	}
	writeReportFiles(t, dir, map[string]string{"report.html": `<html><img class="x" src="insight_1_low_performance/samples/a/data.jpg"></html>`})
	if err := publish(dir); err == nil || !strings.Contains(err.Error(), "inline-html") {
		t.Fatalf("a report with file images is not self-contained, got %v", err)
	}
}
