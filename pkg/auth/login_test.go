package auth

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
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
