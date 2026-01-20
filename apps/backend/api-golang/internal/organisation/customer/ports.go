// Package customer defines ports (interfaces) for the customer sub-domain.
package customer

import "context"

// CreateCustomerInput represents input for creating/finding a customer
type CreateCustomerInput struct {
	OrganisationID string
	Reference      string // Client's internal reference for the customer
	Email          *string
	Name           *string
	CountryCode    string
	State          *string
	PostalCode     *string
	City           *string
}

// Repository defines the port for customer data access (driven adapter)
type Repository interface {
	GetByID(ctx context.Context, id string) (*Entity, error)
	GetByReference(ctx context.Context, organisationID, reference string) (*Entity, error)
	Create(ctx context.Context, customer *Entity) error
}

// Service defines the port for customer business logic (driving port)
type Service interface {
	GetOrCreateCustomer(ctx context.Context, input CreateCustomerInput) (*Entity, error)
}
