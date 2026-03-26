package resources_test

import (
	"context"
	"os"
	"testing"

	letit "github.com/Businnect/letit_go"
)

func TestAdminBlogGet_Integration(t *testing.T) {
	token := os.Getenv("LETIT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: LETIT_API_TOKEN not set")
	}

	client := letit.NewClient(token, "https://api.letit.com")

	res, err := client.AdminBlog.Get(context.Background())
	if err != nil {
		t.Fatalf("integration test failed: %v", err)
	}

	if res != nil {
		if _, ok := res["slug"]; !ok {
			t.Log("admin blog response is not null but has no slug field")
		}
	}
}

func TestAdminBlogList_Integration(t *testing.T) {
	token := os.Getenv("LETIT_API_TOKEN")
	if token == "" {
		t.Skip("Skipping integration test: LETIT_API_TOKEN not set")
	}

	client := letit.NewClient(token, "https://api.letit.com")

	res, err := client.AdminBlog.List(context.Background())
	if err != nil {
		t.Fatalf("integration test failed: %v", err)
	}

	if res == nil {
		t.Fatal("expected non-nil list response")
	}

	if res.TotalList < 0 || res.TotalPages < 0 {
		t.Fatalf("expected non-negative totals, got total_list=%d total_pages=%d", res.TotalList, res.TotalPages)
	}
}
