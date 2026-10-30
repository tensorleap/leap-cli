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

const (
	connectTimeout = 20 * time.Second
	progressGrace  = 5 * time.Second
)

type Server struct {
	signin   func() (*Client, error)
	version  string
	client   *Client
	remote   *sdk.ClientSession
	mu       sync.Mutex
	progress map[any]*sdk.ServerSession
	// the assistant's name and version, once it said them; a client signed in later is told too
	assistant [2]string
	// one tl_status call at a time tries to connect while leap mcp is not connected
	connecting sync.Mutex
}

// NewServer bridges the Tensorleap server client points at
func NewServer(client *Client, version string) *sdk.Server {
	return NewSignedInServer(func() (*Client, error) { return client, nil }, version)
}

// NewSignedInServer bridges the server signin names. An assistant starts leap mcp once per session, often
// before the user signed in or installed a server, so when that fails it still starts with tl_status alone,
// and every tl_status call runs signin and connects again. Once one succeeds, the server's tools replace it
// and the assistant is told the tool list changed: a later `leap auth login` needs no assistant restart.
func NewSignedInServer(signin func() (*Client, error), version string) *sdk.Server {
	s := &Server{signin: signin, version: version, progress: map[any]*sdk.ServerSession{}}
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	b, err := s.dial(ctx)
	if err != nil {
		// the first tl_status call connects again and reports the problem as it is by then
		return s.unavailable()
	}
	srv := sdk.NewServer(&sdk.Implementation{Name: "tensorleap", Title: "Tensorleap", Version: version},
		&sdk.ServerOptions{Instructions: b.instructions, InitializedHandler: s.initialized})
	s.addTo(srv, b)
	return srv
}

// the server was connected before the assistant said who it is, so it's told on every later request
func (s *Server) initialized(_ context.Context, req *sdk.InitializedRequest) {
	info := req.Session.InitializeParams().ClientInfo
	if info == nil {
		return
	}
	s.mu.Lock()
	s.assistant = [2]string{info.Name, info.Version}
	client := s.client
	s.mu.Unlock()
	if client != nil {
		client.setAssistant(info.Name, info.Version)
	}
}

// bridged is what a connected server offers, read in full before any of it is served
type bridged struct {
	client       *Client
	remote       *sdk.ClientSession
	instructions string
	tools        []*sdk.Tool
	resources    []*sdk.Resource
	prompts      []*sdk.Prompt
}

func (s *Server) dial(ctx context.Context) (*bridged, error) {
	client, err := s.signin()
	if err != nil {
		return nil, err
	}
	if client.failure != nil {
		return nil, client.failure
	}
	client.version = headerValue(s.version)
	remote, err := client.connect(ctx, s.version, &sdk.ClientOptions{ProgressNotificationHandler: s.relayProgress})
	if err != nil {
		return nil, connectError(ctx, client, err)
	}
	b := &bridged{client: client, remote: remote, instructions: remote.InitializeResult().Instructions}
	if err := b.list(ctx); err != nil {
		_ = remote.Close()
		return nil, err
	}
	return b, nil
}

func (b *bridged) list(ctx context.Context) error {
	for tool, err := range b.remote.Tools(ctx, nil) {
		if err != nil {
			return err
		}
		b.tools = append(b.tools, tool)
	}
	for res, err := range b.remote.Resources(ctx, nil) {
		if err != nil {
			return err
		}
		b.resources = append(b.resources, res)
	}
	for p, err := range b.remote.Prompts(ctx, nil) {
		if err != nil {
			return err
		}
		b.prompts = append(b.prompts, p)
	}
	return nil
}

func (s *Server) addTo(srv *sdk.Server, b *bridged) {
	s.mu.Lock()
	s.client, s.remote = b.client, b.remote
	b.client.setAssistant(s.assistant[0], s.assistant[1])
	s.mu.Unlock()
	remote := b.remote
	for _, tool := range b.tools {
		srv.AddTool(tool, s.relayTool(tool.Name))
	}
	for _, res := range b.resources {
		srv.AddResource(res, func(ctx context.Context, req *sdk.ReadResourceRequest) (*sdk.ReadResourceResult, error) {
			return remote.ReadResource(ctx, req.Params)
		})
	}
	for _, p := range b.prompts {
		srv.AddPrompt(p, func(ctx context.Context, req *sdk.GetPromptRequest) (*sdk.GetPromptResult, error) {
			return remote.GetPrompt(ctx, req.Params)
		})
	}
	s.addExportTool(srv)
}

func (s *Server) connected() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.remote != nil
}

func (s *Server) relayTool(name string) sdk.ToolHandler {
	return func(ctx context.Context, req *sdk.CallToolRequest) (*sdk.CallToolResult, error) {
		params := &sdk.CallToolParams{Name: name, Arguments: req.Params.Arguments, Meta: req.Params.Meta}
		if token := req.Params.GetProgressToken(); token != nil && req.Session != nil {
			s.mu.Lock()
			s.progress[token] = req.Session
			s.mu.Unlock()
			// progress can arrive just after the result, so the route outlives the call briefly
			defer time.AfterFunc(progressGrace, func() {
				s.mu.Lock()
				delete(s.progress, token)
				s.mu.Unlock()
			})
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

func (c *Client) connect(ctx context.Context, version string, opts *sdk.ClientOptions) (*sdk.ClientSession, error) {
	return sdk.NewClient(&sdk.Implementation{Name: "leap-mcp", Version: version}, opts).Connect(ctx, &sdk.StreamableClientTransport{
		Endpoint:             c.BaseURL + "/mcp",
		HTTPClient:           c.authed(),
		DisableStandaloneSSE: true,
	}, nil)
}

// connectError names the fix: a REST call says whether the key, the network or the server's age is the problem
func connectError(ctx context.Context, c *Client, err error) error {
	if c.failure != nil {
		return c.failure
	}
	var t targetsResponse
	if restErr := c.Post(ctx, "analysis-export/listTargets", map[string]any{}, &t); restErr != nil {
		return explain(restErr)
	}
	return fmt.Errorf("this Tensorleap server (%s) has no MCP endpoint yet; upgrade the server (%v)", c.UIBase(), err)
}

// Probe connects the way an assistant will and counts the tools it would see; leap's own export tool included
func Probe(ctx context.Context, c *Client) (int, error) {
	remote, err := c.connect(ctx, "probe", nil)
	if err != nil {
		return 0, connectError(ctx, c, err)
	}
	defer remote.Close()
	n := 1
	for _, err := range remote.Tools(ctx, nil) {
		if err != nil {
			return 0, err
		}
		n++
	}
	return n, nil
}

// unavailable still starts, so the assistant can tell the user what is wrong instead of only showing that
// the server failed to start, and each tl_status call tries again, so fixing it needs no restart
func (s *Server) unavailable() *sdk.Server {
	srv := sdk.NewServer(&sdk.Implementation{Name: "tensorleap", Title: "Tensorleap", Version: s.version},
		&sdk.ServerOptions{
			Instructions:       "Tensorleap MCP is not connected yet. Call tl_status and relay its message to the user; once they fixed what it names, call tl_status again to connect.",
			InitializedHandler: s.initialized,
			// the server's resources and prompts arrive with its tools, once tl_status connects
			Capabilities: &sdk.ServerCapabilities{Logging: &sdk.LoggingCapabilities{},
				Tools: &sdk.ToolCapabilities{ListChanged: true}, Resources: &sdk.ResourceCapabilities{ListChanged: true}, Prompts: &sdk.PromptCapabilities{ListChanged: true}},
		})
	closed := false
	srv.AddTool(&sdk.Tool{Name: "tl_status", Description: "Why the Tensorleap tools are unavailable and how to fix it; connects once it is fixed.",
		InputSchema: map[string]any{"type": "object"},
		Annotations: &sdk.ToolAnnotations{Title: "Tensorleap status", ReadOnlyHint: true, OpenWorldHint: &closed}},
		func(ctx context.Context, req *sdk.CallToolRequest) (*sdk.CallToolResult, error) {
			s.connecting.Lock()
			defer s.connecting.Unlock()
			if s.connected() {
				return s.relayTool("tl_status")(ctx, req)
			}
			dialCtx, cancel := context.WithTimeout(ctx, connectTimeout)
			b, err := s.dial(dialCtx)
			cancel()
			if err != nil {
				return &sdk.CallToolResult{IsError: true, Content: []sdk.Content{&sdk.TextContent{
					Text: fmt.Sprintf("leap mcp can't serve Tensorleap: %v. Once that is fixed, call tl_status again; leap mcp connects then, without restarting the assistant.", err)}}}, nil
			}
			// the server's own tl_status replaces this one, with the rest of its tools
			s.addTo(srv, b)
			res, err := s.relayTool("tl_status")(ctx, req)
			if err != nil {
				return nil, err
			}
			note := "leap mcp is connected to Tensorleap now and its tools are available (if your assistant doesn't list them, restart it)."
			if b.instructions != "" {
				note += "\n\nServer instructions: " + b.instructions
			}
			res.Content = append([]sdk.Content{&sdk.TextContent{Text: note}}, res.Content...)
			return res, nil
		})
	return srv
}
