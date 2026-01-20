package currency

import (
	"context"
	"fmt"

	"github.com/bilo-mono/packages/common/service"
)

const baseCurrency = "EUR"

// DefaultService implements the Service interface
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new currency service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// ConvertToEUR converts an amount from the given currency to EUR
func (s *DefaultService) ConvertToEUR(ctx context.Context, amount float64, fromCurrency string) (*ConversionResult, error) {
	// If already EUR, no conversion needed
	if fromCurrency == baseCurrency {
		return &ConversionResult{
			OriginalAmount:   amount,
			OriginalCurrency: fromCurrency,
			ConvertedAmount:  amount,
			TargetCurrency:   baseCurrency,
			ExchangeRate:     1.0,
		}, nil
	}

	rate, err := s.Repo.GetExchangeRate(ctx, fromCurrency, baseCurrency)
	if err != nil {
		return nil, fmt.Errorf("getting exchange rate from %s to %s: %w", fromCurrency, baseCurrency, err)
	}

	convertedAmount := amount * rate.Rate

	return &ConversionResult{
		OriginalAmount:   amount,
		OriginalCurrency: fromCurrency,
		ConvertedAmount:  convertedAmount,
		TargetCurrency:   baseCurrency,
		ExchangeRate:     rate.Rate,
	}, nil
}

// ConvertFromEUR converts an amount from EUR to the given currency
func (s *DefaultService) ConvertFromEUR(ctx context.Context, amount float64, toCurrency string) (*ConversionResult, error) {
	// If already EUR, no conversion needed
	if toCurrency == baseCurrency {
		return &ConversionResult{
			OriginalAmount:   amount,
			OriginalCurrency: baseCurrency,
			ConvertedAmount:  amount,
			TargetCurrency:   toCurrency,
			ExchangeRate:     1.0,
		}, nil
	}

	rate, err := s.Repo.GetExchangeRate(ctx, baseCurrency, toCurrency)
	if err != nil {
		return nil, fmt.Errorf("getting exchange rate from %s to %s: %w", baseCurrency, toCurrency, err)
	}

	convertedAmount := amount * rate.Rate

	return &ConversionResult{
		OriginalAmount:   amount,
		OriginalCurrency: baseCurrency,
		ConvertedAmount:  convertedAmount,
		TargetCurrency:   toCurrency,
		ExchangeRate:     rate.Rate,
	}, nil
}
