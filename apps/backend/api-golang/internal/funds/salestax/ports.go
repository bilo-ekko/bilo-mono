// Package salestax defines ports for the sales tax sub-domain.
package salestax

import "context"

// Repository defines the port for tax rate data access
type Repository interface {
	GetTaxRate(ctx context.Context, countryCode, state, postalCode string) (*TaxRate, error)
}

// Service defines the port for sales tax calculation business logic
type Service interface {
	CalculateSalesTax(ctx context.Context, input TaxCalculationInput) (*TaxResult, error)
}
