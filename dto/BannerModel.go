package dto

type Banner struct {
	BaseModel
	Image       string `gorm:"size:255;not null;unique" json:"image"`
	Title       string `json:"title"`
	Description string `json:"description"`
	LinkUrl     string `json:"link_url"`
}

type Gallery struct {
	BaseModel
	ImagePath string `gorm:"size:255;not null;unique" json:"image_path"`
	LinkUrl   string `json:"link_url"`
}
