package mcp

import (
	"context"
	"encoding/json"
	"errors"
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

var errLegacyServer = errors.New("this Tensorleap server predates AI access controls, so its data owner cannot limit what an AI assistant receives; upgrade the server, or start `leap mcp --allow-legacy-server` if your organization accepts that")

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
		if !s.allowLegacy {
			return nil, errLegacyServer
		}
		access = &AiAccess{Stats: true, SampleRows: true, Visuals: true, JobLogs: true, Code: true, legacy: true}
	}
	s.policies.mu.Lock()
	s.policies.entries[projectID] = policyEntry{access: access, at: time.Now()}
	s.policies.mu.Unlock()
	return access, nil
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
	return fmt.Errorf("AI access to %s is turned off for this project. An admin can enable it in Settings > AI access", what)
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
