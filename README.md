# LetIt Go SDK

A professional Go client for the LetIt API, featuring high-performance support for **Microposts** and **Job** management.

## 📖 API Documentation

For detailed information on the underlying REST API, endpoints, and authentication schemas, please visit the official documentation:
* **API Reference**: [http://api.letit.com](https://api.letit.com/docs/client/)

## Features

* **Job Management**: Full support for creating job postings with company logos, descriptions, and metadata.
* **Micropost System**: Create text posts or file-based updates with community-specific targeting.
* **Context Support**: Full `context.Context` support for timeouts and cancellations in every request.

## Installation

```bash
go get github.com/Businnect/letit_go

```

## Quick Start

### Initialize the Client

The client can be initialized with an explicit API key or automatically look for environment variables.

```go
import "github.com/Businnect/letit_go"

func main() {
    // Pass the API key and Base URL
    client := letit.NewClient("your-api-token", "https://api.letit.com")
}

```

### Create a Job with Company Logo

The SDK handles the `multipart` form construction and file streaming automatically.

```go
import (
    "context"
    "os"
    "github.com/Businnect/letit_go/resources"
)

// Open your logo file
file, _ := os.Open("logo.png")
defer file.Close()

req := resources.CreateUserJobRequest{
    CompanyName:    "Acme Corp",
    JobTitle:       "Senior Go Developer",
    JobDescription: "Building next-gen SDKs",
    CompanyLogo: &resources.FilePayload{
        Filename: "logo.png",
        Reader:   file,
    },
}

resp, err := client.Job.CreateWithCompany(context.Background(), req)
if err != nil {
    panic(err)
}
fmt.Printf("Job created successfully: %s\n", resp.Slug)

```

### Create a Micropost

Easily create posts with optional titles and bodies.

```go
title := "New Update"
req := resources.CreateMicropostRequest{
    Title: &title,
    Body:  "The Go SDK is now live!",
}

resp, err := client.Micropost.Create(context.Background(), req)
if err != nil {
    panic(err)
}
fmt.Printf("Post created with ID: %s\n", resp.PublicID)

```

## Environment Variables

The SDK can utilize the following environment variables for testing or default configuration:

* `LETIT_API_TOKEN`: Used by integration tests to authenticate against the live API.

## Testing

Run the test suite using the standard Go toolchain:

```powershell
# In PowerShell
$env:LETIT_API_TOKEN="your-token"; go test -v ./...

```

```bash
# In Bash
LETIT_API_TOKEN="your-token" go test -v ./...

```