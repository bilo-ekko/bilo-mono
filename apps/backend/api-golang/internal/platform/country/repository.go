package country

import (
	"context"
	"fmt"
	"sync"

	"api-golang/internal/shared/errors"
)

const domainName = "country"

// InMemoryRepository implements Repository interface
type InMemoryRepository struct {
	countriesByCode map[string]*Entity // Supports both 2-letter and 3-letter codes
	countriesByID   map[string]*Entity
	mu              sync.RWMutex
}

// NewInMemoryRepository creates a new repository with sample data
func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		countriesByCode: make(map[string]*Entity),
		countriesByID:   make(map[string]*Entity),
	}

	// Seed with sample data - include both ISO2 and ISO3 codes
	countries := []*Entity{
		{ID: "1", Code: "GB", ISO2Code: "GB", ISO3Code: "GBR", Name: "United Kingdom", Currency: "GBP", TaxRate: 0.20, IsEU: false},
		{ID: "2", Code: "IE", ISO2Code: "IE", ISO3Code: "IRL", Name: "Ireland", Currency: "EUR", TaxRate: 0.23, IsEU: true},
		{ID: "3", Code: "DE", ISO2Code: "DE", ISO3Code: "DEU", Name: "Germany", Currency: "EUR", TaxRate: 0.19, IsEU: true},
		{ID: "4", Code: "FR", ISO2Code: "FR", ISO3Code: "FRA", Name: "France", Currency: "EUR", TaxRate: 0.20, IsEU: true},
		{ID: "5", Code: "US", ISO2Code: "US", ISO3Code: "USA", Name: "United States", Currency: "USD", TaxRate: 0.0, IsEU: false},
		{ID: "6", Code: "NL", ISO2Code: "NL", ISO3Code: "NLD", Name: "Netherlands", Currency: "EUR", TaxRate: 0.21, IsEU: true},
		{ID: "7", Code: "ES", ISO2Code: "ES", ISO3Code: "ESP", Name: "Spain", Currency: "EUR", TaxRate: 0.21, IsEU: true},
	}

	for _, c := range countries {
		// Index by both 2-letter and 3-letter codes
		repo.countriesByCode[c.ISO2Code] = c
		repo.countriesByCode[c.ISO3Code] = c
		repo.countriesByID[c.ID] = c
	}

	return repo
}

// GetByCode retrieves a country by ISO code (supports both 2-letter and 3-letter codes)
func (r *InMemoryRepository) GetByCode(_ context.Context, code string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	country, exists := r.countriesByCode[code]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, fmt.Sprintf("country not found for code: %s", code))
	}
	return country, nil
}

// GetByID retrieves a country by ID
func (r *InMemoryRepository) GetByID(_ context.Context, id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	country, exists := r.countriesByID[id]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "country not found")
	}
	return country, nil
}
