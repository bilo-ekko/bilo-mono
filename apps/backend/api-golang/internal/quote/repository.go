package quote

import (
	"context"
	"sync"

	"api-golang/internal/shared/errors"
)

const domainName = "quote"

// InMemoryRepository implements Repository interface
type InMemoryRepository struct {
	quotes map[string]*Entity
	mu     sync.RWMutex
}

// NewInMemoryRepository creates a new repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		quotes: make(map[string]*Entity),
	}
}

// Create stores a new quote
func (r *InMemoryRepository) Create(_ context.Context, quote *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.quotes[quote.ID]; exists {
		return errors.NewValidationError(domainName, "quote already exists")
	}

	r.quotes[quote.ID] = quote
	return nil
}

// GetByID retrieves a quote by ID
func (r *InMemoryRepository) GetByID(_ context.Context, id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	quote, exists := r.quotes[id]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "quote not found")
	}
	return quote, nil
}

// Update updates an existing quote
func (r *InMemoryRepository) Update(_ context.Context, quote *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.quotes[quote.ID]; !exists {
		return errors.NewNotFoundError(domainName, "quote not found")
	}

	r.quotes[quote.ID] = quote
	return nil
}
