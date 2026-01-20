package quote

import (
	"context"
	"testing"

	"api-golang/internal/finance/currency"
	"api-golang/internal/funds/salestax"
	carbonfootprint "api-golang/internal/impact/carbon_footprint"
	"api-golang/internal/impact/fee"
	"api-golang/internal/impact_partner"
	"api-golang/internal/organisation/customer"
	"api-golang/internal/organisation/organisation"
	"api-golang/internal/platform/country"
)

// setupOrchestrator creates an orchestrator with all dependencies for testing
func setupOrchestrator() *Orchestrator {
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

	// Impact domain
	carbonFactorRepo := carbonfootprint.NewInMemoryFactorRepository()
	carbonFootprintRepo := carbonfootprint.NewInMemoryFootprintRepository()
	carbonService := carbonfootprint.NewService(carbonFactorRepo, carbonFootprintRepo)

	feeRepo := fee.NewInMemoryRepository()
	feeService := fee.NewService(feeRepo)

	// Impact Partner domain
	partnerRepo := impact_partner.NewRepository()
	partnerService := impact_partner.NewService(partnerRepo)
	blendedPriceCalc := impact_partner.NewBlendedPriceCalculator(partnerService)

	// Funds domain
	salesTaxRepo := salestax.NewInMemoryRepository()
	salesTaxService := salestax.NewService(salesTaxRepo)

	// Quote domain
	quoteRepo := NewInMemoryRepository()

	return NewOrchestrator(OrchestratorDeps{
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
}

func TestCreateQuote_Success(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "cust-ref-001",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-001",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        100.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote failed: %v", err)
	}

	// Verify response structure
	if response.ID == "" {
		t.Error("Expected quote to have an ID")
	}
	if response.QuoteReference == "" {
		t.Error("Expected quote to have a quoteReference")
	}

	// Verify footprint
	if response.Footprint.Co2eGrams <= 0 {
		t.Error("Expected positive carbon footprint in grams")
	}
	if response.Footprint.Co2eOunces <= 0 {
		t.Error("Expected positive carbon footprint in ounces")
	}
	if len(response.Footprint.Equivalents) == 0 {
		t.Error("Expected at least one equivalent")
	}

	// Verify credits
	if response.Credits.TotalAmount <= 0 {
		t.Error("Expected positive total amount")
	}
	if response.Credits.ImpactAmount <= 0 {
		t.Error("Expected positive impact amount")
	}
	if response.Credits.PricePerTonneCo2e <= 0 {
		t.Error("Expected positive price per tonne CO2e")
	}
	if len(response.Credits.ImpactPartners) == 0 {
		t.Error("Expected at least one impact partner")
	}

	// Verify contribution
	if response.Contribution.ImpactPercentage < 0 || response.Contribution.ImpactPercentage > 1 {
		t.Errorf("Expected impact percentage between 0 and 1, got %f", response.Contribution.ImpactPercentage)
	}
	if len(response.Contribution.ImpactPartners) == 0 {
		t.Error("Expected at least one impact partner in contribution")
	}

	t.Logf("Quote created successfully: ID=%s, Reference=%s, Total=€%.2f",
		response.ID, response.QuoteReference, response.Credits.TotalAmount)
}

func TestCreateQuote_WithChildOrganisation(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	// Create quote for child organisation through parent
	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-child-1", // Child of org-parent-1
		Customer: CustomerRequest{
			Reference: "cust-ref-002",
			Country:   "IRL",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-002",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        50.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote for child org failed: %v", err)
	}

	if response.ID == "" {
		t.Error("Expected quote to have an ID")
	}

	t.Logf("Quote for child org created: ID=%s", response.ID)
}

func TestCreateQuote_UnauthorisedOrganisation(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	// Try to create quote for org that is not a child of header org
	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-child-2", // Different child (not child of org-child-1)
		Customer: CustomerRequest{
			Reference: "cust-ref-003",
			Country:   "DEU",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-003",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        25.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	_, err := orchestrator.CreateQuote(ctx, req, "org-child-1")
	if err == nil {
		t.Fatal("Expected error for unauthorised organisation access")
	}

	t.Logf("Correctly rejected unauthorised access: %v", err)
}

func TestCreateQuote_WithCurrencyConversion(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:         "en-US",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "cust-ref-004",
			Country:   "USA",
			State:     stringPtr("CA"),
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-004",
				Name:     "US Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        100.00,
					CurrencyCode: "USD", // Will be converted to EUR for calculation
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote with USD failed: %v", err)
	}

	// Verify quote was created with USD currency
	if response.Credits.TotalAmount <= 0 {
		t.Error("Expected positive total amount")
	}

	t.Logf("Currency conversion: $%.2f USD quote created, Total=€%.2f",
		100.00, response.Credits.TotalAmount)
}

func TestCreateQuote_WithMerchantDetails(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "cust-ref-005",
			Country:   "GBR",
		},
		Merchant: &MerchantRequest{
			Name: "Test Restaurant",
			MCC:  "5812", // Restaurant MCC
			Address: MerchantAddressRequest{
				Address1:   "123 High Street",
				City:       "London",
				PostalCode: "SW1A 1AA",
				Country:    "GBR",
			},
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-005",
				Name:     "Restaurant Meal",
				Category: "restaurant",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        200.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote with merchant details failed: %v", err)
	}

	// Restaurant MCC (5812) has factor 0.35, which should give higher carbon than default
	// Default factor is 0.23, so restaurant factor should give more carbon
	expectedMinCarbon := 200.00 * 0.30 // At least more than 200 * 0.23
	if response.Footprint.Co2eGrams/1000.0 < expectedMinCarbon {
		t.Errorf("Expected carbon footprint > %.2f kg for restaurant MCC, got %.2f kg",
			expectedMinCarbon, response.Footprint.Co2eGrams/1000.0)
	}

	t.Logf("Merchant-specific carbon: MCC=5812, Carbon=%.2f kg",
		response.Footprint.Co2eGrams/1000.0)
}

func TestCreateQuote_WithFilterByLocation(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "cust-ref-006",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-006",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        150.00,
					CurrencyCode: "EUR",
				},
			},
		},
		Filters: &QuoteFiltersRequest{
			CustomerLocation: true, // Filter projects by customer location
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote with location filter failed: %v", err)
	}

	// Verify location match is indicated
	if response.Credits.CustomerLocationMatch == "" {
		t.Error("Expected customerLocationMatch to be set when filtering by location")
	}

	t.Logf("Location filtered: CustomerLocationMatch=%s, Partners=%d",
		response.Credits.CustomerLocationMatch, len(response.Credits.ImpactPartners))
}

func TestCreateQuote_WithImpactPartnerDetails(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:                      "en-GB",
		OrganisationID:              "org-parent-1",
		IncludeImpactPartnerDetails: true,
		Customer: CustomerRequest{
			Reference: "cust-ref-007",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-007",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        100.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote with partner details failed: %v", err)
	}

	// Verify partner details are included
	if len(response.Credits.ImpactPartners) > 0 {
		partner := response.Credits.ImpactPartners[0]
		if partner.Name == nil {
			t.Error("Expected partner name when IncludeImpactPartnerDetails is true")
		}
	}

	t.Logf("Partner details included: Partners=%d", len(response.Credits.ImpactPartners))
}

func TestGetQuote(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	// First create a quote
	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "cust-ref-008",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-008",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        50.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	createResponse, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote failed: %v", err)
	}

	// Now retrieve it
	quote, err := orchestrator.GetQuote(ctx, createResponse.ID)
	if err != nil {
		t.Fatalf("GetQuote failed: %v", err)
	}

	if quote.ID != createResponse.ID {
		t.Errorf("Expected quote ID %s, got %s", createResponse.ID, quote.ID)
	}
	if quote.QuoteReference != createResponse.QuoteReference {
		t.Errorf("Expected quote reference %s, got %s", createResponse.QuoteReference, quote.QuoteReference)
	}
	if quote.OrganisationID != "org-parent-1" {
		t.Errorf("Expected organisationId org-parent-1, got %s", quote.OrganisationID)
	}
	if quote.Status != StatusPending {
		t.Errorf("Expected status %s, got %s", StatusPending, quote.Status)
	}
	if quote.Currency == "" {
		t.Error("Expected currency to be set")
	}

	t.Logf("Retrieved quote: ID=%s, Reference=%s, Status=%s",
		quote.ID, quote.QuoteReference, quote.Status)
}

func TestGetQuote_NotFound(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	_, err := orchestrator.GetQuote(ctx, "non-existent-id")
	if err == nil {
		t.Fatal("Expected error for non-existent quote")
	}

	t.Logf("Correctly returned error for non-existent quote: %v", err)
}

func TestCreateQuote_WithEmptyCustomerReference(t *testing.T) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	// Note: Customer reference validation happens at the controller level,
	// not the orchestrator level. The orchestrator will create a customer
	// with an empty reference if provided. This test verifies that behavior.
	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			// Empty Reference - orchestrator allows this, controller validates
			Reference: "",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "item-empty-ref",
				Name:     "Test Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        100.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	response, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
	if err != nil {
		t.Fatalf("CreateQuote with empty reference should succeed at orchestrator level: %v", err)
	}

	// Verify quote was created successfully
	if response.ID == "" {
		t.Error("Expected quote to have an ID")
	}

	t.Logf("Quote created with empty customer reference: ID=%s (validation happens at controller level)",
		response.ID)
}

// Helper function for string pointers
func stringPtr(s string) *string {
	return &s
}

// Benchmark for quote creation
func BenchmarkCreateQuote(b *testing.B) {
	orchestrator := setupOrchestrator()
	ctx := context.Background()

	req := &CreateQuoteRequest{
		Locale:         "en-GB",
		OrganisationID: "org-parent-1",
		Customer: CustomerRequest{
			Reference: "bench-cust",
			Country:   "GBR",
		},
		OrderItems: []OrderItemRequest{
			{
				ItemID:   "bench-item",
				Name:     "Benchmark Product",
				Category: "general",
				Quantity: 1,
				UnitPrice: OrderItemPrice{
					Value:        100.00,
					CurrencyCode: "EUR",
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req.Customer.Reference = "bench-cust-" + string(rune(i))
		req.OrderItems[0].ItemID = "bench-item-" + string(rune(i))
		_, err := orchestrator.CreateQuote(ctx, req, "org-parent-1")
		if err != nil {
			b.Fatalf("CreateQuote failed: %v", err)
		}
	}
}
