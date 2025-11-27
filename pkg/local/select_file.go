package local

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/sahilm/fuzzy"
	"github.com/tensorleap/helm-charts/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
)

//
// ---------------------------
// Structs
// ---------------------------
//

// FileSelectSession represents the current state of the interactive file picker.
type FileSelectSession struct {
	Title        string
	CurrentDir   string
	AllowedExt   []string
	AllFiles     []string
	CachedDirs   map[string][]string
	Options      []string
	Filter       string
	Cursor       int
	Hint         string
	ErrorMsg     string
	IsExternal   bool
	SelectedPath *string
}

//
// ---------------------------
// Entry Point
// ---------------------------
//

// SelectFile lets the user browse and select a file interactively.
func SelectFile(allowedExt []string, title string) (string, error) {
	cwd, _ := filepath.Abs(".")
	allFiles, err := collectAllFiles(cwd, allowedExt)
	if err != nil {
		return "", err
	}

	session := &FileSelectSession{
		Title:      title,
		CurrentDir: cwd,
		AllowedExt: allowedExt,
		AllFiles:   allFiles,
		CachedDirs: make(map[string][]string),
		Hint:       "type to filter or enter path",
	}

	// Wrap UI callbacks around our struct methods
	selectedFile, err := log.InteractiveSelectTea(
		func(cursor int, filter string) (string, []string, int, string, string) {
			session.Cursor, session.Filter = cursor, filter
			session.UpdateOnFilter()
			return session.Title, session.Options, session.Cursor, session.ErrorMsg, session.Hint
		},
		func(options []string, cursor int, filter string) (*string, string, string, string) {
			session.Options, session.Cursor, session.Filter = options, cursor, filter
			session.HandleEnterKey()
			return session.SelectedPath, session.Filter, session.ErrorMsg, session.Hint
		},
		func(options []string, cursor int, filter string) string {
			session.Options, session.Cursor, session.Filter = options, cursor, filter
			return session.HandleTabKey()
		},
	)

	if err != nil {
		return "", err
	}
	expanded := expandPath(selectedFile, cwd)
	style := text.Colors{text.Bold, text.FgGreen}
	fmt.Printf("%s %s %s\n", style.Sprint("?"), text.Bold.Sprintf("%s:", title), expanded)
	return selectedFile, nil
}

//
// ---------------------------
// Methods (Main Logic)
// ---------------------------
//

// UpdateOnFilter updates options whenever the user types a new filter.
func (s *FileSelectSession) UpdateOnFilter() {
	s.ErrorMsg = ""

	if s.Filter == "" {
		s.Options = s.AllFiles
		s.Hint = "type to filter or enter path"
		s.IsExternal = false
		return
	}

	isNotInCurrentPath := strings.HasPrefix(s.Filter, "/") || strings.HasPrefix(s.Filter, "~") || strings.HasPrefix(s.Filter, "../") || strings.HasPrefix(s.Filter, "./")

	// 1️⃣ Project-local fuzzy search
	if !isNotInCurrentPath {
		s.Options = fuzzyMatch(s.Filter, s.AllFiles)
		s.Hint = "type to filter or enter path"
		s.IsExternal = false
		return
	}

	// 2️⃣ External path (not in current path: absolute, ~, or ../)
	s.IsExternal = true
	s.Hint = "navigate to file"

	expanded := expandPath(s.Filter, s.CurrentDir)

	dir := filepath.Dir(expanded)
	fileName := ""
	if !strings.HasSuffix(s.Filter, "/") {
		fileName = filepath.Base(expanded)
	} else {
		dir = expanded
	}

	s.Hint = fmt.Sprintf("file: %s, directory: %s, expanded: %s", fileName, dir, expanded)

	dirEntries, err := s.loadDirFromCache(dir)
	if err != nil || len(dirEntries) == 0 {
		// Try fuzzy match one level up (for partial paths like ~/mosh)
		parent := filepath.Dir(dir)
		parentEntries, perr := s.loadDirFromCache(parent)
		if perr == nil {
			s.Options = fuzzyMatch(filepath.Base(dir), parentEntries)
			s.Hint = "type or navigate to path"
			return
		}
		s.Options = []string{}
		s.Hint = "type or navigate to path"
		return
	}

	if fileName == "" {
		s.Options = dirEntries
	} else {
		matched := fuzzyMatch(fileName, dirEntries)
		if len(matched) == 0 {
			// fallback to showing directory itself (helps user see context)
			s.Options = dirEntries
		} else {
			s.Options = matched
		}
	}
}

// HandleEnterKey handles when the user presses ENTER.
func (s *FileSelectSession) HandleEnterKey() {
	isCursorInvalid := s.Cursor < 0 || s.Cursor >= len(s.Options)
	s.ErrorMsg = ""

	if !s.IsExternal {
		if isCursorInvalid {
			s.ErrorMsg = "File not found"
			return
		}
		selected := s.Options[s.Cursor]
		s.SelectedPath = &selected
		s.Hint = "type to filter or enter path"
		return
	}

	s.Hint = "navigate to file"
	expanded := expandPath(s.Filter, s.CurrentDir)

	if isCursorInvalid {
		s.checkTypedFile(expanded)
		return
	}

	s.resolveSelection(expanded)
}

// HandleTabKey autocompletes the current filter.
func (s *FileSelectSession) HandleTabKey() string {
	if s.Cursor < 0 || s.Cursor >= len(s.Options) {
		return s.Filter
	}

	if s.IsExternal {
		dir := filepath.Dir(s.Filter)
		return fmt.Sprintf("%s/", path.Join(dir, s.Options[s.Cursor]))
	}
	return s.Options[s.Cursor]
}

// checkTypedFile validates a directly typed file path.
func (s *FileSelectSession) checkTypedFile(expanded string) {
	exists, _ := local.FileExists(expanded)
	isDir := strings.HasSuffix(expanded, "/")

	switch {
	case exists && !isDir:
		s.SelectedPath = &s.Filter
	case isDir:
		s.ErrorMsg = "cannot select a directory"
	default:
		s.ErrorMsg = "file not found"
	}
}

// resolveSelection handles when user selects a file or folder.
func (s *FileSelectSession) resolveSelection(expanded string) {
	option := s.Options[s.Cursor]
	isDir := strings.HasSuffix(option, "/")
	newPath := path.Join(filepath.Dir(expanded), option)

	if isDir {
		s.Filter = fmt.Sprintf("%s/", path.Join(filepath.Dir(s.Filter), option))
		return
	}

	// Preserve ~ paths
	if strings.HasPrefix(s.Filter, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			if rel, err := filepath.Rel(home, newPath); err == nil {
				readable := filepath.Join("~", rel)
				s.SelectedPath = &readable
				return
			}
		}
	}

	// Preserve relative paths (../ or ./)
	if strings.HasPrefix(s.Filter, "../") || strings.HasPrefix(s.Filter, "./") {
		if rel, err := filepath.Rel(s.CurrentDir, newPath); err == nil {
			// Use relative path, preserving ../ if needed
			if strings.HasPrefix(rel, "..") {
				s.SelectedPath = &rel
				return
			}
			readable := filepath.Join(".", rel)
			s.SelectedPath = &readable
			return
		}
	}

	s.SelectedPath = &newPath
}

//
// ---------------------------
// Helper Methods
// ---------------------------
//

// expandPath expands "~" → user home directory and resolves relative paths (../, ./) relative to currentDir.
func expandPath(p string, currentDir string) string {
	// Handle home directory expansion
	if strings.HasPrefix(p, "~") {
		if home, err := os.UserHomeDir(); err == nil {
			expanded := filepath.Join(home, strings.TrimPrefix(p, "~"))
			abs, err := filepath.Abs(expanded)
			if err == nil {
				return abs
			}
			return expanded
		}
	}

	// Handle relative paths (../, ./)
	if strings.HasPrefix(p, "../") || strings.HasPrefix(p, "./") {
		abs, err := filepath.Abs(filepath.Join(currentDir, p))
		if err == nil {
			return abs
		}
		return p
	}

	// For absolute paths, just return as-is (or make absolute if needed)
	if strings.HasPrefix(p, "/") {
		return p
	}

	// For other paths, resolve relative to current directory
	abs, err := filepath.Abs(filepath.Join(currentDir, p))
	if err == nil {
		return abs
	}
	return p
}

// loadDirFromCache retrieves cached directory listing or scans it.
func (s *FileSelectSession) loadDirFromCache(dir string) ([]string, error) {
	if cached, ok := s.CachedDirs[dir]; ok {
		return cached, nil
	}
	files, err := readDirectory(dir, s.AllowedExt)
	if err != nil {
		return nil, err
	}
	s.CachedDirs[dir] = files
	return files, nil
}

//
// ---------------------------
// Utility Functions
// ---------------------------
//

// collectAllFiles recursively collects all files under root.
func collectAllFiles(root string, allowedExt []string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(p string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		if len(allowedExt) == 0 || hasExtension(allowedExt, strings.ToLower(filepath.Ext(d.Name()))) {
			files = append(files, p)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking path %s: %w", root, err)
	}

	trimmed := make([]string, 0, len(files))
	for _, f := range files {
		relativePath, err := filepath.Rel(root, f)
		if err != nil {
			return nil, fmt.Errorf("error getting relative path for %s: %w", f, err)
		}
		trimmed = append(trimmed, fmt.Sprintf("./%s", relativePath))
	}
	return trimmed, nil
}

// readDirectory lists immediate files and subdirectories.
func readDirectory(dir string, allowedExt []string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	var items []string
	for _, e := range entries {
		if e.IsDir() {
			items = append(items, e.Name()+"/")
		} else if len(allowedExt) == 0 || hasExtension(allowedExt, strings.ToLower(filepath.Ext(e.Name()))) {
			items = append(items, e.Name())
		}
	}
	return items, nil
}

// fuzzyMatch performs fuzzy string matching.
func fuzzyMatch(query string, candidates []string) []string {
	if query == "" {
		return candidates
	}
	matches := fuzzy.Find(query, candidates)
	results := make([]string, 0, len(matches))
	for _, m := range matches {
		results = append(results, candidates[m.Index])
	}
	return results
}

// hasExtension checks if the extension is allowed.
func hasExtension(allowed []string, ext string) bool {
	for _, e := range allowed {
		if strings.EqualFold(e, ext) {
			return true
		}
	}
	return false
}
