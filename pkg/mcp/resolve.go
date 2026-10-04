package mcp

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var objectID = regexp.MustCompile(`^[0-9a-f]{24}$`)

// resolve turns project and version references into ids in place. Assistants often only know
// names ("mnist", "baseline") or want the newest version, so both are accepted besides ids.
func (s *Server) resolve(ctx context.Context, projectRef *string, versionRefs ...*string) error {
	*projectRef = strings.TrimSpace(*projectRef)
	if *projectRef == "" {
		return fmt.Errorf("projectId is required; get it from tl_list_projects")
	}
	if !objectID.MatchString(*projectRef) {
		t, err := s.targets(ctx, "")
		if err != nil {
			return err
		}
		names, matches := []string{}, []string{}
		for _, p := range t.Projects {
			names = append(names, p.Name)
			if strings.EqualFold(p.Name, *projectRef) {
				matches = append(matches, p.Cid)
			}
		}
		switch len(matches) {
		case 0:
			sort.Strings(names)
			return fmt.Errorf("no project named %q; your projects: %s", *projectRef, strings.Join(names, ", "))
		case 1:
			*projectRef = matches[0]
		default:
			return fmt.Errorf("%d projects are named %q; pass the id from tl_list_projects", len(matches), *projectRef)
		}
	}
	for _, ref := range versionRefs {
		if ref == nil || *ref == "" || objectID.MatchString(strings.TrimSpace(*ref)) {
			continue
		}
		id, err := s.resolveVersion(ctx, *projectRef, strings.TrimSpace(*ref))
		if err != nil {
			return err
		}
		*ref = id
	}
	return nil
}

func (s *Server) resolveVersion(ctx context.Context, projectID, ref string) (string, error) {
	t, err := s.targets(ctx, projectID)
	if err != nil {
		return "", err
	}
	if strings.EqualFold(ref, "latest") {
		latest, at := "", ""
		for _, v := range t.Versions {
			if v.Evaluated && v.CreatedAt > at {
				latest, at = v.Cid, v.CreatedAt
			}
		}
		if latest == "" {
			return "", fmt.Errorf("no evaluated version in this project yet; run Evaluate first (tl_list_versions shows their state)")
		}
		return latest, nil
	}
	names, matches := []string{}, []string{}
	for _, v := range t.Versions {
		names = append(names, v.Name)
		if strings.EqualFold(v.Name, ref) {
			matches = append(matches, v.Cid)
		}
	}
	switch len(matches) {
	case 1:
		return matches[0], nil
	case 0:
		return "", fmt.Errorf("no version named %q in this project; versions: %s (or pass \"latest\")", ref, strings.Join(names, ", "))
	}
	return "", fmt.Errorf("%d versions are named %q; pass the id from tl_list_versions", len(matches), ref)
}
