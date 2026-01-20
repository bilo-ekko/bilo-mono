package impact_partner

import (
	"errors"
	"sync"
)

// InMemoryRepository implements Repository interface with in-memory storage
type InMemoryRepository struct {
	partners map[string]*Entity
	mu       sync.RWMutex
}

// NewRepository creates a new repository with sample data
func NewRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		partners: make(map[string]*Entity),
	}

	// Seed with sample data
	shortDesc1 := "Leading carbon offset certification body"
	longDesc1 := "Gold Standard is the global leader in carbon offset certification, ensuring high-quality environmental and social impact."
	logo1 := "https://example.com/goldstandard-logo.png"

	shortDesc2 := "Climate action through reforestation"
	longDesc2 := "Ekko Climate focuses on reforestation and ecosystem restoration projects worldwide."
	logo2 := "https://example.com/ekko-logo.png"

	shortDesc3 := "Renewable energy solutions"
	longDesc3 := "Green Energy Co provides renewable energy certificates and carbon reduction projects."
	logo3 := "https://example.com/greenenergy-logo.png"

	repo.partners["partner-1"] = &Entity{
		ID:               "partner-1",
		Name:             "Green Carbon Trust",
		ShortDescription: &shortDesc1,
		LongDescription:  &longDesc1,
		Logo:             &logo1,
		Website:          "https://greencarbontrust.org",
	}
	repo.partners["partner-2"] = &Entity{
		ID:               "partner-2",
		Name:             "Ocean Conservation Fund",
		ShortDescription: &shortDesc2,
		LongDescription:  &longDesc2,
		Logo:             &logo2,
		Website:          "https://oceanconservation.org",
	}
	repo.partners["partner-3"] = &Entity{
		ID:               "partner-3",
		Name:             "Green Energy Co",
		ShortDescription: &shortDesc3,
		LongDescription:  &longDesc3,
		Logo:             &logo3,
		Website:          "https://greenenergy.co",
	}

	return repo
}

// GetAll returns all impact partners
func (r *InMemoryRepository) GetAll() []*Entity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	partners := make([]*Entity, 0, len(r.partners))
	for _, partner := range r.partners {
		partners = append(partners, partner)
	}
	return partners
}

// GetByID returns a specific impact partner by ID
func (r *InMemoryRepository) GetByID(id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	partner, exists := r.partners[id]
	if !exists {
		return nil, errors.New("partner not found")
	}
	return partner, nil
}

// Create adds a new impact partner
func (r *InMemoryRepository) Create(partner *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.partners[partner.ID]; exists {
		return errors.New("partner already exists")
	}
	r.partners[partner.ID] = partner
	return nil
}
