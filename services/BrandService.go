package services

import (
	"ecommerce/database"
	"ecommerce/dto/brand"
	model "ecommerce/models"
)

func GetAllBrands() []model.Brand {

	var brands []model.Brand

	database.Database.Find(&brands)

	return brands

}

func FindBrandById(id string) (model.Brand, error) {

	b := model.Brand{}

	err := database.Database.Where("id=?", id).Find(&b).Error

	return b, err

}

func CreateBrand(input brand.CreateBrandDTO) (model.Brand, error) {

	b := model.Brand{Title: input.Title}

	err := database.Database.Create(&b).Error

	return b, err

}

func DeleteBrand(id string) {

	b := model.Brand{}

	database.Database.Where("id=?", id).Find(&b)

	database.Database.Delete(&b)
}

func UpdateBrand(id string, input brand.CreateBrandDTO) (model.Brand, error) {

	b := model.Brand{}

	database.Database.Where("id=?", id).First(&b)

	b.Title = input.Title

	err := database.Database.Save(&b).Error

	return b, err

}
