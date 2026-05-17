package api

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

func TestHintVersionMismatch_PassThroughUnrelated(t *testing.T) {
	in := errors.New("disk is on fire")
	got := HintVersionMismatch(in)
	if got != in {
		t.Fatalf("expected unrelated error to pass through unchanged, got %q", got)
	}
}

func TestHintVersionMismatch_Nil(t *testing.T) {
	if HintVersionMismatch(nil) != nil {
		t.Fatalf("expected nil to pass through")
	}
}

func TestHintVersionMismatch_Detects(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{"404 raw status", errors.New("404 Not Found")},
		{"405 raw status", errors.New("405 Method Not Allowed")},
		{"openapi decode missing property", errors.New("no value given for required property foo")},
		{"HTTPError with 404", &HTTPError{StatusCode: http.StatusNotFound, URL: "/x"}},
		{"HTTPError with 405", &HTTPError{StatusCode: http.StatusMethodNotAllowed, URL: "/x"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := HintVersionMismatch(tc.err)
			if got == nil {
				t.Fatalf("expected wrapped error, got nil")
			}
			if !strings.Contains(got.Error(), "leap server upgrade") {
				t.Fatalf("expected hint to mention 'leap server upgrade', got: %s", got)
			}
			// Wrapped error chain must still unwrap to the original cause.
			if !errors.Is(got, tc.err) {
				t.Fatalf("errors.Is failed: wrapped does not chain back to original cause")
			}
		})
	}
}

func TestHintVersionMismatch_DoesNotMatch200LikeError(t *testing.T) {
	in := &HTTPError{StatusCode: http.StatusInternalServerError, URL: "/x"}
	got := HintVersionMismatch(in)
	if got != in {
		t.Fatalf("expected 500 error to pass through unchanged, got %q", got)
	}
}
