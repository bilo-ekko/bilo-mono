package salestax

import (
	"context"
	"sync"
)

// InMemoryRepository implements Repository interface
type InMemoryRepository struct {
	rates map[string]*TaxRate // key: countryCode:state:postalCode (empty parts = *)
	mu    sync.RWMutex
}

// NewInMemoryRepository creates a new repository with sample data
func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		rates: make(map[string]*TaxRate),
	}

	// Seed with sample VAT rates by country
	rates := []*TaxRate{
		// EU VAT rates
		{ID: "1", MerchantLocation: MerchantLocation{Country: "GBR"}, CustomerLocation: CustomerLocation{Country: "GBR"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.20},
		{ID: "2", MerchantLocation: MerchantLocation{Country: "DEU"}, CustomerLocation: CustomerLocation{Country: "DEU"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.19},
		{ID: "3", MerchantLocation: MerchantLocation{Country: "FRA"}, CustomerLocation: CustomerLocation{Country: "FRA"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.20},
		{ID: "4", MerchantLocation: MerchantLocation{Country: "IRL"}, CustomerLocation: CustomerLocation{Country: "IRL"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.23},
		{ID: "5", MerchantLocation: MerchantLocation{Country: "NLD"}, CustomerLocation: CustomerLocation{Country: "NLD"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.21},
		{ID: "6", MerchantLocation: MerchantLocation{Country: "ESP"}, CustomerLocation: CustomerLocation{Country: "ESP"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.21},
		// US - Sales tax varies by state
		{ID: "7", MerchantLocation: MerchantLocation{Country: "USA", State: "CA"}, CustomerLocation: CustomerLocation{Country: "USA", State: "CA"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.0725},
		{ID: "8", MerchantLocation: MerchantLocation{Country: "USA", State: "NY"}, CustomerLocation: CustomerLocation{Country: "USA", State: "NY"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.08},
		{ID: "9", MerchantLocation: MerchantLocation{Country: "USA", State: "TX"}, CustomerLocation: CustomerLocation{Country: "USA", State: "TX"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.0625},
		// Default for US without state
		{ID: "10", MerchantLocation: MerchantLocation{Country: "USA"}, CustomerLocation: CustomerLocation{Country: "USA"}, IsEkkoTaxLiable: false, CarbonCreditRate: 0.0},
	}

	for _, r := range rates {
		key := r.MerchantLocation.Country + ":" + r.MerchantLocation.State + ":" + r.CustomerLocation.Country + ":" + r.CustomerLocation.State
		repo.rates[key] = r
	}

	return repo
}

// GetTaxRate retrieves the tax rate for a location
func (r *InMemoryRepository) GetTaxRate(_ context.Context, countryCode, state, postalCode string) (*TaxRate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Try exact match first (merchant:customer)
	key := countryCode + ":" + state + ":" + countryCode + ":" + state
	if rate, exists := r.rates[key]; exists {
		return rate, nil
	}

	// Try country only (merchant:customer)
	key = countryCode + "::" + countryCode + ":"
	if rate, exists := r.rates[key]; exists {
		return rate, nil
	}

	// Return no tax as default
	return &TaxRate{
		ID:             "default",
		MerchantLocation: MerchantLocation{Country: countryCode, State: state, Zip: postalCode},
		CustomerLocation: CustomerLocation{Country: countryCode, State: state, Zip: postalCode},
		IsEkkoTaxLiable: false,
		CarbonCreditRate: 0,
	}, nil
}
