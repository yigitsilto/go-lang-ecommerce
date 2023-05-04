package model

type Brand struct {
	BaseModel
	Slug        string           `gorm:"size:255;not null;unique" json:"slug"`
	IsActive    int8             `gorm:"not null" json:"is_active"`
	Translation BrandTranslation `json:"translation"`
}

type BrandTranslation struct {
	Name    string `json:"name"`
	BrandId uint   `json:"brand_id"`
}
