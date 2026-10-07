package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/api"
)

type Client struct {
	BaseURL string
	apiKey  string
	http    *http.Client
	failure error
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{BaseURL: strings.TrimRight(baseURL, "/"), apiKey: apiKey, http: api.NewDefaultClient()}
}

// NewFailedClient answers every call with a startup problem, so the assistant can tell the
// user why instead of only showing that the server failed to start
func NewFailedClient(err error) *Client {
	return &Client{failure: err}
}

// UIBase is where the web UI lives: hosted installs serve the API from api.<tenant>.tensorleap.ai
// and the UI from <tenant>.tensorleap.ai
func (c *Client) UIBase() string {
	return strings.TrimSuffix(api.ChangeToUIUrl(c.BaseURL), "/")
}

type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("tensorleap API returned %d: %s", e.Status, e.Body)
}

// authTransport signs every request as the user's leap API key and marks it as leap mcp's
type authTransport struct {
	key  string
	next http.RoundTripper
}

func (t authTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r = r.Clone(r.Context())
	r.Header.Set("X-TL-Client", "leap-mcp")
	if t.key != "" {
		r.Header.Set("Authorization", "Bearer "+t.key)
	}
	return t.next.RoundTrip(r)
}

func (c *Client) authed() *http.Client {
	next := c.http.Transport
	if next == nil {
		next = http.DefaultTransport
	}
	return &http.Client{Transport: authTransport{key: c.apiKey, next: next}}
}

func (c *Client) do(ctx context.Context, path string, body any) (*http.Response, error) {
	if c.failure != nil {
		return nil, c.failure
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/"+path, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.authed().Do(req)
	if err != nil {
		return nil, describeTransportError(c.BaseURL, err)
	}
	if resp.StatusCode >= 300 {
		defer resp.Body.Close()
		raw, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, &APIError{Status: resp.StatusCode, Body: truncate(string(raw), 300)}
	}
	return resp, nil
}

func (c *Client) Post(ctx context.Context, path string, body, out any) error {
	resp, err := c.do(ctx, path, body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if out == nil || len(bytes.TrimSpace(raw)) == 0 {
		return nil
	}
	return json.Unmarshal(raw, out)
}

// PostStream returns the response body unread, for large downloads such as the export bundle
func (c *Client) PostStream(ctx context.Context, path string, body any) (io.ReadCloser, error) {
	resp, err := c.do(ctx, path, body)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func describeTransportError(base string, err error) error {
	if errors.Is(err, api.ErrAuth) {
		return errors.New("the server rejected your API key; run `leap auth login` (or `leap auth select <env>`)")
	}
	var netErr net.Error
	if errors.As(err, &netErr) || strings.Contains(err.Error(), "connection refused") {
		return fmt.Errorf("cannot reach the Tensorleap server at %s: is the server running and, if it is remote, is your SSH/SSM tunnel or VPN up? (%v)", base, err)
	}
	return err
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}
