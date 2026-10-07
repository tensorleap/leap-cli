package mcp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type AiAccess struct {
	Stats      bool `json:"stats"`
	SampleRows bool `json:"sampleRows"`
	Visuals    bool `json:"visuals"`
	JobLogs    bool `json:"jobLogs"`
	Code       bool `json:"code"`
}

func (a AiAccess) blocked() []string {
	var off []string
	for _, c := range []struct {
		on    bool
		label string
	}{{a.Stats, "statistics and insights"}, {a.SampleRows, "per-sample data"}, {a.Visuals, "sample visualizations"}, {a.JobLogs, "job logs"}, {a.Code, "integration code"}} {
		if !c.on {
			off = append(off, c.label)
		}
	}
	return off
}

type targetsResponse struct {
	AiAccess *AiAccess `json:"aiAccess"`
	Me       struct {
		Email string `json:"email"`
	} `json:"me"`
	Projects []struct {
		AiAccess *AiAccess `json:"aiAccess"`
	} `json:"projects"`
}

func explain(err error) error {
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		return err
	}
	switch {
	case apiErr.Status == 401:
		return errors.New("the server rejected your API key; run `leap auth login` (or `leap auth select <env>`)")
	case apiErr.Status == 404 && strings.Contains(apiErr.Body, "Cannot POST"):
		return errors.New("this Tensorleap server is too old for AI assistants; upgrade the server")
	case apiErr.Status >= 500:
		return fmt.Errorf("the Tensorleap server failed on this request (%d): %s", apiErr.Status, serverMessage(apiErr.Body))
	}
	return errors.New(serverMessage(apiErr.Body))
}

func serverMessage(body string) string {
	var payload struct {
		Error string `json:"error"`
	}
	if json.Unmarshal([]byte(body), &payload) == nil && payload.Error != "" {
		return payload.Error
	}
	return body
}

// StatusLines is the one-glance state printed by `leap mcp config` and by a terminal run of `leap mcp`
func StatusLines(ctx context.Context, c *Client) []string {
	var t targetsResponse
	if err := c.Post(ctx, "analysis-export/listTargets", map[string]any{}, &t); err != nil {
		return []string{fmt.Sprintf("Server: %s, not reachable right now: %v", c.UIBase(), explain(err))}
	}
	lines := []string{fmt.Sprintf("Server: %s (%s)", c.UIBase(), t.Me.Email)}
	if t.AiAccess == nil {
		return append(lines, "AI access: this server is too old for AI assistants; upgrade it")
	}
	custom := 0
	for _, p := range t.Projects {
		if p.AiAccess != nil && *p.AiAccess != *t.AiAccess {
			custom++
		}
	}
	access := "everything is on"
	if off := t.AiAccess.blocked(); len(off) > 0 {
		access = "turned off by default: " + strings.Join(off, ", ")
	}
	if custom > 0 {
		access += fmt.Sprintf("; %d project(s) have their own settings", custom)
	}
	return append(lines, "AI access: "+access+" (admins: gear icon > AI ACCESS)")
}
