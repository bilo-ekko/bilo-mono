package customer

import (
	"context"
	"sync"

	"api-golang/internal/shared/errors"
)

const domainName = "customer"

// InMemoryRepository implements Repository interface with in-memory storage
type InMemoryRepository struct {
	customers     map[string]*Entity
	referenceKeys map[string]string // map[orgID:reference]customerID
	mu            sync.RWMutex
}

// NewInMemoryRepository creates a new repository
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		customers:     make(map[string]*Entity),
		referenceKeys: make(map[string]string),
	}
}

// GetByID retrieves a customer by ID
func (r *InMemoryRepository) GetByID(_ context.Context, id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	customer, exists := r.customers[id]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "customer not found")
	}
	return customer, nil
}

// GetByReference retrieves a customer by organisation ID and reference
func (r *InMemoryRepository) GetByReference(_ context.Context, organisationID, reference string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := organisationID + ":" + reference
	customerID, exists := r.referenceKeys[key]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "customer not found")
	}

	customer, exists := r.customers[customerID]
	if !exists {
		return nil, errors.NewNotFoundError(domainName, "customer not found")
	}
	return customer, nil
}

// Create adds a new customer
func (r *InMemoryRepository) Create(_ context.Context, customer *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.customers[customer.ID]; exists {
		return errors.NewValidationError(domainName, "customer already exists")
	}

	r.customers[customer.ID] = customer
	if customer.Reference != "" {
		key := customer.OrganisationID + ":" + customer.Reference
		r.referenceKeys[key] = customer.ID
	}
	return nil
}
