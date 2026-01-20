package country

import (
	"context"
	"fmt"

	"github.com/bilo-mono/packages/common/service"
)

// DefaultService implements the Service interface
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new country service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// GetCountryByCode retrieves a country by its ISO code
func (s *DefaultService) GetCountryByCode(ctx context.Context, code string) (*Entity, error) {
	country, err := s.Repo.GetByCode(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("getting country by code %s: %w", code, err)
	}
	return country, nil
}
