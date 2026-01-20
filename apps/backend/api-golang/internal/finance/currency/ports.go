// Package currency defines ports for the currency sub-domain.
package currency

import "context"

// Repository defines the port for exchange rate data access
type Repository interface {
	GetExchangeRate(ctx context.Context, from, to string) (*ExchangeRate, error)
}

// Service defines the port for currency conversion business logic
type Service interface {
	ConvertToEUR(ctx context.Context, amount float64, fromCurrency string) (*ConversionResult, error)
	ConvertFromEUR(ctx context.Context, amount float64, toCurrency string) (*ConversionResult, error)
}
