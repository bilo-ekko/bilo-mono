package impact_project

import "github.com/bilo-mono/packages/common/service"

// Service handles business logic for impact projects
type Service struct {
	service.BaseService[*Repository]
}

// NewService creates a new service
func NewService(repo *Repository) *Service {
	return &Service{
		BaseService: service.NewBaseService(repo),
	}
}

// GetAllProjects returns all impact projects
func (s *Service) GetAllProjects() []*Entity {
	return s.Repo.GetAll()
}

// GetProjectByID returns a specific project by ID
func (s *Service) GetProjectByID(id string) (*Entity, error) {
	return s.Repo.GetByID(id)
}

// GetProjectsByPartnerID returns all projects for a specific partner
func (s *Service) GetProjectsByPartnerID(partnerID string) []*Entity {
	return s.Repo.GetByPartnerID(partnerID)
}

// CreateProject creates a new impact project
func (s *Service) CreateProject(project *Entity) error {
	// Add business logic/validation here if needed
	return s.Repo.Create(project)
}
