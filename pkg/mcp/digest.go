package mcp

// small cells are hidden the same way the server hides them in an insight's composition
const minCellSize = 10

type Composition struct {
	Field      string  `json:"field"`
	Value      string  `json:"value,omitempty"`
	InGroup    float64 `json:"inGroup"`
	InBaseline float64 `json:"inBaseline"`
	Ratio      float64 `json:"ratio"`
	Kind       string  `json:"kind" jsonschema:"share = fraction of samples with this value; mean = average of a numeric field"`
	Semantics  string  `json:"semantics"`
}

type Contrast struct {
	Metric         string   `json:"metric"`
	GroupMedian    float64  `json:"groupMedian"`
	GroupMean      float64  `json:"groupMean"`
	BaselineMedian *float64 `json:"baselineMedian,omitempty"`
	BaselineMean   *float64 `json:"baselineMean,omitempty"`
	Baseline       string   `json:"baseline"`
}

// GroupSummary is node-server's digest of an insight's sample list (analysis-export exportAnalysis)
type GroupSummary struct {
	GroupSize       int            `json:"groupSize"`
	GroupDefinition string         `json:"groupDefinition"`
	CsvRows         int            `json:"csvRows"`
	Split           map[string]int `json:"split"`
	Composition     []Composition  `json:"composition"`
	Contrast        []Contrast     `json:"contrast"`
	RankedBy        string         `json:"rankedBy,omitempty"`
}

type insightDigest struct {
	GroupSummary
	RankedSampleIDs []string `json:"rankedSampleIds"`
}
