package impact_project

import (
	"errors"
	"sync"

	"api-golang/internal/shared/types"
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
	shortDesc1 := "Preserving the world's largest rainforest"
	longDesc1 := "This project focuses on protecting critical Amazon rainforest areas through community engagement and sustainable practices."
	image1 := "https://example.com/amazon-forest.jpg"
	theme1 := ProjectThemeLandUse
	region1 := "South America"

	shortDesc2 := "Clean energy for rural communities"
	longDesc2 := "Large-scale solar installation providing renewable energy to underserved communities in India."
	image2 := "https://example.com/solar-india.jpg"
	theme2 := ProjectThemeClimateStress
	region2 := "Asia"

	shortDesc3 := "Offshore wind power generation"
	longDesc3 := "State-of-the-art wind farm off the coast of Denmark generating clean electricity."
	image3 := "https://example.com/wind-denmark.jpg"
	region3 := "Europe"

	shortDesc4 := "Coastal ecosystem restoration"
	longDesc4 := "Restoring mangrove ecosystems to protect coastlines and sequester carbon."
	image4 := "https://example.com/mangrove.jpg"
	theme4 := ProjectThemeWaterUse
	region4 := "Southeast Asia"

	repo.projects["project-1"] = &Entity{
		ID:               "project-1",
		Name:             "Amazon Rainforest Conservation",
		ImpactPartnerID:  "partner-1",
		ShortDescription: &shortDesc1,
		LongDescription:  &longDesc1,
		Image:            &image1,
		Type:             ProjectTypeCarbonCredits,
		Theme:            &theme1,
		Location: types.Location{
			Country: "Brazil",
			Region:  &region1,
		},
		Unit: types.ProjectUnit{
			Type:   "tCO2e",
			Symbol: "t",
		},
	}

	repo.projects["project-2"] = &Entity{
		ID:               "project-2",
		Name:             "Solar Farm Initiative India",
		ImpactPartnerID:  "partner-2",
		ShortDescription: &shortDesc2,
		LongDescription:  &longDesc2,
		Image:            &image2,
		Type:             ProjectTypeCarbonCredits,
		Theme:            &theme2,
		Location: types.Location{
			Country: "India",
			Region:  &region2,
		},
		Unit: types.ProjectUnit{
			Type:   "tCO2e",
			Symbol: "t",
		},
	}

	repo.projects["project-3"] = &Entity{
		ID:               "project-3",
		Name:             "Wind Energy Project Denmark",
		ImpactPartnerID:  "partner-2",
		ShortDescription: &shortDesc3,
		LongDescription:  &longDesc3,
		Image:            &image3,
		Type:             ProjectTypeCarbonCredits,
		Theme:            nil, // Optional
		Location: types.Location{
			Country: "Denmark",
			Region:  &region3,
		},
		Unit: types.ProjectUnit{
			Type:   "tCO2e",
			Symbol: "t",
		},
	}

	repo.projects["project-4"] = &Entity{
		ID:               "project-4",
		Name:             "Mangrove Restoration Program",
		ImpactPartnerID:  "partner-1",
		ShortDescription: &shortDesc4,
		LongDescription:  &longDesc4,
		Image:            &image4,
		Type:             ProjectTypeNatureCredits,
		Theme:            &theme4,
		Location: types.Location{
			Country: "Vietnam",
			Region:  &region4,
		},
		Unit: types.ProjectUnit{
			Type:   "hectares",
			Symbol: "ha",
		},
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
		if project.ImpactPartnerID == partnerID {
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
