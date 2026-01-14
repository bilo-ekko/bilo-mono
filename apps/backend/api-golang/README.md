# api-golang

A Go API service implementing climate impact partners and projects management.

## 🏗️ Project Structure

**Feature-First Organization** - Each feature contains all its layers (entity, repository, service, controller).

```
api-golang/
├── cmd/
│   └── api/
│       └── main.go                 # Application entry point
├── internal/
│   ├── impact_partner/             # Impact Partner feature
│   │   ├── entity.go               # Domain model
│   │   ├── repository.go           # Data access layer
│   │   ├── service.go              # Business logic
│   │   └── controller.go           # HTTP handlers
│   └── impact_project/             # Impact Project feature
│       ├── entity.go               # Domain model
│       ├── repository.go           # Data access layer
│       ├── service.go              # Business logic
│       └── controller.go           # HTTP handlers
├── bin/
│   └── api-golang                  # Compiled binary (generated)
├── go.mod
└── moon.yml
```

## 📦 Domain Models

### ImpactPartner
Represents an organization providing climate impact services.

```go
type ImpactPartner struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Category string `json:"category"` // e.g., "carbon-offset", "reforestation", "renewable-energy"
}
```

### ImpactProject
Represents a specific climate impact project.

```go
type ImpactProject struct {
    ID        string `json:"id"`
    Name      string `json:"name"`
    Category  string `json:"category"`  // e.g., "solar", "wind", "forest-conservation"
    PartnerID string `json:"partnerId"` // Foreign key to ImpactPartner
}
```

## 🚀 Getting Started

### Run Development Server
```bash
moon run api-golang:dev
# or
go run ./cmd/api/main.go
```

### Build
```bash
moon run api-golang:build
# or
go build -o bin/api-golang ./cmd/api
```

### Run Binary
```bash
./bin/api-golang
```

## 📡 API Endpoints

**Server runs on:** `http://localhost:8080`

### General Endpoints

| Method | Endpoint           | Description            |
|--------|-------------------|------------------------|
| GET    | `/`               | Welcome message        |
| GET    | `/api/health`     | Health check           |
| GET    | `/api/hello?name=` | Hello endpoint        |

### Impact Partners

| Method | Endpoint                      | Description                    |
|--------|-------------------------------|--------------------------------|
| GET    | `/api/impact-partners`        | Get all impact partners        |
| GET    | `/api/impact-partners/{id}`   | Get specific partner by ID     |

**Sample Response:**
```json
[
  {
    "id": "1",
    "name": "GoldStandard",
    "category": "carbon-offset"
  },
  {
    "id": "2",
    "name": "Ekko Climate",
    "category": "reforestation"
  }
]
```

### Impact Projects

| Method | Endpoint                            | Description                      |
|--------|-------------------------------------|----------------------------------|
| GET    | `/api/impact-projects`              | Get all impact projects          |
| GET    | `/api/impact-projects/{id}`         | Get specific project by ID       |
| GET    | `/api/impact-projects?partnerId={id}` | Get projects by partner ID    |

**Sample Response:**
```json
[
  {
    "id": "1",
    "name": "Amazon Rainforest Conservation",
    "category": "forest-conservation",
    "partnerId": "2"
  },
  {
    "id": "2",
    "name": "Solar Farm Initiative India",
    "category": "solar",
    "partnerId": "3"
  }
]
```

## 🧪 Testing

```bash
# Run all tests
moon run api-golang:test
# or
go test ./...

# Run with verbose output
go test -v ./...
```

## 🔧 Other Commands

```bash
# Format code
moon run api-golang:format
# or
go fmt ./...

# Lint code
moon run api-golang:lint
# or
go vet ./...

# Tidy dependencies
moon run api-golang:tidy
# or
go mod tidy
```

## 📝 Sample Data

The application comes pre-seeded with sample data:

**Partners:**
- GoldStandard (carbon-offset)
- Ekko Climate (reforestation)
- Green Energy Co (renewable-energy)

**Projects:**
- Amazon Rainforest Conservation
- Solar Farm Initiative India
- Wind Energy Project Denmark
- Mangrove Restoration Program

## 🏛️ Architecture

This project follows a **feature-first** clean architecture pattern:

### Feature Organization
Each feature module (`impact_partner`, `impact_project`) contains:

1. **Entity** (`entity.go`) - Core domain model
2. **Repository** (`repository.go`) - Data access layer (currently in-memory, can be swapped with DB)
3. **Service** (`service.go`) - Business logic layer
4. **Controller** (`controller.go`) - HTTP handlers and routing

### Benefits
- **High Cohesion**: Related code stays together
- **Easy Navigation**: All partner-related code is in one place
- **Scalability**: Easy to add new features without affecting others
- **Clear Boundaries**: Each feature is self-contained

The `internal/` directory follows Go conventions, making the packages private to this module.
