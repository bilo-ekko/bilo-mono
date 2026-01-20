package fee

import (
	"context"
	"sync"
)

// InMemoryRepository implements Repository interface
type InMemoryRepository struct {
	configs map[string]*FeeConfig
	mu      sync.RWMutex
}

// NewInMemoryRepository creates a new repository with sample data
func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		configs: make(map[string]*FeeConfig),
	}

	// Seed with sample data
	configs := []*FeeConfig{
		{OrganisationID: "org-parent-1", FeePercentage: 0.10, MinimumFee: 0.01, MaximumFee: 10.00},
		{OrganisationID: "org-child-1", FeePercentage: 0.08, MinimumFee: 0.01, MaximumFee: 5.00},
		{OrganisationID: "org-child-2", FeePercentage: 0.12, MinimumFee: 0.02, MaximumFee: 15.00},
	}

	for _, c := range configs {
		repo.configs[c.OrganisationID] = c
	}

	return repo
}

// GetFeeConfig retrieves the fee configuration for an organisation
func (r *InMemoryRepository) GetFeeConfig(_ context.Context, organisationID string) (*FeeConfig, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	config, exists := r.configs[organisationID]
	if !exists {
		// Return default config
		return &FeeConfig{
			OrganisationID: organisationID,
			FeePercentage:  0.10, // 10% default
			MinimumFee:     0.01,
			MaximumFee:     0,
		}, nil
	}
	return config, nil
}
