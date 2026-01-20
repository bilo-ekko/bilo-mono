// Package country handles country-related lookups.
package country

// Entity represents a country
// Matches Country data model
// See: https://www.notion.so/ekko-earth/Currency-and-country-2b7f93807de480d1a1accf1400743413
type Entity struct {
	ISO3Code   string `json:"iso3Code"`           // Required - ISO 3166-1 alpha-3
	Name       string `json:"name"`               // Required
	ISO2Code   string `json:"iso2Code,omitempty"` // ISO 3166-1 alpha-2
	Region     string `json:"region,omitempty"`
	Continent  string `json:"continent,omitempty"`
	Status     string `json:"status,omitempty"` // active, inactive, etc.
	ISONumeric string `json:"isoNumeric,omitempty"`

	// Legacy fields for backward compatibility
	ID       string  `json:"id,omitempty"`       // Can map to ISO3Code
	Code     string  `json:"code,omitempty"`     // Can map to ISO2Code or ISO3Code
	Currency string  `json:"currency,omitempty"` // Not in data model but useful
	TaxRate  float64 `json:"taxRate,omitempty"`  // Not in data model but useful
	IsEU     bool    `json:"isEU,omitempty"`     // Not in data model but useful
}
