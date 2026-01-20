package organisation

import (
	"context"
	"sync"

	"api-golang/internal/shared/errors"
	"api-golang/internal/shared/types"
)

const domainName = "organisation"

// InMemoryRepository implements Repository interface with in-memory storage
type InMemoryRepository struct {
	organisations map[string]*Entity
	mu            sync.RWMutex
}

// NewInMemoryRepository creates a new repository with sample data
func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		organisations: make(map[string]*Entity),
	}

	// Seed with sample data
	parentID := "org-parent-1"
	mcc6011 := "6011" // Banks
	serviceFee := 0.05

	repo.organisations["org-parent-1"] = &Entity{
		OrganisationID:       "org-parent-1",
		ParentOrganisationID: nil,
		TradingName:          "Acme Bank",
		LegalName:            "Acme Bank PLC",
		MCC:                  &mcc6011,
		CurrencyCode:         "EUR",
		Address: types.Address{
			Line1:       "100 Bank Street",
			City:        "London",
			PostalCode:  "EC1A 1BB",
			CountryCode: "GBR",
		},
		ServiceFeePercentage: serviceFee,
		Status: types.OrganisationStatus{
			Value: "active",
		},
		ImpactPartners: []types.OrganisationImpactPartner{
			{ID: "partner-1", Name: "Green Carbon Trust"},
			{ID: "partner-2", Name: "Ocean Conservation Fund"},
		},
		CalculationTypes: []string{"carbon", "nature"},
	}

	repo.organisations["org-child-1"] = &Entity{
		OrganisationID:       "org-child-1",
		ParentOrganisationID: &parentID,
		TradingName:          "Acme Bank Ireland",
		LegalName:            "Acme Bank Ireland Ltd",
		MCC:                  &mcc6011,
		CurrencyCode:         "EUR",
		Address: types.Address{
			Line1:       "50 Finance Street",
			City:        "Dublin",
			PostalCode:  "D02",
			CountryCode: "IRL",
		},
		ServiceFeePercentage: serviceFee,
		Status: types.OrganisationStatus{
			Value: "active",
		},
		ImpactPartners: []types.OrganisationImpactPartner{
			{ID: "partner-1", Name: "Green Carbon Trust"},
		},
		CalculationTypes: []string{"carbon"},
	}

	repo.organisations["org-child-2"] = &Entity{
		OrganisationID:       "org-child-2",
		ParentOrganisationID: &parentID,
		TradingName:          "Acme Bank Germany",
		LegalName:            "Acme Bank Germany GmbH",
		MCC:                  &mcc6011,
		CurrencyCode:         "EUR",
		Address: types.Address{
			Line1:       "25 Bankstraße",
			City:        "Berlin",
			PostalCode:  "10115",
			CountryCode: "DEU",
		},
		ServiceFeePercentage: serviceFee,
		Status: types.OrganisationStatus{
			Value: "active",
		},
		ImpactPartners: []types.OrganisationImpactPartner{
			{ID: "partner-1", Name: "Green Carbon Trust"},
			{ID: "partner-2", Name: "Ocean Conservation Fund"},
		},
		CalculationTypes: []string{"carbon", "nature"},
	}

	return repo
}

// GetByID retrieves an organisation by ID
func (r *InMemoryRepository) GetByID(_ context.Context, id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	org, exists := r.organisations[id]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "organisation not found")
	}
	return org, nil
}

// GetChildren retrieves all child organisations of a parent
func (r *InMemoryRepository) GetChildren(_ context.Context, parentID string) ([]*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	children := make([]*Entity, 0)
	for _, org := range r.organisations {
		if org.ParentOrganisationID != nil && *org.ParentOrganisationID == parentID {
			children = append(children, org)
		}
	}
	return children, nil
}
