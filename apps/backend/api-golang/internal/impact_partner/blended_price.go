package impact_partner

import (
	"api-golang/internal/shared/types"
)

// ProjectWithPrice extends project info with pricing
type ProjectWithPrice struct {
	ID              string         `json:"id"`
	Name            string         `json:"name"`
	ImpactPartnerID string         `json:"impactPartnerId"`
	Type            string         `json:"type"`      // carbonCredits, natureCredits, contribution
	UnitPrice       float64        `json:"unitPrice"` // Price per kg CO2e in EUR
	Location        types.Location `json:"location,omitempty"`
}

// BlendedPriceCalculator calculates blended prices across projects
type BlendedPriceCalculator struct {
	partnerService *Service
}

// NewBlendedPriceCalculator creates a new calculator
func NewBlendedPriceCalculator(partnerService *Service) *BlendedPriceCalculator {
	return &BlendedPriceCalculator{
		partnerService: partnerService,
	}
}

// CalculateBlendedPrice calculates the blended unit price across all projects for an organisation
func (c *BlendedPriceCalculator) CalculateBlendedPrice(organisationID string, filterByLocation bool, locationCountry string) (*types.BlendedPriceResult, error) {
	// Get all partners for the organisation
	// In a real implementation, this would filter by organisation
	partners := c.partnerService.GetAllPartners()

	// Sample project prices (in real implementation, these would come from a projects repository)
	projectPrices := map[string][]ProjectWithPrice{
		"partner-1": { // Green Carbon Trust
			{ID: "project-1", Name: "Amazon Rainforest Conservation", ImpactPartnerID: "partner-1", Type: "carbonCredits", UnitPrice: 15.00, Location: types.Location{Country: "Brazil"}},
			{ID: "project-4", Name: "Mangrove Restoration Program", ImpactPartnerID: "partner-1", Type: "natureCredits", UnitPrice: 22.00, Location: types.Location{Country: "Vietnam"}},
		},
		"partner-2": { // Ocean Conservation Fund
			{ID: "project-2", Name: "Solar Farm Initiative India", ImpactPartnerID: "partner-2", Type: "carbonCredits", UnitPrice: 8.50, Location: types.Location{Country: "India"}},
			{ID: "project-3", Name: "Wind Energy Project Denmark", ImpactPartnerID: "partner-2", Type: "carbonCredits", UnitPrice: 10.00, Location: types.Location{Country: "Denmark"}},
		},
	}

	// Collect all projects
	var allProjects []types.BlendedProject
	for _, partner := range partners {
		projects, ok := projectPrices[partner.ID]
		if !ok {
			continue
		}

		for _, p := range projects {
			// Filter by location if required
			if filterByLocation && locationCountry != "" && p.Location.Country != locationCountry {
				continue
			}

			allProjects = append(allProjects, types.BlendedProject{
				ProjectID:   p.ID,
				ProjectName: p.Name,
				PartnerID:   p.ImpactPartnerID,
				UnitPrice:   p.UnitPrice,
				Allocation:  0, // Will be calculated below
				Location:    p.Location,
			})
		}
	}

	// If no projects, return zero
	if len(allProjects) == 0 {
		return &types.BlendedPriceResult{
			BlendedUnitPrice: 0,
			Projects:         []types.BlendedProject{},
		}, nil
	}

	// Calculate equal allocation and blended price
	allocation := 1.0 / float64(len(allProjects))
	var totalPrice float64
	for i := range allProjects {
		allProjects[i].Allocation = allocation
		totalPrice += allProjects[i].UnitPrice * allocation
	}

	return &types.BlendedPriceResult{
		BlendedUnitPrice: totalPrice,
		Projects:         allProjects,
	}, nil
}
