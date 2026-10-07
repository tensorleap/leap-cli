package mcp

import (
	"context"
	"fmt"
	"sync"
	"time"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// leap mcp is a bridge: the Tensorleap server hosts the MCP tools (/api/v2/mcp) and this process
// relays them to an assistant over stdio, signed with the user's leap login. It adds the one tool a
// server cannot offer, tl_export_analysis, which writes the analysis to the user's disk.

const connectTimeout = 20 * time.Second

type Server struct {
	client   *Client
	remote   *sdk.ClientSession
	mu       sync.Mutex
	progress map[any]*sdk.ServerSession
}

func NewServer(client *Client, version string) *sdk.Server {
	s := &Server{client: client, progress: map[any]*sdk.ServerSession{}}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	srv, err := s.bridge(ctx, version)
	if err != nil {
		return unavailable(version, err)
	}
	return srv
}

func (s *Server) bridge(ctx context.Context, version string) (*sdk.Server, error) {
	if s.client.failure != nil {
		return nil, s.client.failure
	}
	c := sdk.NewClient(&sdk.Implementation{Name: "leap-mcp", Version: version}, &sdk.ClientOptions{
		ProgressNotificationHandler: s.relayProgress,
	})
	remote, err := c.Connect(ctx, &sdk.StreamableClientTransport{
		Endpoint:             s.client.BaseURL + "/mcp",
		HTTPClient:           s.client.authed(),
		DisableStandaloneSSE: true,
	}, nil)
	if err != nil {
		return nil, s.connectError(ctx, err)
	}
	s.remote = remote
	init := remote.InitializeResult()
	srv := sdk.NewServer(&sdk.Implementation{Name: "tensorleap", Title: "Tensorleap", Version: version},
		&sdk.ServerOptions{Instructions: init.Instructions})

	for tool, err := range remote.Tools(ctx, nil) {
		if err != nil {
			return nil, err
		}
		srv.AddTool(tool, s.relayTool(tool.Name))
	}
	for res, err := range remote.Resources(ctx, nil) {
		if err != nil {
			return nil, err
		}
		srv.AddResource(res, func(ctx context.Context, req *sdk.ReadResourceRequest) (*sdk.ReadResourceResult, error) {
			return remote.ReadResource(ctx, req.Params)
		})
	}
	for p, err := range remote.Prompts(ctx, nil) {
		if err != nil {
			return nil, err
		}
		srv.AddPrompt(p, func(ctx context.Context, req *sdk.GetPromptRequest) (*sdk.GetPromptResult, error) {
			return remote.GetPrompt(ctx, req.Params)
		})
	}
	s.addExportTool(srv)
	return srv, nil
}

func (s *Server) relayTool(name string) sdk.ToolHandler {
	return func(ctx context.Context, req *sdk.CallToolRequest) (*sdk.CallToolResult, error) {
		params := &sdk.CallToolParams{Name: name, Arguments: req.Params.Arguments, Meta: req.Params.Meta}
		if token := req.Params.GetProgressToken(); token != nil && req.Session != nil {
			s.mu.Lock()
			s.progress[token] = req.Session
			s.mu.Unlock()
			defer func() {
				s.mu.Lock()
				delete(s.progress, token)
				s.mu.Unlock()
			}()
		}
		res, err := s.remote.CallTool(ctx, params)
		if err != nil {
			return nil, fmt.Errorf("the Tensorleap server did not answer %s: %w", name, err)
		}
		return res, nil
	}
}

// progress from a long tool (tl_wait_for_job) goes back to the assistant that asked
func (s *Server) relayProgress(ctx context.Context, req *sdk.ProgressNotificationClientRequest) {
	s.mu.Lock()
	session := s.progress[req.Params.ProgressToken]
	s.mu.Unlock()
	if session != nil {
		_ = session.NotifyProgress(ctx, req.Params)
	}
}

// connectError names the fix: a REST call says whether the key, the network or the server's age is the problem
func (s *Server) connectError(ctx context.Context, err error) error {
	var t targetsResponse
	if restErr := s.client.Post(ctx, "analysis-export/listTargets", map[string]any{}, &t); restErr != nil {
		return explain(restErr)
	}
	return fmt.Errorf("this Tensorleap server (%s) has no MCP endpoint yet; upgrade the server (%v)", s.client.UIBase(), err)
}

// unavailable still starts, so the assistant can tell the user what is wrong instead of only
// showing that the server failed to start
func unavailable(version string, cause error) *sdk.Server {
	srv := sdk.NewServer(&sdk.Implementation{Name: "tensorleap", Title: "Tensorleap", Version: version},
		&sdk.ServerOptions{Instructions: "Tensorleap MCP is not connected. Call tl_status and relay its message to the user."})
	closed := false
	srv.AddTool(&sdk.Tool{Name: "tl_status", Description: "Why the Tensorleap tools are unavailable and how to fix it.",
		InputSchema: map[string]any{"type": "object"},
		Annotations: &sdk.ToolAnnotations{Title: "Tensorleap status", ReadOnlyHint: true, OpenWorldHint: &closed}},
		func(context.Context, *sdk.CallToolRequest) (*sdk.CallToolResult, error) {
			return &sdk.CallToolResult{IsError: true, Content: []sdk.Content{&sdk.TextContent{
				Text: fmt.Sprintf("leap mcp can't serve Tensorleap: %v. Fix it in a terminal, then restart the assistant.", cause)}}}, nil
		})
	return srv
}
