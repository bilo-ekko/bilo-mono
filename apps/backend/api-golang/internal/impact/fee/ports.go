// Package fee defines ports for the fee sub-domain.
package fee

import "context"

// Repository defines the port for fee config data access
type Repository interface {
	GetFeeConfig(ctx context.Context, organisationID string) (*FeeConfig, error)
}

// Service defines the port for fee calculation business logic
type Service interface {
	CalculateServiceFee(ctx context.Context, organisationID string, compensationAmount float64) (*FeeResult, error)
}
