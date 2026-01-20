package organisation

import (
	"context"
	"fmt"

	"api-golang/internal/shared/errors"

	"github.com/bilo-mono/packages/common/service"
)

// DefaultService implements the Service interface
type DefaultService struct {
	service.BaseService[Repository]
}

// NewService creates a new organisation service
func NewService(repo Repository) *DefaultService {
	return &DefaultService{
		BaseService: service.NewBaseService(repo),
	}
}

// ValidateOrganisation validates organisation access based on header and body org IDs
// If headerOrgID == bodyOrgID, access is allowed
// If bodyOrgID is a child of headerOrgID, access is allowed
// Otherwise, access is forbidden
func (s *DefaultService) ValidateOrganisation(ctx context.Context, headerOrgID, bodyOrgID string) (*Entity, error) {
	// If they match, simply return the organisation
	if headerOrgID == bodyOrgID {
		return s.GetOrganisation(ctx, bodyOrgID)
	}

	// Check if body org is a child of header org
	bodyOrg, err := s.Repo.GetByID(ctx, bodyOrgID)
	if err != nil {
		return nil, fmt.Errorf("validating organisation: %w", err)
	}

	if !bodyOrg.IsChildOf(headerOrgID) {
		return nil, errors.NewForbiddenError(domainName,
			fmt.Sprintf("organisation %s is not a child of %s", bodyOrgID, headerOrgID))
	}

	return bodyOrg, nil
}

// GetOrganisation retrieves an organisation by ID
func (s *DefaultService) GetOrganisation(ctx context.Context, id string) (*Entity, error) {
	org, err := s.Repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting organisation: %w", err)
	}
	return org, nil
}
