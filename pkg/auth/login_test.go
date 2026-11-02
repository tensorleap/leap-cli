package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLoginToKeycloakWrongPassword(t *testing.T) {
	keycloak := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"invalid_grant","error_description":"Invalid user credentials"}`))
	}))
	defer keycloak.Close()

	_, err := loginToKeycloak("tensorleap-client", "jane@acme.com", "wrong", "tensorleap", keycloak.URL)
	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("want ErrInvalidCredentials, got %v", err)
	}
}

func TestBrowserLoginWithoutOpeningABrowser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/v2/auth/getApiKeyByCode" {
			t.Errorf("unexpected request %s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"apiKey":"key-1"}`))
	}))
	defer server.Close()

	key, err := LoginAndGetAuthTokenWithBrowser(context.Background(), server.URL+"/api/v2", false)
	if err != nil || key != "key-1" {
		t.Fatalf("got %q, %v", key, err)
	}
}

// G92: inside a sandbox with no network the first poll fails; the CLI says so plainly instead of a raw dial error
func TestBrowserLoginSaysWhenTheServerCantBeReached(t *testing.T) {
	down := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	down.Close() // nothing listens there any more

	_, err := LoginAndGetAuthTokenWithBrowser(context.Background(), down.URL+"/api/v2", false)
	var unreachable *ServerUnreachableError
	if !errors.As(err, &unreachable) {
		t.Fatalf("want ServerUnreachableError, got %v", err)
	}
	msg := err.Error()
	if !strings.HasPrefix(msg, "Can't reach the Tensorleap server at "+down.URL+" from here (dial tcp") ||
		!strings.Contains(msg, "If you're an AI assistant running in a sandbox, run this command outside it.") {
		t.Fatalf("message: %q", msg)
	}
	if strings.Contains(msg, "getApiKeyByCode") || strings.Contains(msg, "failed to get API key") {
		t.Fatalf("the raw request error reached the user: %q", msg)
	}
}

// Before the user authorizes, the server answers the poll with an error status; the CLI keeps waiting
func TestBrowserLoginKeepsWaitingUntilAuthorized(t *testing.T) {
	prev := loginPollEvery
	loginPollEvery = time.Millisecond
	t.Cleanup(func() { loginPollEvery = prev })
	polls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		polls++
		if polls < 2 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(`{"message":"Invalid code verifier"}`))
			return
		}
		_, _ = w.Write([]byte(`{"apiKey":"key-2"}`))
	}))
	defer server.Close()

	key, err := LoginAndGetAuthTokenWithBrowser(context.Background(), server.URL+"/api/v2", false)
	if err != nil || key != "key-2" || polls != 2 {
		t.Fatalf("got %q, %v after %d polls", key, err, polls)
	}
}
