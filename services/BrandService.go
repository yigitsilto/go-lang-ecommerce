package services

import (
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllBrands() []model.Brand {

	var brands []model.Brand

	database.Database.Table("brands").Select("*").Joins(" INNER JOIN brand_translations bt on bt.brand_id = brands.id " +
		"INNER JOIN entity_files ef on ef.entity_type = 'Modules\\\\Brand\\\\Entities\\\\Brand' and ef.entity_id = brands.id" +
		" INNER JOIN files f on f.id = ef.file_id").Find(&brands)

	return brands

}

func FindBrandById(id string) (model.Brand, error) {

	b := model.Brand{}

	err := database.Database.Where("id=?", id).Preload("Translation").Find(&b).Error

	return b, err

}
