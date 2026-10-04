package mcp

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
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
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{BaseURL: strings.TrimRight(baseURL, "/"), apiKey: apiKey, http: api.NewDefaultClient()}
}

func (c *Client) UIBase() string {
	return strings.TrimSuffix(c.BaseURL, "/api/v2")
}

type APIError struct {
	Status int
	Body   string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("tensorleap API returned %d: %s", e.Status, e.Body)
}

func (c *Client) Post(ctx context.Context, path string, body, out any) error {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/"+path, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TL-Client", "leap-mcp")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return describeTransportError(c.BaseURL, err)
	}
	defer resp.Body.Close()
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return &APIError{Status: resp.StatusCode, Body: truncate(string(raw), 300)}
	}
	if out == nil || len(bytes.TrimSpace(raw)) == 0 {
		return nil
	}
	return json.Unmarshal(raw, out)
}

func (c *Client) Download(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("cannot download a stored artifact (%v); signed URLs point at the server's storage path, so point `leap auth` at the ingress URL (port 4589), not at node-server directly", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		return nil, fmt.Errorf("artifact download returned %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func decompress(b []byte) ([]byte, error) {
	switch {
	case bytes.HasPrefix(b, []byte("PK")):
		zr, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
		if err != nil {
			return nil, err
		}
		for _, f := range zr.File {
			if f.FileInfo().IsDir() {
				continue
			}
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
		return nil, errors.New("empty zip archive")
	case bytes.HasPrefix(b, []byte{0x1f, 0x8b}):
		gr, err := gzip.NewReader(bytes.NewReader(b))
		if err != nil {
			return nil, err
		}
		defer gr.Close()
		return io.ReadAll(gr)
	}
	return b, nil
}

func describeTransportError(base string, err error) error {
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
