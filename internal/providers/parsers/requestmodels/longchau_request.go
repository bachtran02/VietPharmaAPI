package requestmodels

/* LongChauSearchRequest represents the search request payload for Long Chau API */
type LongChauSearchRequest struct {
	Keyword        string   `json:"keyword"`
	MaxResultCount int      `json:"maxResultCount"`
	SkipCount      int      `json:"skipCount"`
	SortType       int      `json:"sortType"`
	Codes          []string `json:"codes"`
	SuggestSize    int      `json:"suggestSize"`
}
