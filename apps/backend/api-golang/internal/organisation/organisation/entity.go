// Package organisation handles organisation-related business logic.
package organisation

import (
	"api-golang/internal/shared/types"
)

// Entity represents an organisation in the system (matches Ekko API v3 schema)
// See: https://docs.ekko.earth/v3/reference/get_organisations-organisationid
type Entity struct {
	OrganisationID          string                            `json:"organisationId"`
	ParentOrganisationID    *string                           `json:"parentOrganisationId,omitempty"`
	OrganisationReference   *string                           `json:"organisationReference,omitempty"` // Optional metadata for internal references
	TradingName             string                            `json:"tradingName"`
	LegalName               string                            `json:"legalName"`
	Address                 types.Address                     `json:"address"`
	CurrencyCode            string                            `json:"currencyCode"` // ISO 4217 (3 chars)
	Website                 *string                           `json:"website,omitempty"`
	Billing                 *types.BillingConfig              `json:"billing,omitempty"`
	MCC                     *string                           `json:"mcc,omitempty"`           // Merchant Category Code
	RelativeProfitShare     float64                           `json:"relativeProfitShare"`     // 0-1, percentage of service fee from parent
	ProportionalProfitShare float64                           `json:"proportionalProfitShare"` // 0-1, effective profit share after hierarchy
	ServiceFeePercentage    float64                           `json:"serviceFeePercentage"`    // 0-1, set during onboarding
	Status                  types.OrganisationStatus          `json:"status"`
	ImpactPartners          []types.OrganisationImpactPartner `json:"impactPartners,omitempty"`   // Inherited from parent
	CalculationTypes        []string                          `json:"calculationTypes,omitempty"` // e.g., ["carbon", "nature"]
}

// IsActive checks if the organisation is active
func (e *Entity) IsActive() bool {
	return e.Status.Value == "active"
}

// IsChildOf checks if this organisation is a child of another organisation
func (e *Entity) IsChildOf(parentID string) bool {
	return e.ParentOrganisationID != nil && *e.ParentOrganisationID == parentID
}

// GetMCC returns the MCC or empty string if not set
func (e *Entity) GetMCC() string {
	if e.MCC != nil {
		return *e.MCC
	}
	return ""
}

// HasImpactPartner checks if the organisation has access to a specific impact partner
func (e *Entity) HasImpactPartner(partnerID string) bool {
	for _, p := range e.ImpactPartners {
		if p.ID == partnerID {
			return true
		}
	}
	return false
}
