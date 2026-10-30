package analytics

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

// sends events to a local stand-in for Mixpanel, with the saved ids in a temp dir, and returns what it received
func captureEvents(t *testing.T) *[]map[string]interface{} {
	t.Helper()
	var got []map[string]interface{}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			t.Error(err)
		}
		var event struct {
			Properties map[string]interface{} `json:"properties"`
		}
		if err := json.Unmarshal([]byte(r.PostForm.Get("data")), &event); err != nil {
			t.Errorf("data is not JSON: %v", err)
		}
		got = append(got, event.Properties)
	}))
	t.Cleanup(ts.Close)
	dir := t.TempDir()
	oldEndpoint, oldUser, oldDevice, oldEmail := endpoint, userIDFile, deviceIDFile, emailFile
	t.Cleanup(func() { endpoint, userIDFile, deviceIDFile, emailFile = oldEndpoint, oldUser, oldDevice, oldEmail })
	endpoint = ts.URL
	userIDFile, deviceIDFile, emailFile = filepath.Join(dir, "user_id"), filepath.Join(dir, "device_id"), filepath.Join(dir, "email")
	return &got
}

func TestEmailKeepsItsPlus(t *testing.T) {
	got := captureEvents(t)
	SendEvent(EventServerInstallStarted, map[string]interface{}{"email": "tom.koren+e2e1@tensorleap.ai"})
	if len(*got) != 1 || (*got)[0]["email"] != "tom.koren+e2e1@tensorleap.ai" {
		t.Fatalf("got %v", *got)
	}
}

func TestEveryEventCarriesTheLastKnownEmailAndUser(t *testing.T) {
	got := captureEvents(t)
	SendEvent(EventCliProjectsPushStarted, nil)
	if _, ok := (*got)[0]["email"]; ok {
		t.Fatalf("no email is known yet: %v", (*got)[0])
	}

	SendEvent(EventMcpConfig, map[string]interface{}{"email": "a+b@c.d"})
	if err := SaveUserIDIfMissing("cid-admin"); err != nil {
		t.Fatal(err)
	}
	SendEvent(EventCliProjectsPushStarted, nil)
	push := (*got)[2]
	if push["email"] != "a+b@c.d" || push["$user_id"] != "cid-admin" || push["distinct_id"] != "cid-admin" || push["$device_id"] == "" {
		t.Fatalf("push: %v", push)
	}

	// leap auth login names the user it signed in; that wins over the license's admin
	SendEvent(EventAuthLoginSuccess, map[string]interface{}{"email": "e@f.g", "user_id": "cid-login"})
	if err := SaveUserIDIfMissing("cid-other"); err != nil {
		t.Fatal(err)
	}
	SendEvent(EventAuthLoginFailed, map[string]interface{}{"success": false})
	failed := (*got)[4]
	if failed["email"] != "e@f.g" || failed["$user_id"] != "cid-login" {
		t.Fatalf("failed login: %v", failed)
	}
}
