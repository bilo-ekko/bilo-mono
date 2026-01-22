package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"api-golang/internal/impact_partner/impact_partner"
)

// HTTPClientAdapter implements impact_partner.Service interface
// This adapter makes HTTP calls to the impact_partner microservice
//
// This is a CONCRETE EXAMPLE of a driven adapter (outbound).
// When impact_partner is extracted to a microservice, this adapter
// replaces the local DefaultService implementation.
type HTTPClientAdapter struct {
	baseURL    string
	httpClient *http.Client
}

// NewHTTPClientAdapter creates a new HTTP client adapter
func NewHTTPClientAdapter(baseURL string) *HTTPClientAdapter {
	return &HTTPClientAdapter{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetAllPartners implements impact_partner.Service interface
// Makes HTTP GET request to /api/impact-partners
func (a *HTTPClientAdapter) GetAllPartners(ctx context.Context) ([]*impact_partner.Entity, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", a.baseURL+"/api/impact-partners", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var partners []*impact_partner.Entity
	if err := json.NewDecoder(resp.Body).Decode(&partners); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return partners, nil
}

// GetPartnerByID implements impact_partner.Service interface
// Makes HTTP GET request to /api/impact-partners/{id}
func (a *HTTPClientAdapter) GetPartnerByID(ctx context.Context, id string) (*impact_partner.Entity, error) {
	url := fmt.Sprintf("%s/api/impact-partners/%s", a.baseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("partner not found")
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	var partner impact_partner.Entity
	if err := json.NewDecoder(resp.Body).Decode(&partner); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &partner, nil
}

// CreatePartner implements impact_partner.Service interface
// Makes HTTP POST request to /api/impact-partners
func (a *HTTPClientAdapter) CreatePartner(ctx context.Context, partner *impact_partner.Entity) error {
	body, err := json.Marshal(partner)
	if err != nil {
		return fmt.Errorf("failed to marshal partner: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/api/impact-partners",
		bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
