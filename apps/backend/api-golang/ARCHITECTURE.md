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

**Cross-Domain Communication**: The orchestrator communicates with external domains through **Service interfaces**, not repositories:
- Example: [`impact_partner.Service`](internal/impact_partner/ports.go) interface is used by the orchestrator
- Service implementations (e.g., `DefaultService`) can be replaced with HTTP/gRPC clients when extracting to microservices
- This pattern prepares the architecture for microservices extraction - inter-service communication can be abstracted behind service interfaces

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
	blendedPriceCalc      *impact_partner.BlendedPriceCalculator
	impactPartnerService  impact_partner.Service

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
	BlendedPriceCalc:      blendedPriceCalc,
	ImpactPartnerService:   partnerService,
	SalesTaxService:        salesTaxService,
	QuoteRepo:              quoteRepo,
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
- **Service interfaces enable microservices extraction**:
  - Orchestrator communicates with external domains through `Service` interfaces (e.g., `impact_partner.Service`)
  - Service implementations can be swapped with HTTP/gRPC clients when extracting to microservices
  - Example: [`impact_partner.Service`](internal/impact_partner/ports.go) interface allows replacing `DefaultService` with a remote client
  - Inter-service communication layer can be abstracted behind service interfaces
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
- **Note**: The orchestrator communicates with external domains through `Service` interfaces (e.g., `impact_partner.Service`), not repositories, which prepares it for microservices extraction

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

## Architectural Questions & Answers

This section addresses key architectural concerns and design decisions for the quote creation system.

### **Can any steps run in parallel?**

**Answer**: Currently, most steps run sequentially due to data dependencies. However, some steps could potentially run in parallel:

- **Parallelizable**: Fetching organisation data and country data (if both are needed independently), or fetching multiple impact partners/projects
- **Sequential**: Customer creation depends on organisation validation, carbon calculation depends on customer data, and quote creation depends on all previous calculations

The orchestrator pattern makes it straightforward to identify dependencies. If we wanted to parallelize, we could:
- Use Go's `sync.WaitGroup` and goroutines for independent operations
- Refactor the orchestrator to separate independent steps from dependent ones
- Consider using a workflow engine for complex parallel execution

**Current State**: All steps run sequentially in [`orchestrator.go`](internal/quote/orchestrator.go) for simplicity and correctness.

### **What happens if an intermediate step fails?**

**Answer**: Errors bubble up through the call stack and are returned to the controller, which maps them to HTTP status codes. However, there's **no rollback mechanism** for partial operations:

- If customer creation succeeds but carbon calculation fails, the customer remains created
- If quote calculation succeeds but quote storage fails, the calculation is lost
- No compensating actions are taken for completed steps

**Current Error Flow**:
1. Service/domain returns a domain error (e.g., `NotFoundError`, `ValidationError`)
2. Orchestrator propagates the error
3. Controller maps domain errors to HTTP status codes (see [`controller.go`](internal/quote/controller.go))
4. HTTP response is sent with error details

**Future Consideration**: Implement a Saga pattern or compensating transactions for multi-step operations that need atomicity.

### **How do you handle partial rollback?**

**Answer**: Currently, **partial rollback is not handled**. This is identified as a weakness in the architecture (see [Transaction Management](#transaction-management) section).

**Options for Future Implementation**:

1. **Saga Pattern**: Each step has a compensating action that undoes its work
   - Example: If quote creation fails after customer creation, delete the customer
   - Pros: Works across distributed systems
   - Cons: Complex to implement and test

2. **Database Transactions**: If all operations use the same database, wrap them in a transaction
   - Pros: Simple, atomic
   - Cons: Only works for single-database operations

3. **Event Sourcing**: Store events instead of state, allowing replay/rollback
   - Pros: Full audit trail, can replay to any point
   - Cons: Significant architectural change

4. **Two-Phase Commit**: For distributed systems requiring strong consistency
   - Pros: Guarantees atomicity
   - Cons: Complex, can block on failures

**Current State**: Accepts eventual consistency - if a step fails, previous steps remain completed. This is acceptable for MVP but should be addressed for production.

### **Who owns the schema migrations?**

**Answer**: Schema ownership follows domain boundaries:

- **Domain-specific schemas**: Each domain owns its entity structure (defined in `entity.go` files)
  - Quote domain owns quote table/collection schema
  - Organisation domain owns organisation and customer schemas
  - Impact domain owns impact partner and project schemas

- **Shared schemas**: Located in [`internal/shared/types/types.go`](internal/shared/types/types.go)
  - `Address`, `Money`, `BillingConfig`, etc.
  - Changes require coordination across domains

- **Migration Strategy**: Not yet implemented (using in-memory storage). When database-backed repositories are added:
  - Each domain should own its migration files
  - Shared types migrations should be coordinated through a shared migration directory
  - Consider using a migration tool (e.g., `golang-migrate`, `atlas`) with domain-specific folders

**Recommendation**: Create a `migrations/` directory structure mirroring the domain structure:
```
migrations/
├── quote/
│   └── 001_create_quotes.sql
├── organisation/
│   ├── 001_create_organisations.sql
│   └── 002_create_customers.sql
└── shared/
    └── 001_create_types.sql
```

### **Can you deploy domain A without domain B?**

**Answer**: **Yes**, due to the vertical slice architecture and dependency injection:

- Each domain is self-contained in its own directory
- Dependencies are injected through interfaces (ports)
- No compile-time coupling between domains (only through interfaces)
- Domains communicate through the orchestrator, not directly

**Example**: You could deploy a new version of the `organisation` domain without redeploying `impact` or `finance` domains, as long as:
- The interfaces (`ports.go`) remain compatible
- The orchestrator is updated to handle any interface changes
- Shared types in [`internal/shared/types/types.go`](internal/shared/types/types.go) remain backward compatible

**Current Limitation**: All domains are in a single binary (`cmd/api/main.go`), so they deploy together. To enable independent deployment:
- Extract domains into separate services/microservices
- Use gRPC or HTTP APIs for inter-domain communication
- Implement API versioning for interface changes

### **How do you handle schema changes that affect multiple domains?**

**Answer**: Through **shared types** and **interface versioning**:

1. **Shared Types**: Common structures live in [`internal/shared/types/types.go`](internal/shared/types/types.go)
   - Examples: `Address`, `Money`, `BillingConfig`
   - Changes require coordination and potentially breaking changes

2. **Interface Contracts**: Domain interfaces (`ports.go`) define contracts
   - Changes to interfaces require updating both implementer and consumer
   - The orchestrator acts as the coordinator, so interface changes propagate through it

3. **Versioning Strategy**: Not yet implemented, but recommended:
   - **Backward Compatible Changes**: Add optional fields, extend enums
   - **Breaking Changes**: Create new versions of interfaces/types
   - **Deprecation**: Mark old versions as deprecated, migrate gradually

**Current Process**: 
- Changes to shared types require updating all consuming domains
- Interface changes require updating implementers and the orchestrator
- No formal versioning or deprecation process yet

**Recommendation**: When extracting to microservices, use API versioning (e.g., `/v1/quotes`, `/v2/quotes`) to allow gradual migration.

### **How many files does a developer need to touch?**

**Answer**: Depends on the scope of the change:

**Simple Domain Change** (e.g., add a field to ImpactPartner):
- 1-2 files: `impact_partner/entity.go`, possibly `impact_partner/repository.go`

**Cross-Domain Change** (e.g., add customer location filtering):
- 3-5 files: 
  - `quote/entity.go` (request DTO)
  - `quote/orchestrator.go` (business logic)
  - `customer/entity.go` (if customer schema changes)
  - `quote/controller.go` (if validation changes)
  - `quote/orchestrator_test.go` (tests)

**Shared Type Change** (e.g., modify Address structure):
- 5+ files: All domains using the shared type
  - `shared/types/types.go` (definition)
  - `organisation/organisation/entity.go` (usage)
  - `quote/entity.go` (usage in merchant address)
  - All repositories/services using Address
  - Tests

**New Feature** (e.g., add round-up calculation):
- 3-6 files:
  - `quote/entity.go` (request/response DTOs)
  - `quote/orchestrator.go` (calculation logic)
  - `quote/controller.go` (validation)
  - `quote/orchestrator_test.go` (tests)
  - Possibly `funds/` domain if it involves fees/taxes

**Average**: For a typical feature change, expect to touch **3-5 files**. The vertical slice architecture keeps related changes together, reducing the number of files compared to a layered architecture.

### **How easy is it for a new developer to understand the flow?**

**Answer**: **Moderately easy** - the architecture provides clear entry points but the orchestrator is complex:

**Strengths**:
- **Clear Entry Point**: Start at [`controller.go`](internal/quote/controller.go) → `HandleCreateQuote`
- **Single Orchestration Flow**: All logic flows through [`orchestrator.go`](internal/quote/orchestrator.go)
- **Domain Boundaries**: Easy to understand what each domain does
- **Self-Contained**: Each domain has all its code in one place

**Challenges**:
- **Large Orchestrator**: ~530 lines handling 11 steps (see [Orchestrator Complexity](#orchestrator-complexity))
- **Many Dependencies**: Orchestrator depends on 9 different services
- **No Visual Documentation**: Flow is only in code comments

**Onboarding Path**:
1. Read [`ARCHITECTURE.md`](ARCHITECTURE.md) (this file)
2. Trace a request: `controller.go` → `orchestrator.go` → services
3. Understand domain boundaries by reading `entity.go` files
4. Read tests in `orchestrator_test.go` for examples

**Recommendation**: 
- Add sequence diagrams for complex flows
- Split orchestrator into smaller, named functions
- Add more inline documentation explaining "why" not just "what"

### **Can you trace a request through the system?**

**Answer**: **Yes**, but tracing is currently manual (no distributed tracing yet):

**Request Flow** (see [Request Flow](#request-flow) section):
```
HTTP Request → Controller → Orchestrator → Domain Services → Repository → Response
```

**Tracing Points**:
1. **Controller**: [`controller.go:207`](internal/quote/controller.go#L207) - `HandleCreateQuote`
2. **Orchestrator**: [`orchestrator.go:100`](internal/quote/orchestrator.go#L100) - `CreateQuote`
3. **Domain Services**: Each step calls a service (e.g., `organisationService.GetByID`)
4. **Repository**: Services call repositories (e.g., `quoteRepo.Create`)

**Current State**: 
- No request IDs or correlation IDs
- No structured logging with context
- No distributed tracing (OpenTelemetry, Jaeger, etc.)

**Future Enhancement**: Add:
- Request ID middleware to generate and propagate IDs
- Structured logging with request context
- OpenTelemetry instrumentation for distributed tracing
- Log aggregation (e.g., ELK stack, Datadog)

**Manual Tracing**: To trace a request manually:
1. Add log statements at each step
2. Use Go's `context.Context` to pass request metadata
3. Check logs for the request path

### **If you change the footprint calculation, how many places break?**

**Answer**: **Minimal impact** - changes are isolated to the carbon footprint domain:

**Affected Files**:
1. [`internal/impact/carbon_footprint/service.go`](internal/impact/carbon_footprint/service.go) - Calculation logic
2. [`internal/impact/carbon_footprint/entity.go`](internal/impact/carbon_footprint/entity.go) - If response structure changes
3. [`internal/quote/orchestrator.go`](internal/quote/orchestrator.go) - If orchestrator needs to pass different parameters or handle different response

**Isolated Impact**: 
- The carbon footprint service exposes an interface (`carbonfootprint.Service`)
- Other domains depend on the interface, not the implementation
- Changes to calculation logic don't break consumers unless the interface changes

**Example Scenarios**:
- **Change calculation algorithm**: Only `carbon_footprint/service.go` changes
- **Change response structure**: `carbon_footprint/entity.go` + `orchestrator.go` (if it uses the response)
- **Change input parameters**: `carbon_footprint/ports.go` (interface) + `orchestrator.go` (caller)

**Domain Isolation**: This demonstrates the strength of domain boundaries - changes to one domain's internals don't cascade to others.

### **Who owns each domain?**

**Answer**: Domain ownership is **not explicitly defined** in the codebase, but domains are structured for clear ownership:

**Domain Structure** (see [Directory Structure](#directory-structure)):
- `quote/` - Quote creation and management
- `organisation/` - Organisation and customer management
- `impact/` - Carbon footprint and fee calculation
- `finance/` - Currency conversion
- `funds/` - Sales tax calculation
- `platform/` - Country lookups
- `impact_partner/` - Impact partner management
- `impact_project/` - Impact project management

**Ownership Model** (Recommended):
- Each domain should have a **domain owner** (team or individual)
- Domain owner is responsible for:
  - Entity schema (`entity.go`)
  - Business logic (services, orchestrator steps)
  - Repository implementation
  - Domain tests
- **Shared types** (`shared/types/`) require coordination across owners

**Current State**: All domains are in a single repository, so ownership is implicit. For microservices:
- Each domain could be a separate service/repository
- Clear ownership boundaries through service boundaries
- API contracts define interaction points

**Recommendation**: Document domain ownership in a `DOMAIN_OWNERSHIP.md` file or in team documentation.

### **How do teams coordinate changes?**

**Answer**: Through **interfaces**, **shared types**, and **orchestrator coordination**:

**Coordination Mechanisms**:

1. **Interface Contracts** (`ports.go` files):
   - Domains define `Service` interfaces for cross-domain communication
   - The orchestrator communicates with external domains through service interfaces (e.g., `impact_partner.Service`), not repositories
   - Changes to service interfaces require coordination
   - Example: If `impact_partner.Service` interface changes, `quote` domain (orchestrator) must be updated
   - This pattern enables microservices extraction - service implementations can be swapped with HTTP/gRPC clients

2. **Shared Types** ([`internal/shared/types/types.go`](internal/shared/types/types.go)):
   - Common structures used across domains
   - Changes require cross-team coordination
   - Example: Changing `Address` structure affects `organisation` and `quote` domains

3. **Orchestrator as Coordinator**:
   - The orchestrator (`quote/orchestrator.go`) coordinates cross-domain workflows
   - Changes to orchestration logic may require input from multiple domain owners
   - Example: Adding a new step that calls `finance` domain requires finance team input

**Current Process** (Implicit):
- Code review for interface changes
- Discussion for shared type changes
- Orchestrator changes reviewed by affected domain owners

**Recommended Process**:
- **RFC/ADR Process**: Document significant changes (Architecture Decision Records)
- **Interface Versioning**: Version interfaces to allow gradual migration
- **Shared Type Governance**: Establish a process for shared type changes
- **Orchestrator Changes**: Require approval from affected domain owners

**Future (Microservices)**:
- API contracts define coordination points (already established through `Service` interfaces)
- Service interfaces (e.g., `impact_partner.Service`) can be replaced with HTTP/gRPC clients
- Inter-service communication layer can be abstracted behind service interfaces
- API versioning allows gradual migration
- Service mesh for inter-service communication
- Contract testing (Pact, etc.) to ensure compatibility

### **When we extract to microservices, are we comfortable with eventual consistency (data replication) or do we need strong consistency (RPC)?**

**Answer**: **Depends on the operation** - quote creation requires **strong consistency**, while some read operations could use eventual consistency:

**Strong Consistency Required**:
- **Quote Creation**: All steps must complete atomically or fail together
  - Customer creation → Carbon calculation → Quote storage
  - Cannot have partial quotes or inconsistent state
  - **Solution**: Use RPC calls with transaction coordination (Saga pattern, two-phase commit)

- **Financial Calculations**: Amounts must be accurate and consistent
  - Currency conversion, tax calculation, fee calculation
  - **Solution**: Synchronous RPC calls, validate totals

**Eventual Consistency Acceptable**:
- **Read Operations**: Organisation data, impact partner/project listings
  - Can tolerate slight staleness (cache invalidation delays)
  - **Solution**: Caching, read replicas, eventual consistency

- **Analytics/Reporting**: Historical quote data, aggregated statistics
  - Can be computed asynchronously
  - **Solution**: Event streaming, data replication, batch processing

**Hybrid Approach** (Recommended):
- **Write Path**: Strong consistency via RPC
  - Quote creation: Synchronous calls, transaction coordination
  - Customer updates: Synchronous, immediate consistency
- **Read Path**: Eventual consistency where acceptable
  - Organisation listings: Cached, eventual consistency
  - Impact projects: Cached, eventual consistency
- **Async Processing**: For non-critical operations
  - Email notifications, analytics updates, audit logs

**Current State**: All operations are synchronous (single binary), so consistency is guaranteed. When extracting to microservices:
- Use **synchronous RPC** (gRPC, HTTP) for quote creation flow
- Use **eventual consistency** (event streaming, message queues) for non-critical operations
- Implement **Saga pattern** for distributed transactions in quote creation

**Trade-offs**:
- **Strong Consistency**: Higher latency, more complex error handling, potential for blocking
- **Eventual Consistency**: Lower latency, simpler error handling, but potential for temporary inconsistencies

## Architecture Ratings

The following table provides ratings (out of 5.0) for key architectural dimensions:

| Dimension | Rating | Justification |
|-----------|--------|---------------|
| **Code Maintainability** | 3.8/5.0 | Vertical slice architecture keeps related code together, making features easy to locate and modify. However, the orchestrator (~530 lines) is large and handles multiple concerns, making it harder to maintain. Domain boundaries are clear, reducing coupling. |
| **Schema Maintainability** | 3.2/5.0 | Domains own their schemas (good), but shared types require cross-domain coordination. No migration strategy or versioning yet. Schema changes can cascade through multiple domains. |
| **Implementation Complexity** | 3.5/5.0 | Medium complexity. The orchestrator pattern is conceptually clear but requires managing many dependencies (9 services). Domain structure is straightforward. Dependency injection adds some boilerplate but improves testability. |
| **Loose Coupling** | 4.2/5.0 | Excellent - domains communicate only through interfaces (ports), no direct dependencies. Dependency injection enables swapping implementations. Orchestrator is the only coupling point, but it's intentional and explicit. |
| **Logical Separation of Concerns** | 4.5/5.0 | Excellent - vertical slices ensure each domain is self-contained with clear boundaries. Business logic separated from HTTP handling. Data access abstracted through ports. Orchestrator handles cross-domain coordination cleanly. |
| **Infra Abstraction Fit** | 4.0/5.0 | Good - ports/adapters pattern allows swapping implementations (in-memory → database) without changing domain logic. Currently using in-memory repos, but can easily swap to database-backed implementations. |
| **Testability** | 4.0/5.0 | Good - dependency injection makes mocking straightforward. Each domain can be tested in isolation. In-memory repositories enable fast unit tests. However, integration tests are minimal and error scenarios could be better covered. |
| **Debugability** | 3.0/5.0 | Moderate - clear flow through orchestrator makes tracing possible, but large orchestrator makes debugging complex operations harder. No request IDs, structured logging, or distributed tracing yet. Error messages could be more contextual. |
| **Ability to Isolate to Microservices** | 4.5/5.0 | Excellent - domains are already separated with clear boundaries. Communication through `Service` interfaces (not repositories) makes extraction straightforward. The orchestrator uses service interfaces (e.g., `impact_partner.Service`), which can be replaced with HTTP/gRPC clients. Each domain could become a microservice with minimal refactoring. Orchestrator could become an API gateway or workflow engine. |
| **Performance** | 3.0/5.0 | Not yet optimized - in-memory storage is fast but not persistent. No caching layer for frequently accessed data (organisations, countries). No performance benchmarks or optimization yet. Acceptable for MVP but needs work for production scale. |
| **Velocity of Development** | 3.7/5.0 | Good - vertical slices make features easy to add (3-5 files typically). Clear domain boundaries reduce coordination overhead. However, orchestrator changes can affect multiple domains, and shared type changes require coordination. |
| **Ease of Onboarding** | 3.5/5.0 | Moderate - clear entry points (controller → orchestrator) and domain structure help. However, large orchestrator (~530 lines) can be overwhelming. Documentation exists but could use more examples and sequence diagrams. |
| **Development Autonomy** | 4.0/5.0 | Good - domains are self-contained, allowing teams to work independently on their domains. Shared types require coordination, but this is manageable. Interface contracts define clear boundaries. Orchestrator changes may require cross-team input. |
| **Architectural Scalability** | 3.8/5.0 | Good foundation - vertical slices and domain boundaries support scaling. Can extract to microservices easily. However, orchestrator could become a bottleneck. No horizontal scaling strategy yet (single binary). Needs caching, connection pooling, and load balancing for production scale. |

### Rating Summary

**Strengths** (4.0+):
- Logical Separation of Concerns (4.5)
- Ability to Isolate to Microservices (4.5)
- Loose Coupling (4.2)
- Infra Abstraction Fit (4.0)
- Testability (4.0)
- Development Autonomy (4.0)

**Areas for Improvement** (<3.5):
- Debugability (3.0) - Add tracing, structured logging, request IDs
- Performance (3.0) - Add caching, optimize queries, benchmark
- Schema Maintainability (3.2) - Add versioning, migration strategy
- Debugability (3.0) - Improve observability

**Overall Architecture Score: 3.7/5.0**

The architecture provides a solid foundation with clear domain boundaries and good separation of concerns. The main areas for improvement are observability (debugging/tracing), performance optimization, and schema management strategy.

## Related Documentation

- [README.md](README.md) - Setup and usage instructions
- [.structure.md](.structure.md) - Directory structure details
- [internal/quote/orchestrator_test.go](internal/quote/orchestrator_test.go) - Test examples
/