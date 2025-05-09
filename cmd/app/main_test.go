package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AGiantSquid/golang-template/internal/handler"
)

func TestMainIntegration(t *testing.T) {
	// Create a new mux and register our handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.HelloHandler)

	// Start a test server with our mux
	server := httptest.NewServer(mux)
	defer server.Close()

	// Make a request to the server
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("Could not send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK; got %v", resp.Status)
	}

	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Could not read response body: %v", err)
	}

	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("Expected response body %q; got %q", expected, string(body))
	}
}
