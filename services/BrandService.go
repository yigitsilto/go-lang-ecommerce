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

func FindBrandById(id string) model.Brand {

	brand := model.Brand{}

	database.Database.Where("id=?", id).Find(&brand)

	return brand

}

func CreateBrand(input brand.CreateBrandDTO) model.Brand {

	b := model.Brand{Title: input.Title}

	database.Database.Create(&b)

	// TODO go error araştırması yapılacak!!

	return b

}

func DeleteBrand(id string) {

	b := model.Brand{}

	database.Database.Where("id=?", id).Find(&b)

	database.Database.Delete(&b)
}
