# chirpy

A simple HTTP server written in Go that serves static files and provides a health check endpoint.

## Features

- **Static file server**: Serves files from the root directory at `/app/*` endpoint
- **Health check endpoint**: `/healthz` endpoint returns HTTP 200 OK status
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
- **Description**: Serves static files from the project root directory
- **Example**: `http://localhost:8080/app/index.html`

### Health Check
- **Path**: `/healthz`
- **Method**: GET
- **Description**: Returns server status
- **Response**: `200 OK` with plain text "OK"

## Project Structure

```
.
├── main.go          # Main application with HTTP handlers
├── index.html       # Sample HTML page
├── assets/          # Static assets directory
│   └── logo.png     # Project logo
├── go.mod           # Go module definition
└── README.md        # This file
```

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
