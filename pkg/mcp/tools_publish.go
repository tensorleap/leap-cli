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
	Dir      string `json:"dir" jsonschema:"absolute path of the export directory tl_export_analysis wrote, holding report.html (and report.md, report.json when written)"`
	ReportID string `json:"reportId,omitempty" jsonschema:"the id of a report this export was published as before; that report is updated instead of a new one being added"`
}

type PublishResult struct {
	Link     string `json:"link" jsonschema:"opens the report in the Tensorleap UI"`
	ReportID string `json:"reportId"`
	Title    string `json:"title"`
	Updated  bool   `json:"updated" jsonschema:"true when the report passed as reportId was updated, false when a new report was added"`
	Note     string `json:"note,omitempty"`
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
		Description: "Publish the finished report of an analysis export (report.html, plus report.md and report.json when present) to the Tensorleap server: the team finds it in the project's Reports tab in the UI and the returned link opens it. Every publish adds a new report in that tab, unless reportId names a report this export was published as before (the reportId of an earlier publish, or one from tl_list_reports): that report is updated in place and keeps its link. Call it after build-report and inline-html; dir is the export directory tl_export_analysis wrote.",
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
		LinkPath string `json:"linkPath"`
	}
	if in.ReportID != "" {
		req["reportId"] = in.ReportID
	}
	err = s.client.Post(ctx, "analysis-reports/publish", req, &out)
	var apiErr *APIError
	if in.ReportID != "" && errors.As(err, &apiErr) && apiErr.Status == 400 && strings.Contains(apiErr.Body, "reportId") && strings.Contains(apiErr.Body, "excess property") {
		// a server that predates updates refuses the field it does not know; it can still add a report
		delete(req, "reportId")
		err = s.client.Post(ctx, "analysis-reports/publish", req, &out)
	}
	if err != nil {
		if routeMissing(err) {
			return nil, PublishResult{}, fmt.Errorf("this Tensorleap server predates published reports; upgrade the server (the report stays at %s)", filepath.Join(dir, "report.html"))
		}
		return nil, PublishResult{}, explain(err)
	}
	res := PublishResult{Link: s.client.UIBase() + out.LinkPath, ReportID: out.Report.Cid, Title: out.Report.Title, Updated: in.ReportID != "" && out.Report.Cid == in.ReportID}
	if in.ReportID != "" && !res.Updated {
		res.Note = fmt.Sprintf("this Tensorleap server predates updating a report, so it added a new report (%s) and left %s as it was; upgrade the server to update a report in place", res.ReportID, in.ReportID)
	}
	return nil, res, nil
}

// plainText flattens one of report.json's HTML fragments into one line of text
func plainText(fragment string) string {
	return strings.Join(strings.Fields(html.UnescapeString(tagRe.ReplaceAllString(fragment, " "))), " ")
}
