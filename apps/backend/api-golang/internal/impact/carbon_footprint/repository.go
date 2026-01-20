package carbonfootprint

import (
	"context"
	"sync"

	"api-golang/internal/shared/errors"
)

const domainName = "carbon_footprint"

// InMemoryFactorRepository implements FactorRepository interface
type InMemoryFactorRepository struct {
	factors map[string]*CarbonFactor // key: mcc:countryId
	mu      sync.RWMutex
}

// NewInMemoryFactorRepository creates a new repository with sample data
func NewInMemoryFactorRepository() *InMemoryFactorRepository {
	repo := &InMemoryFactorRepository{
		factors: make(map[string]*CarbonFactor),
	}

	// Default factor for any MCC/country combination
	defaultFactor := &CarbonFactor{
		ID:          "default",
		MCC:         "*",
		CountryID:   "*",
		Factor:      0.23, // kg CO2e per EUR (average)
		Description: "Default carbon factor",
	}
	repo.factors["*:*"] = defaultFactor

	// Specific factors by MCC
	factors := []*CarbonFactor{
		// Airlines (high carbon)
		{ID: "1", MCC: "4511", CountryID: "*", Factor: 1.2, Description: "Airlines"},
		// Restaurants (medium)
		{ID: "2", MCC: "5812", CountryID: "*", Factor: 0.35, Description: "Restaurants"},
		// Gas stations (high)
		{ID: "3", MCC: "5541", CountryID: "*", Factor: 2.5, Description: "Gas stations"},
		// Grocery stores (lower)
		{ID: "4", MCC: "5411", CountryID: "*", Factor: 0.18, Description: "Grocery stores"},
		// Banks (low)
		{ID: "5", MCC: "6011", CountryID: "*", Factor: 0.05, Description: "Banks/Financial"},
		// Electronics
		{ID: "6", MCC: "5732", CountryID: "*", Factor: 0.45, Description: "Electronics"},
		// Clothing
		{ID: "7", MCC: "5651", CountryID: "*", Factor: 0.40, Description: "Clothing"},
	}

	for _, f := range factors {
		key := f.MCC + ":" + f.CountryID
		repo.factors[key] = f
	}

	return repo
}

// GetFactor retrieves a carbon factor for MCC and country
func (r *InMemoryFactorRepository) GetFactor(_ context.Context, mcc, countryID string) (*CarbonFactor, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Try exact match first
	key := mcc + ":" + countryID
	if factor, exists := r.factors[key]; exists {
		return factor, nil
	}

	// Try MCC with wildcard country
	key = mcc + ":*"
	if factor, exists := r.factors[key]; exists {
		return factor, nil
	}

	// Fall back to default
	if factor, exists := r.factors["*:*"]; exists {
		return factor, nil
	}

	return nil, errors.NewNotFoundError(domainName, "no carbon factor found")
}

// InMemoryFootprintRepository implements FootprintRepository interface
type InMemoryFootprintRepository struct {
	footprints map[string]*Footprint
	mu         sync.RWMutex
}

// NewInMemoryFootprintRepository creates a new repository
func NewInMemoryFootprintRepository() *InMemoryFootprintRepository {
	return &InMemoryFootprintRepository{
		footprints: make(map[string]*Footprint),
	}
}

// Create stores a new footprint
func (r *InMemoryFootprintRepository) Create(_ context.Context, footprint *Footprint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.footprints[footprint.ID] = footprint
	return nil
}

// GetByID retrieves a footprint by ID
func (r *InMemoryFootprintRepository) GetByID(_ context.Context, id string) (*Footprint, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	footprint, exists := r.footprints[id]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "footprint not found")
	}
	return footprint, nil
}
