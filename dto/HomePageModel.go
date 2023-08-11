package dto

type BlogModel struct {
	BaseModel
	Slug             string `gorm:"size:255;not null;unique" json:"slug"`
	Title            string `json:"title"`
	ShortDescription string `json:"short_description"`
	CoverImage       string `json:"cover_image"`
}

type HomePageModel struct {
	Products          []Product              `json:"products"`
	BlogModel         []BlogModel            `json:"blogs"`
	Slider            []Slider               `json:"sliders"`
	PopularCategories []PopularCategoryModel `json:"populerCategories"`
}

type Slider struct {
	Id     int64  `json:"id"`
	FileId int64  `json:"file_id"`
	Path   string `json:"path"`
}
