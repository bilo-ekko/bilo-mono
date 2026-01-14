package impact_partner

import (
	"errors"
	"sync"
)

// Repository handles data access for ImpactPartners
type Repository struct {
	partners map[string]*Entity
	mu       sync.RWMutex
}

// NewRepository creates a new repository with sample data
func NewRepository() *Repository {
	repo := &Repository{
		partners: make(map[string]*Entity),
	}
	
	// Seed with sample data
	repo.partners["1"] = &Entity{
		ID:       "1",
		Name:     "GoldStandard",
		Category: "carbon-offset",
	}
	repo.partners["2"] = &Entity{
		ID:       "2",
		Name:     "Ekko Climate",
		Category: "reforestation",
	}
	repo.partners["3"] = &Entity{
		ID:       "3",
		Name:     "Green Energy Co",
		Category: "renewable-energy",
	}
	
	return repo
}

// GetAll returns all impact partners
func (r *Repository) GetAll() []*Entity {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	partners := make([]*Entity, 0, len(r.partners))
	for _, partner := range r.partners {
		partners = append(partners, partner)
	}
	return partners
}

// GetByID returns a specific impact partner by ID
func (r *Repository) GetByID(id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	
	partner, exists := r.partners[id]
	if !exists {
		return nil, errors.New("partner not found")
	}
	return partner, nil
}

// Create adds a new impact partner
func (r *Repository) Create(partner *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	
	if _, exists := r.partners[partner.ID]; exists {
		return errors.New("partner already exists")
	}
	r.partners[partner.ID] = partner
	return nil
}
