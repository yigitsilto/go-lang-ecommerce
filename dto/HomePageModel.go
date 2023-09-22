package dto

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
