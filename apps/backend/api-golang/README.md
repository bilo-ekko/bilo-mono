# Go API

A Go Lang REST API built with standard library HTTP server.

## Features

- RESTful API endpoints
- Impact Partners management
- Impact Projects management
- Shared packages integration:
  - Logger for structured logging
  - Calculator for mathematical operations

## Shared Packages

This application uses shared Go packages from the monorepo:

- `packages/go/common/logger` - Structured logging with context
- `packages/go/common/calculator` - Mathematical calculation utilities

## Development

### Prerequisites

- Go 1.25.3 or higher
- Moon CLI (for task running)

### Running Tests

```bash
# Run all tests
moon run api-golang:test

# Run tests with coverage
moon run api-golang:test-coverage

# Run tests directly with go
go test -v ./...

# Run specific package tests
go test -v ./cmd/api
go test -v github.com/bilo-mono/packages/common/calculator
go test -v github.com/bilo-mono/packages/common/logger
```

### Development Server

```bash
# Start dev server
moon run api-golang:dev

# Or with Go directly
go run ./cmd/api/main.go
```

### Building

```bash
# Build the application
moon run api-golang:build

# Or with Go directly
go build -o bin/api-golang ./cmd/api
```

### Other Tasks

```bash
# Format code
moon run api-golang:format

# Lint code
moon run api-golang:lint

# Tidy dependencies
moon run api-golang:tidy
```

## API Endpoints

### Health & Info

- `GET /` - Welcome message
- `GET /api/health` - Health check
- `GET /api/hello?name=World` - Hello endpoint

### Impact Partners

- `GET /api/impact-partners` - List all partners
- `GET /api/impact-partners/{id}` - Get partner by ID

### Impact Projects

- `GET /api/impact-projects` - List all projects
- `GET /api/impact-projects/{id}` - Get project by ID
- `GET /api/impact-projects?partnerId={id}` - List projects by partner

## Testing

The application includes comprehensive tests:

### Integration Tests (`cmd/api/main_test.go`)

Tests that verify shared packages work correctly in the application:

- `TestCalculatorIntegration` - Verifies calculator package functionality
- `TestCalculatorBatch` - Verifies batch calculation operations
- `TestLoggerIntegration` - Verifies logger package functionality
- `TestPtrFloat64` - Tests helper function
- `TestSharedPackagesAccessibility` - Verifies packages can be imported

### Package Tests

- `packages/go/common/calculator/calculator_test.go` - Calculator package unit tests
- `packages/go/common/logger/logger_test.go` - Logger package unit tests

All tests include:
- Unit tests for core functionality
- Edge case testing (zero values, negative values, etc.)
- Benchmark tests for performance monitoring

## Architecture

```
apps/backend/api-golang/
├── cmd/
│   └── api/
│       ├── main.go          # Application entry point
│       └── main_test.go     # Integration tests
├── internal/
│   ├── impact_partner/      # Partner domain logic
│   └── impact_project/      # Project domain logic
├── go.mod                   # Go module definition
└── moon.yml                 # Moon task configuration
```

## Dependencies

- Standard library (net/http, encoding/json, etc.)
- `github.com/bilo-mono/packages/common` - Shared monorepo packages
