package carbonfootprint

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// DefaultService implements the Service interface
type DefaultService struct {
	factorRepo    FactorRepository
	footprintRepo FootprintRepository
}

// NewService creates a new carbon footprint service
func NewService(factorRepo FactorRepository, footprintRepo FootprintRepository) *DefaultService {
	return &DefaultService{
		factorRepo:    factorRepo,
		footprintRepo: footprintRepo,
	}
}

// Calculate calculates the carbon footprint for a transaction
func (s *DefaultService) Calculate(ctx context.Context, input CalculateInput) (*Footprint, error) {
	// Get the carbon factor
	factor, err := s.factorRepo.GetFactor(ctx, input.MCC, input.CountryID)
	if err != nil {
		return nil, fmt.Errorf("getting carbon factor: %w", err)
	}

	// Calculate carbon footprint: amount * factor
	carbonKg := input.AmountEUR * factor.Factor
	carbonGrams := carbonKg * 1000.0
	carbonOunces := carbonKg * 35.274

	// Create footprint record
	footprint := &Footprint{
		ID:               uuid.New().String(),
		CustomerID:       input.CustomerID,
		OrganisationID:   input.OrganisationID,
		MCC:              input.MCC,
		MerchantCountry: input.CountryID, // Using CountryID as merchant country for now
		Amount:           input.AmountEUR,
		Currency:         "EUR",
		CarbonCo2eGrams:  carbonGrams,
		CarbonCo2eOunces: carbonOunces,
		Factor:           factor.Factor,
		CalculationMethod: "MCC-based calculation",
		CreatedAt:        time.Now(),
	}

	// Store the footprint
	if err := s.footprintRepo.Create(ctx, footprint); err != nil {
		return nil, fmt.Errorf("storing footprint: %w", err)
	}

	return footprint, nil
}
