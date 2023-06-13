package Repositories

import (
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"os"
)

func GetAllRelatedProducts() ([]model.PopularProductsModel, error) {

	popularProducts := []model.PopularProductsModel{}

	err := database.Database.Table("popular_products").
		Select(
			"products.id, products.slug, products.short_desc, products.price, products.special_price, products.qty, products.in_stock," +
				" brt.name AS brand_name, pt.name, " +
				" f.path AS path, products.is_active, popular_products.created_at, popular_products.updated_at",
		).
		Joins(
			"INNER JOIN products ON products.id = popular_products.product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id WHERE products.is_active = true order by rand()",
		).
		Limit(20).
		Find(&popularProducts).Error

	buildPopularProducts(popularProducts)

	return popularProducts, err

}

func buildPopularProducts(popularProducts []model.PopularProductsModel) {
	for index, product := range popularProducts {
		popularProducts[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		popularProducts[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		popularProducts[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}
