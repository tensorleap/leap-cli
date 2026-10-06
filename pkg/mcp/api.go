package mcp

import (
	"context"
	"errors"
	"fmt"
	"strings"
)

type targetsResponse struct {
	ContractVersion int       `json:"contractVersion"`
	AiAccess        *AiAccess `json:"aiAccess"`
	Me              struct {
		Email, Name, TeamID, Role string
	} `json:"me"`
	Projects []struct {
		Cid      string    `json:"cid"`
		Name     string    `json:"name"`
		AiAccess *AiAccess `json:"aiAccess"`
	} `json:"projects"`
	Versions []struct {
		Cid          string   `json:"cid"`
		Name         string   `json:"name"`
		SerialNumber *float64 `json:"serialNumber"`
		CreatedAt    string   `json:"createdAt"`
		Evaluated    bool     `json:"evaluated"`
		HasInsights  bool     `json:"hasInsights"`
	} `json:"versions"`
}

func (s *Server) targets(ctx context.Context, projectID string) (*targetsResponse, error) {
	body := map[string]any{}
	if projectID != "" {
		body["projectId"] = projectID
	}
	var out targetsResponse
	if err := s.client.Post(ctx, "analysis-export/listTargets", body, &out); err != nil {
		return nil, explain(err)
	}
	if out.ContractVersion != supportedContract {
		return nil, fmt.Errorf("this server speaks analysis-export contract v%d but this leap CLI expects v%d; upgrade whichever side is older", out.ContractVersion, supportedContract)
	}
	return &out, nil
}

func explain(err error) error {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		switch apiErr.Status {
		case 401:
			return errors.New("the server rejected your API key; run `leap auth login` (or `leap auth select <env>`)")
		case 400, 403:
			return errors.New(serverMessage(apiErr.Body))
		case 404:
			if strings.Contains(apiErr.Body, "Cannot POST") {
				return errors.New("this Tensorleap server has no analysis-export API; upgrade the server")
			}
			return fmt.Errorf("%s Check the ids with tl_list_projects / tl_list_versions.", strings.TrimSpace(serverMessage(apiErr.Body)))
		}
		if apiErr.Status >= 500 {
			return fmt.Errorf("the Tensorleap server failed on this request (%d): %s", apiErr.Status, serverMessage(apiErr.Body))
		}
	}
	return err
}

type Visualizer struct {
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	ArgNames []string `json:"argNames"`
}

type exportedInsight struct {
	Cid             string         `json:"cid"`
	Index           float64        `json:"index"`
	Status          string         `json:"status"`
	Description     string         `json:"description"`
	InsightType     map[string]any `json:"insightType"`
	CsvURL          string         `json:"csvUrl"`
	ClusterBlobURL  string         `json:"clusterBlobUrl"`
	TopPanelURL     string         `json:"topPanelUrl"`
	FixingCsvURL    string         `json:"fixingCsvUrl"`
	AnalyzeLinkPath string         `json:"analyzeLinkPath"`
}

type exportResponse struct {
	ContractVersion      int                 `json:"contractVersion"`
	DeepLinkPath         string              `json:"deepLinkPath"`
	PopulationCsvURL     string              `json:"populationCsvUrl"`
	IntegrationCodeURL   string              `json:"integrationCodeUrl"`
	IntegrationEntryFile string              `json:"integrationEntryFile"`
	PredictionLabels     map[string][]string `json:"predictionLabels"`
	Visualizers          []Visualizer        `json:"visualizers"`
	Version              struct {
		Name         string   `json:"name"`
		SerialNumber *float64 `json:"serialNumber"`
	} `json:"version"`
	Insights []exportedInsight `json:"insights"`
}

func (s *Server) export(ctx context.Context, projectID, versionID string) (*exportResponse, error) {
	var e exportResponse
	if err := s.client.Post(ctx, "analysis-export/exportAnalysis", map[string]any{"projectId": projectID, "versionId": versionID}, &e); err != nil {
		return nil, explain(err)
	}
	return &e, nil
}

// internal storage paths and UI filter state; everything else the engine computed is worth reading
func (s *Server) population(ctx context.Context, versionID, url string) (*Population, error) {
	if url == "" {
		return nil, errors.New("the version has no population file")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if p, ok := s.pops[versionID]; ok {
		return p, nil
	}
	b, err := s.client.Download(ctx, url)
	if err != nil {
		return nil, err
	}
	if b, err = decompress(b); err != nil {
		return nil, err
	}
	p, err := NewPopulation(b)
	if err != nil {
		return nil, err
	}
	s.pops[versionID] = p
	return p, nil
}

// StatusLines is the one-glance state printed by `leap mcp config` and by a terminal run of `leap mcp`
func StatusLines(ctx context.Context, c *Client) []string {
	s := &Server{client: c}
	t, err := s.targets(ctx, "")
	if err != nil {
		return []string{fmt.Sprintf("Server: %s, not reachable right now: %v", c.UIBase(), err)}
	}
	lines := []string{fmt.Sprintf("Server: %s (%s)", c.UIBase(), t.Me.Email)}
	if t.AiAccess == nil {
		return append(lines, "AI access: this server predates AI access controls; assistants can read everything")
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
