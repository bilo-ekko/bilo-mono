package main

import (
	"database/sql"
	"os"

	// Adapters
	"api-golang/internal/adapters/impact_partner/http"
	postgresAdapter "api-golang/internal/adapters/quote/persistence/postgres"

	// Domains
	"api-golang/internal/impact_partner/impact_partner"
	"api-golang/internal/quote"
)

// This file demonstrates how to use adapters in the main application.
// It shows how to swap between different adapter implementations.

func setupQuoteOrchestrator() *quote.Orchestrator {
	// ============================================
	// ADAPTER SELECTION BASED ON ENVIRONMENT
	// ============================================

	// 1. Choose Quote Repository Adapter (Driven Adapter)
	var quoteRepo quote.Repository
	if os.Getenv("DATABASE_URL") != "" {
		// Use PostgreSQL adapter
		db, _ := sql.Open("postgres", os.Getenv("DATABASE_URL"))
		quoteRepo = postgresAdapter.NewPostgresRepository(db)
	} else {
		// Use in-memory adapter (for development/testing)
		quoteRepo = quote.NewInMemoryRepository()
	}

	// 2. Choose Impact Partner Service Adapter (Driven Adapter)
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

	// 3. Create Orchestrator with Adapters
	// The orchestrator doesn't know or care which adapters are used!
	// It only depends on interfaces (ports).
	orchestrator := quote.NewOrchestrator(quote.OrchestratorDeps{
		// ... other dependencies
		ImpactPartnerService: partnerService, // Same interface, different implementation!
		QuoteRepo:            quoteRepo,      // Same interface, different implementation!
		// ...
	})

	return orchestrator
}

// Key Points:
//
// 1. ADAPTERS ARE INJECTED AT APPLICATION STARTUP
//    - Main application decides which adapters to use
//    - Based on environment variables, config, etc.
//
// 2. DOMAIN DOESN'T KNOW ABOUT ADAPTERS
//    - Orchestrator only knows about interfaces (ports)
//    - Can swap adapters without changing domain code
//
// 3. ADAPTERS IMPLEMENT PORTS
//    - HTTPClientAdapter implements impact_partner.Service interface
//    - PostgresRepository implements quote.Repository interface
//
// 4. EASY TO TEST
//    - Can inject mock adapters for testing
//    - Domain logic tested independently of infrastructure
//
// 5. MICROSERVICES READY
//    - HTTP client adapters call remote services
//    - No changes needed to domain code
