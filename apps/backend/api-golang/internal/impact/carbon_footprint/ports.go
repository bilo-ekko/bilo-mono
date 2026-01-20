// Package carbonfootprint defines ports for the carbon footprint sub-domain.
package carbonfootprint

import "context"

// CalculateInput contains parameters for carbon footprint calculation
type CalculateInput struct {
	TransactionID  string
	OrganisationID string
	CustomerID     string
	AmountEUR      float64 // Amount in EUR (already converted)
	MCC            string
	CountryID      string
}

// FactorRepository defines the port for carbon factor data access
type FactorRepository interface {
	GetFactor(ctx context.Context, mcc, countryID string) (*CarbonFactor, error)
}

// FootprintRepository defines the port for footprint data access
type FootprintRepository interface {
	Create(ctx context.Context, footprint *Footprint) error
	GetByID(ctx context.Context, id string) (*Footprint, error)
}

// Service defines the port for carbon footprint business logic
type Service interface {
	Calculate(ctx context.Context, input CalculateInput) (*Footprint, error)
}
