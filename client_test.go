package letit

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

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
