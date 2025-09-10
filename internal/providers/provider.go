package providers

import (
	"net/http"
	"vietpharma-api/internal/models"
)

// PharmacyProvider defines the interface that all pharmacy providers must implement
type PharmacyProvider interface {
	// SearchProduct searches for products based on the query
	SearchProduct(query string) ([]models.Product, error)

	// GetProviderInfo returns basic information about the provider
	GetProviderInfo() ProviderInfo

	// IsAvailable checks if the provider's service is currently available
	IsAvailable() bool
}

// ProviderInfo contains basic information about a provider
type ProviderInfo struct {
	Name    string
	BaseURL string
	APIURL  string
}

// BasePharmacyProvider provides common functionality for pharmacy providers
type BasePharmacyProvider interface {
	// GetHTTPClient returns the HTTP client to use for requests
	GetHTTPClient() *http.Client

	// GetProviderInfo returns basic provider information
	GetProviderInfo() ProviderInfo

	// IsAvailable checks if the provider's service is currently available
	IsAvailable() bool
}

// defaultBasePharmacyProvider implements BasePharmacyProvider interface with common functionality
type defaultBasePharmacyProvider struct {
	info   ProviderInfo
	client *http.Client
}

// NewBasePharmacyProvider creates a new instance of BasePharmacyProvider
func NewBasePharmacyProvider(name, baseURL, apiURL string) BasePharmacyProvider {
	return &defaultBasePharmacyProvider{
		info: ProviderInfo{
			Name:    name,
			BaseURL: baseURL,
			APIURL:  apiURL,
		},
		client: &http.Client{},
	}
}

func (p *defaultBasePharmacyProvider) GetHTTPClient() *http.Client {
	return p.client
}

func (p *defaultBasePharmacyProvider) GetProviderInfo() ProviderInfo {
	return p.info
}

func (p *defaultBasePharmacyProvider) IsAvailable() bool {
	resp, err := p.client.Get(p.info.BaseURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
