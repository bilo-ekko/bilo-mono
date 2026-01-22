# Adapters in Ports & Adapters Architecture

This document explains where adapters belong in the hexagonal architecture and provides concrete examples.

## Visual Overview

```
┌─────────────────────────────────────────────────────────────┐
│                    APPLICATION LAYER                        │
│                  (cmd/api/main.go)                          │
│                                                             │
│  ┌──────────────────────────────────────────────────────┐   │
│  │         DRIVING ADAPTERS (Inbound)                   │   │
│  │  ┌──────────────┐  ┌──────────────┐  ┌───────────┐   │   │
│  │  │ HTTP         │  │ gRPC         │  │ CLI       │   │   │
│  │  │ Controller   │  │ Handler      │  │ Command   │   │   │
│  │  └──────┬───────┘  └──────┬───────┘  └─────┬─────┘   │   │
│  └─────────┼──────────────────┼─────────────────┼───────┘   │
│            │                  │                 │           │
│            └──────────────────┼─────────────────┘           │
│                               │                             │
│            ┌──────────────────▼─────────────────┐           │
│            │      DOMAIN (Ports + Logic)         │          │
│            │  ┌──────────────────────────────┐  │           │
│            │  │  Orchestrator                 │  │          │
│            │  │  (Business Logic)             │  │          │
│            │  └───────────┬──────────────────┘  │          │
│            │              │                      │          │
│            │  ┌───────────▼──────────────────┐  │          │
│            │  │  Ports (Interfaces)          │  │          │
│            │  │  - Repository interface      │  │          │
│            │  │  - Service interface         │  │          │
│            │  └───────────┬──────────────────┘  │          │
│            └──────────────┼──────────────────────┘          │
│                           │                                 │
│  ┌────────────────────────▼────────────────────────────┐  │
│  │         DRIVEN ADAPTERS (Outbound)                    │  │
│  │  ┌──────────────┐  ┌──────────────┐  ┌───────────┐  │  │
│  │  │ PostgreSQL   │  │ HTTP Client  │  │ In-Memory │  │  │
│  │  │ Repository   │  │ (Microservice│  │ Repository│  │  │
│  │  └──────────────┘  │  Client)      │  └───────────┘  │  │
│  │                    └──────────────┘                  │  │
│  └──────────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘

Key Points:
- Domain only knows about Ports (interfaces)
- Adapters implement Ports
- Domain doesn't know which adapter is used
- Easy to swap adapters without changing domain code
```

## Current Structure

Currently, adapters are co-located with the domain:

```
internal/quote/
├── ports.go          # Ports (interfaces)
├── repository.go     # Adapter: InMemoryRepository (driven adapter)
├── controller.go     # Adapter: HTTP Controller (driving adapter)
├── orchestrator.go   # Domain logic
└── entity.go         # Domain entities
```

**This works**, but for better separation (especially for microservices), adapters should be in separate packages.

## Recommended Structure

### Option 1: Adapters in Separate Sub-packages (Recommended for Monolith)

```
internal/quote/
├── ports.go          # Ports (interfaces)
├── orchestrator.go   # Domain logic
└── entity.go         # Domain entities

internal/quote/adapters/
├── http/             # Driving adapters (inbound)
│   └── controller.go # HTTP controller adapter
└── persistence/       # Driven adapters (outbound)
    ├── inmemory/
    │   └── repository.go  # InMemoryRepository adapter
    └── postgres/
        └── repository.go  # PostgresRepository adapter
```

### Option 2: Adapters at Root Level (Better for Microservices)

```
internal/
├── quote/            # Domain (ports + logic only)
│   ├── ports.go
│   ├── orchestrator.go
│   └── entity.go
│
├── adapters/         # All adapters
│   ├── quote/
│   │   ├── http/
│   │   │   └── controller.go
│   │   └── persistence/
│   │       ├── inmemory/
│   │       └── postgres/
│   └── impact_partner/
│       └── http/
│           └── client.go  # HTTP client adapter
```

## Concrete Example: HTTP Client Adapter

Here's a concrete example of an **HTTP client adapter** that implements the `impact_partner.Service` interface. This would be used when the impact_partner domain is extracted to a microservice.

### Port (Interface)

**File**: `internal/impact_partner/ports.go`

```go
package impact_partner

import "context"

// Service defines the port for impact partner business logic
type Service interface {
	GetAllPartners(ctx context.Context) ([]*Entity, error)
	GetPartnerByID(ctx context.Context, id string) (*Entity, error)
	CreatePartner(ctx context.Context, partner *Entity) error
}
```

### Adapter: HTTP Client Implementation

**File**: `internal/adapters/impact_partner/http/client.go`

**See**: [`internal/adapters/impact_partner/http/client.go`](internal/adapters/impact_partner/http/client.go) for the full implementation.

```go
package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"api-golang/internal/impact_partner/impact_partner"
)

// HTTPClientAdapter implements impact_partner.Service interface
// This adapter makes HTTP calls to the impact_partner microservice
type HTTPClientAdapter struct {
	baseURL    string
	httpClient *http.Client
}

// NewHTTPClientAdapter creates a new HTTP client adapter
func NewHTTPClientAdapter(baseURL string) *HTTPClientAdapter {
	return &HTTPClientAdapter{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAllPartners implements impact_partner.Service interface
// Makes HTTP GET request to /api/impact-partners
func (a *HTTPClientAdapter) GetAllPartners(ctx context.Context) ([]*impact_partner.Entity, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", a.baseURL+"/api/impact-partners", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var partners []*impact_partner.Entity
	if err := json.NewDecoder(resp.Body).Decode(&partners); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return partners, nil
}

// GetPartnerByID implements impact_partner.Service interface
// Makes HTTP GET request to /api/impact-partners/{id}
func (a *HTTPClientAdapter) GetPartnerByID(ctx context.Context, id string) (*impact_partner.Entity, error) {
	url := fmt.Sprintf("%s/api/impact-partners/%s", a.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("partner not found")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var partner impact_partner.Entity
	if err := json.NewDecoder(resp.Body).Decode(&partner); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &partner, nil
}

// CreatePartner implements impact_partner.Service interface
// Makes HTTP POST request to /api/impact-partners
func (a *HTTPClientAdapter) CreatePartner(ctx context.Context, partner *impact_partner.Entity) error {
	body, err := json.Marshal(partner)
	if err != nil {
		return fmt.Errorf("failed to marshal partner: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/api/impact-partners", 
		bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
```

### Usage in Main Application

**File**: `examples/adapter_usage.go`

**See**: [`examples/adapter_usage.go`](examples/adapter_usage.go) for a complete example.

```go
package main

import (
	"os"
	
	"api-golang/internal/adapters/impact_partner/http"
	"api-golang/internal/impact_partner/impact_partner"
	"api-golang/internal/quote"
)

func main() {
	// Determine which adapter to use based on environment
	var partnerService impact_partner.Service
	
	if os.Getenv("IMPACT_PARTNER_SERVICE_URL") != "" {
		// Use HTTP client adapter (microservice mode)
		partnerService = http.NewHTTPClientAdapter(
			os.Getenv("IMPACT_PARTNER_SERVICE_URL"),
		)
	} else {
		// Use local service (monolith mode)
		partnerRepo := impact_partner.NewRepository()
		partnerService = impact_partner.NewService(partnerRepo)
	}

	// Orchestrator doesn't know or care which adapter is used!
	quoteOrchestrator := quote.NewOrchestrator(quote.OrchestratorDeps{
		// ... other dependencies
		ImpactPartnerService: partnerService, // Same interface!
		// ...
	})
	
	// ... rest of setup
}
```

### Adapter: PostgreSQL Repository Implementation

**File**: `internal/adapters/quote/persistence/postgres/repository.go`

**See**: [`internal/adapters/quote/persistence/postgres/repository.go`](internal/adapters/quote/persistence/postgres/repository.go) for the full implementation.

This shows another driven adapter - a PostgreSQL repository that implements the `quote.Repository` interface.

## Adapter Types

### 1. Driving Adapters (Inbound)
These adapters call INTO the domain:
- **HTTP Controllers**: `internal/quote/adapters/http/controller.go`
- **gRPC Handlers**: `internal/quote/adapters/grpc/handler.go`
- **CLI Commands**: `internal/quote/adapters/cli/command.go`
- **Message Queue Consumers**: `internal/quote/adapters/messaging/consumer.go`

### 2. Driven Adapters (Outbound)
These adapters are called BY the domain:
- **Repositories**: `internal/quote/adapters/persistence/postgres/repository.go`
- **External Service Clients**: `internal/adapters/impact_partner/http/client.go`
- **Email Services**: `internal/adapters/email/smtp/client.go`
- **File Storage**: `internal/adapters/storage/s3/client.go`

## Key Principles

1. **Domain doesn't know about adapters**: The domain only knows about ports (interfaces)
2. **Adapters implement ports**: Adapters satisfy the interface contracts
3. **Dependency injection**: Adapters are injected at application startup
4. **Easy to swap**: Change from `InMemoryRepository` to `PostgresRepository` without changing domain code
5. **Testability**: Mock adapters for testing domain logic

## Migration Path

### Phase 1: Current (Monolith)
```
internal/quote/
├── ports.go
├── repository.go      # InMemoryRepository adapter
└── controller.go      # HTTP controller adapter
```

### Phase 2: Refactored (Still Monolith)
```
internal/quote/
├── ports.go
└── entity.go

internal/quote/adapters/
├── http/controller.go
└── persistence/inmemory/repository.go
```

### Phase 3: Microservices
```
internal/quote/
├── ports.go
└── entity.go

internal/adapters/quote/
├── http/controller.go
└── persistence/postgres/repository.go

internal/adapters/impact_partner/
└── http/client.go  # Calls impact_partner microservice
```

## Benefits

1. **Clear separation**: Domain logic is separate from infrastructure concerns
2. **Easy testing**: Mock adapters for unit tests
3. **Flexibility**: Swap implementations without changing domain code
4. **Microservices ready**: HTTP client adapters can call remote services
5. **Multiple interfaces**: Same domain can have HTTP, gRPC, and CLI adapters
