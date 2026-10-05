package mcp

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

var inFlightStatuses = []string{"UNSTARTED", "QUEUED", "PENDING", "INITIALIZING", "STARTED"}
var terminalStatuses = map[string]bool{"FINISHED": true, "FAILED": true, "STOPPED": true, "TERMINATED": true}

type rawJob struct {
	Cid       string `json:"cid"`
	Type      string `json:"type"`
	SubType   string `json:"subType"`
	Status    string `json:"status"`
	ProjectID string `json:"projectId"`
	VersionID string `json:"versionId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Job struct {
	ID        string `json:"id"`
	Kind      string `json:"kind" jsonschema:"what the job does, e.g. Evaluate, Push, Population Exploration"`
	Status    string `json:"status"`
	InFlight  bool   `json:"inFlight"`
	ProjectID string `json:"projectId,omitempty"`
	VersionID string `json:"versionId,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func toJob(r rawJob) Job {
	kind := r.SubType
	if kind == "" {
		kind = r.Type
	}
	return Job{ID: r.Cid, Kind: kind, Status: r.Status, InFlight: !terminalStatuses[r.Status], ProjectID: r.ProjectID,
		VersionID: r.VersionID, CreatedAt: r.CreatedAt, UpdatedAt: r.UpdatedAt}
}

type JobsIn struct {
	ProjectID string `json:"projectId" jsonschema:"project id or name"`
	VersionID string `json:"versionId,omitempty" jsonschema:"optional: only this version's jobs (id or name)"`
	Status    string `json:"status,omitempty" jsonschema:"in-flight | failed | all (default all)"`
	Limit     int    `json:"limit,omitempty" jsonschema:"max jobs, newest first (default 20)"`
}

type JobsOut struct {
	Server string `json:"server"`
	Jobs   []Job  `json:"jobs"`
	Reason string `json:"reason,omitempty"`
}

func (s *Server) listJobs(ctx context.Context, _ *sdk.CallToolRequest, in JobsIn) (*sdk.CallToolResult, JobsOut, error) {
	if err := s.resolve(ctx, &in.ProjectID, &in.VersionID); err != nil {
		return nil, JobsOut{}, err
	}
	var resp struct {
		Jobs []rawJob `json:"jobs"`
	}
	var err error
	if in.VersionID != "" {
		err = s.client.Post(ctx, "jobs/getVersionJobs", map[string]any{"projectId": in.ProjectID, "versionId": in.VersionID}, &resp)
	} else {
		body := map[string]any{"projectId": in.ProjectID}
		switch in.Status {
		case "in-flight":
			body["status"] = inFlightStatuses
		case "failed":
			body["status"] = []string{"FAILED"}
		}
		err = s.client.Post(ctx, "jobs/getSlimJobs", body, &resp)
	}
	if err != nil {
		return nil, JobsOut{}, explain(err)
	}
	limit := in.Limit
	if limit <= 0 {
		limit = 20
	}
	out := JobsOut{Server: s.client.UIBase(), Jobs: []Job{}}
	for _, r := range resp.Jobs {
		j := toJob(r)
		if (in.Status == "in-flight" && !j.InFlight) || (in.Status == "failed" && j.Status != "FAILED") {
			continue
		}
		out.Jobs = append(out.Jobs, j)
	}
	sortNewestFirst(out.Jobs)
	if len(out.Jobs) > limit {
		out.Jobs = out.Jobs[:limit]
	}
	if len(out.Jobs) == 0 {
		out.Reason = "no-jobs"
	}
	return nil, out, nil
}

func sortNewestFirst(js []Job) {
	for i := 1; i < len(js); i++ {
		for j := i; j > 0 && js[j].CreatedAt > js[j-1].CreatedAt; j-- {
			js[j], js[j-1] = js[j-1], js[j]
		}
	}
}

func (s *Server) ownedJob(ctx context.Context, jobID string) (*Job, error) {
	var resp struct {
		Jobs []rawJob `json:"jobs"`
	}
	if err := s.client.Post(ctx, "jobs/getSlimJobs", map[string]any{"cid": []string{jobID}}, &resp); err != nil {
		return nil, explain(err)
	}
	if len(resp.Jobs) == 0 {
		return nil, errors.New("job not found in your team; list jobs with tl_list_jobs")
	}
	j := toJob(resp.Jobs[0])
	return &j, nil
}

type WaitIn struct {
	JobID          string `json:"jobId"`
	TimeoutSeconds int    `json:"timeoutSeconds,omitempty" jsonschema:"how long to wait for a status change (default 60, max 300)"`
}

type WaitOut struct {
	Job      Job    `json:"job"`
	Changed  bool   `json:"changed"`
	Waited   int    `json:"waitedSeconds"`
	NextStep string `json:"nextStep,omitempty"`
}

var pollInterval = 10 * time.Second

func (s *Server) waitForJob(ctx context.Context, req *sdk.CallToolRequest, in WaitIn) (*sdk.CallToolResult, WaitOut, error) {
	timeout := in.TimeoutSeconds
	if timeout <= 0 {
		timeout = 60
	}
	if timeout > 300 {
		timeout = 300
	}
	start := time.Now()
	first, err := s.ownedJob(ctx, in.JobID)
	if err != nil {
		return nil, WaitOut{}, err
	}
	current := first
	for !terminalStatuses[current.Status] && time.Since(start) < time.Duration(timeout)*time.Second {
		select {
		case <-ctx.Done():
			return nil, WaitOut{}, ctx.Err()
		case <-time.After(pollInterval):
		}
		if current, err = s.ownedJob(ctx, in.JobID); err != nil {
			return nil, WaitOut{}, err
		}
		if token := req.Params.GetProgressToken(); token != nil && req.Session != nil {
			_ = req.Session.NotifyProgress(ctx, &sdk.ProgressNotificationParams{ProgressToken: token,
				Progress: time.Since(start).Seconds(), Total: float64(timeout), Message: current.Kind + ": " + current.Status})
		}
		if current.Status != first.Status {
			break
		}
	}
	out := WaitOut{Job: *current, Changed: current.Status != first.Status, Waited: int(time.Since(start).Seconds())}
	switch {
	case current.Status == "FAILED":
		out.NextStep = "read the failure with tl_get_job_logs"
	case current.InFlight:
		out.NextStep = "still running; call tl_wait_for_job again later instead of polling repeatedly"
	}
	return nil, out, nil
}

type LogsIn struct {
	JobID string `json:"jobId"`
	Tail  int    `json:"tail,omitempty" jsonschema:"last N lines per pod (default 80, max 400)"`
}

type PodLog struct {
	Pod   string `json:"pod"`
	Lines string `json:"lines"`
}

type LogsOut struct {
	Job      Job      `json:"job"`
	Errors   []string `json:"errorLines" jsonschema:"lines that look like the failure (tracebacks, errors)"`
	Pods     []PodLog `json:"pods"`
	Redacted string   `json:"redaction"`
}

func (s *Server) getJobLogs(ctx context.Context, _ *sdk.CallToolRequest, in LogsIn) (*sdk.CallToolResult, LogsOut, error) {
	job, err := s.ownedJob(ctx, in.JobID)
	if err != nil {
		return nil, LogsOut{}, err
	}
	tail := in.Tail
	if tail <= 0 {
		tail = 80
	}
	if tail > 400 {
		tail = 400
	}
	var resp struct {
		PodsLogs []struct {
			Name string `json:"name"`
			Logs string `json:"logs"`
		} `json:"podsLogs"`
	}
	access, err := s.allowed(ctx, job.ProjectID, jobLogsClass)
	if err != nil {
		return nil, LogsOut{}, err
	}
	path, body := "analysis-export/getJobLogs", map[string]any{"projectId": job.ProjectID, "jobId": in.JobID}
	if access.legacy {
		path, body = "jobs/getJobLogs", map[string]any{"jobId": in.JobID}
	}
	if err := s.client.Post(ctx, path, body, &resp); err != nil {
		return nil, LogsOut{}, explain(err)
	}
	out := LogsOut{Job: *job, Errors: []string{}, Pods: []PodLog{}, Redacted: "credentials, tokens and signed-URL parameters are removed"}
	for _, p := range resp.PodsLogs {
		if strings.HasPrefix(p.Name, "describe-") {
			continue
		}
		lines := strings.Split(strings.TrimRight(Scrub(p.Logs), "\n"), "\n")
		for _, l := range lines {
			if isErrorLine(l) && len(out.Errors) < 30 {
				out.Errors = append(out.Errors, truncate(l, 400))
			}
		}
		if len(lines) > tail {
			lines = lines[len(lines)-tail:]
		}
		out.Pods = append(out.Pods, PodLog{Pod: p.Name, Lines: strings.Join(lines, "\n")})
	}
	return nil, out, nil
}

var errorLine = regexp.MustCompile(`\b(ERROR|CRITICAL|FATAL)\b|Traceback \(most recent call last\)|\b[A-Z][A-Za-z]*(Error|Exception)\b:|OOMKilled|exit code 137`)

func isErrorLine(l string) bool {
	return errorLine.MatchString(l)
}
