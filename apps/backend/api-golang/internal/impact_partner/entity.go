// Package impact_partner handles impact partner business logic.
package impact_partner

// Entity represents an impact partner (matches Ekko API v3 schema)
// See: https://docs.ekko.earth/v3/reference/get_impact-partners
type Entity struct {
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	LongDescription  *string `json:"longDescription,omitempty"`
	Logo             *string `json:"logo,omitempty"` // URL to logo image
	Website          string  `json:"website"`        // URI
}
