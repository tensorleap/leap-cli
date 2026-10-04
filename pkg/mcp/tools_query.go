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

func (s *Server) describeFields(ctx context.Context, _ *sdk.CallToolRequest, in VersionIn) (*sdk.CallToolResult, FieldsOut, error) {
	if err := s.requireStats(ctx, in.ProjectID); err != nil {
		return nil, FieldsOut{}, err
	}
	v, err := s.slimVersion(ctx, in.ProjectID, in.VersionID)
	if err != nil {
		return nil, FieldsOut{}, err
	}
	out := FieldsOut{Server: s.client.UIBase(), Fields: []Field{}}
	if v.Resources.EsMetricsIndex == "" {
		out.Reason, out.NextStep = "no-evaluation", "evaluate this version before querying its metrics"
		return nil, out, nil
	}
	var mapping struct {
		Aggregatable []string `json:"aggregatableFields"`
		Numeric      []string `json:"numericFields"`
		Boolean      []string `json:"booleanFields"`
	}
	body := map[string]any{"projectId": in.ProjectID, "versionIds": []string{in.VersionID}, "inferenceArtifactIds": []string{v.Resources.InferenceArtifactID}}
	if err := s.client.Post(ctx, "dashboards/getDashletFields", body, &mapping); err != nil {
		return nil, FieldsOut{}, explain(err)
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
	ProjectID  string    `json:"projectId"`
	VersionIDs []string  `json:"versionIds" jsonschema:"one or more evaluated version ids; several versions are compared row by row"`
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

var aggregations = map[string]bool{"Average": true, "Min": true, "Max": true, "Median": true, "Count": true}

func (s *Server) query(ctx context.Context, _ *sdk.CallToolRequest, in QueryIn) (*sdk.CallToolResult, QueryOut, error) {
	if len(in.VersionIDs) == 0 || len(in.Measures) == 0 {
		return nil, QueryOut{}, errors.New("versionIds and measures are required")
	}
	if err := s.requireStats(ctx, in.ProjectID); err != nil {
		return nil, QueryOut{}, err
	}
	if len(in.GroupBy) > 2 {
		return nil, QueryOut{}, errors.New("groupBy supports at most 2 fields")
	}
	for _, m := range in.Measures {
		if !aggregations[m.Aggregation] {
			return nil, QueryOut{}, fmt.Errorf("unsupported aggregation %q; use Average, Min, Max, Median or Count", m.Aggregation)
		}
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
	sort.SliceStable(out.Rows, func(i, j int) bool { return out.Rows[i].Version < out.Rows[j].Version })
	if sawEmpty {
		out.Notes = append(out.Notes, "null values mean no data for that group, or exactly 0 on servers that render zero as empty")
	}
	if len(out.Rows) == 0 {
		out.Reason, out.NextStep = "no-rows", "check field names with tl_describe_fields, loosen the filters, or confirm the versions are evaluated (tl_list_versions)"
	}
	return nil, out, nil
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
