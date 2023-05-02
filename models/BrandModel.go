package model

type Brand struct {
	BaseModel
	Title string `gorm:"size:255;not null;unique" json:"Title"`
}
