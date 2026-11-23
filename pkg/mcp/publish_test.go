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

func TestPublishSendsTheReportContentAndReturnsItsLink(t *testing.T) {
	dir := t.TempDir()
	writeReportFiles(t, dir, map[string]string{
		"manifest.json": `{"projectId":"p1","versionId":"v1","version":"model"}`,
		"report.md":     "# report",
		"report.json": `{"schemaVersion":1,"project":"mnist","version":"regenrate",
			"summary":"The <a href=\"#insight-1\">test split</a> is a copy &amp;\n 100% means nothing.",
			"groups":[{"title":"Copies across splits","findings":[{"id":"insight-1","severity":3}]}]}`,
	})
	f := &fakeTensorleap{}
	ts := f.start(t)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "k")}
	_, res, err := s.publishReport(context.Background(), nil, PublishIn{Dir: dir})
	if err != nil {
		t.Fatal(err)
	}
	if res.ReportID != "r1" || res.Title != "mnist / regenrate" || !strings.HasSuffix(res.Link, "/project/p1/dashboard?report=r1") || strings.Contains(res.Link, "/api/v2") {
		t.Fatalf("result: %+v", res)
	}
	p := f.published
	if p["projectId"] != "p1" || p["versionId"] != "v1" || p["title"] != "mnist / regenrate" || p["summary"] != "The test split is a copy & 100% means nothing." {
		t.Fatalf("published: %+v", p)
	}
	content, _ := p["content"].(map[string]any)
	groups, _ := content["groups"].([]any)
	if content["schemaVersion"] != float64(1) || content["project"] != "mnist" || len(groups) != 1 {
		t.Fatalf("report.json must be sent as the report's content, an object and not a string, got %#v", p["content"])
	}
	for _, gone := range []string{"html", "markdown", "json"} {
		if _, sent := p[gone]; sent {
			t.Fatalf("%q is no longer part of a published report: %+v", gone, p)
		}
	}

	writeReportFiles(t, dir, map[string]string{"report.json": `{"schemaVersion":1}`})
	if _, _, err = s.publishReport(context.Background(), nil, PublishIn{Dir: dir}); err != nil || f.published["title"] != "model" || f.published["summary"] != "" {
		t.Fatalf("a report that does not name its project is titled by the version: %v %+v", err, f.published)
	}
}

func TestPublishRefusesWhatIsNotAFinishedReport(t *testing.T) {
	f := &fakeTensorleap{}
	ts := f.start(t)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "k")}
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
	if err := publish(dir); err == nil || !strings.Contains(err.Error(), "no report.json") {
		t.Fatalf("an export without a report must be refused, got %v", err)
	}
	for _, broken := range []string{"not json", `["a list"]`, `{"summary": 3}`} {
		writeReportFiles(t, dir, map[string]string{"report.json": broken})
		if err := publish(dir); err == nil || !strings.Contains(err.Error(), "is not a report") {
			t.Fatalf("report.json %q must be refused, got %v", broken, err)
		}
	}
	// what the analysis skill wrote while a report was still an HTML file
	writeReportFiles(t, dir, map[string]string{"report.json": `{"project":"mnist","version":"regenrate","executive_summary":"s","overview":[],"groups":[]}`})
	if err := publish(dir); err == nil || !strings.Contains(err.Error(), "older analysis skill") || !strings.Contains(err.Error(), "leap mcp config") {
		t.Fatalf("a report from a skill copy older than this CLI must say how to update the skills, got %v", err)
	}
	if f.published != nil {
		t.Fatalf("nothing may reach the server for a refused report, got %+v", f.published)
	}
}

func TestPublishRelaysWhatTheServerSays(t *testing.T) {
	dir := t.TempDir()
	writeReportFiles(t, dir, map[string]string{
		"manifest.json": `{"projectId":"p1","versionId":"v1","version":"model"}`,
		"report.json":   `{"schemaVersion":1,"project":"mnist","version":"regenrate","summary":"s"}`,
	})
	f := &fakeTensorleap{}
	ts := f.start(t)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "k")}
	publish := func(refusal string) string {
		f.publishRefusal = refusal
		_, _, err := s.publishReport(context.Background(), nil, PublishIn{Dir: dir})
		if err == nil {
			t.Fatalf("a refused publish must fail (%s)", refusal)
		}
		return err.Error()
	}
	// what a server that keeps reports as HTML files answers, recorded from one
	old := publish(`{"error":"Schema validation failed! interfaceName: PublishAnalysisReportParams\n ValidationError: \npath: , keyword: required, msg: must have required property 'html', schemaPath: #/required\n"}`)
	if !strings.Contains(old, "upgrade the server") || !strings.Contains(old, filepath.Join(dir, "report.json")) {
		t.Fatalf("an older server must be named as the cause, with where the report stays, got %q", old)
	}
	if got := publish(`{"error":"finding id \"failure-1\" is not valid: use insight-<number> or insight-<number>-sub-<number>"}`); !strings.Contains(got, `finding id "failure-1" is not valid`) {
		t.Fatalf("what the server finds wrong with the content must reach the assistant as written, got %q", got)
	}
}
