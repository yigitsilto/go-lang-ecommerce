package Repositories

import (
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllRelatedProducts() ([]model.RelatedProductsModel, error) {

	popularProducts := []model.RelatedProductsModel{}

	err := database.Database.Table("popular_products").
		Select(
			"products.id, products.slug, products.short_desc, products.price, products.special_price, products.qty, products.in_stock," +
				"GROUP_CONCAT(DISTINCT brt.name) AS brand_name, pt.name, " +
				"GROUP_CONCAT(DISTINCT f.path) AS path, products.is_active, popular_products.created_at, popular_products.updated_at",
		).
		Joins(
			"INNER JOIN products ON products.id = popular_products.product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id",
		).
		Group(
			"products.id, products.slug, products.short_desc, products.price, products.special_price, products.qty, products.in_stock, " +
				"pt.name, products.is_active, popular_products.created_at, popular_products.updated_at",
		).
		Limit(20).
		Find(&popularProducts).Error

	return popularProducts, err

}
