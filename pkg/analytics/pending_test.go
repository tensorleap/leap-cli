package analytics

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// writePending makes a project folder holding code-loader's buffered events and returns it
func writePending(t *testing.T, content string) (dir, path string) {
	t.Helper()
	dir = t.TempDir()
	path = filepath.Join(dir, PendingEventsFile)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir, path
}

// fakeMixpanel answers each message with answer(data) as code-loader's Mixpanel endpoint would, and records what it got
func fakeMixpanel(t *testing.T, answer func(w http.ResponseWriter, data string)) *[]string {
	t.Helper()
	var got []string
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
			t.Errorf("%s %s", r.Method, r.Header.Get("Content-Type"))
		}
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		if r.PostForm.Get("verbose") != "1" || r.PostForm.Get("ip") != "0" {
			t.Errorf("form %v", r.PostForm)
		}
		data := r.PostForm.Get("data")
		got = append(got, data)
		answer(w, data)
	}))
	t.Cleanup(ts.Close)
	old := endpoint
	t.Cleanup(func() { endpoint = old })
	endpoint = ts.URL
	return &got
}

func accept(w http.ResponseWriter, _ string) { _, _ = w.Write([]byte(`{"status": 1, "error": null}`)) }

const (
	ev1 = `{"event":"code_loader_event","properties":{"token":"t","distinct_id":"d","$insert_id":"1","email":"a+b@c.d"}}`
	ev2 = `{"event":"code_loader_event","properties":{"token":"t","distinct_id":"d","$insert_id":"2"}}`
	ev3 = `{"event":"code_loader_event","properties":{"token":"t","distinct_id":"d","$insert_id":"3"}}`
)

func TestPendingEventsAreSentAsIsAndTheFileRemoved(t *testing.T) {
	t.Setenv("TL_DISABLE_ANALYTICS", "")
	got := fakeMixpanel(t, accept)
	dir, path := writePending(t, ev1+"\n\n"+ev2+"\n")

	SendPendingEvents(context.Background(), dir)

	if strings.Join(*got, "|") != ev1+"|"+ev2 {
		t.Fatalf("sent %q", *got)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("file still there: %v", err)
	}
}

func TestPendingEventsKeepWhatWasNotSent(t *testing.T) {
	t.Setenv("TL_DISABLE_ANALYTICS", "")
	fakeMixpanel(t, func(w http.ResponseWriter, data string) {
		switch {
		case strings.Contains(data, `"$insert_id":"1"`):
			accept(w, data)
		case strings.Contains(data, `"$insert_id":"2"`):
			_, _ = w.Write([]byte(`{"status": 0, "error": "bad event"}`)) // rejected: dropped, a resend won't help
		default:
			w.WriteHeader(http.StatusServiceUnavailable) // kept for the next push
		}
	})
	dir, path := writePending(t, ev1+"\n"+ev2+"\n"+ev3+"\n"+`{"partial`)

	SendPendingEvents(context.Background(), dir)

	if b, _ := os.ReadFile(path); string(b) != ev3+"\n"+`{"partial` {
		t.Fatalf("left %q", b)
	}
}

func TestPendingEventsStopAtTheFirstNetworkError(t *testing.T) {
	t.Setenv("TL_DISABLE_ANALYTICS", "")
	got := fakeMixpanel(t, accept)
	old := endpoint
	endpoint = "http://127.0.0.1:1" // nothing listens there
	t.Cleanup(func() { endpoint = old })
	content := ev1 + "\n" + ev2 + "\n"
	dir, path := writePending(t, content)

	SendPendingEvents(context.Background(), dir)

	if b, _ := os.ReadFile(path); string(b) != content || len(*got) != 0 {
		t.Fatalf("left %q, sent %q", b, *got)
	}
}

func TestPendingEventsKeepLinesAppendedWhileSending(t *testing.T) {
	t.Setenv("TL_DISABLE_ANALYTICS", "")
	var path string
	fakeMixpanel(t, func(w http.ResponseWriter, data string) {
		if strings.Contains(data, `"$insert_id":"1"`) {
			f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0)
			if err != nil {
				t.Fatal(err)
			}
			_, _ = f.WriteString(ev3 + "\n")
			_ = f.Close()
		}
		accept(w, data)
	})
	var dir string
	dir, path = writePending(t, ev1+"\n")

	SendPendingEvents(context.Background(), dir)

	if b, _ := os.ReadFile(path); string(b) != ev3+"\n" {
		t.Fatalf("left %q", b)
	}
}

func TestPendingEventsHonorTheAnalyticsOptOut(t *testing.T) {
	t.Setenv("TL_DISABLE_ANALYTICS", "true")
	got := fakeMixpanel(t, accept)
	content := ev1 + "\n"
	dir, path := writePending(t, content)

	SendPendingEvents(context.Background(), dir)

	if b, _ := os.ReadFile(path); string(b) != content || len(*got) != 0 {
		t.Fatalf("left %q, sent %q", b, *got)
	}
}

func TestNoPendingEventsFileSendsNothing(t *testing.T) {
	got := fakeMixpanel(t, accept)
	SendPendingEvents(context.Background(), t.TempDir())
	if len(*got) != 0 {
		t.Fatalf("sent %q", *got)
	}
}
