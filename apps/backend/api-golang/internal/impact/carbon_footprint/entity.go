// Package carbonfootprint handles carbon footprint calculations.
package carbonfootprint

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// CarbonFactor represents emission factors for a specific MCC and country
type CarbonFactor struct {
	ID          string  `json:"id"`
	MCC         string  `json:"mcc"` // Merchant Category Code
	CountryID   string  `json:"countryId"`
	Factor      float64 `json:"factor"` // kg CO2e per EUR
	Description string  `json:"description"`
}

// Footprint represents a calculated carbon footprint (matches Impact calculation data model)
// See: https://www.notion.so/ekko-earth/Impact-calculation-2b7f93807de48020b101c3d3de0f1286
type Footprint struct {
	ID string `json:"id"` // UUID

	// Customer and Organisation
	CustomerID     string `json:"customerId"`     // UUID
	OrganisationID string `json:"organisationId"` // UUID

	// Merchant details
	MCC                string  `json:"mcc"` // Merchant category code
	MerchantName       *string `json:"merchantName,omitempty"`
	MerchantStreet     *string `json:"merchantStreet,omitempty"`
	MerchantCity       *string `json:"merchantCity,omitempty"`
	MerchantPostalCode *string `json:"merchantPostalCode,omitempty"`
	MerchantState      *string `json:"merchantState,omitempty"`
	MerchantCountry    string  `json:"merchantCountry"` // ISO-3

	// Transaction details
	Amount   float64 `json:"amount"`   // decimal
	Currency string  `json:"currency"` // ISO-3 code

	// Carbon footprint (high level metrics)
	CarbonCo2eGrams  float64 `json:"carbonCo2eGrams"`  // High level metric
	CarbonCo2eOunces float64 `json:"carbonCo2eOunces"` // High level metric

	// Nature footprint (for future use)
	NatureTotalMsa *float64 `json:"natureTotalMsa,omitempty"` // High level metric

	// Pressure points (pressure point details)
	PressurePoints *PressurePoints `json:"pressurePoints,omitempty"`

	// Equivalents (stored as JSON blobs)
	CarbonEquivalents *Equivalents `json:"carbonEquivalents,omitempty"`
	NatureEquivalents *Equivalents `json:"natureEquivalents,omitempty"`

	// Metadata
	Factor            float64   `json:"factor"` // The factor used
	CalculationMethod string    `json:"calculationMethod"`
	CreatedAt         time.Time `json:"createdAt"`
}

// PressurePoints represents pressure point details
type PressurePoints struct {
	// Structure can be expanded based on nature impact requirements
	Data map[string]interface{} `json:"data"`
}

// Equivalents represents equivalents array (stored as JSON blob)
type Equivalents []Equivalent

// Equivalent represents a single equivalent
type Equivalent struct {
	Key      string  `json:"key"`
	Value    float64 `json:"value"`
	Template string  `json:"template"`
	Unit     string  `json:"unit,omitempty"`
}

// Value implements driver.Valuer for database storage
func (e Equivalents) Value() (driver.Value, error) {
	if len(e) == 0 {
		return nil, nil
	}
	return json.Marshal(e)
}

// Scan implements sql.Scanner for database retrieval
func (e *Equivalents) Scan(value interface{}) error {
	if value == nil {
		*e = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), e)
	}
	return json.Unmarshal(bytes, e)
}

// Value implements driver.Valuer for database storage
func (pp PressurePoints) Value() (driver.Value, error) {
	return json.Marshal(pp.Data)
}

// Scan implements sql.Scanner for database retrieval
func (pp *PressurePoints) Scan(value interface{}) error {
	if value == nil {
		*pp = PressurePoints{Data: make(map[string]interface{})}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), &pp.Data)
	}
	return json.Unmarshal(bytes, &pp.Data)
}

// CarbonKg returns carbon footprint in kg (convenience method)
func (f *Footprint) CarbonKg() float64 {
	return f.CarbonCo2eGrams / 1000.0
}
