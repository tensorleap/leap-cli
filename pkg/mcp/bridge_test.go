package mcp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type fakeTensorleap struct {
	auth    []string
	bundle  []byte
	noMCP   bool
	badKey  bool
	targets string
}

// a node-server stand-in: the MCP endpoint served by go-sdk's own stateless handler, plus the REST calls leap mcp makes
func (f *fakeTensorleap) start(t *testing.T) *httptest.Server {
	t.Helper()
	remote := sdk.NewServer(&sdk.Implementation{Name: "tensorleap", Version: "test"}, &sdk.ServerOptions{Instructions: "Call tl_status first."})
	type statusOut struct {
		User string `json:"user"`
	}
	sdk.AddTool(remote, &sdk.Tool{Name: "tl_status", Description: "who you are"}, func(context.Context, *sdk.CallToolRequest, struct{}) (*sdk.CallToolResult, statusOut, error) {
		return nil, statusOut{User: "a@b.c"}, nil
	})
	sdk.AddTool(remote, &sdk.Tool{Name: "tl_wait_for_job", Description: "waits"}, func(ctx context.Context, req *sdk.CallToolRequest, _ struct{}) (*sdk.CallToolResult, statusOut, error) {
		if token := req.Params.GetProgressToken(); token != nil {
			if err := req.Session.NotifyProgress(ctx, &sdk.ProgressNotificationParams{ProgressToken: token, Progress: 1, Total: 2, Message: "Evaluate: STARTED"}); err != nil {
				panic(err)
			}
		}
		return nil, statusOut{User: "done"}, nil
	})
	remote.AddResource(&sdk.Resource{URI: "tensorleap://glossary", Name: "glossary", MIMEType: "text/markdown"},
		func(context.Context, *sdk.ReadResourceRequest) (*sdk.ReadResourceResult, error) {
			return &sdk.ReadResourceResult{Contents: []*sdk.ResourceContents{{URI: "tensorleap://glossary", Text: "Failure Mode"}}}, nil
		})
	remote.AddPrompt(&sdk.Prompt{Name: "analyze_version"}, func(context.Context, *sdk.GetPromptRequest) (*sdk.GetPromptResult, error) {
		return &sdk.GetPromptResult{Messages: []*sdk.PromptMessage{{Role: "user", Content: &sdk.TextContent{Text: "start with tl_status"}}}}, nil
	})
	mcpHandler := sdk.NewStreamableHTTPHandler(func(*http.Request) *sdk.Server { return remote }, &sdk.StreamableHTTPOptions{Stateless: true})
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.auth = append(f.auth, r.Header.Get("Authorization")+"|"+r.Header.Get("X-TL-Client"))
		if f.badKey {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		switch r.URL.Path {
		case "/api/v2/mcp":
			if f.noMCP {
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Cannot POST /api/v2/mcp"))
				return
			}
			mcpHandler.ServeHTTP(w, r)
		case "/api/v2/analysis-export/listTargets":
			_, _ = w.Write([]byte(f.targets))
		case "/api/v2/analysis-export/exportBundle":
			_, _ = w.Write(f.bundle)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	t.Cleanup(ts.Close)
	return ts
}

func connectBridge(t *testing.T, srv *sdk.Server) *sdk.ClientSession {
	t.Helper()
	a, b := sdk.NewInMemoryTransports()
	if _, err := srv.Connect(context.Background(), a, nil); err != nil {
		t.Fatal(err)
	}
	session, err := sdk.NewClient(&sdk.Implementation{Name: "assistant"}, nil).Connect(context.Background(), b, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = session.Close() })
	return session
}

func textOf(r *sdk.CallToolResult) string {
	var b strings.Builder
	for _, c := range r.Content {
		if t, ok := c.(*sdk.TextContent); ok {
			b.WriteString(t.Text)
		}
	}
	return b.String()
}

func TestBridgeRelaysTheServersToolsResourcesAndPrompts(t *testing.T) {
	f := &fakeTensorleap{}
	ts := f.start(t)
	session := connectBridge(t, NewServer(NewClient(ts.URL+"/api/v2", "key-1"), "test"))
	if got := session.InitializeResult().Instructions; got != "Call tl_status first." {
		t.Fatalf("the server's instructions must reach the assistant, got %q", got)
	}
	names := []string{}
	for tool, err := range session.Tools(context.Background(), nil) {
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, tool.Name)
	}
	if strings.Join(names, ",") != "tl_export_analysis,tl_status,tl_wait_for_job" {
		t.Fatalf("tools: %v", names)
	}
	res, err := session.CallTool(context.Background(), &sdk.CallToolParams{Name: "tl_status", Arguments: map[string]any{}})
	if err != nil || !strings.Contains(textOf(res), "a@b.c") {
		t.Fatalf("tl_status relayed: %v %+v", err, res)
	}
	g, err := session.ReadResource(context.Background(), &sdk.ReadResourceParams{URI: "tensorleap://glossary"})
	if err != nil || g.Contents[0].Text != "Failure Mode" {
		t.Fatalf("glossary relayed: %v %+v", err, g)
	}
	p, err := session.GetPrompt(context.Background(), &sdk.GetPromptParams{Name: "analyze_version"})
	if err != nil || len(p.Messages) != 1 {
		t.Fatalf("prompt relayed: %v %+v", err, p)
	}
	for _, a := range f.auth {
		if a != "Bearer key-1|leap-mcp" {
			t.Fatalf("every request must carry the user's key and mark leap mcp, got %q", a)
		}
	}
}

func TestBridgeRelaysProgressOfLongTools(t *testing.T) {
	f := &fakeTensorleap{}
	ts := f.start(t)
	srv := NewServer(NewClient(ts.URL+"/api/v2", "k"), "test")
	a, b := sdk.NewInMemoryTransports()
	if _, err := srv.Connect(context.Background(), a, nil); err != nil {
		t.Fatal(err)
	}
	got := make(chan string, 1)
	session, err := sdk.NewClient(&sdk.Implementation{Name: "assistant"}, &sdk.ClientOptions{
		ProgressNotificationHandler: func(_ context.Context, req *sdk.ProgressNotificationClientRequest) {
			got <- req.Params.Message
		}}).Connect(context.Background(), b, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	// go-sdk v1.4.0's SetProgressToken drops the token when Meta is nil, so it is set directly
	params := &sdk.CallToolParams{Name: "tl_wait_for_job", Arguments: map[string]any{}, Meta: sdk.Meta{"progressToken": "tok-1"}}
	if _, err := session.CallTool(context.Background(), params); err != nil {
		t.Fatal(err)
	}
	if msg := <-got; msg != "Evaluate: STARTED" {
		t.Fatalf("progress: %q", msg)
	}
}

func TestBridgeExplainsWhyItCannotServe(t *testing.T) {
	cases := map[string]struct {
		f    *fakeTensorleap
		want string
	}{
		"server too old": {&fakeTensorleap{noMCP: true, targets: `{"aiAccess":{"stats":true}}`}, "has no MCP endpoint yet; upgrade the server"},
		"bad API key":    {&fakeTensorleap{badKey: true}, "rejected your API key"},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			ts := c.f.start(t)
			session := connectBridge(t, NewServer(NewClient(ts.URL+"/api/v2", "k"), "test"))
			res, err := session.CallTool(context.Background(), &sdk.CallToolParams{Name: "tl_status", Arguments: map[string]any{}})
			if err != nil || !res.IsError || !strings.Contains(textOf(res), c.want) {
				t.Fatalf("want %q, got %v %q", c.want, err, textOf(res))
			}
		})
	}
}

func bundleOf(t *testing.T, files map[string]string, order ...string) []byte {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	for _, name := range order {
		body := files[name]
		if err := tw.WriteHeader(&tar.Header{Name: name, Mode: 0o644, Size: int64(len(body)), Typeflag: tar.TypeReg}); err != nil {
			t.Fatal(err)
		}
		_, _ = tw.Write([]byte(body))
	}
	_ = tw.Close()
	_ = gz.Close()
	return buf.Bytes()
}

const bundleManifest = `{"projectId":"p1","versionId":"v1","version":"model","insightsPanelLink":"http://ui/p","filesWritten":3,
 "insights":[{"index":1,"type":"low_performance","name":"Failure Mode","dir":"insight_1_low_performance","summary":{"groupSize":2},
   "samples":[{"id":"training_1","rank":1,"rendered":true,"files":["insight_1_low_performance/samples/training_1/image/Image/assets/data.png"]},{"id":"training_2","rank":2,"rendered":false}]}]}`

func goodBundle(t *testing.T, manifest string) []byte {
	return bundleOf(t, map[string]string{
		"population.csv": "sample_id\ntraining_1\n",
		"insight_1_low_performance/samples/training_1/image/Image/assets/data.png": "png",
		"../../escape.txt": "x",
		"manifest.json":    manifest,
	}, "population.csv", "insight_1_low_performance/samples/training_1/image/Image/assets/data.png", "../../escape.txt", "manifest.json")
}

func exportTo(t *testing.T, f *fakeTensorleap, dir string) (ExportResult, error) {
	t.Helper()
	ts := f.start(t)
	s := &Server{client: NewClient(ts.URL+"/api/v2", "k")}
	_, res, err := s.exportAnalysis(context.Background(), nil, ExportIn{ProjectID: "mnist", VersionID: "latest", Dir: dir})
	return res, err
}

func TestExportUnpacksTheServerBundle(t *testing.T) {
	root := t.TempDir()
	dir := filepath.Join(root, "out")
	res, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, bundleManifest)}, dir)
	if err != nil {
		t.Fatal(err)
	}
	if res.Insights[0].GroupSize != 2 || res.Insights[0].Exported != 2 || res.Insights[0].Rendered != 1 || res.Version != "model" {
		t.Fatalf("result: %+v", res)
	}
	if _, err := os.Stat(filepath.Join(dir, "insight_1_low_performance/samples/training_1/image/Image/assets/data.png")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "escape.txt")); err == nil {
		t.Fatal("a bundle entry must never land outside the export dir")
	}
	var m map[string]any
	b, _ := os.ReadFile(res.Manifest)
	if json.Unmarshal(b, &m) != nil || m["dir"] != dir || m["populationCsv"] != nil && filepath.IsAbs(m["populationCsv"].(string)) {
		t.Fatalf("manifest must keep the server's relative paths and gain the local dir: %s", b)
	}
	if _, err := os.Stat(filepath.Join(dir, stagingDir)); err == nil {
		t.Fatal("the staging dir must be removed")
	}
}

func TestExportKeepsThePreviousExportWhenTheNewOneIsCutShort(t *testing.T) {
	dir := t.TempDir()
	if _, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, bundleManifest)}, dir); err != nil {
		t.Fatal(err)
	}
	truncated := bundleOf(t, map[string]string{"population.csv": "new"}, "population.csv")
	if _, err := exportTo(t, &fakeTensorleap{bundle: truncated}, dir); err == nil || !strings.Contains(err.Error(), "cut short") {
		t.Fatalf("a bundle without manifest.json is an interrupted export, got %v", err)
	}
	if b, _ := os.ReadFile(filepath.Join(dir, "population.csv")); string(b) != "sample_id\ntraining_1\n" {
		t.Fatalf("the previous export must survive a failed refresh, got %q", b)
	}
	other := strings.Replace(bundleManifest, `"versionId":"v1"`, `"versionId":"v2"`, 1)
	if _, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, other)}, dir); err == nil || !strings.Contains(err.Error(), "another version") {
		t.Fatalf("another version's export must not be overwritten, got %v", err)
	}
	if _, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, bundleManifest)}, dir); err != nil {
		t.Fatalf("refreshing the same version: %v", err)
	}
}

func TestExportRefusesDirectoriesItDidNotWrite(t *testing.T) {
	web := t.TempDir()
	_ = os.WriteFile(filepath.Join(web, "manifest.json"), []byte(`{"name":"my pwa"}`), 0o644)
	if _, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, bundleManifest)}, web); err == nil || !strings.Contains(err.Error(), "not empty") {
		t.Fatalf("a foreign manifest.json must not pass as a previous export, got %v", err)
	}
	old := t.TempDir()
	_ = os.WriteFile(filepath.Join(old, "insights.json"), []byte("{}"), 0o644)
	_ = os.MkdirAll(filepath.Join(old, "insight_9_duplication"), 0o755)
	if _, err := exportTo(t, &fakeTensorleap{bundle: goodBundle(t, bundleManifest)}, old); err != nil {
		t.Fatalf("the pre-MCP skill's output dir must be accepted: %v", err)
	}
	if _, err := os.Stat(filepath.Join(old, "insight_9_duplication")); err == nil {
		t.Fatal("stale insight dirs must be removed")
	}
	if _, _, err := (&Server{client: NewClient("http://x", "k")}).exportAnalysis(context.Background(), nil, ExportIn{Dir: "relative"}); err == nil {
		t.Fatal("a relative dir must be refused")
	}
}
