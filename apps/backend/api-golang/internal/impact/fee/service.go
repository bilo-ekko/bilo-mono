package fee

import (
	"context"
	"fmt"
	"math"
)

// DefaultService implements the Service interface
type DefaultService struct {
	repo Repository
}

// NewService creates a new fee service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

// CalculateServiceFee calculates the service fee for a compensation amount
func (s *DefaultService) CalculateServiceFee(ctx context.Context, organisationID string, compensationAmount float64) (*FeeResult, error) {
	config, err := s.repo.GetFeeConfig(ctx, organisationID)
	if err != nil {
		return nil, fmt.Errorf("getting fee config: %w", err)
	}

	// Calculate fee
	feeAmount := compensationAmount * config.FeePercentage

	// Apply minimum
	if feeAmount < config.MinimumFee {
		feeAmount = config.MinimumFee
	}

	// Apply maximum (if set)
	if config.MaximumFee > 0 && feeAmount > config.MaximumFee {
		feeAmount = config.MaximumFee
	}

	// Round to 2 decimal places
	feeAmount = math.Round(feeAmount*100) / 100

	return &FeeResult{
		CompensationAmount: compensationAmount,
		FeeAmount:          feeAmount,
		FeePercentage:      config.FeePercentage,
	}, nil
}
