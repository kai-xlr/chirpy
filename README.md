# chirpy

A simple HTTP server written in Go.

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
