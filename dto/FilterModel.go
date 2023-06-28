package dto

type FilterModel struct {
	Id     int                 `json:"id"`
	Slug   string              `json:"slug"`
	Title  string              `json:"title"`
	Values []FilterValuesModel `json:"values"`
}

type FilterValuesModel struct {
	Id    int    `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}
