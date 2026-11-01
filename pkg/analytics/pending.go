package analytics

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// PendingEventsFile is where code-loader keeps the events it could not send (no network, e.g. inside an AI
// assistant's sandbox), relative to the folder holding leap.yaml. Each line is the JSON message code-loader
// would have posted to Mixpanel's /track.
var PendingEventsFile = filepath.Join(".tensorleap", "pending_events.jsonl")

// pendingEventsTimeout bounds the whole send, so a push never waits long for it
const pendingEventsTimeout = 3 * time.Second

// SendPendingEvents sends the events code-loader buffered in projectDir the way code-loader sends them, and
// removes the ones Mixpanel answered (sent or rejected). Best effort: it stops at the first network error or
// after pendingEventsTimeout, keeps what it could not send for next time, and never returns an error.
// It does nothing when TL_DISABLE_ANALYTICS is set, as code-loader does.
func SendPendingEvents(ctx context.Context, projectDir string) {
	if os.Getenv("TL_DISABLE_ANALYTICS") != "" {
		return
	}
	path := filepath.Join(projectDir, PendingEventsFile)
	data, err := os.ReadFile(path)
	if err != nil {
		return
	}
	// code-loader writes whole lines; anything after the last newline is still being written
	complete := data[:bytes.LastIndexByte(data, '\n')+1]
	if len(complete) == 0 {
		return
	}

	ctx, cancel := context.WithTimeout(ctx, pendingEventsTimeout)
	defer cancel()
	client := &http.Client{}
	var keep []string
	offline := false
	for _, line := range strings.Split(strings.TrimSuffix(string(complete), "\n"), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if offline || ctx.Err() != nil {
			keep = append(keep, line)
			continue
		}
		answered, err := sendPendingEvent(ctx, client, line)
		if err != nil {
			offline = true
		}
		if !answered {
			keep = append(keep, line)
		}
	}

	// keep what code-loader appended while the events were being sent
	now, err := os.ReadFile(path)
	if err != nil || !bytes.HasPrefix(now, complete) {
		return
	}
	var rest bytes.Buffer
	for _, line := range keep {
		rest.WriteString(line + "\n")
	}
	rest.Write(now[len(complete):])
	if rest.Len() == 0 {
		_ = os.Remove(path)
		return
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, rest.Bytes(), 0o644); err != nil {
		return
	}
	_ = os.Rename(tmp, path)
}

// sendPendingEvent posts one buffered message as code-loader's Mixpanel client does: a form with
// data=<message>, verbose=1 and ip=0. answered is true once Mixpanel has decided on it: it accepted it
// ({"status": 1}) or rejected it, which a resend would not change. err is a network error.
func sendPendingEvent(ctx context.Context, client *http.Client, message string) (answered bool, err error) {
	body := url.Values{"data": {message}, "verbose": {"1"}, "ip": {"0"}}.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(body))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() { _ = resp.Body.Close() }()
	switch {
	case resp.StatusCode >= 500:
		return false, nil
	case resp.StatusCode != http.StatusOK:
		return true, nil // a 4xx: resending the same message won't help
	}
	// {"status": 1} accepted, {"status": 0} rejected; anything else isn't Mixpanel's answer (e.g. a captive portal)
	var res struct {
		Status *int `json:"status"`
	}
	return json.NewDecoder(resp.Body).Decode(&res) == nil && res.Status != nil, nil
}
