package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const maxReportBytes = 32 << 20

type PublishIn struct {
	Dir string `json:"dir" jsonschema:"absolute path of the export directory tl_export_analysis wrote, holding report.html (and report.md, report.json when written)"`
}

type PublishResult struct {
	Link     string `json:"link" jsonschema:"opens the report in the Tensorleap UI"`
	ReportID string `json:"reportId"`
	Title    string `json:"title"`
	Replaced bool   `json:"replaced" jsonschema:"true when your previous report of this evaluation was replaced"`
}

// the report.json fields that name the report in the UI
type reportSpec struct {
	Project          string `json:"project"`
	Version          string `json:"version"`
	ExecutiveSummary string `json:"executive_summary"`
}

func (s *Server) addPublishTool(srv *sdk.Server) {
	closed := false
	sdk.AddTool(srv, &sdk.Tool{Name: "tl_publish_report",
		Description: "Publish the finished report of an analysis export (report.html, plus report.md and report.json when present) to the Tensorleap server: the team finds it in the project's Reports tab in the UI and the returned link opens it. Publishing the same evaluation again replaces your previous report. Call it after build-report and inline-html; dir is the export directory tl_export_analysis wrote.",
		Annotations: &sdk.ToolAnnotations{Title: "Publish report", ReadOnlyHint: false, IdempotentHint: true, OpenWorldHint: &closed}},
		s.publishReport)
}

var (
	imgSrcRe = regexp.MustCompile(`(?is)<img[^>]*?\ssrc="([^"]*)"`)
	tagRe    = regexp.MustCompile(`<[^>]+>`)
)

func (s *Server) publishReport(ctx context.Context, _ *sdk.CallToolRequest, in PublishIn) (*sdk.CallToolResult, PublishResult, error) {
	if !filepath.IsAbs(in.Dir) {
		return nil, PublishResult{}, errors.New("dir must be an absolute path")
	}
	dir := filepath.Clean(in.Dir)
	var m manifest
	if raw, err := os.ReadFile(filepath.Join(dir, "manifest.json")); err != nil || json.Unmarshal(raw, &m) != nil || m.ProjectID == "" || m.VersionID == "" {
		return nil, PublishResult{}, fmt.Errorf("%s holds no manifest.json written by tl_export_analysis; pass the export directory", dir)
	}
	page, err := os.ReadFile(filepath.Join(dir, "report.html"))
	if err != nil {
		return nil, PublishResult{}, fmt.Errorf("no report.html in %s; run build-report and inline-html first", dir)
	}
	if len(page) > maxReportBytes {
		return nil, PublishResult{}, errors.New("report.html is larger than 32 MB; re-run inline-html so its images are downscaled")
	}
	for _, match := range imgSrcRe.FindAllSubmatch(page, -1) {
		if src := string(match[1]); !strings.HasPrefix(src, "data:") {
			return nil, PublishResult{}, fmt.Errorf("report.html still references an image file (%s); run inline-html first", truncate(src, 60))
		}
	}
	req := map[string]any{"projectId": m.ProjectID, "versionId": m.VersionID, "html": string(page), "title": m.Version, "summary": ""}
	if b, err := os.ReadFile(filepath.Join(dir, "report.json")); err == nil {
		var spec reportSpec
		if json.Unmarshal(b, &spec) == nil {
			if spec.Project != "" && spec.Version != "" {
				req["title"] = spec.Project + " / " + spec.Version
			}
			req["summary"] = plainText(spec.ExecutiveSummary)
		}
		req["json"] = string(b)
	}
	if req["title"] == "" {
		req["title"] = "Analysis report"
	}
	if b, err := os.ReadFile(filepath.Join(dir, "report.md")); err == nil {
		req["markdown"] = string(b)
	}
	var out struct {
		Report struct {
			Cid   string `json:"cid"`
			Title string `json:"title"`
		} `json:"report"`
		Replaced bool   `json:"replaced"`
		LinkPath string `json:"linkPath"`
	}
	if err := s.client.Post(ctx, "analysis-reports/publish", req, &out); err != nil {
		var apiErr *APIError
		if errors.As(err, &apiErr) && apiErr.Status == 404 && strings.Contains(apiErr.Body, "Cannot POST") {
			return nil, PublishResult{}, fmt.Errorf("this Tensorleap server predates published reports; upgrade the server (the report stays at %s)", filepath.Join(dir, "report.html"))
		}
		return nil, PublishResult{}, explain(err)
	}
	return nil, PublishResult{Link: s.client.UIBase() + out.LinkPath, ReportID: out.Report.Cid, Title: out.Report.Title, Replaced: out.Replaced}, nil
}

// plainText flattens one of report.json's HTML fragments into one line of text
func plainText(fragment string) string {
	return strings.Join(strings.Fields(html.UnescapeString(tagRe.ReplaceAllString(fragment, " "))), " ")
}
