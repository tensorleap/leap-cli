package auth

import (
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
