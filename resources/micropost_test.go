package resources_test

import (
	"context"
	"os"
	"testing"

	"github.com/Businnect/letit_go"
	"github.com/Businnect/letit_go/resources"
)

func TestClientCreateMicropost_Integration(t *testing.T) {

	token := os.Getenv("LETIT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: LETIT_API_TOKEN not set")
	}

	client := letit.NewClient(token, "https://api.letit.com")

	title := "Test"
	req := resources.CreateMicropostRequest{
		Title: &title,
		Body:  "Hello",
	}

	ctx := context.Background()
	response, err := client.Micropost.Create(ctx, req)

	if err != nil {
		t.Fatalf("Integration test failed: %v", err)
	}

	if response.PublicID == "" {
		t.Error("Expected a non-empty PublicID in the response")
	}
}
