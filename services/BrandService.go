package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllBrands() ([]model.Brand, error) {

	brands, err := Repositories.FindAllBrands()
	return brands, err

}

func FindBrandById(id string) (model.Brand, error) {

	b := model.Brand{}

	err := database.Database.Where("id=?", id).Preload("Translation").Find(&b).Error

	return b, err

}
