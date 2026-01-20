// Package customer handles customer-related business logic.
package customer

import (
	"time"
)

// Entity represents a stored customer in the system
// This is the internal representation, different from the quote request customer object
type Entity struct {
	ID             string    `json:"id"`
	OrganisationID string    `json:"organisationId"`
	Reference      string    `json:"reference"` // Client's internal reference for the customer
	Email          *string   `json:"email,omitempty"`
	Name           *string   `json:"name,omitempty"`
	PostalCode     *string   `json:"postalCode,omitempty"`
	City           *string   `json:"city,omitempty"`
	State          *string   `json:"state,omitempty"`
	CountryCode    string    `json:"countryCode"` // ISO 3166-1 alpha-3 (3 chars)
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// QuoteCustomer represents the customer object used in quote requests
// See: https://docs.ekko.earth/v3/reference/post_quotes-carbon
type QuoteCustomer struct {
	Reference   string  `json:"reference"` // Required - client's internal reference
	PostalCode  *string `json:"postalCode,omitempty"`
	City        *string `json:"city,omitempty"`
	State       *string `json:"state,omitempty"`
	CountryCode string  `json:"countryCode"` // Required - ISO 3166-1 alpha-3 (3 chars)
}

// ToQuoteCustomer converts an Entity to a QuoteCustomer
func (e *Entity) ToQuoteCustomer() QuoteCustomer {
	return QuoteCustomer{
		Reference:   e.Reference,
		PostalCode:  e.PostalCode,
		City:        e.City,
		State:       e.State,
		CountryCode: e.CountryCode,
	}
}
