// Package salestax handles sales tax calculations.
package salestax

// TaxRate represents a sales tax rate configuration
// Matches Sales tax data model
// See: https://www.notion.so/ekko-earth/Sales-tax-2b7f93807de480a69495c23d832f98a8
type TaxRate struct {
	ID string `json:"id"` // Required

	// Merchant location (zip, state, country ISO-3)
	MerchantLocation MerchantLocation `json:"merchantLocation"` // Required

	// Customer location (zip, state, country ISO-3)
	CustomerLocation CustomerLocation `json:"customerLocation"` // Required

	// Tax liability
	IsEkkoTaxLiable bool `json:"isEkkoTaxLiable"` // YES/NO - who is liable for sales tax

	// Tax rates
	ServiceFeeRate   float64 `json:"serviceFeeRate,omitempty"`
	CarbonCreditRate float64 `json:"carbonCreditRate,omitempty"`
	CharityRate      float64 `json:"charityRate,omitempty"`
	NonCharityRate   float64 `json:"nonCharityRate,omitempty"`

	// External tax service
	TaxJurisdictionID string `json:"taxJurisdictionId,omitempty"` // From Avalara
}

// MerchantLocation represents merchant location for tax calculation
type MerchantLocation struct {
	Zip     string `json:"zip,omitempty"` // Postal code
	State   string `json:"state,omitempty"`
	Country string `json:"country"` // ISO-3 (required)
}

// CustomerLocation represents customer location for tax calculation
type CustomerLocation struct {
	Zip     string `json:"zip,omitempty"` // Postal code
	State   string `json:"state,omitempty"`
	Country string `json:"country"` // ISO-3 (required)
}

// TaxCalculationInput represents input for sales tax calculation
type TaxCalculationInput struct {
	MerchantCountry    string
	MerchantState      string
	MerchantPostalCode string
	CustomerCountry    string
	CustomerState      string
	CustomerPostalCode string
	Amount             float64
}

// TaxResult represents the calculated sales tax
type TaxResult struct {
	TaxableAmount float64 `json:"taxableAmount"`
	TaxRate       float64 `json:"taxRate"`
	TaxAmount     float64 `json:"taxAmount"`
	TaxName       string  `json:"taxName"`
	IsApplicable  bool    `json:"isApplicable"`
}
