// Package fee handles service fee calculations.
package fee

// FeeConfig represents the fee configuration for an organisation
type FeeConfig struct {
	OrganisationID   string  `json:"organisationId"`
	FeePercentage    float64 `json:"feePercentage"`    // Service fee as percentage (e.g., 0.05 for 5%)
	MinimumFee       float64 `json:"minimumFee"`       // Minimum fee in EUR
	MaximumFee       float64 `json:"maximumFee"`       // Maximum fee in EUR (0 = no max)
}

// FeeResult represents the calculated service fee
type FeeResult struct {
	CompensationAmount float64 `json:"compensationAmount"`
	FeeAmount          float64 `json:"feeAmount"`
	FeePercentage      float64 `json:"feePercentage"`
}
