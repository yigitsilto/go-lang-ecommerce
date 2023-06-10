package model

type Product struct {
	BaseModel
	Slug         string  `gorm:"size:255;not null;unique" json:"slug"`
	IsActive     int8    `gorm:"not null" json:"is_active"`
	Name         string  `json:"name"`
	Path         string  `json:"path"`
	Qty          int     `json:"qty"`
	InStock      bool    `json:"in_stock"`
	Price        float32 `json:"price"`
	SpecialPrice float32 `json:"special_price"`
	BrandName    string  `json:"brand_name"`
}
