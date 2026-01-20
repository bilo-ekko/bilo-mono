package impact_partner

import "context"

// DefaultService implements the Service interface defined in ports.go
type DefaultService struct {
	repo Repository
}

// NewService creates a new service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		repo: repo,
	}
}

// GetAllPartners returns all impact partners
// Implements the Service interface
func (s *DefaultService) GetAllPartners(ctx context.Context) ([]*Entity, error) {
	return s.repo.GetAll(), nil
}

// GetPartnerByID returns a specific partner by ID
// Implements the Service interface
func (s *DefaultService) GetPartnerByID(ctx context.Context, id string) (*Entity, error) {
	return s.repo.GetByID(id)
}

// CreatePartner creates a new impact partner
// Implements the Service interface
func (s *DefaultService) CreatePartner(ctx context.Context, partner *Entity) error {
	// Add business logic/validation here if needed
	return s.repo.Create(partner)
}
