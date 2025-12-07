# chirpy

A simple HTTP server written in Go that serves static files with request tracking and monitoring capabilities.

## Features

- **Static file server**: Serves files from the root directory at `/app/*` endpoint
- **Request tracking**: Automatically counts all file server requests using thread-safe atomic operations
- **Metrics endpoint**: View the total number of requests to the file server
- **Reset functionality**: Reset the request counter back to zero
- **Health check endpoint**: `/healthz` endpoint returns HTTP 200 OK status
- **Lightweight**: Built using only Go's standard library

## Getting Started

### Prerequisites

- Go 1.25.2 or later

### Building

```bash
go build -o out ./cmd/chirpy
```

### Running

```bash
./out
```

The server will start on port 8080.

Alternatively, build and run in one command:

```bash
go run ./cmd/chirpy
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

### Metrics
- **Path**: `/api/metrics`
- **Method**: GET
- **Description**: Returns the total number of requests made to the file server
- **Response**: `200 OK` with plain text showing hit count (e.g., "Hits: 42")

### Reset Metrics
- **Path**: `/api/reset`
- **Method**: POST
- **Description**: Resets the file server request counter to zero
- **Response**: `200 OK` with plain text "Hits reset to 0"

## Project Structure

```
.
├── cmd/
│   └── chirpy/
│       └── main.go              # Application entry point
├── internal/
│   ├── handlers/
│   │   ├── readiness.go         # Health check handler
│   │   └── reset.go             # Reset handler and APIConfig type
│   └── middleware/
│       └── metrics.go           # Metrics middleware and handler
├── assets/                      # Static assets directory
│   └── logo.png                 # Project logo
├── index.html                   # Sample HTML page
├── go.mod                       # Go module definition
└── README.md                    # This file
```

This project follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout):
- `cmd/chirpy/` - Main application entry point
- `internal/` - Private application code (not importable by external projects)
- `internal/handlers/` - HTTP request handlers
- `internal/middleware/` - HTTP middleware functions

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
