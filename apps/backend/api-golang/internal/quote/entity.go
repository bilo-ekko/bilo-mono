// Package quote handles quote creation and management.
package quote

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"api-golang/internal/shared/types"
)

// Status represents the status of a quote
type Status string

const (
	StatusPending   Status = "pending"
	StatusAccepted  Status = "accepted"
	StatusRejected  Status = "rejected"
	StatusExpired   Status = "expired"
	StatusCompleted Status = "completed"
)

// Entity represents a stored carbon offset quote (matches Notion data model)
// See: https://www.notion.so/ekko-earth/Carbon-quote-2b7f93807de480e997e6e7b73ea9194a
type Entity struct {
	// Primary identifiers
	ID                   string `json:"id"`
	QuoteReference       string `json:"quoteReference"`       // Client-facing reference
	CalculationReference string `json:"calculationReference"` // Links to calculation

	// Organisation and customer
	OrganisationID string `json:"organisationId"`
	CustomerID     string `json:"customerId"`

	// Currency
	Currency string `json:"currency"` // ISO-3 currency code

	// Carbon credit amounts (stored in quote currency)
	CarbonCreditTotal              float64 `json:"carbonCreditTotal"`
	CarbonCreditImpact             float64 `json:"carbonCreditImpact"`
	CarbonCreditImpactSalesTax     float64 `json:"carbonCreditImpactSalesTax"`
	ImpactTaxRate                  float64 `json:"impactTaxRate"` // Rate at time of quote (for funds)
	CarbonCreditServiceFee         float64 `json:"carbonCreditServiceFee"`
	CarbonCreditServiceFeeSalesTax float64 `json:"carbonCreditServiceFeeSalesTax"`
	ServiceFeeTaxRate              float64 `json:"serviceFeeTaxRate"` // Rate at time of quote (for funds)

	// Pricing
	PricePerTonneCo2e float64 `json:"pricePerTonneCo2e"` // In quote currency

	// Payment processing
	PaymentServiceProviderID string   `json:"paymentServiceProviderId,omitempty"` // UUID
	CarbonCreditProcessorFee *float64 `json:"carbonCreditProcessorFee,omitempty"` // If applicable

	// Contribution details (stored as JSON blob)
	ContributionDetails ContributionDetails `json:"contributionDetails"`

	// Tax liability
	IsMerchantTaxLiable bool `json:"isMerchantTaxLiable"` // YES/NO

	// Filters and options (stored from request)
	CustomerLocationFilter bool `json:"customerLocationFilter"`
	IncludePartnerDetail   bool `json:"includePartnerDetail"`
	IncludeProjectDetail   bool `json:"includeProjectDetail"`

	// Product tracking
	EkkoProduct string `json:"ekkoProduct"` // Product name (embedded SDK, takeover SDK, checkout SDK, ImpactPay, API)

	// Service fee share (percentage at time of quote)
	ServiceFeeShare float64 `json:"serviceFeeShare"`

	// Order items (stored as JSON blob)
	OrderItems OrderItems `json:"orderItems"`

	// Metadata
	Status    Status    `json:"status"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ContributionDetails represents contribution breakdown (stored as JSON)
type ContributionDetails struct {
	ImpactPercentage             float64                     `json:"impactPercentage"`
	ImpactSalesTaxPercentage     float64                     `json:"impactSalesTaxPercentage"`
	ServiceFeePercentage         float64                     `json:"serviceFeePercentage"`
	ServiceFeeSalesTaxPercentage float64                     `json:"serviceFeeSalesTaxPercentage"`
	ImpactPartners               []ContributionImpactPartner `json:"impactPartners"`
}

// ContributionImpactPartner represents impact partner contribution breakdown
type ContributionImpactPartner struct {
	ID                           string   `json:"id"`
	ImpactPercentage             float64  `json:"impactPercentage"`
	ImpactSalesTaxPercentage     float64  `json:"impactSalesTaxPercentage"`
	ServiceFeePercentage         float64  `json:"serviceFeePercentage"`
	ServiceFeeSalesTaxPercentage float64  `json:"serviceFeeSalesTaxPercentage"`
	ProjectIDs                   []string `json:"projectIds"`
}

// OrderItems represents order items array (stored as JSON blob)
type OrderItems []OrderItem

// OrderItem represents a single order item
type OrderItem struct {
	ItemID    string      `json:"itemId"`
	SKU       string      `json:"sku,omitempty"`
	Name      string      `json:"name"`
	Category  string      `json:"category"`
	Quantity  int         `json:"quantity"`
	UnitPrice types.Money `json:"unitPrice"`
}

// Value implements driver.Valuer for database storage
func (oi OrderItems) Value() (driver.Value, error) {
	if len(oi) == 0 {
		return nil, nil
	}
	return json.Marshal(oi)
}

// Scan implements sql.Scanner for database retrieval
func (oi *OrderItems) Scan(value interface{}) error {
	if value == nil {
		*oi = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), oi)
	}
	return json.Unmarshal(bytes, oi)
}

// Value implements driver.Valuer for database storage
func (cd ContributionDetails) Value() (driver.Value, error) {
	return json.Marshal(cd)
}

// Scan implements sql.Scanner for database retrieval
func (cd *ContributionDetails) Scan(value interface{}) error {
	if value == nil {
		*cd = ContributionDetails{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return json.Unmarshal([]byte(value.(string)), cd)
	}
	return json.Unmarshal(bytes, cd)
}

// ========================================
// Request/Response DTOs (matching Notion API spec)
// See: https://www.notion.so/ekko-earth/Carbon-quote-2b7f93807de480e997e6e7b73ea9194a
// ========================================

// CustomerRequest represents the customer object in a quote request
type CustomerRequest struct {
	ID         *string `json:"id,omitempty"` // Optional - existing customer ID
	Reference  string  `json:"reference"`    // Required - client's internal reference
	PostalCode *string `json:"postalCode,omitempty"`
	City       *string `json:"city,omitempty"`
	State      *string `json:"state,omitempty"`
	Country    string  `json:"country"` // Required - ISO 3166-1 alpha-3 (3 chars)
}

// MerchantRequest represents the merchant object in a quote request
type MerchantRequest struct {
	MCC     string                 `json:"mcc"`     // Required - Merchant Category Code
	Name    string                 `json:"name"`    // Required
	Address MerchantAddressRequest `json:"address"` // Required
}

// MerchantAddressRequest represents the merchant address in a quote request
type MerchantAddressRequest struct {
	Address1   string  `json:"address1"`
	Address2   *string `json:"address2,omitempty"`
	Address3   *string `json:"address3,omitempty"`
	City       string  `json:"city"`
	State      *string `json:"state,omitempty"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"` // ISO 3166-1 alpha-3 (3 chars)
}

// OrderItemRequest represents an order item in a quote request
type OrderItemRequest struct {
	ItemID    string         `json:"itemId"`
	SKU       string         `json:"sku,omitempty"`
	Name      string         `json:"name"`
	Category  string         `json:"category"`
	Quantity  int            `json:"quantity"`
	UnitPrice OrderItemPrice `json:"unitPrice"`
}

// OrderItemPrice represents the unit price of an order item
type OrderItemPrice struct {
	Value        float64 `json:"value"`
	CurrencyCode string  `json:"currencyCode"` // ISO 4217 (3 chars)
}

// QuoteFiltersRequest represents filtering options
type QuoteFiltersRequest struct {
	CustomerLocation bool `json:"customerLocation"` // Best effort: state, country, region, world (only for credits)
}

// CreateQuoteRequest represents the request body for creating a carbon quote
type CreateQuoteRequest struct {
	Locale                      string               `json:"locale,omitempty"`     // API only, not needed for sessions
	OrganisationID              string               `json:"organisationId"`       // Required
	Customer                    CustomerRequest      `json:"customer"`             // Required
	Merchant                    *MerchantRequest     `json:"merchant,omitempty"`   // Optional if provided during org onboarding
	OrderItems                  []OrderItemRequest   `json:"orderItems,omitempty"` // Optional order objects, stored in DW
	IncludeImpactPartnerDetails bool                 `json:"includeImpactPartnerDetails,omitempty"`
	Filters                     *QuoteFiltersRequest `json:"filters,omitempty"` // Advanced options
}

// CreateQuoteResponse represents the response from creating a carbon quote
type CreateQuoteResponse struct {
	ID             string               `json:"id"`             // Entity ID for GET requests
	QuoteReference string               `json:"quoteReference"` // Client-facing reference
	Footprint      FootprintResponse    `json:"footprint"`
	Credits        CreditsResponse      `json:"credits"`
	Contribution   ContributionResponse `json:"contribution"`
}

// FootprintResponse represents the carbon footprint in the quote response
type FootprintResponse struct {
	Co2eGrams   float64              `json:"co2eGrams"`
	Co2eOunces  float64              `json:"co2eOunces"`
	Equivalents []EquivalentResponse `json:"equivalents"`
}

// EquivalentResponse represents a localised equivalent in the quote response
type EquivalentResponse struct {
	Key      string  `json:"key"`
	Value    float64 `json:"value"`
	Template string  `json:"template"` // Language based on locale
}

// CreditsResponse represents the credits section in the quote response
type CreditsResponse struct {
	TotalAmount              float64                 `json:"totalAmount"`
	ImpactAmount             float64                 `json:"impactAmount"`
	ImpactSalesTaxAmount     float64                 `json:"impactSalesTaxAmount"`
	ServiceFeeAmount         float64                 `json:"serviceFeeAmount"`
	ServiceFeeSalesTaxAmount float64                 `json:"serviceFeeSalesTaxAmount"`
	PricePerTonneCo2e        float64                 `json:"pricePerTonneCo2e"`     // In quote currency
	ImpactPartners           []ImpactPartnerResponse `json:"impactPartners"`        // Min 1, max 10 projects
	CustomerLocationMatch    string                  `json:"customerLocationMatch"` // state, country, region, world
}

// ImpactPartnerResponse represents an impact partner in the credits/contribution response
type ImpactPartnerResponse struct {
	ID          string            `json:"id"`
	Name        *string           `json:"name,omitempty"`        // Optional - enriched if includeImpactPartnerDetails is true
	Description *string           `json:"description,omitempty"` // Optional - enriched if includeImpactPartnerDetails is true
	Logo        *string           `json:"logo,omitempty"`        // Optional - enriched if includeImpactPartnerDetails is true
	Projects    []ProjectResponse `json:"projects"`
}

// ProjectResponse represents a project in the quote response
type ProjectResponse struct {
	ID string `json:"id"`
}

// ContributionResponse represents the contribution section in the quote response
type ContributionResponse struct {
	ImpactPercentage             float64                             `json:"impactPercentage"`
	ImpactSalesTaxPercentage     float64                             `json:"impactSalesTaxPercentage"`
	ServiceFeePercentage         float64                             `json:"serviceFeePercentage"`
	ServiceFeeSalesTaxPercentage float64                             `json:"serviceFeeSalesTaxPercentage"`
	ImpactPartners               []ContributionImpactPartnerResponse `json:"impactPartners"`
}

// ContributionImpactPartnerResponse represents an impact partner in the contribution response
type ContributionImpactPartnerResponse struct {
	ID                           string            `json:"id"`
	ImpactPercentage             float64           `json:"impactPercentage"`
	ImpactSalesTaxPercentage     float64           `json:"impactSalesTaxPercentage"`
	ServiceFeePercentage         float64           `json:"serviceFeePercentage"`
	ServiceFeeSalesTaxPercentage float64           `json:"serviceFeeSalesTaxPercentage"`
	Name                         *string           `json:"name,omitempty"`        // Optional - enriched if includeImpactPartnerDetails is true
	Description                  *string           `json:"description,omitempty"` // Optional - enriched if includeImpactPartnerDetails is true
	Logo                         *string           `json:"logo,omitempty"`        // Optional - enriched if includeImpactPartnerDetails is true
	Projects                     []ProjectResponse `json:"projects"`
}
