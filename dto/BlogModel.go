package dto

type BlogModel struct {
	BaseModel
	Slug             string `gorm:"size:255;not null;unique" json:"slug"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	CoverImage       string `json:"cover_image"`
}

type BlogLongModel struct {
	BaseModel
	Slug             string `gorm:"size:255;not null;unique" json:"slug"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	CoverImage       string `json:"cover_image"`
	Description      string `json:"description"`
}
