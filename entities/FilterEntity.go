package entities

type Filters struct {
	ID     uint           `gorm:"primarykey" json:"id"`
	Slug   string         `json:"slug"`
	Title  string         `json:"title"`
	Values []FilterValues `gorm:"foreignKey:FilterID" json:"values"`
}

type FilterValues struct {
	Id       uint   `json:"id"`
	FilterID uint   `json:"filter_id"`
	Slug     string `json:"slug"`
	Title    string `json:"title"`
}
