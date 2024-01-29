package dto

type Product struct {
	BaseModel
	Slug                  string  `gorm:"size:255;not null;unique" json:"slug"`
	IsActive              int8    `gorm:"not null" json:"is_active"`
	Name                  string  `json:"name"`
	ShortDescription      string  `json:"short_description"`
	Path                  string  `json:"path"`
	SecondImage           string  `json:"second_image"`
	Qty                   int     `json:"qty"`
	InStock               bool    `json:"in_stock"`
	CompanyPriceId        int     `json:"company_price_id"`
	Price                 float64 `gorm:"type:decimal(18,4) unsigned" json:"price"`
	Price2                float64 `gorm:"type:decimal(18,4) unsigned" json:"price2"`
	Price3                float64 `gorm:"type:decimal(18,4) unsigned" json:"price3"`
	Price4                float64 `gorm:"type:decimal(18,4) unsigned" json:"price4"`
	Price5                float64 `gorm:"type:decimal(18,4) unsigned" json:"price5"`
	PriceFormatted        string  `json:"price_formatted"`
	SpecialPrice          float64 `gorm:"type:decimal(18,4) unsigned" json:"special_price"`
	SpecialPriceFormatted string  `json:"special_price_formatted"`
	BrandName             string  `json:"brand_name"`
	Tax                   float64 `gorm:"type:decimal(18,4) unsigned" json:"tax"`
	VideoUrl              string  `json:"video_url"`
}

type FilterIdValues struct {
	FilterId string `gorm:"not null" json:"filter_id"`
	Id       string `gorm:"not null" json:"id"`
}
