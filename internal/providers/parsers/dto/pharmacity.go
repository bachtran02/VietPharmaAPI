package dto

// PharmacitySearchResult represents the search result structure from Pharmacity
type PharmacitySearchResult struct {
	Success  bool                `json:"success"`
	Products []PharmacityProduct `json:"products"`
	Metadata PharmacityMetadata  `json:"metadata"`
}

type PharmacityMetadata struct {
	Total       int `json:"total"`
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
}

type PharmacityProduct struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Slug         string   `json:"slug"`
	SKU          string   `json:"sku"`
	Description  string   `json:"description"`
	Price        Price    `json:"price"`
	Images       []string `json:"images"`
	Category     Category `json:"category"`
	Manufacturer string   `json:"manufacturer"`
	Store        Store    `json:"store"`
}

type Price struct {
	Regular     float64 `json:"regular"`
	Sale        float64 `json:"sale"`
	HasDiscount bool    `json:"hasDiscount"`
}

type Category struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Store struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Province string   `json:"province"`
	District string   `json:"district"`
	Location Location `json:"location"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}
