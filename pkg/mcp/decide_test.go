package mcp

import (
	"context"
	"strings"
	"testing"
)

func TestDecideRecordsThePersonsDecision(t *testing.T) {
	f := &fakeTensorleap{decision: `{"status":"Approved","applied":true,"decision":{"insightId":"i1","decision":"acknowledge","status":"Approved",
	 "by":{"name":"Dana Levi","kind":"member","userId":"u1","via":"assistant"},"reportId":"r1","note":"seen in the field","at":"2026-11-16T10:00:00.000Z"}}`}
	s := &Server{client: NewClient(f.start(t).URL+"/api/v2", "k")}
	_, res, err := s.decideFinding(context.Background(), nil, DecideIn{ProjectID: "p1", InsightID: "i1", Decision: "acknowledge", Note: " seen in\n the field ", ReportID: "r1"})
	if err != nil {
		t.Fatal(err)
	}
	if res != (DecideResult{Status: "Approved", Applied: true, By: "Dana Levi", Note: "seen in the field"}) {
		t.Fatalf("result: %+v", res)
	}
	d := f.decided
	if len(d) != 5 || d["projectId"] != "p1" || d["insightId"] != "i1" || d["decision"] != "acknowledge" || d["note"] != "seen in the field" || d["reportId"] != "r1" {
		t.Fatalf("sent: %+v", d)
	}

	f.decision = `{"status":"Archived","applied":false}`
	_, res, err = s.decideFinding(context.Background(), nil, DecideIn{ProjectID: "p1", InsightID: "i1", Decision: "archive", Note: strings.Repeat("x", maxDecisionNote)})
	if err != nil || res != (DecideResult{Status: "Archived"}) {
		t.Fatalf("a finding that already had the status: %v %+v", err, res)
	}
	if _, sent := f.decided["reportId"]; sent || len(f.decided["note"].(string)) != maxDecisionNote {
		t.Fatalf("an unset reportId must be omitted and a note of the longest length sent whole: %+v", f.decided)
	}
}

func TestDecideRefusesWhatTheServerWouldNotRecord(t *testing.T) {
	f := &fakeTensorleap{decision: `{"status":"Approved","applied":true}`}
	s := &Server{client: NewClient(f.start(t).URL+"/api/v2", "k")}
	for want, in := range map[string]DecideIn{
		"decision must be": {ProjectID: "p1", InsightID: "i1", Decision: "approve"},
		"note is 301":      {ProjectID: "p1", InsightID: "i1", Decision: "archive", Note: strings.Repeat("x", maxDecisionNote+1)},
	} {
		if _, _, err := s.decideFinding(context.Background(), nil, in); err == nil || !strings.Contains(err.Error(), want) {
			t.Fatalf("want %q, got %v", want, err)
		}
	}
	if f.decided != nil {
		t.Fatalf("a refused decision must not reach the server, got %+v", f.decided)
	}
}

func TestDecideTellsAnOldServerFromAMissingFinding(t *testing.T) {
	decide := func(f *fakeTensorleap, insightID string) error {
		s := &Server{client: NewClient(f.start(t).URL+"/api/v2", "k")}
		_, _, err := s.decideFinding(context.Background(), nil, DecideIn{ProjectID: "p1", InsightID: insightID, Decision: "reopen"})
		return err
	}
	if err := decide(&fakeTensorleap{}, "i1"); err == nil || !strings.Contains(err.Error(), "predates decisions on findings; upgrade the server") {
		t.Fatalf("a server without the route must be named as too old, got %v", err)
	}
	if err := decide(&fakeTensorleap{decision: `{}`}, "gone"); err == nil || err.Error() != "finding not found" {
		t.Fatalf("the server's own 404 is its message, got %v", err)
	}
}
