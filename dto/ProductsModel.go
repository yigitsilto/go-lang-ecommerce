package dto

type Product struct {
	BaseModel
	Slug                  string  `gorm:"size:255;not null;unique" json:"slug"`
	IsActive              int8    `gorm:"not null" json:"is_active"`
	Name                  string  `json:"name"`
	ShortDescription      string  `json:"short_description"`
	Path                  string  `json:"path"`
	Qty                   int     `json:"qty"`
	InStock               bool    `json:"in_stock"`
	CompanyPriceId        int     `json:"company_price_id"`
	Price                 float64 `gorm:"type:decimal(18,4) unsigned" json:"price"`
	PriceFormatted        string  `json:"price_formatted"`
	SpecialPrice          float64 `gorm:"type:decimal(18,4) unsigned" json:"special_price"`
	SpecialPriceFormatted string  `json:"special_price_formatted"`
	BrandName             string  `json:"brand_name"`
}
