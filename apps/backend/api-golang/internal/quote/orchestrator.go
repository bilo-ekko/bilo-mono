package quote

import (
	"context"
	"fmt"
	"math"
	"time"

	"api-golang/internal/finance/currency"
	"api-golang/internal/funds/salestax"
	carbonfootprint "api-golang/internal/impact/carbon_footprint"
	"api-golang/internal/impact/fee"
	"api-golang/internal/impact_partner"
	"api-golang/internal/organisation/customer"
	"api-golang/internal/organisation/organisation"
	"api-golang/internal/platform/country"
	"api-golang/internal/shared/types"

	"github.com/google/uuid"
)

// convertOrderItems converts OrderItemRequest slice to OrderItems
func convertOrderItems(items []OrderItemRequest) OrderItems {
	result := make(OrderItems, len(items))
	for i, item := range items {
		result[i] = OrderItem{
			ItemID:   item.ItemID,
			SKU:      item.SKU,
			Name:     item.Name,
			Category: item.Category,
			Quantity: item.Quantity,
			UnitPrice: types.Money{
				Amount:   item.UnitPrice.Value,
				Currency: item.UnitPrice.CurrencyCode,
			},
		}
	}
	return result
}

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

// OrchestratorDeps contains all dependencies for the orchestrator
type OrchestratorDeps struct {
	OrganisationService organisation.Service
	CustomerService     customer.Service
	CountryService      country.Service
	CurrencyService     currency.Service
	CarbonService       carbonfootprint.Service
	FeeService          fee.Service
	BlendedPriceCalc    *impact_partner.BlendedPriceCalculator
	ImpactPartnerRepo   *impact_partner.Repository
	SalesTaxService     salestax.Service
	QuoteRepo           Repository
}

// NewOrchestrator creates a new quote orchestrator
func NewOrchestrator(deps OrchestratorDeps) *Orchestrator {
	return &Orchestrator{
		organisationService: deps.OrganisationService,
		customerService:     deps.CustomerService,
		countryService:      deps.CountryService,
		currencyService:     deps.CurrencyService,
		carbonService:       deps.CarbonService,
		feeService:          deps.FeeService,
		blendedPriceCalc:    deps.BlendedPriceCalc,
		impactPartnerRepo:   deps.ImpactPartnerRepo,
		salesTaxService:     deps.SalesTaxService,
		quoteRepo:           deps.QuoteRepo,
	}
}

// CreateQuote implements the full quote creation flow following the sequence:
// 1. Validate Organisation
// 2. Get or Create Customer
// 3. Calculate Carbon Footprint
// 4. Get Blended Project Unit Price
// 5. Calculate Compensation Amount
// 6. Calculate Round Up
// 7. Calculate Service Fee
// 8. Calculate Sales Tax
// 9. Write Quote
func (o *Orchestrator) CreateQuote(ctx context.Context, req *CreateQuoteRequest, headerOrgID string) (*CreateQuoteResponse, error) {
	// ============================================
	// Step 1: Validate Organisation
	// ============================================
	org, err := o.organisationService.ValidateOrganisation(ctx, headerOrgID, req.OrganisationID)
	if err != nil {
		return nil, fmt.Errorf("step 1 - validate organisation: %w", err)
	}

	// ============================================
	// Step 2: Get or Create Customer
	// ============================================
	// Always use reference to get or create customer
	cust, err := o.customerService.GetOrCreateCustomer(ctx, customer.CreateCustomerInput{
		OrganisationID: org.OrganisationID,
		Reference:      req.Customer.Reference,
		CountryCode:    req.Customer.Country, // API uses "country", we store as "countryCode"
		State:          req.Customer.State,
		PostalCode:     req.Customer.PostalCode,
		City:           req.Customer.City,
	})
	if err != nil {
		return nil, fmt.Errorf("step 2 - get/create customer: %w", err)
	}

	// ============================================
	// Step 3: Calculate Carbon Footprint
	// ============================================
	// 3.1: Use merchant details from request or fall back to organisation defaults
	merchantMCC := org.GetMCC()
	merchantCountryCode := org.Address.CountryCode
	if req.Merchant != nil {
		if req.Merchant.MCC != "" {
			merchantMCC = req.Merchant.MCC
		}
		if req.Merchant.Address.Country != "" {
			merchantCountryCode = req.Merchant.Address.Country
		}
	}

	// 3.2: Get merchant country ID
	merchantCountry, err := o.countryService.GetCountryByCode(ctx, merchantCountryCode)
	if err != nil {
		return nil, fmt.Errorf("step 3.2 - get merchant country: %w", err)
	}

	// 3.3: Calculate transaction amount from orderItems or use a default
	// In the new API, we need to sum orderItems or use a transaction amount
	// For now, we'll calculate from orderItems if provided
	var transactionAmount float64
	var transactionCurrency string = "EUR" // Default
	if len(req.OrderItems) > 0 {
		for _, item := range req.OrderItems {
			transactionAmount += item.UnitPrice.Value * float64(item.Quantity)
			if transactionCurrency == "EUR" {
				transactionCurrency = item.UnitPrice.CurrencyCode
			}
		}
	} else {
		// If no orderItems, we need a transaction amount - this should be provided
		// For demo purposes, we'll use a default
		transactionAmount = 100.0
		transactionCurrency = "EUR"
	}

	// 3.4: Convert currency to EUR if needed
	var amountEUR float64
	if transactionCurrency != "EUR" {
		conversionResult, err := o.currencyService.ConvertToEUR(ctx, transactionAmount, transactionCurrency)
		if err != nil {
			return nil, fmt.Errorf("step 3.4 - convert currency: %w", err)
		}
		amountEUR = conversionResult.ConvertedAmount
	} else {
		amountEUR = transactionAmount
	}

	// 3.5: Calculate carbon footprint using MCC and country
	transactionID := uuid.New().String()
	footprint, err := o.carbonService.Calculate(ctx, carbonfootprint.CalculateInput{
		TransactionID:  transactionID,
		OrganisationID: org.OrganisationID,
		CustomerID:     cust.ID,
		AmountEUR:      amountEUR,
		MCC:            merchantMCC,
		CountryID:      merchantCountry.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("step 3.5 - calculate carbon footprint: %w", err)
	}

	// ============================================
	// Step 4: Get Blended Project Unit Price
	// ============================================
	// Check if customer location filter is enabled
	filterByLocation := false
	locationFilter := ""
	if req.Filters != nil && req.Filters.CustomerLocation {
		filterByLocation = true
		locationFilter = req.Customer.Country // Use customer country for filtering
	}

	blendedPrice, err := o.blendedPriceCalc.CalculateBlendedPrice(
		org.OrganisationID,
		filterByLocation,
		locationFilter,
	)
	if err != nil {
		return nil, fmt.Errorf("step 4 - get blended price: %w", err)
	}

	// Convert blended price to quote currency if needed
	// BlendedUnitPrice is per kg CO2e, convert to per tonne (multiply by 1000)
	quoteCurrency := transactionCurrency
	pricePerKgCo2e := blendedPrice.BlendedUnitPrice
	pricePerTonneCo2e := pricePerKgCo2e * 1000.0 // Convert from per kg to per tonne
	if quoteCurrency != "EUR" {
		// Convert price from EUR to quote currency
		conversionResult, err := o.currencyService.ConvertFromEUR(ctx, pricePerTonneCo2e, quoteCurrency)
		if err != nil {
			return nil, fmt.Errorf("step 4.1 - convert price to quote currency: %w", err)
		}
		pricePerTonneCo2e = conversionResult.ConvertedAmount
	}

	// ============================================
	// Step 5: Calculate Compensation Amount (Impact Amount)
	// ============================================
	// Convert carbon footprint from kg to tonnes
	carbonTonnes := footprint.CarbonKg() / 1000.0
	impactAmount := carbonTonnes * pricePerTonneCo2e
	impactAmount = math.Round(impactAmount*100) / 100

	// ============================================
	// Step 6: Calculate Round Up (disabled in new API)
	// ============================================
	roundUpAmount := 0.0
	totalBeforeFees := impactAmount + roundUpAmount

	// ============================================
	// Step 7: Calculate Service Fee
	// ============================================
	feeResult, err := o.feeService.CalculateServiceFee(ctx, org.OrganisationID, totalBeforeFees)
	if err != nil {
		return nil, fmt.Errorf("step 7 - calculate service fee: %w", err)
	}
	serviceFeeAmount := feeResult.FeeAmount

	// ============================================
	// Step 8: Calculate Sales Tax
	// ============================================
	// Calculate tax on impact amount
	merchantAddress := types.Address{
		CountryCode: org.Address.CountryCode,
		State:       org.Address.State,
		PostalCode:  org.Address.PostalCode,
	}
	if req.Merchant != nil && req.Merchant.Address.Country != "" {
		merchantAddress.CountryCode = req.Merchant.Address.Country
		if req.Merchant.Address.State != nil {
			merchantAddress.State = req.Merchant.Address.State
		}
		merchantAddress.PostalCode = req.Merchant.Address.PostalCode
	}

	var customerState, customerPostalCode string
	if req.Customer.State != nil {
		customerState = *req.Customer.State
	}
	if req.Customer.PostalCode != nil {
		customerPostalCode = *req.Customer.PostalCode
	}
	var merchantState, merchantPostalCode string
	if merchantAddress.State != nil {
		merchantState = *merchantAddress.State
	}
	merchantPostalCode = merchantAddress.PostalCode

	// Calculate tax on impact amount
	impactTaxResult, err := o.salesTaxService.CalculateSalesTax(ctx, salestax.TaxCalculationInput{
		MerchantCountry:    merchantAddress.CountryCode,
		MerchantState:      merchantState,
		MerchantPostalCode: merchantPostalCode,
		CustomerCountry:    req.Customer.Country,
		CustomerState:      customerState,
		CustomerPostalCode: customerPostalCode,
		Amount:             impactAmount,
	})
	if err != nil {
		return nil, fmt.Errorf("step 8.1 - calculate impact sales tax: %w", err)
	}
	impactSalesTaxAmount := impactTaxResult.TaxAmount
	impactTaxRate := impactTaxResult.TaxRate

	// Calculate tax on service fee
	serviceFeeTaxResult, err := o.salesTaxService.CalculateSalesTax(ctx, salestax.TaxCalculationInput{
		MerchantCountry:    merchantAddress.CountryCode,
		MerchantState:      merchantState,
		MerchantPostalCode: merchantPostalCode,
		CustomerCountry:    req.Customer.Country,
		CustomerState:      customerState,
		CustomerPostalCode: customerPostalCode,
		Amount:             serviceFeeAmount,
	})
	if err != nil {
		return nil, fmt.Errorf("step 8.2 - calculate service fee sales tax: %w", err)
	}
	serviceFeeSalesTaxAmount := serviceFeeTaxResult.TaxAmount
	serviceFeeTaxRate := serviceFeeTaxResult.TaxRate

	// ============================================
	// Step 9: Calculate Totals and Build Response
	// ============================================
	totalAmount := impactAmount + impactSalesTaxAmount + serviceFeeAmount + serviceFeeSalesTaxAmount
	totalAmount = math.Round(totalAmount*100) / 100

	// Determine customer location match (best effort: state, country, region, world)
	customerLocationMatch := "world" // Default
	if filterByLocation {
		// Try to match projects by customer location
		// For demo, we'll use a simple check
		for _, project := range blendedPrice.Projects {
			if project.Location.Country == req.Customer.Country {
				customerLocationMatch = "country"
				if customerState != "" && project.Location.Region != nil && *project.Location.Region == customerState {
					customerLocationMatch = "state"
				}
				break
			}
		}
	}

	// Build impact partners response
	impactPartnersMap := make(map[string]*ImpactPartnerResponse)
	for _, project := range blendedPrice.Projects {
		partner, exists := impactPartnersMap[project.PartnerID]
		if !exists {
			partner = &ImpactPartnerResponse{
				ID:       project.PartnerID,
				Projects: []ProjectResponse{},
			}
			// Add partner details if requested
			if req.IncludeImpactPartnerDetails {
				partnerEntity, err := o.impactPartnerRepo.GetByID(project.PartnerID)
				if err == nil {
					partner.Name = &partnerEntity.Name
					if partnerEntity.ShortDescription != nil {
						partner.Description = partnerEntity.ShortDescription
					}
					partner.Logo = partnerEntity.Logo
				}
			}
			impactPartnersMap[project.PartnerID] = partner
		}
		partner.Projects = append(partner.Projects, ProjectResponse{
			ID: project.ProjectID,
		})
	}

	// Convert map to slice
	impactPartners := make([]ImpactPartnerResponse, 0, len(impactPartnersMap))
	for _, partner := range impactPartnersMap {
		impactPartners = append(impactPartners, *partner)
	}

	// Build contribution breakdown
	// For demo, we'll split evenly across partners
	numPartners := float64(len(impactPartners))
	var impactPercentagePerPartner, serviceFeePercentagePerPartner float64
	if numPartners > 0 {
		impactPercentagePerPartner = 1.0 / numPartners
		serviceFeePercentagePerPartner = 1.0 / numPartners
	}

	contributionImpactPartners := make([]ContributionImpactPartnerResponse, 0, len(impactPartners))
	for _, partner := range impactPartners {
		var impactSalesTaxPerPartner, serviceFeeSalesTaxPerPartner float64
		if numPartners > 0 {
			impactSalesTaxPerPartner = impactTaxRate / numPartners
			serviceFeeSalesTaxPerPartner = serviceFeeTaxRate / numPartners
		}
		contributionImpactPartners = append(contributionImpactPartners, ContributionImpactPartnerResponse{
			ID:                           partner.ID,
			ImpactPercentage:             impactPercentagePerPartner,
			ImpactSalesTaxPercentage:     impactSalesTaxPerPartner,
			ServiceFeePercentage:         serviceFeePercentagePerPartner,
			ServiceFeeSalesTaxPercentage: serviceFeeSalesTaxPerPartner,
			Name:                         partner.Name,
			Description:                  partner.Description,
			Logo:                         partner.Logo,
			Projects:                     partner.Projects,
		})
	}

	// Calculate contribution totals
	// Avoid division by zero
	var totalImpactPercentage, totalImpactSalesTaxPercentage, totalServiceFeePercentage, totalServiceFeeSalesTaxPercentage float64
	if totalAmount > 0 {
		totalImpactPercentage = impactAmount / totalAmount
		totalImpactSalesTaxPercentage = impactSalesTaxAmount / totalAmount
		totalServiceFeePercentage = serviceFeeAmount / totalAmount
		totalServiceFeeSalesTaxPercentage = serviceFeeSalesTaxAmount / totalAmount
	}

	// ============================================
	// Step 10: Write Quote Entity
	// ============================================
	quoteReference := uuid.New().String()
	now := time.Now()
	quote := &Entity{
		ID:                   uuid.New().String(),
		QuoteReference:       quoteReference,
		CalculationReference: footprint.ID,
		OrganisationID:       org.OrganisationID,
		CustomerID:           cust.ID,
		Currency:             quoteCurrency,

		CarbonCreditTotal:              totalAmount,
		CarbonCreditImpact:             impactAmount,
		CarbonCreditImpactSalesTax:     impactSalesTaxAmount,
		ImpactTaxRate:                  impactTaxRate,
		CarbonCreditServiceFee:         serviceFeeAmount,
		CarbonCreditServiceFeeSalesTax: serviceFeeSalesTaxAmount,
		ServiceFeeTaxRate:              serviceFeeTaxRate,

		PricePerTonneCo2e: pricePerTonneCo2e,

		ContributionDetails: ContributionDetails{
			ImpactPercentage:             totalImpactPercentage,
			ImpactSalesTaxPercentage:     totalImpactSalesTaxPercentage,
			ServiceFeePercentage:         totalServiceFeePercentage,
			ServiceFeeSalesTaxPercentage: totalServiceFeeSalesTaxPercentage,
			ImpactPartners:               make([]ContributionImpactPartner, len(contributionImpactPartners)),
		},

		CustomerLocationFilter: filterByLocation,
		IncludePartnerDetail:   req.IncludeImpactPartnerDetails,
		IncludeProjectDetail:   false, // Project details never available in quote response

		EkkoProduct:     "API", // Track through endpoint
		ServiceFeeShare: feeResult.FeePercentage,

		OrderItems: convertOrderItems(req.OrderItems),

		Status:    StatusPending,
		ExpiresAt: now.Add(24 * time.Hour),
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Convert contribution impact partners for storage
	for i, cp := range contributionImpactPartners {
		projectIDs := make([]string, len(cp.Projects))
		for j, p := range cp.Projects {
			projectIDs[j] = p.ID
		}
		quote.ContributionDetails.ImpactPartners[i] = ContributionImpactPartner{
			ID:                           cp.ID,
			ImpactPercentage:             cp.ImpactPercentage,
			ImpactSalesTaxPercentage:     cp.ImpactSalesTaxPercentage,
			ServiceFeePercentage:         cp.ServiceFeePercentage,
			ServiceFeeSalesTaxPercentage: cp.ServiceFeeSalesTaxPercentage,
			ProjectIDs:                   projectIDs,
		}
	}

	if err := o.quoteRepo.Create(ctx, quote); err != nil {
		return nil, fmt.Errorf("step 10 - save quote: %w", err)
	}

	// ============================================
	// Step 11: Build Response
	// ============================================
	// Use stored grams and ounces from footprint
	co2eGrams := footprint.CarbonCo2eGrams
	co2eOunces := footprint.CarbonCo2eOunces

	// Build equivalents (simplified for demo)
	carbonKg := footprint.CarbonKg()
	equivalents := []EquivalentResponse{
		{
			Key:      "tree",
			Value:    carbonKg / 40.0, // Roughly 40kg CO2 per tree
			Template: fmt.Sprintf("That's like planting %.1f trees", carbonKg/40.0),
		},
	}

	response := &CreateQuoteResponse{
		ID:             quote.ID,
		QuoteReference: quoteReference,
		Footprint: FootprintResponse{
			Co2eGrams:   co2eGrams,
			Co2eOunces:  co2eOunces,
			Equivalents: equivalents,
		},
		Credits: CreditsResponse{
			TotalAmount:              totalAmount,
			ImpactAmount:             impactAmount,
			ImpactSalesTaxAmount:     impactSalesTaxAmount,
			ServiceFeeAmount:         serviceFeeAmount,
			ServiceFeeSalesTaxAmount: serviceFeeSalesTaxAmount,
			PricePerTonneCo2e:        pricePerTonneCo2e,
			ImpactPartners:           impactPartners,
			CustomerLocationMatch:    customerLocationMatch,
		},
		Contribution: ContributionResponse{
			ImpactPercentage:             totalImpactPercentage,
			ImpactSalesTaxPercentage:     totalImpactSalesTaxPercentage,
			ServiceFeePercentage:         totalServiceFeePercentage,
			ServiceFeeSalesTaxPercentage: totalServiceFeeSalesTaxPercentage,
			ImpactPartners:               contributionImpactPartners,
		},
	}

	return response, nil
}

// GetQuote retrieves a quote by ID
func (o *Orchestrator) GetQuote(ctx context.Context, id string) (*Entity, error) {
	return o.quoteRepo.GetByID(ctx, id)
}
