package model

import "gorm.io/gorm"

type Brand struct {
	gorm.Model
	Title string `gorm:"size:255;not null;unique" json:"Title"`
}
