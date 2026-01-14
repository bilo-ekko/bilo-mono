package impact_partner

// Service handles business logic for impact partners
type Service struct {
	repo *Repository
}

// NewService creates a new service
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAllPartners returns all impact partners
func (s *Service) GetAllPartners() []*Entity {
	return s.repo.GetAll()
}

// GetPartnerByID returns a specific partner by ID
func (s *Service) GetPartnerByID(id string) (*Entity, error) {
	return s.repo.GetByID(id)
}

// CreatePartner creates a new impact partner
func (s *Service) CreatePartner(partner *Entity) error {
	// Add business logic/validation here if needed
	return s.repo.Create(partner)
}
