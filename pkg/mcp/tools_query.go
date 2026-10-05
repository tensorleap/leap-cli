package mcp

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

const splitField = "dataset_state.keyword"

type slimVersion struct {
	Cid       string `json:"cid"`
	Notes     string `json:"notes"`
	Resources struct {
		InferenceArtifactID string `json:"inference_artifact_id"`
		EsMetricsIndex      string `json:"es_metrics_index"`
	} `json:"resources"`
}

func (s *Server) slimVersion(ctx context.Context, projectID, versionID string) (*slimVersion, error) {
	var out struct {
		Versions []slimVersion `json:"versions"`
	}
	if err := s.client.Post(ctx, "versions/getProjectSlimVersions", map[string]any{"projectId": projectID}, &out); err != nil {
		return nil, explain(err)
	}
	for i := range out.Versions {
		if out.Versions[i].Cid == versionID {
			return &out.Versions[i], nil
		}
	}
	return nil, fmt.Errorf("version %s not found in project %s; check the ids with tl_list_versions", versionID, projectID)
}

type Field struct {
	Name string `json:"name"`
	Kind string `json:"kind" jsonschema:"metric | metadata | split"`
	Type string `json:"type" jsonschema:"numeric | string | boolean"`
}

type FieldsOut struct {
	Server   string  `json:"server"`
	Fields   []Field `json:"fields"`
	Reason   string  `json:"reason,omitempty"`
	NextStep string  `json:"nextStep,omitempty"`
}

type fieldMapping struct {
	Aggregatable []string `json:"aggregatableFields"`
	Numeric      []string `json:"numericFields"`
	Boolean      []string `json:"booleanFields"`
}

func (s *Server) fieldMapping(ctx context.Context, projectID, versionID string) (*fieldMapping, error) {
	v, err := s.slimVersion(ctx, projectID, versionID)
	if err != nil {
		return nil, err
	}
	if v.Resources.EsMetricsIndex == "" {
		return nil, nil
	}
	var m fieldMapping
	body := map[string]any{"projectId": projectID, "versionIds": []string{versionID}, "inferenceArtifactIds": []string{v.Resources.InferenceArtifactID}}
	if err := s.client.Post(ctx, "dashboards/getDashletFields", body, &m); err != nil {
		return nil, explain(err)
	}
	return &m, nil
}

// knownFields caches a version's queryable field names; an evaluated version's fields don't change
func (s *Server) knownFields(ctx context.Context, projectID, versionID string) (map[string]bool, error) {
	s.mu.Lock()
	known, ok := s.fields[versionID]
	s.mu.Unlock()
	if ok {
		return known, nil
	}
	m, err := s.fieldMapping(ctx, projectID, versionID)
	if err != nil || m == nil {
		return nil, err
	}
	known = toSet(append(append(append([]string{}, m.Aggregatable...), m.Numeric...), m.Boolean...))
	s.mu.Lock()
	s.fields[versionID] = known
	s.mu.Unlock()
	return known, nil
}

func (s *Server) describeFields(ctx context.Context, _ *sdk.CallToolRequest, in VersionIn) (*sdk.CallToolResult, FieldsOut, error) {
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, FieldsOut{}, err
	}
	if _, err := s.allowed(ctx, in.ProjectID, statsClass); err != nil {
		return nil, FieldsOut{}, err
	}
	mapping, err := s.fieldMapping(ctx, in.ProjectID, in.VersionID)
	if err != nil {
		return nil, FieldsOut{}, err
	}
	out := FieldsOut{Server: s.client.UIBase(), Fields: []Field{}}
	if mapping == nil {
		out.Reason, out.NextStep = "no-evaluation", "evaluate this version before querying its metrics"
		return nil, out, nil
	}
	numeric, boolean := toSet(mapping.Numeric), toSet(mapping.Boolean)
	for _, name := range mapping.Aggregatable {
		kind := ""
		switch {
		case strings.HasPrefix(name, "metrics."):
			kind = "metric"
		case strings.HasPrefix(name, "metadata."):
			kind = "metadata"
		case name == splitField:
			kind = "split"
		default:
			continue
		}
		typ := "string"
		if numeric[name] {
			typ = "numeric"
		} else if boolean[name] {
			typ = "boolean"
		}
		out.Fields = append(out.Fields, Field{Name: name, Kind: kind, Type: typ})
	}
	if len(out.Fields) == 0 {
		out.Reason, out.NextStep = "no-fields", "the evaluation wrote no metric or metadata fields; check the integration's metrics and metadata functions"
	}
	return nil, out, nil
}

func toSet(l []string) map[string]bool {
	m := make(map[string]bool, len(l))
	for _, s := range l {
		m[s] = true
	}
	return m
}

type Measure struct {
	Field       string `json:"field" jsonschema:"metric or metadata field name from tl_describe_fields"`
	Aggregation string `json:"aggregation" jsonschema:"Average | Min | Max | Median | Count"`
}

type Filter struct {
	Field    string   `json:"field" jsonschema:"field name from tl_describe_fields"`
	Operator string   `json:"operator" jsonschema:"equal | not-equal | in | not-in | greater-than | less-than | between"`
	Value    any      `json:"value,omitempty" jsonschema:"value for equal, not-equal, greater-than, less-than"`
	Values   []any    `json:"values,omitempty" jsonschema:"values for in / not-in"`
	Min      *float64 `json:"min,omitempty" jsonschema:"lower bound for between"`
	Max      *float64 `json:"max,omitempty" jsonschema:"upper bound for between"`
}

type QueryIn struct {
	ProjectID  string    `json:"projectId" jsonschema:"project id or name"`
	VersionIDs []string  `json:"versionIds" jsonschema:"one or more evaluated version ids or names (or \"latest\"); several versions are compared row by row"`
	GroupBy    []string  `json:"groupBy,omitempty" jsonschema:"up to 2 field names to group by (empty = whole population)"`
	Measures   []Measure `json:"measures" jsonschema:"what to compute per group; a sample count n is always included"`
	Filters    []Filter  `json:"filters,omitempty"`
	Limit      int       `json:"limit,omitempty" jsonschema:"max groups per dimension (default 20, max 100)"`
}

type Row struct {
	Version string         `json:"version"`
	Group   map[string]any `json:"group,omitempty"`
	N       *float64       `json:"n"`
	Values  map[string]any `json:"values"`
}

type QueryOut struct {
	Server   string   `json:"server"`
	Rows     []Row    `json:"rows"`
	Notes    []string `json:"notes,omitempty"`
	Reason   string   `json:"reason,omitempty"`
	NextStep string   `json:"nextStep,omitempty"`
}

var aggregations = map[string]string{
	"average": "Average", "avg": "Average", "mean": "Average",
	"min": "Min", "minimum": "Min",
	"max": "Max", "maximum": "Max",
	"median": "Median", "p50": "Median",
	"count": "Count",
}

func (s *Server) query(ctx context.Context, _ *sdk.CallToolRequest, in QueryIn) (*sdk.CallToolResult, QueryOut, error) {
	if len(in.VersionIDs) == 0 || len(in.Measures) == 0 {
		return nil, QueryOut{}, errors.New("versionIds and measures are required")
	}
	refs := make([]*string, len(in.VersionIDs))
	for i := range in.VersionIDs {
		refs[i] = &in.VersionIDs[i]
	}
	if err := s.resolve(ctx, &in.ProjectID, refs...); err != nil {
		return nil, QueryOut{}, err
	}
	access, err := s.allowed(ctx, in.ProjectID, statsClass)
	if err != nil {
		return nil, QueryOut{}, err
	}
	if len(in.GroupBy) > 2 {
		return nil, QueryOut{}, errors.New("groupBy supports at most 2 fields")
	}
	for i, m := range in.Measures {
		canonical, ok := aggregations[strings.ToLower(m.Aggregation)]
		if !ok {
			return nil, QueryOut{}, fmt.Errorf("unsupported aggregation %q; use Average, Min, Max, Median or Count", m.Aggregation)
		}
		in.Measures[i].Aggregation = canonical
	}
	if err := s.checkFields(ctx, in); err != nil {
		return nil, QueryOut{}, err
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	filters, err := toEsFilters(in.Filters)
	if err != nil {
		return nil, QueryOut{}, err
	}
	buckets := make([]map[string]any, 0, len(in.GroupBy))
	for _, f := range in.GroupBy {
		buckets = append(buckets, map[string]any{"field": f, "distribution": "distinct", "order": "desc", "limit": limit})
	}
	if len(buckets) == 0 {
		// table charts return nothing without a bucket; the model id has one value per version
		buckets = append(buckets, map[string]any{"field": "model.extId.keyword", "distribution": "distinct", "order": "desc", "limit": 1})
	}
	batches := measureBatches(append([]Measure{{Field: "sample_id", Aggregation: "Count"}}, in.Measures...))
	out := QueryOut{Server: s.client.UIBase(), Rows: []Row{}}
	sawEmpty := false
	for _, versionID := range in.VersionIDs {
		rows := map[string]*Row{}
		var order []string
		for _, batch := range batches {
			aggs := make([]map[string]any, 0, len(batch))
			for _, m := range batch {
				aggs = append(aggs, map[string]any{"field": m.Field, "aggregation": m.Aggregation})
			}
			body := map[string]any{"projectId": in.ProjectID, "versionIds": []string{versionID}, "showAllEpochs": false,
				"filters": filters, "aggregations": aggs, "buckets": buckets}
			var resp struct {
				Charts []struct {
					Data struct {
						Data []struct {
							Data map[string]any `json:"data"`
						} `json:"data"`
					} `json:"data"`
				} `json:"charts"`
			}
			if err := s.client.Post(ctx, "sessionmetrics/getTableChart", body, &resp); err != nil {
				return nil, QueryOut{}, explain(err)
			}
			for _, chart := range resp.Charts {
				for _, item := range chart.Data.Data {
					key, group := groupKey(item.Data, in.GroupBy)
					r, ok := rows[key]
					if !ok {
						r = &Row{Version: versionID, Group: group, Values: map[string]any{}}
						rows[key] = r
						order = append(order, key)
					}
					for _, m := range batch {
						v := item.Data[m.Field]
						if s, isStr := v.(string); isStr && s == "" {
							v, sawEmpty = nil, true
						}
						if m.Field == "sample_id" && m.Aggregation == "Count" {
							if f, ok := v.(float64); ok {
								r.N = &f
							}
							continue
						}
						r.Values[m.Aggregation+"("+m.Field+")"] = v
					}
				}
			}
		}
		for _, k := range order {
			out.Rows = append(out.Rows, *rows[k])
		}
	}
	if !access.SampleRows {
		kept, hidden := out.Rows[:0], 0
		for _, r := range out.Rows {
			if r.N != nil && *r.N < minCellSize {
				hidden++
				continue
			}
			kept = append(kept, r)
		}
		out.Rows = kept
		if hidden > 0 {
			// with per-sample data off, tiny groups would leak individual samples' values
			out.Notes = append(out.Notes, fmt.Sprintf("%d group(s) with fewer than %d samples are hidden because per-sample data is turned off for this project", hidden, minCellSize))
		}
	}
	sort.SliceStable(out.Rows, func(i, j int) bool { return out.Rows[i].Version < out.Rows[j].Version })
	if sawEmpty {
		note := "null means the group has no values for that field (e.g. unlabeled samples have no loss)"
		if access.legacy {
			note += ", or exactly 0 on this older server"
		}
		out.Notes = append(out.Notes, note)
	}
	if len(out.Rows) == 0 {
		out.Reason, out.NextStep = "no-rows", "check field names with tl_describe_fields, loosen the filters, or confirm the versions are evaluated (tl_list_versions)"
	}
	return nil, out, nil
}

// checkFields rejects unknown field names up front; the metrics API would silently return nulls
func (s *Server) checkFields(ctx context.Context, in QueryIn) error {
	known, err := s.knownFields(ctx, in.ProjectID, in.VersionIDs[0])
	if err != nil || known == nil {
		return err
	}
	names := append([]string{}, in.GroupBy...)
	for _, m := range in.Measures {
		names = append(names, m.Field)
	}
	for _, f := range in.Filters {
		names = append(names, f.Field)
	}
	for _, n := range names {
		if !known[n] && n != "sample_id" {
			return fmt.Errorf("unknown field %q for this version%s; tl_describe_fields lists the exact names", n, suggest(n, known))
		}
	}
	return nil
}

func suggest(name string, known map[string]bool) string {
	leaf := strings.ToLower(name[strings.LastIndex(name, ".")+1:])
	var close []string
	for k := range known {
		kl := strings.ToLower(k)
		if strings.Contains(kl, leaf) || strings.Contains(leaf, kl[strings.LastIndex(kl, ".")+1:]) {
			close = append(close, k)
		}
	}
	if len(close) == 0 {
		return ""
	}
	sort.Strings(close)
	if len(close) > 5 {
		close = close[:5]
	}
	return " (did you mean " + strings.Join(close, ", ") + "?)"
}

func measureBatches(ms []Measure) [][]Measure {
	var batches [][]Measure
	for _, m := range ms {
		placed := false
		for i := range batches {
			clash := false
			for _, b := range batches[i] {
				if b.Field == m.Field {
					clash = true
					break
				}
			}
			if !clash {
				batches[i] = append(batches[i], m)
				placed = true
				break
			}
		}
		if !placed {
			batches = append(batches, []Measure{m})
		}
	}
	return batches
}

func groupKey(data map[string]any, groupBy []string) (string, map[string]any) {
	if len(groupBy) == 0 {
		return "", nil
	}
	group := map[string]any{}
	parts := make([]string, 0, len(groupBy))
	for _, g := range groupBy {
		group[g] = data[g]
		parts = append(parts, fmt.Sprint(data[g]))
	}
	return strings.Join(parts, "\x00"), group
}

func toEsFilters(fs []Filter) ([]map[string]any, error) {
	out := make([]map[string]any, 0, len(fs))
	for _, f := range fs {
		value := map[string]any{}
		switch f.Operator {
		case "equal", "not-equal":
			value["eq"] = f.Value
		case "in", "not-in":
			value["lst"] = f.Values
		case "greater-than":
			value["gt"] = f.Value
		case "less-than":
			value["lt"] = f.Value
		case "between":
			if f.Min == nil || f.Max == nil {
				return nil, fmt.Errorf("filter on %s: between needs min and max", f.Field)
			}
			value["gte"], value["lte"] = *f.Min, *f.Max
		default:
			return nil, fmt.Errorf("filter on %s: unsupported operator %q", f.Field, f.Operator)
		}
		out = append(out, map[string]any{"field": f.Field, "operator": f.Operator, "value": value})
	}
	return out, nil
}

type ListSamplesIn struct {
	ProjectID string   `json:"projectId" jsonschema:"project id or name"`
	VersionID string   `json:"versionId" jsonschema:"evaluated version id or name, or \"latest\""`
	Filters   []Filter `json:"filters,omitempty" jsonschema:"which samples, e.g. a label equal to 2 and a prediction equal to 7"`
	SortBy    string   `json:"sortBy,omitempty" jsonschema:"field to rank by, e.g. metrics.loss (default: sample id)"`
	Ascending bool     `json:"ascending,omitempty" jsonschema:"lowest first instead of highest first"`
	Fields    []string `json:"fields,omitempty" jsonschema:"fields to return per sample (default: the sort and filter fields)"`
	Limit     int      `json:"limit,omitempty" jsonschema:"max samples (default 20, max 100)"`
}

type ListedSample struct {
	ID       string         `json:"id" jsonschema:"pass to tl_view_samples"`
	Rendered *bool          `json:"rendered,omitempty" jsonschema:"whether tl_view_samples can show it; unrendered samples need Visualize in the UI first"`
	Values   map[string]any `json:"values"`
}

type ListSamplesOut struct {
	Server   string         `json:"server"`
	Matching int            `json:"matching" jsonschema:"how many samples match the filters"`
	Samples  []ListedSample `json:"samples"`
	Reason   string         `json:"reason,omitempty"`
	NextStep string         `json:"nextStep,omitempty"`
}

func (s *Server) listSamples(ctx context.Context, _ *sdk.CallToolRequest, in ListSamplesIn) (*sdk.CallToolResult, ListSamplesOut, error) {
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, ListSamplesOut{}, err
	}
	access, err := s.allowed(ctx, in.ProjectID, sampleRowsClass)
	if err != nil {
		return nil, ListSamplesOut{}, err
	}
	fields := append([]string{}, in.Fields...)
	if len(fields) == 0 {
		if in.SortBy != "" {
			fields = append(fields, in.SortBy)
		}
		for _, f := range in.Filters {
			fields = appendOnce(fields, f.Field)
		}
	}
	names := append([]string{}, fields...)
	if in.SortBy != "" {
		names = appendOnce(names, in.SortBy)
	}
	if err := s.checkFields(ctx, QueryIn{ProjectID: in.ProjectID, VersionIDs: []string{in.VersionID}, GroupBy: names, Filters: in.Filters}); err != nil {
		return nil, ListSamplesOut{}, err
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 20
	}
	if limit > 100 {
		limit = 100
	}
	filters, kept, err := toSampleFilters(in.Filters)
	if err != nil {
		return nil, ListSamplesOut{}, err
	}
	body := map[string]any{"projectId": in.ProjectID, "versionId": in.VersionID, "filters": filters, "pageSize": limit}
	if in.SortBy != "" {
		dir := "desc"
		if in.Ascending {
			dir = "asc"
		}
		body["sort"] = map[string]any{"field": in.SortBy, "dir": dir}
	}
	var resp struct {
		Rows  []map[string]any `json:"rows"`
		Total int              `json:"total"`
	}
	if err := s.client.Post(ctx, "sample-collection/getVersionSampleOrder", body, &resp); err != nil {
		return nil, ListSamplesOut{}, explain(err)
	}
	out := ListSamplesOut{Server: s.client.UIBase(), Matching: resp.Total, Samples: []ListedSample{}}
	for _, row := range resp.Rows {
		for field, values := range kept {
			if !values[fmt.Sprint(row[field])] {
				// servers without include filters return every sample; never pass that off as the case
				return nil, ListSamplesOut{}, errors.New("this Tensorleap server can't list samples by value yet; upgrade it, or filter with greater-than / less-than / between")
			}
		}
		smp := ListedSample{ID: fmt.Sprintf("%v_%v", row["state"], row["index"]), Values: map[string]any{}}
		for _, f := range fields {
			smp.Values[f] = row[f]
		}
		out.Samples = append(out.Samples, smp)
	}
	if len(out.Samples) == 0 {
		out.Reason, out.NextStep = "no-matching-samples", "loosen the filters, or check values with tl_query grouped by the same fields"
	}
	if access.Visuals && len(out.Samples) > 0 {
		// saves the assistant from probing tl_view_samples sample by sample for one it can see
		ids := make([]string, len(out.Samples))
		for i, smp := range out.Samples {
			ids[i] = smp.ID
		}
		if rendered, err := s.sampleAssets(ctx, in.ProjectID, in.VersionID, ids); err == nil {
			for i := range out.Samples {
				r := len(rendered[out.Samples[i].ID]) > 0
				out.Samples[i].Rendered = &r
			}
		}
	}
	return nil, out, nil
}

// toSampleFilters maps tl_query-style filters onto the sample list API; kept holds the values an
// include filter must match, to detect servers that ignore them
func toSampleFilters(fs []Filter) ([]map[string]any, map[string]map[string]bool, error) {
	out := make([]map[string]any, 0, len(fs))
	kept := map[string]map[string]bool{}
	for _, f := range fs {
		spec := map[string]any{"field": f.Field}
		switch f.Operator {
		case "equal", "in":
			values := f.Values
			if f.Operator == "equal" {
				values = []any{f.Value}
			}
			spec["values"] = values
			kept[f.Field] = map[string]bool{}
			for _, v := range values {
				kept[f.Field][fmt.Sprint(v)] = true
			}
		case "not-equal":
			spec["hiddenValues"] = []any{f.Value}
		case "not-in":
			spec["hiddenValues"] = f.Values
		case "greater-than":
			spec["range"] = map[string]any{"min": f.Value}
		case "less-than":
			spec["range"] = map[string]any{"max": f.Value}
		case "between":
			if f.Min == nil || f.Max == nil {
				return nil, nil, fmt.Errorf("filter on %s: between needs min and max", f.Field)
			}
			spec["range"] = map[string]any{"min": *f.Min, "max": *f.Max}
		default:
			return nil, nil, fmt.Errorf("filter on %s: unsupported operator %q", f.Field, f.Operator)
		}
		out = append(out, spec)
	}
	return out, kept, nil
}
