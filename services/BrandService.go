package services

import (
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllBrands() []model.Brand {

	var brands []model.Brand

	database.Database.Preload("Translation").Find(&brands)

	return brands

}

func FindBrandById(id string) (model.Brand, error) {

	b := model.Brand{}

	err := database.Database.Where("id=?", id).Preload("Translation").Find(&b).Error

	return b, err

}
