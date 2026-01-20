# Quote Creation Architecture

This document describes the architectural patterns and design decisions used in the Go API for quote creation, focusing on the orchestration pattern and domain-driven design approach.

## Overview of Architectural Patterns

The quote creation system implements several architectural patterns working together:

> Resources
> - [DDD with VSA](https://medium.com/@evgeni.n.rusev/net-domain-driven-design-template-with-a-vertical-slice-architecture-33812c22b509)
> - []()

### 1. **Vertical Slice Architecture**
The codebase is organized by **feature/domain** rather than by technical layers. Each domain (quote, organisation, impact, finance, etc.) contains all its layers (entity, repository, service, controller) within a single directory. This approach:
- Keeps related code together
- Makes features easier to understand and modify
- Reduces coupling between unrelated features
- Aligns with team ownership boundaries

### 2. **Domain-Driven Design (DDD)**
The system is organized into **bounded contexts** representing distinct business domains:
- **Quote Domain**: Quote creation and management
- **Organisation Domain**: Organisation and customer management
- **Impact Domain**: Carbon footprint calculation and fee management
- **Finance Domain**: Currency conversion
- **Funds Domain**: Sales tax calculation
- **Platform Domain**: Country lookups

Each domain encapsulates its own business logic and data models.

### 3. **Hexagonal Architecture (Ports & Adapters)**
Each domain defines **ports** (interfaces) for its dependencies:
- **Driving Ports**: Interfaces that the domain exposes (e.g., `quote.Service` interface in [`ports.go`](internal/quote/ports.go))
- **Driven Ports**: Repositories that the domain depends on (e.g., `quote.Repository` interface in [`ports.go`](internal/quote/ports.go))

The **adapters** (implementations) can be swapped without changing the domain logic:
- `quote.Repository` interface is implemented by [`repository.go`](internal/quote/repository.go) (InMemoryRepository)
- `quote.Service` interface exists but is **not implemented** - the `Orchestrator` is used directly instead
- Currently using in-memory repositories, but these can be replaced with database implementations

### 4. **Orchestration Pattern**
The `Orchestrator` coordinates complex workflows that span multiple domains. Instead of domains calling each other directly, the orchestrator:
- Coordinates the multi-step quote creation process
- Manages cross-domain dependencies
- Handles transaction-like workflows
- Provides a single entry point for complex operations

### 5. **Dependency Injection**
All dependencies are injected through constructors, making the system:
- Testable (easy to mock dependencies)
- Flexible (can swap implementations)
- Explicit (dependencies are clear)

## Directory Structure

```
apps/backend/api-golang/
├── cmd/
│   └── api/
│       ├── main.go              # Application entry point & dependency wiring
│       └── main_test.go         # Integration tests
│
├── internal/
│   ├── quote/                   # Quote Domain (Vertical Slice)
│   │   ├── entity.go           # Domain models (Entity, DTOs)
│   │   ├── ports.go            # Port interfaces (Repository, Service)
│   │   ├── repository.go       # Adapter: InMemoryRepository implements Repository port
│   │   ├── orchestrator.go     # Cross-domain workflow coordinator (used instead of Service)
│   │   ├── controller.go       # HTTP handlers (adapter)
│   │   └── orchestrator_test.go # Domain tests
│   │
│   ├── organisation/           # Organisation Domain
│   │   ├── organisation/
│   │   │   ├── entity.go
│   │   │   ├── ports.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── customer/
│   │       ├── entity.go
│   │       ├── ports.go
│   │       ├── repository.go
│   │       └── service.go
│   │
│   ├── impact/                  # Impact Domain
│   │   ├── carbon_footprint/
│   │   │   ├── entity.go
│   │   │   ├── ports.go
│   │   │   ├── repository.go
│   │   │   └── service.go
│   │   └── fee/
│   │       ├── entity.go
│   │       ├── ports.go
│   │       ├── repository.go
│   │       └── service.go
│   │
│   ├── finance/                 # Finance Domain
│   │   └── currency/
│   │       ├── entity.go
│   │       ├── ports.go
│   │       ├── repository.go
│   │       └── service.go
│   │
│   ├── funds/                   # Funds Domain
│   │   └── salestax/
│   │       ├── entity.go
│   │       ├── ports.go
│   │       ├── repository.go
│   │       └── service.go
│   │
│   ├── platform/                # Platform Domain
│   │   └── country/
│   │       ├── entity.go
│   │       ├── ports.go
│   │       ├── repository.go
│   │       └── service.go
│   │
│   ├── impact_partner/          # Impact Partner Domain
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   ├── blended_price.go    # Specialized calculator
│   │   └── controller.go
│   │
│   ├── impact_project/          # Impact Project Domain
│   │   ├── entity.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── controller.go
│   │
│   └── shared/                  # Shared utilities
│       ├── errors/
│       │   └── errors.go        # Domain error types
│       └── types/
│           └── types.go         # Shared types (Money, Address, etc.)
│
├── go.mod
├── go.sum
├── moon.yml
└── README.md
```

## Key Components

### Quote Orchestrator

The orchestrator is the heart of the quote creation flow. It coordinates 11 steps across multiple domains:

**Location**: [`internal/quote/orchestrator.go`](internal/quote/orchestrator.go)

**Key Code Sample**:

```go
// Orchestrator coordinates the quote creation flow across multiple domains.
// This implements the Vertical Slice Architecture pattern - handling the entire
// request flow from validation through to quote creation.
type Orchestrator struct {
	// Organisation domain
	organisationService organisation.Service
	customerService     customer.Service

	// Platform domain
	countryService country.Service

	// Finance domain
	currencyService currency.Service

	// Impact domain
	carbonService carbonfootprint.Service
	feeService    fee.Service

	// Impact Partner domain
	blendedPriceCalc  *impact_partner.BlendedPriceCalculator
	impactPartnerRepo *impact_partner.Repository

	// Funds domain
	salesTaxService salestax.Service

	// Quote domain
	quoteRepo Repository
}
```

**Workflow Steps** (from [`orchestrator.go:100-110`](internal/quote/orchestrator.go#L100-L110)):

1. Validate Organisation
2. Get or Create Customer
3. Calculate Carbon Footprint
4. Get Blended Project Unit Price
5. Calculate Compensation Amount (Impact Amount)
6. Calculate Round Up (disabled in new API)
7. Calculate Service Fee
8. Calculate Sales Tax
9. Calculate Totals and Build Response
10. Write Quote Entity
11. Build Response

### Quote Controller

The controller handles HTTP requests and delegates to the orchestrator:

**Location**: [`internal/quote/controller.go`](internal/quote/controller.go)

**Key Code Sample**:

```go
// HandleCreateQuote handles POST /api/quotes
func (c *Controller) HandleCreateQuote(w http.ResponseWriter, r *http.Request) {
	// Parse and validate request
	var req CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.writeError(w, http.StatusBadRequest, "INVALID_REQUEST", "Invalid request body: "+err.Error())
		return
	}

	// Get organisation ID from header (simulating auth)
	headerOrgID := r.Header.Get("X-Organisation-ID")
	
	// Validate required fields
	if req.Customer.Reference == "" {
		c.writeError(w, http.StatusBadRequest, "MISSING_FIELD", "customer.reference is required")
		return
	}

	// Delegate to orchestrator
	ctx := r.Context()
	response, err := c.orchestrator.CreateQuote(ctx, &req, headerOrgID)
	if err != nil {
		// Map domain errors to HTTP status codes
		c.writeError(w, http.StatusNotFound, "NOT_FOUND", err.Error())
		return
	}

	c.writeJSON(w, http.StatusCreated, response)
}
```

### Quote Entity

The domain model representing a quote:

**Location**: [`internal/quote/entity.go`](internal/quote/entity.go)

**Key Code Sample**:

```go
// Entity represents a stored carbon offset quote (matches Notion data model)
type Entity struct {
	// Primary identifiers
	ID                   string `json:"id"`
	QuoteReference       string `json:"quoteReference"`       // Client-facing reference
	CalculationReference string `json:"calculationReference"` // Links to calculation

	// Organisation and customer
	OrganisationID string `json:"organisationId"`
	CustomerID     string `json:"customerId"`

	// Currency
	Currency string `json:"currency"` // ISO-3 currency code

	// Carbon credit amounts (stored in quote currency)
	CarbonCreditTotal              float64 `json:"carbonCreditTotal"`
	CarbonCreditImpact             float64 `json:"carbonCreditImpact"`
	CarbonCreditImpactSalesTax     float64 `json:"carbonCreditImpactSalesTax"`
	// ... more fields

	// Contribution details (stored as JSON blob)
	ContributionDetails ContributionDetails `json:"contributionDetails"`

	// Metadata
	Status    Status    `json:"status"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
```

### Ports (Interfaces)

Domain interfaces define contracts without implementation details:

**Location**: [`internal/quote/ports.go`](internal/quote/ports.go)

**Key Code Sample**:

```go
// Repository defines the port for quote data access
type Repository interface {
	Create(ctx context.Context, quote *Entity) error
	GetByID(ctx context.Context, id string) (*Entity, error)
	Update(ctx context.Context, quote *Entity) error
}

// Service defines the port for quote business logic
// NOTE: This interface exists but is not currently implemented.
// The Orchestrator is used directly instead of implementing this interface.
type Service interface {
	CreateQuote(ctx context.Context, req *CreateQuoteRequest) (*CreateQuoteResponse, error)
	GetQuote(ctx context.Context, id string) (*Entity, error)
}
```

**Note**: The `Service` interface is defined but not implemented. The `Orchestrator` provides the same functionality and is used directly by the controller. This is a design decision - the orchestrator pattern is preferred over a traditional service layer for complex cross-domain workflows.

### Repository Adapter

In-memory implementation of the `Repository` port interface:

**Interface**: [`internal/quote/ports.go`](internal/quote/ports.go) - `Repository` interface  
**Implementation**: [`internal/quote/repository.go`](internal/quote/repository.go) - `InMemoryRepository`

**Key Code Sample**:

```go
// InMemoryRepository implements the Repository port interface
type InMemoryRepository struct {
	quotes map[string]*Entity
	mu     sync.RWMutex
}

func (r *InMemoryRepository) Create(_ context.Context, quote *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.quotes[quote.ID]; exists {
		return errors.NewValidationError(domainName, "quote already exists")
	}

	r.quotes[quote.ID] = quote
	return nil
}
```

**Note**: The `Repository` interface is defined in [`ports.go`](internal/quote/ports.go) and implemented by `InMemoryRepository` in [`repository.go`](internal/quote/repository.go). This follows the Ports & Adapters pattern - the interface can be swapped with a database-backed implementation without changing the domain logic.

### Dependency Wiring

All dependencies are wired together in the main application:

**Location**: [`cmd/api/main.go`](cmd/api/main.go)

**Key Code Sample**:

```go
// Quote domain - Orchestrator
quoteRepo := quote.NewInMemoryRepository()
quoteOrchestrator := quote.NewOrchestrator(quote.OrchestratorDeps{
	OrganisationService: orgService,
	CustomerService:     customerService,
	CountryService:      countryService,
	CurrencyService:     currencyService,
	CarbonService:       carbonService,
	FeeService:          feeService,
	BlendedPriceCalc:    blendedPriceCalc,
	ImpactPartnerRepo:   partnerRepo,
	SalesTaxService:     salesTaxService,
	QuoteRepo:           quoteRepo,
})
quoteController := quote.NewController(quoteOrchestrator)
```

### Error Handling

Domain-specific errors with proper context:

**Location**: [`internal/shared/errors/errors.go`](internal/shared/errors/errors.go)

**Key Code Sample**:

```go
// DomainError represents a domain-specific error
type DomainError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Cause   error  `json:"-"`
}

func NewNotFoundError(domain, message string) *DomainError {
	return &DomainError{
		Code:    ErrCodeNotFound,
		Message: message,
		Domain:  domain,
	}
}
```

## Request Flow

```
HTTP Request (POST /api/quotes)
    ↓
Controller (controller.go)
    ├─ Parse JSON
    ├─ Validate required fields
    └─ Extract headerOrgID
    ↓
Orchestrator (orchestrator.go)
    ├─ Step 1: Validate Organisation → organisation.Service
    ├─ Step 2: Get/Create Customer → customer.Service
    ├─ Step 3: Calculate Carbon → carbonfootprint.Service
    ├─ Step 4: Get Blended Price → impact_partner.BlendedPriceCalculator
    ├─ Step 5: Calculate Impact Amount
    ├─ Step 6: Calculate Service Fee → fee.Service
    ├─ Step 7: Calculate Sales Tax → salestax.Service
    ├─ Step 8: Build Response
    ├─ Step 9: Write Quote → quote.Repository (InMemoryRepository in repository.go)
    └─ Step 10: Return CreateQuoteResponse
    ↓
Controller
    ├─ Map domain errors to HTTP status codes
    └─ Write JSON response
    ↓
HTTP Response (201 Created)
```

## Strengths

### ✅ **Separation of Concerns**
- Each domain is self-contained with clear boundaries
- Business logic is separated from HTTP handling
- Data access is abstracted through ports (`Repository` interface in `ports.go`, implemented in `repository.go`)
- Orchestrator handles cross-domain coordination (no separate Service layer needed)

### ✅ **Testability**
- Dependencies are injected, making mocking straightforward
- Each domain can be tested in isolation
- In-memory repositories enable fast unit tests
- See [`orchestrator_test.go`](internal/quote/orchestrator_test.go) for examples

### ✅ **Maintainability**
- Vertical slices make features easy to locate and modify
- Changes to one domain don't affect others
- Clear ownership boundaries

### ✅ **Flexibility**
- Repository implementations can be swapped (in-memory → database)
  - `Repository` interface defined in [`ports.go`](internal/quote/ports.go)
  - `InMemoryRepository` implementation in [`repository.go`](internal/quote/repository.go)
  - Can be replaced with `PostgresRepository` or `MySQLRepository` without changing domain logic
- Orchestrator can be extended or refactored without affecting controllers
- New domains can be added without modifying existing code

### ✅ **Explicit Dependencies**
- All dependencies are visible in constructors
- No hidden global state or singletons
- Easy to understand what each component needs

### ✅ **Domain-Driven Design**
- Code structure reflects business domains
- Business logic is centralized in orchestrator (no separate Service layer - orchestrator handles coordination)
- Entities represent business concepts clearly

## Weaknesses

### ⚠️ **Orchestrator Complexity**
- The orchestrator has many dependencies (9 services)
- Large method (~530 lines) handling multiple concerns
- Could benefit from breaking into smaller, focused orchestrators
- **Note**: The `Service` interface exists in [`ports.go`](internal/quote/ports.go) but is not implemented - the orchestrator is used directly instead

**Mitigation**: Consider splitting into:
- `QuoteCreationOrchestrator` (core flow)
- `QuoteCalculationOrchestrator` (pricing calculations)
- `QuoteResponseBuilder` (response assembly)
- Alternatively, implement the `Service` interface and have it delegate to the orchestrator for better abstraction

### ⚠️ **Error Handling**
- Error mapping in controller uses string matching (`strings.Contains`)
- Could miss edge cases or be fragile to error message changes
- No structured error types for cross-domain errors

**Mitigation**: Consider:
- Domain-specific error types with error codes
- Error wrapping with context
- Centralized error mapping

### ⚠️ **Transaction Management**
- No explicit transaction boundaries
- If step 9 (Write Quote) fails, previous steps aren't rolled back
- Could lead to inconsistent state

**Mitigation**: Consider:
- Saga pattern for distributed transactions
- Compensating actions for rollback
- Event sourcing for auditability

### ⚠️ **In-Memory Storage**
- Current repositories are in-memory only
- Data is lost on restart
- Not suitable for production

**Mitigation**: This is intentional for MVP/demo purposes. For production:
- Implement database-backed repositories
- Add connection pooling
- Add migration support

### ⚠️ **Validation Split**
- Some validation in controller (required fields)
- Some validation in orchestrator (business rules)
- Could be inconsistent

**Mitigation**: Consider:
- Moving all validation to orchestrator
- Using a validation framework
- Creating a validation middleware

### ⚠️ **Testing Coverage**
- Integration tests are minimal
- No end-to-end tests
- Error scenarios could be better covered

**Mitigation**: Add:
- Integration tests for full workflows
- E2E tests with test containers
- Error scenario tests

## Future Improvements

1. **Split Orchestrator**: Break into smaller, focused components
2. **Add Database Layer**: Implement SQL repositories
3. **Add Caching**: Cache frequently accessed data (countries, organisations)
4. **Add Observability**: Structured logging, metrics, tracing
5. **Add Event System**: Publish domain events for quote creation
6. **Add Validation Framework**: Centralized, reusable validation
7. **Add API Versioning**: Support multiple API versions
8. **Add Rate Limiting**: Protect against abuse
9. **Add Retry Logic**: Handle transient failures
10. **Add Circuit Breaker**: Prevent cascading failures

## Related Documentation

- [README.md](README.md) - Setup and usage instructions
- [.structure.md](.structure.md) - Directory structure details
- [internal/quote/orchestrator_test.go](internal/quote/orchestrator_test.go) - Test examples
/