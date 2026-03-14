package resources_test

import (
	"context"
	"os"
	"testing"

	"github.com/Businnect/letit_go"
	"github.com/Businnect/letit_go/resources"
)

func TestCreateUserJobWithCompany_Integration(t *testing.T) {

	token := os.Getenv("LETIT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: LETIT_API_TOKEN not set")
	}

	client := letit.NewClient(token, "https://api.letit.com")

	logoPath := "../.git/logo.png"
	file, err := os.Open(logoPath)
	if err != nil {
		t.Fatalf("Failed to open test logo at %s: %v", logoPath, err)
	}
	defer file.Close()

	req := resources.CreateUserJobRequest{
		CompanyName:        "LetIt Go SDK Test",
		CompanyDescription: "A professional Go SDK integration test.",
		CompanyWebsite:     "https://letit.com",
		JobTitle:           "Software Engineer (Go)",
		JobDescription:     "Developing high-performance SDKs in Go.",
		JobHowToApply:      "https://letit.com/apply",
		JobSkills:          ptr("Go, PostgreSQL, Docker"),
		CompanyLogo: &resources.FilePayload{
			Filename: "logo.png",
			Reader:   file,
			MimeType: "image/png",
		},
	}

	ctx := context.Background()
	response, err := client.Job.CreateWithCompany(ctx, req)

	if err != nil {
		t.Fatalf("Integration test failed: %v", err)
	}

	if response.Slug == "" {
		t.Error("Expected a non-empty slug in the response")
	}
}

func ptr(s string) *string {
	return &s
}