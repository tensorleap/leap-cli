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
	s.policies.mu.Lock()
	s.policies.entries[projectID] = policyEntry{access: access, at: time.Now()}
	s.policies.mu.Unlock()
	return access, nil
}

// servers that predate AI access controls serve everything, like an unrestricted new server
func legacyAccess() *AiAccess {
	return &AiAccess{Stats: true, SampleRows: true, Visuals: true, JobLogs: true, Code: true, legacy: true}
}

func (a *AiAccess) blocked() []string {
	var off []string
	for _, c := range []struct {
		on   bool
		name string
	}{{a.Stats, "statistics and insights"}, {a.SampleRows, "per-sample data"}, {a.Visuals, "sample visualizations"}, {a.JobLogs, "job logs"}, {a.Code, "integration code"}} {
		if !c.on {
			off = append(off, c.name)
		}
	}
	return off
}

func (s *Server) requireStats(ctx context.Context, projectID string) error {
	access, err := s.policy(ctx, projectID)
	if err != nil {
		return err
	}
	if !access.Stats {
		return disabled("statistics and insights")
	}
	return nil
}

func disabled(what string) error {
	return fmt.Errorf("A Tensorleap admin turned off AI access to %s for this project. It can be turned back on in the gear menu > AI access.", what)
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
