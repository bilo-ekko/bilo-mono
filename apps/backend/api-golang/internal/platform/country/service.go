package country

import (
	"context"
	"fmt"
)

// DefaultService implements the Service interface
type DefaultService struct {
	repo Repository
}

// NewService creates a new country service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{repo: repo}
}

// GetCountryByCode retrieves a country by its ISO code
func (s *DefaultService) GetCountryByCode(ctx context.Context, code string) (*Entity, error) {
	country, err := s.repo.GetByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("getting country by code %s: %w", code, err)
	}
	return country, nil
}
