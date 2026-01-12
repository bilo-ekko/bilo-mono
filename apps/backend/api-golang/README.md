# api-golang

A basic Golang REST API project.

## Features

- Simple HTTP server with JSON responses
- Health check endpoint
- Example API endpoints

## Prerequisites

- Go 1.16 or higher

## Getting Started

### Install Dependencies

```bash
go mod tidy
```

### Run the Application

```bash
go run main.go
```

The server will start on `http://localhost:8080`

### Build the Application

```bash
go build -o api-golang
./api-golang
```

## API Endpoints

### Root
- **GET** `/`
  - Returns a welcome message

### Health Check
- **GET** `/api/health`
  - Returns the health status of the API

### Hello
- **GET** `/api/hello?name=YourName`
  - Returns a personalized greeting
  - Query parameter `name` is optional (defaults to "World")

## Project Structure

```
api-golang/
├── main.go          # Main application entry point
├── go.mod           # Go module dependencies
└── README.md        # This file
```

## Example Requests

```bash
# Root endpoint
curl http://localhost:8080/

# Health check
curl http://localhost:8080/api/health

# Hello endpoint
curl http://localhost:8080/api/hello?name=Bilo
```

## Development

To add more dependencies:

```bash
go get <package-name>
go mod tidy
```
