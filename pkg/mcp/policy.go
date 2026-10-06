package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type AiAccess struct {
	Stats      bool `json:"stats"`
	SampleRows bool `json:"sampleRows"`
	Visuals    bool `json:"visuals"`
	JobLogs    bool `json:"jobLogs"`
	Code       bool `json:"code"`
	legacy     bool
	admin      bool
}

const policyTTL = time.Minute

type policyCache struct {
	mu      sync.Mutex
	entries map[string]policyEntry
}

type policyEntry struct {
	access *AiAccess
	at     time.Time
}

func (s *Server) policy(ctx context.Context, projectID string) (*AiAccess, error) {
	s.policies.mu.Lock()
	if e, ok := s.policies.entries[projectID]; ok && time.Since(e.at) < policyTTL {
		s.policies.mu.Unlock()
		return e.access, nil
	}
	s.policies.mu.Unlock()
	t, err := s.targets(ctx, projectID)
	if err != nil {
		return nil, err
	}
	access := t.AiAccess
	if access == nil {
		access = legacyAccess()
	}
	access.admin = t.Me.Role == "admin"
	s.policies.mu.Lock()
	s.policies.entries[projectID] = policyEntry{access: access, at: time.Now()}
	s.policies.mu.Unlock()
	return access, nil
}

// allowed never refuses from the cache: right after an admin turns a class on, the retry must see it
func (s *Server) allowed(ctx context.Context, projectID string, class aiClass) (*AiAccess, error) {
	s.policies.mu.Lock()
	e, cached := s.policies.entries[projectID]
	cached = cached && time.Since(e.at) < policyTTL
	s.policies.mu.Unlock()
	access, err := s.policy(ctx, projectID)
	if err != nil || class.on(access) {
		return access, err
	}
	if cached {
		s.policies.mu.Lock()
		delete(s.policies.entries, projectID)
		s.policies.mu.Unlock()
		if access, err = s.policy(ctx, projectID); err != nil {
			return nil, err
		}
	}
	if !class.on(access) {
		return access, class.refusal(access.admin)
	}
	return access, nil
}

type aiClass struct {
	label, title string
	on           func(*AiAccess) bool
}

var (
	statsClass      = aiClass{"statistics and insights", "Statistics and insights", func(a *AiAccess) bool { return a.Stats }}
	jobLogsClass    = aiClass{"job logs", "Job logs", func(a *AiAccess) bool { return a.JobLogs }}
	sampleRowsClass = aiClass{"per-sample data", "Per-sample data", func(a *AiAccess) bool { return a.SampleRows }}
	visualsClass    = aiClass{"sample visualizations", "Sample visualizations", func(a *AiAccess) bool { return a.Visuals }}
	codeClass       = aiClass{"integration code", "Integration code", func(a *AiAccess) bool { return a.Code }}
	allClasses      = []aiClass{statsClass, sampleRowsClass, visualsClass, jobLogsClass, codeClass}
)

// refusal matches the server's wording so assistants see one message whichever side refused
func (c aiClass) refusal(admin bool) error {
	what := fmt.Sprintf("AI access to %s is turned off for this project, so Tensorleap did not share it.", c.label)
	where := fmt.Sprintf("%q for this project (gear icon, top right > AI ACCESS)", c.title)
	if admin {
		return fmt.Errorf("%s You are an admin: turn on %s.", what, where)
	}
	return fmt.Errorf("%s Ask a Tensorleap admin to turn on %s.", what, where)
}

// servers that predate AI access controls serve everything, like an unrestricted new server
func legacyAccess() *AiAccess {
	return &AiAccess{Stats: true, SampleRows: true, Visuals: true, JobLogs: true, Code: true, legacy: true}
}

func (a *AiAccess) blocked() []string {
	var off []string
	for _, c := range allClasses {
		if !c.on(a) {
			off = append(off, c.label)
		}
	}
	return off
}

func serverMessage(body string) string {
	var payload struct {
		Error string `json:"error"`
		Code  string `json:"code"`
	}
	if json.Unmarshal([]byte(body), &payload) == nil && payload.Error != "" {
		return payload.Error
	}
	return body
}
