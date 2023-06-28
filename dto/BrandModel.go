package dto

type Brand struct {
	BaseModel
	Slug     string `gorm:"size:255;not null;unique" json:"slug"`
	IsActive int8   `gorm:"not null" json:"is_active"`
	Name     string `json:"name"`
	Path     string `json:"path"`
}
