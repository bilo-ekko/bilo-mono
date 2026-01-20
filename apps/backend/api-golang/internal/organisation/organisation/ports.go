// Package organisation defines ports (interfaces) for the organisation sub-domain.
package organisation

import "context"

// Repository defines the port for organisation data access (driven adapter)
type Repository interface {
	GetByID(ctx context.Context, id string) (*Entity, error)
	GetChildren(ctx context.Context, parentID string) ([]*Entity, error)
}

// Service defines the port for organisation business logic (driving port)
type Service interface {
	ValidateOrganisation(ctx context.Context, headerOrgID, bodyOrgID string) (*Entity, error)
	GetOrganisation(ctx context.Context, id string) (*Entity, error)
}
