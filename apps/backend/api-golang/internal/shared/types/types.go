// Package types contains shared domain types used across the application.
package types

import "time"

// Money represents a monetary amount with currency
type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"` // ISO 4217 currency code (3 chars)
}

// Address represents a physical address (matches Ekko API schema)
type Address struct {
	Line1       string  `json:"line1"`
	Line2       *string `json:"line2,omitempty"`
	Line3       *string `json:"line3,omitempty"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postalCode"`
	State       *string `json:"state,omitempty"`
	CountryCode string  `json:"countryCode"` // ISO 3166-1 alpha-3 (3 chars)
}

// BillingConfig represents billing configuration for an organisation
type BillingConfig struct {
	CompanyRegistrationNumber string  `json:"companyRegistrationNumber"`
	CurrencyCode              string  `json:"currencyCode"` // ISO 4217 (3 chars)
	Email                     string  `json:"email"`
	TaxNumber                 *string `json:"taxNumber,omitempty"`
}

// OrganisationStatus represents the status of an organisation
type OrganisationStatus struct {
	Value   string  `json:"value"` // "active" or "inactive"
	Message *string `json:"message,omitempty"`
}

// OrganisationImpactPartner represents a simplified impact partner reference in an organisation
type OrganisationImpactPartner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// MerchantDetails represents merchant information for a transaction
type MerchantDetails struct {
	Name        string  `json:"name,omitempty"`
	MCC         string  `json:"mcc"` // Merchant Category Code
	CountryCode string  `json:"countryCode"`
	Address     Address `json:"address,omitempty"`
}

// TransactionDetails represents a financial transaction
type TransactionDetails struct {
	ID              string          `json:"id"`
	Amount          Money           `json:"amount"`
	MerchantDetails MerchantDetails `json:"merchantDetails,omitempty"`
	Timestamp       time.Time       `json:"timestamp"`
}

// CarbonFootprint represents a calculated carbon footprint
type CarbonFootprint struct {
	ID                string    `json:"id"`
	Amount            float64   `json:"amount"` // in kg CO2e
	TransactionID     string    `json:"transactionId"`
	CalculationMethod string    `json:"calculationMethod"`
	CreatedAt         time.Time `json:"createdAt"`
}

// Location represents a geographic location for projects
type Location struct {
	Country     string   `json:"country,omitempty"`
	Region      *string  `json:"region,omitempty"`
	Coordinates *LatLong `json:"coordinates,omitempty"`
}

// LatLong represents latitude and longitude coordinates
type LatLong struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// ProjectUnit represents the unit of measurement for a project
type ProjectUnit struct {
	Type   string `json:"type"`   // e.g., "tCO2e", "hectares"
	Symbol string `json:"symbol"` // e.g., "t", "ha"
}

// BlendedProject represents a project with its blended price
type BlendedProject struct {
	ProjectID   string   `json:"projectId"`
	ProjectName string   `json:"projectName"`
	PartnerID   string   `json:"partnerId"`
	UnitPrice   float64  `json:"unitPrice"`  // Price per kg CO2e
	Allocation  float64  `json:"allocation"` // Percentage allocation (0-1)
	Location    Location `json:"location,omitempty"`
}

// BlendedPriceResult contains the blended price calculation result
type BlendedPriceResult struct {
	BlendedUnitPrice float64          `json:"blendedUnitPrice"`
	Projects         []BlendedProject `json:"projects"`
}
