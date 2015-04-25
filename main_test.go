package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntegration(t *testing.T) {
	// Make request to the server.
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	HelloWorldServer(w, r)

	// Body must contain "Hello, World!" message.
	if b := w.Body.String(); !strings.Contains(b, "Hello, World!") {
		t.Fatalf("body = %s, want Hello, World!", b)
	}
}
