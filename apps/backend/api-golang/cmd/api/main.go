package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	// Shared packages
	"github.com/bilo-mono/packages/common/calculator"
	"github.com/bilo-mono/packages/common/logger"

	// Domain imports
	"api-golang/internal/finance/currency"
	"api-golang/internal/funds/salestax"
	carbonfootprint "api-golang/internal/impact/carbon_footprint"
	"api-golang/internal/impact/fee"
	"api-golang/internal/impact_partner"
	"api-golang/internal/impact_project"
	"api-golang/internal/organisation/customer"
	"api-golang/internal/organisation/organisation"
	"api-golang/internal/platform/country"
	"api-golang/internal/quote"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

type HealthResponse struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

func main() {
	// Initialize shared logger
	appLogger := logger.NewLogger("Go-API")
	appLogger.Info("Initializing Go API application...")

	// Demonstrate calculator usage
	calcResult := calculator.CalculateX(calculator.CalculateXInput{
		Value:      100.0,
		Multiplier: ptrFloat64(1.5),
		Offset:     ptrFloat64(25.0),
	})
	appLogger.Infof("Calculation result: %s", calcResult.Formula)

	// ============================================
	// Initialize all domain services
	// ============================================

	// Organisation domain
	orgRepo := organisation.NewInMemoryRepository()
	orgService := organisation.NewService(orgRepo)

	customerRepo := customer.NewInMemoryRepository()
	customerService := customer.NewService(customerRepo)

	// Platform domain
	countryRepo := country.NewInMemoryRepository()
	countryService := country.NewService(countryRepo)

	// Finance domain
	currencyRepo := currency.NewInMemoryRepository()
	currencyService := currency.NewService(currencyRepo)

	// Impact domain - Carbon Footprint
	carbonFactorRepo := carbonfootprint.NewInMemoryFactorRepository()
	carbonFootprintRepo := carbonfootprint.NewInMemoryFootprintRepository()
	carbonService := carbonfootprint.NewService(carbonFactorRepo, carbonFootprintRepo)

	// Impact domain - Fee
	feeRepo := fee.NewInMemoryRepository()
	feeService := fee.NewService(feeRepo)

	// Impact Partner domain (existing)
	partnerRepo := impact_partner.NewRepository()
	partnerService := impact_partner.NewService(partnerRepo)
	partnerController := impact_partner.NewController(partnerService)
	blendedPriceCalc := impact_partner.NewBlendedPriceCalculator(partnerService)

	// Impact Project domain (existing)
	projectRepo := impact_project.NewRepository()
	projectService := impact_project.NewService(projectRepo)
	projectController := impact_project.NewController(projectService)

	// Funds domain - Sales Tax
	salesTaxRepo := salestax.NewInMemoryRepository()
	salesTaxService := salestax.NewService(salesTaxRepo)

	// Quote domain - Orchestrator
	quoteRepo := quote.NewInMemoryRepository()
	quoteOrchestrator := quote.NewOrchestrator(quote.OrchestratorDeps{
		OrganisationService:  orgService,
		CustomerService:      customerService,
		CountryService:       countryService,
		CurrencyService:      currencyService,
		CarbonService:        carbonService,
		FeeService:           feeService,
		BlendedPriceCalc:     blendedPriceCalc,
		ImpactPartnerService: partnerService,
		SalesTaxService:      salesTaxService,
		QuoteRepo:            quoteRepo,
	})
	quoteController := quote.NewController(quoteOrchestrator)

	appLogger.Info("All domain services initialized successfully")

	// ============================================
	// Register routes
	// ============================================

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)

	// Impact Partners routes
	http.HandleFunc("/api/impact-partners", partnerController.HandleGetAll)
	http.HandleFunc("/api/impact-partners/", func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/impact-partners/") != "" {
			partnerController.HandleGetByID(w, r)
		} else {
			partnerController.HandleGetAll(w, r)
		}
	})

	// Impact Projects routes
	http.HandleFunc("/api/impact-projects", projectController.HandleGetAll)
	http.HandleFunc("/api/impact-projects/", func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/impact-projects/") != "" {
			projectController.HandleGetByID(w, r)
		} else {
			projectController.HandleGetAll(w, r)
		}
	})

	// Quote routes (NEW)
	http.HandleFunc("/api/quotes", quoteController.HandleCreateQuote)
	http.HandleFunc("/api/quotes/", func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/api/quotes/") != "" {
			quoteController.HandleGetQuote(w, r)
		} else {
			quoteController.HandleCreateQuote(w, r)
		}
	})

	// ============================================
	// Start server
	// ============================================
	port := ":8080"
	fmt.Printf("🚀 Go API Server starting on http://localhost%s\n", port)
	fmt.Println("\nAvailable endpoints:")
	fmt.Println("  - GET  http://localhost" + port + "/")
	fmt.Println("  - GET  http://localhost" + port + "/api/health")
	fmt.Println("  - GET  http://localhost" + port + "/api/hello?name=World")
	fmt.Println("\nImpact Partners:")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-partners")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-partners/{id}")
	fmt.Println("\nImpact Projects:")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects/{id}")
	fmt.Println("  - GET  http://localhost" + port + "/api/impact-projects?partnerId={id}")
	fmt.Println("\nQuotes:")
	fmt.Println("  - POST http://localhost" + port + "/api/quotes")
	fmt.Println("  - GET  http://localhost" + port + "/api/quotes/{id}")
	fmt.Println()

	appLogger.Infof("Server listening on http://localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		appLogger.Error("Server failed to start", err)
	}
}

// ptrFloat64 returns a pointer to a float64 value
func ptrFloat64(v float64) *float64 {
	return &v
}

// homeHandler handles requests to the root path
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to api-golang! 🚀",
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

// healthHandler handles health check requests
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthResponse{
		Status: "ok",
		Time:   time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

// helloHandler handles hello requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := Response{
		Message: fmt.Sprintf("Hello, %s!", name),
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}
