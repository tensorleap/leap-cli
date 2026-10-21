package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestForbiddenCarriesServerMessage(t *testing.T) {
	const body = `{"error":"activate first","code":"LICENSE_LIMIT"}`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(body))
	}))
	defer srv.Close()
	res, err := NewDefaultClient().Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	got, _ := io.ReadAll(res.Body)
	if res.Status != "activate first" || string(got) != body {
		t.Fatalf("status %q body %q", res.Status, got)
	}
}
