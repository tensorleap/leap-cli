package api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

// HTTPError represents an error with details about an HTTP response.
type HTTPError struct {
	URL        string
	StatusCode int
	Body       string
}

// ConcurrentEvaluateLimitCode is the machine-readable code the node-server
// returns on the structured 429 payload when the per-user concurrent
// Evaluate/Continue-Evaluate jobs limit is reached. Mirrors the constant in
// node-server's evaluate logic.
const ConcurrentEvaluateLimitCode = "CONCURRENT_EVALUATE_LIMIT"

// RunningEvaluateJobInfo mirrors the entries the backend sends inside
// `runningJobs` on a CONCURRENT_EVALUATE_LIMIT error. Used by the CLI to
// offer an interactive terminate-and-retry flow.
type RunningEvaluateJobInfo struct {
	JobId       string `json:"jobId"`
	ProjectId   string `json:"projectId"`
	ProjectName string `json:"projectName,omitempty"`
	VersionId   string `json:"versionId"`
	VersionName string `json:"versionName,omitempty"`
	StartedAt   string `json:"startedAt"`
	SubType     string `json:"subType"`
}

// ConcurrentEvaluateLimitError is returned by CheckRes when an Evaluate or
// Continue-Evaluate request is rejected because the user is already at their
// per-user concurrent-jobs limit. Callers can type-assert on this error to
// offer a friendly recovery flow (terminate one of the running jobs, then
// retry the original action).
type ConcurrentEvaluateLimitError struct {
	URL         string
	Message     string
	Limit       int
	RunningJobs []RunningEvaluateJobInfo
}

func (e *ConcurrentEvaluateLimitError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return fmt.Sprintf(
		"Concurrent Evaluate jobs limit reached (limit=%d, running=%d)",
		e.Limit, len(e.RunningJobs),
	)
}

const LEAP_SKIP_SSL_VERIFY = "LEAP_SKIP_SSL_VERIFY"

var ErrAuth = errors.New("user is not currently authenticated, to authenticate - please run 'leap auth login'")

func (e *HTTPError) Error() string {
	if e.StatusCode == 401 {
		return ErrAuth.Error()
	}
	if e.StatusCode == 403 {
		return extractServerMessage(e.Body)
	}
	return fmt.Sprintf("Error: HTTP status code %d\nURL: %s\nResponse body:\n%s", e.StatusCode, e.URL, e.Body)
}

func extractServerMessage(body string) string {
	var parsed struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal([]byte(body), &parsed); err == nil {
		if parsed.Error != "" {
			return parsed.Error
		}
		if parsed.Message != "" {
			return parsed.Message
		}
	}
	if body != "" {
		return body
	}
	return "Forbidden"
}

type CustomRoundTripper struct {
	originalRoundTripper http.Transport
}

func (c *CustomRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	thereIsNoDomain := strings.HasPrefix(req.URL.String(), "/")
	if thereIsNoDomain {
		return nil, ErrAuth
	}
	res, err := c.originalRoundTripper.RoundTrip(req)
	if err != nil {
		return res, err
	}
	if res.StatusCode == 401 {
		return res, ErrAuth
	}
	return res, nil
}

// NewDefaultClient returns a new http.Client with default settings.
func NewDefaultClient() *http.Client {
	customHttpClient := &http.Client{}

	roundTripper := &CustomRoundTripper{
		originalRoundTripper: http.Transport{},
	}
	customHttpClient.Transport = roundTripper

	if os.Getenv(LEAP_SKIP_SSL_VERIFY) == "true" {
		// skip SSL verification
		roundTripper.originalRoundTripper.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	return customHttpClient
}

func CheckRes(response *http.Response, err error) error {
	if err != nil {
		if response != nil && response.StatusCode == 403 {
			body := extractBodyFromError(err)
			return &HTTPError{
				URL:        response.Request.URL.String(),
				StatusCode: 403,
				Body:       body,
			}
		}
		// The auto-generated OpenAPI client returns a non-nil error for any
		// non-2xx response (e.g. 429), so the structured-error detection has
		// to live here, not just below in the `err == nil` branch.
		if response != nil && response.StatusCode == http.StatusTooManyRequests {
			body := extractBodyFromError(err)
			if cle := parseConcurrentEvaluateLimitError(response.Request.URL.String(), []byte(body)); cle != nil {
				return cle
			}
		}
		return err
	}
	if IsValidStatus(response) {
		return nil
	}

	if response.StatusCode == 401 {
		return &HTTPError{
			URL:        response.Request.URL.String(),
			StatusCode: response.StatusCode,
			Body:       ErrAuth.Error(),
		}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		body = []byte(fmt.Sprintf("Error reading response body: %s", err))
	}

	// Detect the structured concurrent-evaluate-limit error so the caller can
	// offer a friendly recovery flow instead of dumping the raw JSON.
	if response.StatusCode == http.StatusTooManyRequests {
		if cle := parseConcurrentEvaluateLimitError(response.Request.URL.String(), body); cle != nil {
			return cle
		}
	}

	return &HTTPError{
		URL:        response.Request.URL.String(),
		StatusCode: response.StatusCode,
		Body:       string(body),
	}
}

// parseConcurrentEvaluateLimitError returns a typed error if `body` is the
// node-server's structured CONCURRENT_EVALUATE_LIMIT payload, otherwise nil
// (so the caller can fall back to a generic HTTPError).
func parseConcurrentEvaluateLimitError(url string, body []byte) *ConcurrentEvaluateLimitError {
	var parsed struct {
		Error       string                   `json:"error"`
		Code        string                   `json:"code"`
		Limit       int                      `json:"limit"`
		RunningJobs []RunningEvaluateJobInfo `json:"runningJobs"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil
	}
	if parsed.Code != ConcurrentEvaluateLimitCode {
		return nil
	}
	return &ConcurrentEvaluateLimitError{
		URL:         url,
		Message:     parsed.Error,
		Limit:       parsed.Limit,
		RunningJobs: parsed.RunningJobs,
	}
}

func extractBodyFromError(err error) string {
	type bodyProvider interface {
		Body() []byte
	}
	var bp bodyProvider
	if errors.As(err, &bp) {
		if b := bp.Body(); len(b) > 0 {
			return string(b)
		}
	}
	return err.Error()
}

func IsValidStatus(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300
}

// HintVersionMismatch returns err unchanged unless its shape suggests a
// CLI/server version mismatch, in which case it wraps err with an actionable
// hint telling the user to upgrade the server or the CLI. The detection is
// duck-typed on the error string so it works for both *HTTPError and
// *tensorleapapi.GenericOpenAPIError without coupling to the generated code.
func HintVersionMismatch(err error) error {
	if err == nil || !looksLikeVersionMismatch(err) {
		return err
	}
	return fmt.Errorf(
		"%w\n\nThis may indicate a CLI/server version mismatch. "+
			"Run 'leap server upgrade' on the server host to upgrade the server, "+
			"or 'leap cli upgrade -s | bash' to upgrade the CLI",
		err,
	)
}

func looksLikeVersionMismatch(err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	if strings.Contains(msg, "404 Not Found") || strings.Contains(msg, "405 Method Not Allowed") {
		return true
	}
	// OpenAPI client emits this when a response is missing a required field —
	// the typical shape of a server-too-old response decode failure.
	if strings.Contains(msg, "no value given for required property") {
		return true
	}
	var httpErr *HTTPError
	if errors.As(err, &httpErr) {
		if httpErr.StatusCode == http.StatusNotFound || httpErr.StatusCode == http.StatusMethodNotAllowed {
			return true
		}
	}
	return false
}

func init() {
	http.DefaultClient = NewDefaultClient()
}

const API_PATH = "/api/v2"
const API_CLOUD_SUBDOMAIN = "://api."

func isCloudUrl(url string) bool {
	return strings.Contains(url, ".tensorleap.ai")
}

func NormalizeAPIUrl(url string) (string, error) {
	url = strings.TrimSuffix(url, "/")
	hasProtocol := strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://")
	if !hasProtocol {
		return "", fmt.Errorf("URL must start with http:// or https://")
	}
	isCloud := isCloudUrl(url)
	hasApiSubdomain := strings.Contains(url, API_CLOUD_SUBDOMAIN)
	if isCloud && !hasApiSubdomain {
		url = strings.Replace(url, "://", API_CLOUD_SUBDOMAIN, 1)
	}

	if strings.HasSuffix(url, API_PATH) {
		return url, nil
	}
	return fmt.Sprintf("%s%s", url, API_PATH), nil
}

func ChangeToUIUrl(apiUrl string) string {
	url := strings.Replace(apiUrl, API_PATH, "", 1)
	return strings.Replace(url, API_CLOUD_SUBDOMAIN, "://", 1)
}

func StepsFromJob(job *tensorleapapi.Job) []log.Step {
	jobSteps := []log.Step{}
	if job.EventsSnapshot == nil {
		return jobSteps
	}

	mapStatus := func(status tensorleapapi.StatusEnum) log.StepStatus {
		switch status {
		case tensorleapapi.STATUSENUM_UNSTARTED:
			return log.StepStatusPending
		case tensorleapapi.STATUSENUM_STARTED:
			return log.StepStatusRunning
		case tensorleapapi.STATUSENUM_FINISHED:
			return log.StepStatusDone
		case tensorleapapi.STATUSENUM_FAILED:
			return log.StepStatusFailed
		default:
			return log.StepStatusWaiting
		}
	}
	for _, event := range job.EventsSnapshot.Events {
		var current, total float64
		if event.Progress != nil {
			current = event.Progress.Current
			total = event.Progress.Total
		}
		jobSteps = append(jobSteps, log.Step{
			Name:    event.Name,
			Status:  mapStatus(event.Status),
			Current: current,
			Total:   total,
		})
	}
	return jobSteps
}

func IsJobFailed(jobStatus tensorleapapi.JobStatus) bool {
	return jobStatus == tensorleapapi.JOBSTATUS_FAILED || jobStatus == tensorleapapi.JOBSTATUS_TERMINATED || jobStatus == tensorleapapi.JOBSTATUS_STOPPED
}

func IsJobRunning(jobStatus tensorleapapi.JobStatus) bool {
	return jobStatus == tensorleapapi.JOBSTATUS_STARTED || jobStatus == tensorleapapi.JOBSTATUS_PENDING || jobStatus == tensorleapapi.JOBSTATUS_UNSTARTED
}

func IsJobFinished(jobStatus tensorleapapi.JobStatus) bool {
	return jobStatus == tensorleapapi.JOBSTATUS_FINISHED
}

func ParseAndFormatDateToLocalTime(raw string) (string, error) {
	// Trim the (Coordinated Universal Time) part
	clean := strings.SplitN(raw, " (", 2)[0]

	// JS-style format: Mon Sep 08 2025 14:46:56 GMT+0000
	const layoutIn = "Mon Jan 02 2006 15:04:05 GMT-0700"

	t, err := time.Parse(layoutIn, clean)
	if err != nil {
		return "", fmt.Errorf("failed to parse time %q: %w", clean, err)
	}

	return FormatDateToLocalTime(t), nil
}

func FormatDateToLocalTime(t time.Time) string {

	// Convert to local time zone
	localTime := t.Local()

	// Desired output: Mon, 08 Sep 2025 14:46 (local)
	return localTime.Format("Mon, 02 Jan 2006 15:04")
}
