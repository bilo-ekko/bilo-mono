// Package quote defines ports for the quote domain.
package quote

import (
	"context"
)

// Repository defines the port for quote data access
type Repository interface {
	Create(ctx context.Context, quote *Entity) error
	GetByID(ctx context.Context, id string) (*Entity, error)
	Update(ctx context.Context, quote *Entity) error
}

// Service defines the port for quote business logic
type Service interface {
	CreateQuote(ctx context.Context, req *CreateQuoteRequest) (*CreateQuoteResponse, error)
	GetQuote(ctx context.Context, id string) (*Entity, error)
}
