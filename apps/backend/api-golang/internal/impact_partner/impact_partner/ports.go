// Package impact_partner defines ports (interfaces) for the impact partner domain.
package impact_partner

import "context"

// Repository defines the port for impact partner data access (driven adapter)
type Repository interface {
	GetAll() []*Entity
	GetByID(id string) (*Entity, error)
	Create(partner *Entity) error
}

// Service defines the port for impact partner business logic (driving port)
type Service interface {
	GetAllPartners(ctx context.Context) ([]*Entity, error)
	GetPartnerByID(ctx context.Context, id string) (*Entity, error)
	CreatePartner(ctx context.Context, partner *Entity) error
}
