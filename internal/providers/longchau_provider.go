package providers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"vietpharma-api/internal/models"
	"vietpharma-api/internal/providers/parsers/dto"
	"vietpharma-api/internal/providers/parsers/requestmodels"
)

const (
	LONGCHAU_SEARCH_ENDPOINT = "/lccus/search-product-service/api/products/ecom/product/search"
)

// LongChauProvider implements PharmacyProvider interface
type LongChauProvider struct {
	base BasePharmacyProvider // Embed BaseProvider for common functionality
	// parser parsers.HTMLParser   // Parser for LongChau-specific HTML/JSON
}

// NewLongChauProvider creates a new instance of LongChauProvider
func NewLongChauProvider() PharmacyProvider {
	return &LongChauProvider{
		base: NewBasePharmacyProvider(
			"Nhà Thuốc Long Châu",
			"https://nhathuoclongchau.com.vn",
			"https://api.nhathuoclongchau.com.vn",
		),
		// parser: parsers.NewLongChauParser(),
	}
}

// SearchProduct implements the Provider interface
func (p *LongChauProvider) SearchProduct(searchQuery string) ([]models.Product, error) {
	/* Get the HTTP client and provider info from base provider */
	client := p.base.GetHTTPClient()
	providerInfo := p.base.GetProviderInfo()

	/* Construct the search URL */
	searchURL := providerInfo.APIURL + LONGCHAU_SEARCH_ENDPOINT

	/* Create search request payload with web default params */
	searchRequest := requestmodels.LongChauSearchRequest{
		Keyword:        searchQuery,
		MaxResultCount: 16,
		SkipCount:      0,
		SortType:       4,
		Codes: []string{
			"category",
			"objectUse",
			"indications",
			"prescription",
			"skin",
			"flavor",
			"manufactor",
			"brand",
			"brandOrigin",
		},
		SuggestSize: 6,
	}

	/* Convert request to JSON */
	requestBody, err := json.Marshal(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal search request: %w", err)
	}

	/* Create POST request */
	req, err := http.NewRequest("POST", searchURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	/* Set headers */
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	/* Make the request */
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch from Long Chau: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	/* Parse the JSON response */
	var searchResult dto.LongChauSearchResult
	if err := json.Unmarshal(bodyBytes, &searchResult); err != nil {
		return nil, fmt.Errorf("failed to parse Long Chau search response: %w", err)
	}

	fmt.Println(searchResult.TotalCount)

	// Convert LongChau products to common Product model
	var products []models.Product
	for _, item := range searchResult.SearchResultData {
		product := models.Product{
			ID:          fmt.Sprintf("lc-%d", item.ID),
			Name:        item.Name,
			Description: item.Ingredients,
			Price:       item.Price.Price,
			Pharmacy: models.Pharmacy{
				Name: p.base.GetProviderInfo().Name,
			},
		}
		products = append(products, product)
	}

	return products, nil
}

// GetProviderInfo implements the Provider interface
// Delegates to base provider
func (p *LongChauProvider) GetProviderInfo() ProviderInfo {
	return p.base.GetProviderInfo()
}

// IsAvailable implements the Provider interface
// Uses base provider's implementation
func (p *LongChauProvider) IsAvailable() bool {
	return p.base.IsAvailable()
}
