package models

// Brand represents the core attributes of a brand in the business domain.
type Brand struct {
	BrandUuid string `json:"brandUuid"`
	BrandName string `json:"brandName"`
}
