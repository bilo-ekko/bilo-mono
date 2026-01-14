package impact_partner

// Entity represents an organization or entity that provides climate impact services
type Entity struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"` // e.g., "carbon-offset", "reforestation", "renewable-energy"
}
