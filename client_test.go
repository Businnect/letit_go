package letit

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_Do_Errors(t *testing.T) {
	tests := []struct {
		name           string
		statusResponse int
		bodyResponse   string
		expectedError  string
	}{
		{
			name:           "Internal Server Error 500",
			statusResponse: http.StatusInternalServerError,
			bodyResponse:   `{"error": "something went wrong"}`,
			expectedError:  "api error: status 500",
		},
		{
			name:           "Invalid API Token 401",
			statusResponse: http.StatusUnauthorized,
			bodyResponse:   `{"invalid_api_user_token": "USER-API-TOKEN header is not valid"}`,
			expectedError:  "USER-API-TOKEN header is not valid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusResponse)
				w.Write([]byte(tt.bodyResponse))
			}))
			defer server.Close()

			client := NewClient("fake-key", server.URL)
			req, _ := http.NewRequest("GET", "/test", nil)

			_, err := client.Do(req)

			if err == nil {
				t.Fatal("expected an error but got nil")
			}

			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Errorf("expected error to contain '%s', got '%v'", tt.expectedError, err)
			}
		})
	}
}