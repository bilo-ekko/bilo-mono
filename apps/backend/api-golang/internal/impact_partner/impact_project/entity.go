// Package impact_project handles impact project business logic.
package impact_project

import "api-golang/internal/shared/types"

// ProjectType represents the type of impact project
type ProjectType string

const (
	ProjectTypeNatureCredits  ProjectType = "natureCredits"
	ProjectTypeCarbonCredits  ProjectType = "carbonCredits"
	ProjectTypeContribution   ProjectType = "contribution"
)

// ProjectTheme represents the environmental theme of a project
type ProjectTheme string

const (
	ProjectThemePollution     ProjectTheme = "pollution"
	ProjectThemeClimateStress ProjectTheme = "climateStress"
	ProjectThemeLandUse       ProjectTheme = "landUse"
	ProjectThemeWaterUse      ProjectTheme = "waterUse"
)

// Entity represents an impact project (matches Ekko API v3 schema)
// See: https://docs.ekko.earth/v3/reference/get_impact-partners-projects
type Entity struct {
	ID               string          `json:"id"`
	Name             string          `json:"name"`
	ImpactPartnerID  string          `json:"impactPartnerId"`
	ShortDescription *string         `json:"shortDescription,omitempty"`
	LongDescription  *string         `json:"longDescription,omitempty"`
	Image            *string         `json:"image,omitempty"` // URL to project image
	Type             ProjectType     `json:"type"`            // natureCredits, carbonCredits, contribution
	Subtype          *string         `json:"subtype,omitempty"`
	Location         types.Location  `json:"location"`
	Theme            *ProjectTheme   `json:"theme,omitempty"` // pollution, climateStress, landUse, waterUse
	SDGs             []int           `json:"sdg,omitempty"`   // Sustainable Development Goals (numbers 1-17)
	Unit             types.ProjectUnit `json:"unit"`
	Status           string          `json:"status"`         // active/inactive
	TaxType          *string         `json:"taxType,omitempty"`
}

// IsCarbonProject checks if this is a carbon credits project
func (e *Entity) IsCarbonProject() bool {
	return e.Type == ProjectTypeCarbonCredits
}

// IsNatureProject checks if this is a nature credits project
func (e *Entity) IsNatureProject() bool {
	return e.Type == ProjectTypeNatureCredits
}
