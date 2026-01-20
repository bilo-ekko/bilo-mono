// Package country defines ports for the country sub-domain.
package country

import "context"

// Repository defines the port for country data access
type Repository interface {
	GetByCode(ctx context.Context, code string) (*Entity, error)
	GetByID(ctx context.Context, id string) (*Entity, error)
}

// Service defines the port for country business logic
type Service interface {
	GetCountryByCode(ctx context.Context, code string) (*Entity, error)
}
