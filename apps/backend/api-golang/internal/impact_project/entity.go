package impact_project

// Entity represents a specific climate impact project
type Entity struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`  // e.g., "solar", "wind", "forest-conservation"
	PartnerID string `json:"partnerId"` // Foreign key to ImpactPartner
}
