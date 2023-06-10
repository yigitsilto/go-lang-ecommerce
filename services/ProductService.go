package services

import (
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllProducts() ([]model.Product, error) {

	var products []model.Product

	err := database.Database.Table("products").Select("products.id, products.slug, " +
		"products.price, products.special_price, products.qty, products.in_stock, brt.name as brand_name," +
		" pt.name, f.path, products.is_active").Joins(" INNER JOIN product_translations pt on pt.product_id = products.id " +
		"INNER JOIN entity_files ef on ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' and ef.entity_id = products.id" +
		" INNER JOIN files f on f.id = ef.file_id INNER JOIN brands br ON br.id = products.brand_id " +
		"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ").Find(&products).Error

	return products, err

}

func FindProductById(id string) (model.Product, error) {

	product := model.Product{}

	err := database.Database.Table("products").Select("*").Joins(" INNER JOIN product_translations pt on pt.product_id = products.id "+
		"INNER JOIN entity_files ef on ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' and ef.entity_id = products.id"+
		" INNER JOIN files f on f.id = ef.file_id").Where("products.id=?", id).Find(&product).Error

	return product, err

}
