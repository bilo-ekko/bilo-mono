package salestax

import (
	"context"
	"fmt"
	"math"

	"github.com/bilo-mono/packages/common/service"
)

// DefaultService implements the Service interface
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new sales tax service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// CalculateSalesTax calculates sales tax based on merchant and customer locations
// Tax jurisdiction is typically based on the customer's location for digital services
func (s *DefaultService) CalculateSalesTax(ctx context.Context, input TaxCalculationInput) (*TaxResult, error) {
	// For B2C digital services, use customer location for tax calculation
	taxRate, err := s.Repo.GetTaxRate(ctx, input.CustomerCountry, input.CustomerState, input.CustomerPostalCode)
	if err != nil {
		return nil, fmt.Errorf("getting tax rate: %w", err)
	}

	// Use carbon credit rate (for carbon offset quotes)
	taxRateValue := taxRate.CarbonCreditRate
	if taxRateValue == 0 {
		// Fallback to service fee rate if carbon credit rate not available
		taxRateValue = taxRate.ServiceFeeRate
	}

	// If no applicable tax, return zero
	if taxRateValue == 0 {
		return &TaxResult{
			TaxableAmount: input.Amount,
			TaxRate:       0,
			TaxAmount:     0,
			TaxName:       "N/A",
			IsApplicable:  false,
		}, nil
	}

	// Calculate tax amount
	taxAmount := input.Amount * taxRateValue

	// Round to 2 decimal places
	taxAmount = math.Round(taxAmount*100) / 100

	return &TaxResult{
		TaxableAmount: input.Amount,
		TaxRate:       taxRateValue,
		TaxAmount:     taxAmount,
		TaxName:       "Sales Tax",
		IsApplicable:  true,
	}, nil
}
