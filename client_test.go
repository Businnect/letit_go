package letit

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestClient_Connection(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth != "Bearer test-api-key" {
			t.Errorf("expected Authorization header 'Bearer test-api-key', got '%s'", auth)
		}

		ua := r.Header.Get("User-Agent")
		if ua != "LetIt-Go-SDK/1.0" {
			t.Errorf("expected User-Agent 'LetIt-Go-SDK/1.0', got '%s'", ua)
		}

		if r.URL.Path != "/api/v1/ping" {
			t.Errorf("expected path '/api/v1/ping', got '%s'", r.URL.Path)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "connected"}`))
	}))
	defer server.Close()

	client := NewClient("test-api-key", server.URL)

	req, err := http.NewRequest("GET", "/api/v1/ping", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	respBody, err := client.Do(req)
	if err != nil {
		t.Fatalf("client.Do failed: %v", err)
	}
	defer respBody.Close()

	content, err := io.ReadAll(respBody)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	if !strings.Contains(string(content), "connected") {
		t.Errorf("expected response to contain 'connected', got '%s'", string(content))
	}
}

func TestClient_Connection_Error(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	apiKey := os.Getenv("LETIT_API_KEY")
	client := NewClient(apiKey, server.URL)
	req, _ := http.NewRequest("GET", "/fail", nil)

	_, err := client.Do(req)
	if err == nil {
		t.Error("expected error for HTTP 500 status code, but got nil")
	}

	expectedErrorMsg := "api error: status 500"
	if err != nil && !strings.Contains(err.Error(), expectedErrorMsg) {
		t.Errorf("expected error message to contain '%s', got '%v'", expectedErrorMsg, err)
	}
}