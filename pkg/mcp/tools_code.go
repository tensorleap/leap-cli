package mcp

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"path"
	"sort"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	maxCodeArchive = 32 << 20
	maxCodeFile    = 64 << 10
)

type CodeIn struct {
	ProjectID string `json:"projectId" jsonschema:"project id or name"`
	VersionID string `json:"versionId" jsonschema:"version id or name, or \"latest\""`
	File      string `json:"file,omitempty" jsonschema:"path inside the integration code to read (default: the entry file)"`
}

type CodeFile struct {
	Path string `json:"path"`
	Size int    `json:"size"`
}

type CodeOut struct {
	EntryFile string     `json:"entryFile"`
	Files     []CodeFile `json:"files"`
	File      string     `json:"file,omitempty" jsonschema:"the file returned in content"`
	Content   string     `json:"content,omitempty"`
	Note      string     `json:"note,omitempty"`
	Reason    string     `json:"reason,omitempty"`
	NextStep  string     `json:"nextStep,omitempty"`
}

type codeArchive struct {
	entry string
	files map[string][]byte
}

// integrationCode always asks the server first: a withheld URL means the policy changed, and the cache must not outlive that
func (s *Server) integrationCode(ctx context.Context, e *exportResponse, versionID string, admin bool) (*codeArchive, error) {
	if e.IntegrationCodeURL == "" && e.IntegrationEntryFile != "" {
		return nil, codeClass.refusal(admin)
	}
	if e.IntegrationCodeURL == "" {
		return nil, errors.New("this version has no integration code on the server (it predates code snapshots)")
	}
	s.mu.Lock()
	cached, ok := s.code[versionID]
	s.mu.Unlock()
	if ok {
		return cached, nil
	}
	raw, err := s.client.Download(ctx, e.IntegrationCodeURL)
	if err != nil {
		return nil, err
	}
	arc, err := readCodeArchive(raw)
	if err != nil {
		return nil, fmt.Errorf("integration code archive unreadable: %w", err)
	}
	arc.entry = e.IntegrationEntryFile
	s.mu.Lock()
	// ponytail: one archive cached (up to 32MB); a session that hops versions re-downloads
	s.code = map[string]*codeArchive{versionID: arc}
	s.mu.Unlock()
	return arc, nil
}

func readCodeArchive(raw []byte) (*codeArchive, error) {
	if len(raw) > maxCodeArchive {
		return nil, fmt.Errorf("archive is %d MB, larger than the %d MB limit", len(raw)>>20, maxCodeArchive>>20)
	}
	var r io.Reader = bytes.NewReader(raw)
	if gz, err := gzip.NewReader(bytes.NewReader(raw)); err == nil {
		r = gz
	}
	arc := &codeArchive{files: map[string][]byte{}}
	tr := tar.NewReader(r)
	total := 0
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		name := path.Clean(h.Name)
		// an archive is customer-controlled input: keep every member inside its own tree
		if h.Typeflag != tar.TypeReg || path.IsAbs(name) || name == "." || name == ".." || strings.HasPrefix(name, "../") || strings.ContainsRune(name, '\\') {
			continue
		}
		b, err := io.ReadAll(io.LimitReader(tr, int64(maxCodeArchive-total)))
		if err != nil {
			return nil, err
		}
		if total += len(b); total >= maxCodeArchive {
			return nil, fmt.Errorf("archive expands beyond the %d MB limit", maxCodeArchive>>20)
		}
		arc.files[name] = b
	}
	if len(arc.files) == 0 {
		return nil, errors.New("no files in the archive")
	}
	return arc, nil
}

func (s *Server) getIntegrationCode(ctx context.Context, _ *sdk.CallToolRequest, in CodeIn) (*sdk.CallToolResult, CodeOut, error) {
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, CodeOut{}, err
	}
	access, err := s.allowed(ctx, in.ProjectID, codeClass)
	if err != nil {
		return nil, CodeOut{}, err
	}
	e, err := s.export(ctx, in.ProjectID, in.VersionID)
	if err != nil {
		return nil, CodeOut{}, err
	}
	arc, err := s.integrationCode(ctx, e, in.VersionID, access.admin)
	if err != nil {
		return nil, CodeOut{}, err
	}
	out := CodeOut{EntryFile: arc.entry, Files: []CodeFile{}}
	for name, b := range arc.files {
		out.Files = append(out.Files, CodeFile{Path: name, Size: len(b)})
	}
	sort.Slice(out.Files, func(i, j int) bool { return out.Files[i].Path < out.Files[j].Path })
	want := strings.TrimSpace(in.File)
	if want == "" {
		want = arc.entry
	}
	b, ok := arc.files[path.Clean(want)]
	if !ok {
		out.Reason, out.NextStep = "file-not-found", fmt.Sprintf("%q is not in the archive; pick one of files", want)
		return nil, out, nil
	}
	out.File = path.Clean(want)
	// scrub before truncating so a secret cut at the boundary cannot escape its pattern
	text := Scrub(string(b))
	if len(text) > maxCodeFile {
		out.Note = fmt.Sprintf("showing the first %d KB of %d KB", maxCodeFile>>10, len(text)>>10)
		text = strings.ToValidUTF8(text[:maxCodeFile], "")
	}
	out.Content = text
	return nil, out, nil
}
