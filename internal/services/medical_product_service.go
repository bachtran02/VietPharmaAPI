package services

import (
	"sync"
	"vietpharma-api/internal/models"
	"vietpharma-api/internal/providers"
)

type MedicalProductService interface {
	SearchProduct(query string) ([]models.Product, error)
}

type medicalProductService struct {
	providers []providers.PharmacyProvider
}

func NewMedicalProductService() MedicalProductService {
	return &medicalProductService{
		providers: []providers.PharmacyProvider{
			providers.NewLongChauProvider(),
		},
	}
}

/* SearchProduct searches for products across providers */
func (s *medicalProductService) SearchProduct(query string) ([]models.Product, error) {
	var (
		wg       sync.WaitGroup
		mu       sync.Mutex
		results  []models.Product
		numProvs = len(s.providers)
	)

	/* Create a channel to collect results from all providers */
	resultsChan := make(chan []models.Product, numProvs)

	/* Query all providers concurrently */
	for _, provider := range s.providers {
		wg.Add(1)
		go func(p providers.PharmacyProvider) {
			defer wg.Done()

			/* Skip if provider is not available */
			if !p.IsAvailable() {
				return
			}

			/* Search products from this provider */
			products, err := p.SearchProduct(query)
			if err != nil {
				return
			}

			resultsChan <- products
		}(provider)
	}

	/* Wait for all providers to complete in a separate goroutine */
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	/* Collect results from the channel */
	for products := range resultsChan {
		mu.Lock()
		results = append(results, products...)
		mu.Unlock()
	}

	return results, nil
}
