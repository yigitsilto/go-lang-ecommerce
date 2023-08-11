package dto

type PopularCategoryModel struct {
	BaseModel
	Slug string `gorm:"size:255;not null;unique" json:"slug"`
	Name string `json:"name"`
	Path string `json:"path"`
}
