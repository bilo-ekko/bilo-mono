package currency

import (
	"context"
	"sync"
	"time"

	"api-golang/internal/shared/errors"
)

const domainName = "currency"

// InMemoryRepository implements Repository interface
type InMemoryRepository struct {
	rates map[string]*ExchangeRate // key: sourceCurrency:targetCurrency
	mu    sync.RWMutex
}

// NewInMemoryRepository creates a new repository with sample exchange rates
func NewInMemoryRepository() *InMemoryRepository {
	repo := &InMemoryRepository{
		rates: make(map[string]*ExchangeRate),
	}

	now := time.Now()
	validFrom := now.Add(-24 * time.Hour)
	validTo := now.Add(24 * time.Hour)

	// Sample exchange rates (to EUR)
	rates := []*ExchangeRate{
		{ID: "GBP:EUR", SourceCurrency: "EUR", TargetCurrency: "GBP", ConversionDate: now, Rate: 1.17, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "USD:EUR", SourceCurrency: "EUR", TargetCurrency: "USD", ConversionDate: now, Rate: 0.92, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "EUR:EUR", SourceCurrency: "EUR", TargetCurrency: "EUR", ConversionDate: now, Rate: 1.0, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "CHF:EUR", SourceCurrency: "EUR", TargetCurrency: "CHF", ConversionDate: now, Rate: 1.05, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "SEK:EUR", SourceCurrency: "EUR", TargetCurrency: "SEK", ConversionDate: now, Rate: 0.088, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "NOK:EUR", SourceCurrency: "EUR", TargetCurrency: "NOK", ConversionDate: now, Rate: 0.086, ValidFrom: validFrom, ValidTo: validTo},
		{ID: "DKK:EUR", SourceCurrency: "EUR", TargetCurrency: "DKK", ConversionDate: now, Rate: 0.134, ValidFrom: validFrom, ValidTo: validTo},
	}

	for _, r := range rates {
		key := r.SourceCurrency + ":" + r.TargetCurrency
		repo.rates[key] = r
	}

	return repo
}

// GetExchangeRate retrieves the exchange rate between two currencies
func (r *InMemoryRepository) GetExchangeRate(_ context.Context, from, to string) (*ExchangeRate, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := from + ":" + to
	rate, exists := r.rates[key]
	if !exists {
		// Try reverse lookup (to:from) and calculate inverse
		reverseKey := to + ":" + from
		if reverseRate, exists := r.rates[reverseKey]; exists {
			// Create inverse rate
			inverseRate := &ExchangeRate{
				ID:             from + ":" + to,
				SourceCurrency: from,
				TargetCurrency: to,
				ConversionDate: reverseRate.ConversionDate,
				Rate:           1.0 / reverseRate.Rate,
				ValidFrom:      reverseRate.ValidFrom,
				ValidTo:        reverseRate.ValidTo,
			}
			return inverseRate, nil
		}
		return nil, errors.NewNotFoundError(domainName, "exchange rate not found for "+from+" to "+to)
	}
	return rate, nil
}
