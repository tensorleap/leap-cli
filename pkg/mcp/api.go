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

// ServerStatus is what `leap mcp config` shows about the server before connecting an assistant
type ServerStatus struct {
	UI             string
	Email          string
	Off            []string
	CustomProjects int
	Projects       []string
}

func Status(ctx context.Context, c *Client) (*ServerStatus, error) {
	var t struct {
		targetsResponse
		Projects []struct {
			Name     string    `json:"name"`
			AiAccess *AiAccess `json:"aiAccess"`
		} `json:"projects"`
	}
	if err := c.Post(ctx, "analysis-export/listTargets", map[string]any{}, &t); err != nil {
		return nil, explain(err)
	}
	st := &ServerStatus{UI: c.UIBase(), Email: t.Me.Email}
	for _, p := range t.Projects {
		st.Projects = append(st.Projects, p.Name)
		if t.AiAccess != nil && p.AiAccess != nil && *p.AiAccess != *t.AiAccess {
			st.CustomProjects++
		}
	}
	if t.AiAccess != nil {
		st.Off = t.AiAccess.blocked()
	}
	return st, nil
}
