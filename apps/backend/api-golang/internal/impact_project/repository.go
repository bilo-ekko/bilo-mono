package impact_project

import (
	"errors"
	"sync"
)

// Repository handles data access for ImpactProjects
type Repository struct {
	projects map[string]*Entity
	mu       sync.RWMutex
}

// NewRepository creates a new repository with sample data
func NewRepository() *Repository {
	repo := &Repository{
		projects: make(map[string]*Entity),
	}

	// Seed with sample data
	repo.projects["1"] = &Entity{
		ID:        "1",
		Name:      "Amazon Rainforest Conservation",
		Category:  "forest-conservation",
		PartnerID: "2",
	}
	repo.projects["2"] = &Entity{
		ID:        "2",
		Name:      "Solar Farm Initiative India",
		Category:  "solar",
		PartnerID: "3",
	}
	repo.projects["3"] = &Entity{
		ID:        "3",
		Name:      "Wind Energy Project Denmark",
		Category:  "wind",
		PartnerID: "3",
	}
	repo.projects["4"] = &Entity{
		ID:        "4",
		Name:      "Mangrove Restoration Program",
		Category:  "reforestation",
		PartnerID: "2",
	}

	return repo
}

// GetAll returns all impact projects
func (r *Repository) GetAll() []*Entity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	projects := make([]*Entity, 0, len(r.projects))
	for _, project := range r.projects {
		projects = append(projects, project)
	}
	return projects
}

// GetByID returns a specific impact project by ID
func (r *Repository) GetByID(id string) (*Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	project, exists := r.projects[id]
	if !exists {
		return nil, errors.New("project not found")
	}
	return project, nil
}

// GetByPartnerID returns all projects for a specific partner
func (r *Repository) GetByPartnerID(partnerID string) []*Entity {
	r.mu.RLock()
	defer r.mu.RUnlock()

	projects := make([]*Entity, 0)
	for _, project := range r.projects {
		if project.PartnerID == partnerID {
			projects = append(projects, project)
		}
	}
	return projects
}

// Create adds a new impact project
func (r *Repository) Create(project *Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[project.ID]; exists {
		return errors.New("project already exists")
	}
	r.projects[project.ID] = project
	return nil
}
