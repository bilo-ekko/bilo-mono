// Package currency handles currency conversion.
package currency

import "time"

// ExchangeRate represents an exchange rate between two currencies
// Matches Currency conversion rate data model
// See: https://www.notion.so/ekko-earth/Currency-and-country-2b7f93807de480d1a1accf1400743413
type ExchangeRate struct {
	ID             string    `json:"id"`             // currency_code
	SourceCurrency string    `json:"sourceCurrency"` // EUR (required)
	TargetCurrency string    `json:"targetCurrency"` // Required
	ConversionDate time.Time `json:"conversionDate"` // Required
	Rate           float64   `json:"rate"`
	ValidFrom      time.Time `json:"validFrom,omitempty"`
	ValidTo        time.Time `json:"validTo,omitempty"`
}

// Currency represents a currency entity
// Matches Currency data model
// See: https://www.notion.so/ekko-earth/Currency-and-country-2b7f93807de480d1a1accf1400743413
type Currency struct {
	Code           string `json:"code"` // Required
	Name           string `json:"name"` // Required (e.g., "EUR")
	Symbol         string `json:"symbol,omitempty"`
	DecimalPlaces  int    `json:"decimalPlaces,omitempty"`
	Region         string `json:"region,omitempty"`
	ISONumericCode string `json:"isoNumericCode,omitempty"`
	Status         string `json:"status,omitempty"` // active, inactive, etc.
}

// ConversionResult represents the result of a currency conversion
type ConversionResult struct {
	OriginalAmount   float64 `json:"originalAmount"`
	OriginalCurrency string  `json:"originalCurrency"`
	ConvertedAmount  float64 `json:"convertedAmount"`
	TargetCurrency   string  `json:"targetCurrency"`
	ExchangeRate     float64 `json:"exchangeRate"`
}
