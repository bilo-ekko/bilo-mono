package customer

import (
	"context"
	"fmt"
	"time"

	"api-golang/internal/shared/errors"

	"github.com/bilo-mono/packages/common/service"

	"github.com/google/uuid"
)

// DefaultService implements the Service interface
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new customer service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// GetOrCreateCustomer finds an existing customer or creates a new one
func (s *DefaultService) GetOrCreateCustomer(ctx context.Context, input CreateCustomerInput) (*Entity, error) {
	// Try to find existing customer by reference
	if input.Reference != "" {
		existing, err := s.Repo.GetByReference(ctx, input.OrganisationID, input.Reference)
		if err == nil {
			return existing, nil
		}
		// If error is not "not found", return it
		var domainErr *errors.DomainError
		if !errors.IsDomainError(err, &domainErr) || domainErr.Code != errors.ErrCodeNotFound {
			return nil, fmt.Errorf("checking existing customer: %w", err)
		}
	}

	// Create new customer
	now := time.Now()
	customer := &Entity{
		ID:             uuid.New().String(),
		OrganisationID: input.OrganisationID,
		Reference:      input.Reference,
		Email:          input.Email,
		Name:           input.Name,
		PostalCode:     input.PostalCode,
		City:           input.City,
		State:          input.State,
		CountryCode:    input.CountryCode,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err := s.Repo.Create(ctx, customer); err != nil {
		return nil, fmt.Errorf("creating customer: %w", err)
	}

	return customer, nil
}
