package impact_project

// Service handles business logic for impact projects
type Service struct {
	repo *Repository
}

// NewService creates a new service
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAllProjects returns all impact projects
func (s *Service) GetAllProjects() []*Entity {
	return s.repo.GetAll()
}

// GetProjectByID returns a specific project by ID
func (s *Service) GetProjectByID(id string) (*Entity, error) {
	return s.repo.GetByID(id)
}

// GetProjectsByPartnerID returns all projects for a specific partner
func (s *Service) GetProjectsByPartnerID(partnerID string) []*Entity {
	return s.repo.GetByPartnerID(partnerID)
}

// CreateProject creates a new impact project
func (s *Service) CreateProject(project *Entity) error {
	// Add business logic/validation here if needed
	return s.repo.Create(project)
}
