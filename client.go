package letit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Businnect/letit_go/resources"
)

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client

	Micropost *resources.MicropostResource
	Job       *resources.JobResource
}

func NewClient(apiKey string, baseURL string) *Client {
	c := &Client{
		baseURL: strings.TrimSuffix(baseURL, "/"),
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	c.Micropost = resources.NewMicropostResource(c)
	c.Job = resources.NewJobResource(c)

	return c
}

func (c *Client) Do(req *http.Request) (io.ReadCloser, error) {
	if !strings.HasPrefix(req.URL.String(), "http") {
		req.URL, _ = req.URL.Parse(c.baseURL + req.URL.Path)
	}

	req.Header.Set("USER-API-TOKEN", c.apiKey)
	req.Header.Set("User-Agent", "LetIt-Go-SDK/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		resp.Body.Close()

		var errorResp struct {
			InvalidToken string `json:"invalid_api_user_token"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&errorResp); err == nil && errorResp.InvalidToken != "" {
			return nil, fmt.Errorf("api error (401): %s", errorResp.InvalidToken)
		}

		return nil, fmt.Errorf("api error: status %d", resp.StatusCode)
	}

	return resp.Body, nil
}
