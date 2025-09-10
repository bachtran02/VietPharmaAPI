package dto

// LongChauSearchResult represents the search result structure from Long Chau
type LongChauSearchResult struct {
	AggregationData  any                     `json:"aggregations"`
	SearchResultData []LongChauSearchProduct `json:"products"`
	TotalCount       int                     `json:"totalCount"`
}

type LongChauSearchProduct struct {
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	WebName        string             `json:"webName"`
	Slug           string             `json:"slug"`
	SKU            string             `json:"sku"`
	Brand          string             `json:"brand"`
	DisplayCode    int                `json:"displayCode"`
	DosageForm     string             `json:"dosageForm"`
	Image          string             `json:"image"`
	Ingredients    string             `json:"ingredients"`
	IsActive       bool               `json:"isActive"`
	IsPublish      bool               `json:"isPublish"`
	Category       []LongChauCategory `json:"category"`
	Price          LongChauPrice      `json:"price"`
	Prices         []LongChauPrice    `json:"prices"`
	ProductRanking float64            `json:"productRanking"`
	SearchScoring  float64            `json:"searchScoring"`
	Specification  string             `json:"specification"`
}

type LongChauPrice struct {
	ID              int     `json:"id"`
	CurrencySymbol  string  `json:"currencySymbol"`
	IsDefault       bool    `json:"isDefault"`
	IsInventory     bool    `json:"isInventory"`
	IsSellDefault   bool    `json:"isSellDefault"`
	Level           int     `json:"level"`
	MeasureUnitCode int     `json:"measureUnitCode"`
	MeasureUnitName string  `json:"measureUnitName"`
	Price           float64 `json:"price"`
	ProductSpecs    any     `json:"productSpecs"`
}

type LongChauCategory struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	IsActive   bool    `json:"isActive"`
	Level      int     `json:"level"`
	ParentName *string `json:"parentName"`
	Slug       string  `json:"slug"`
}
