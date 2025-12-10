# chirpy

A simple HTTP server written in Go that serves static files with request tracking, health monitoring, and chirp validation.

## Features

- **Static file server**: Serves files from the root directory at `/app/*` endpoint
- **Request tracking**: Automatically counts all file server requests using thread-safe atomic operations
- **Metrics dashboard**: HTML dashboard displaying total file server visits
- **Chirp validation**: Validates chirp messages (max 140 characters)
- **Health check endpoint**: `/api/healthz` endpoint returns HTTP 200 OK status
- **Admin reset**: Reset the request counter back to zero
- **Lightweight**: Built using only Go's standard library

## Getting Started

### Prerequisites

- Go 1.25.2 or later

### Building

```bash
go build -o out
```

### Running

```bash
./out
```

The server will start on port 8080.

Alternatively, build and run in one command:

```bash
go run main.go
```

## API Endpoints

### Static Files
- **Path**: `/app/*`
- **Method**: GET
- **Description**: Serves static files from the project root directory. Each request is automatically tracked.
- **Example**: `http://localhost:8080/app/index.html`

### Health Check
- **Path**: `/api/healthz`
- **Method**: GET
- **Description**: Returns server status
- **Response**: `200 OK` with plain text "OK"

### Validate Chirp
- **Path**: `/api/validate_chirp`
- **Method**: POST
- **Description**: Validates a chirp message (must be 140 characters or less)
- **Request Body**: `{"body": "Your chirp message here"}`
- **Response**: `200 OK` with JSON `{"valid": true}` or `400 Bad Request` with error message

### Admin Metrics
- **Path**: `/admin/metrics`
- **Method**: GET
- **Description**: Returns an HTML page displaying the total number of file server visits
- **Response**: `200 OK` with HTML dashboard

### Admin Reset
- **Path**: `/admin/reset`
- **Method**: POST
- **Description**: Resets the file server request counter to zero
- **Response**: `200 OK` with plain text "Hits reset to 0"

## Project Structure

```
.
├── main.go                      # Application entry point and server setup
├── handler_validate.go          # Chirp validation handler
├── json.go                      # JSON response utilities
├── metrics.go                   # Metrics middleware and handler
├── readiness.go                 # Health check handler
├── reset.go                     # Reset handler
├── index.html                   # Sample HTML page
├── go.mod                       # Go module definition
└── README.md                    # This file
```

The project currently uses a single-package structure with all code in the `main` package. As the application grows, consider refactoring into a more structured layout with `cmd/` and `internal/` directories

## Development

### Testing

```bash
go test ./...
```

### Formatting

```bash
go fmt ./...
```

### Vetting

```bash
go vet ./...
```
