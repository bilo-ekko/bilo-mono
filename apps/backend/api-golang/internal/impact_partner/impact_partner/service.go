package impact_partner

import (
	"context"

	"github.com/bilo-mono/packages/common/service"
)

// DefaultService implements the Service interface defined in ports.go
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// GetAllPartners returns all impact partners
// Implements the Service interface
func (s *DefaultService) GetAllPartners(ctx context.Context) ([]*Entity, error) {
	return s.Repo.GetAll(), nil
}

// GetPartnerByID returns a specific partner by ID
// Implements the Service interface
func (s *DefaultService) GetPartnerByID(ctx context.Context, id string) (*Entity, error) {
	return s.Repo.GetByID(id)
}

// CreatePartner creates a new impact partner
// Implements the Service interface
func (s *DefaultService) CreatePartner(ctx context.Context, partner *Entity) error {
	// Add business logic/validation here if needed
	return s.Repo.Create(partner)
}
