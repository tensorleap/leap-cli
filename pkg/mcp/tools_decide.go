package mcp

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"unicode/utf16"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const maxDecisionNote = 300

type DecideIn struct {
	ProjectID string `json:"projectId" jsonschema:"project id, not its name (report.projectId in tl_get_report's result)"`
	InsightID string `json:"insightId" jsonschema:"the finding's insightId from tl_get_report; a finding without one is no longer on the platform and cannot be decided"`
	Decision  string `json:"decision" jsonschema:"acknowledge, archive or reopen"`
	Note      string `json:"note,omitempty" jsonschema:"the person's reason, one line, at most 300 characters"`
	ReportID  string `json:"reportId,omitempty" jsonschema:"the id of the report the finding was read from"`
}

type DecideResult struct {
	Status  string `json:"status" jsonschema:"the finding's status now: Approved, Archived or InReview"`
	Applied bool   `json:"applied" jsonschema:"false when nothing was recorded: the finding already had that status, or someone changed it at the same moment"`
	By      string `json:"by" jsonschema:"the name the decision was recorded under; empty when applied is false"`
	Note    string `json:"note,omitempty" jsonschema:"the note as recorded"`
}

func (s *Server) addDecideTool(srv *sdk.Server) {
	closed := false
	sdk.AddTool(srv, &sdk.Tool{Name: "tl_decide_finding",
		Description: "Record a person's decision on a finding of a published report on the Tensorleap server. Use it only when the person asked for that decision on that finding; never decide for them. It is recorded under their name and changes the finding's status for the whole team. acknowledge means they agree the finding is real and worth acting on (status Approved); archive sets it aside (Archived); reopen puts it back in review (InReview). Read the finding with tl_get_report first.",
		Annotations: &sdk.ToolAnnotations{Title: "Decide on a finding", ReadOnlyHint: false, IdempotentHint: true, OpenWorldHint: &closed}},
		s.decideFinding)
}

func (s *Server) decideFinding(ctx context.Context, _ *sdk.CallToolRequest, in DecideIn) (*sdk.CallToolResult, DecideResult, error) {
	switch in.Decision {
	case "acknowledge", "archive", "reopen":
	default:
		return nil, DecideResult{}, fmt.Errorf("decision must be acknowledge, archive or reopen, not %q", truncate(in.Decision, 40))
	}
	note := strings.Join(strings.Fields(in.Note), " ")
	// counted the way the server does, which cuts a longer note without saying so
	if n := len(utf16.Encode([]rune(note))); n > maxDecisionNote {
		return nil, DecideResult{}, fmt.Errorf("note is %d characters long; shorten it to one line of at most %d", n, maxDecisionNote)
	}
	req := map[string]any{"projectId": in.ProjectID, "insightId": in.InsightID, "decision": in.Decision}
	if note != "" {
		req["note"] = note
	}
	if in.ReportID != "" {
		req["reportId"] = in.ReportID
	}
	var out struct {
		Status   string `json:"status"`
		Applied  bool   `json:"applied"`
		Decision struct {
			By struct {
				Name string `json:"name"`
			} `json:"by"`
			Note string `json:"note"`
		} `json:"decision"`
	}
	if err := s.client.Post(ctx, "analysis-reports/decide", req, &out); err != nil {
		if routeMissing(err) {
			return nil, DecideResult{}, errors.New("this Tensorleap server predates decisions on findings; upgrade the server")
		}
		return nil, DecideResult{}, explain(err)
	}
	return nil, DecideResult{Status: out.Status, Applied: out.Applied, By: out.Decision.By.Name, Note: out.Decision.Note}, nil
}
