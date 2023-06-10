package services

import (
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"os"
)

func GetHomePage() (model.HomePageModel, error) {

	popularProducts := []model.RelatedProductsModel{}

	err := database.Database.Table("popular_products").Select(
		"products.id, products.slug, products.short_desc, " +
			"products.price, products.special_price, products.qty, products.in_stock, brt.name as brand_name," +
			" pt.name, f.path, products.is_active, popular_products.created_at, popular_products.updated_at ",
	).Joins(
		"INNER JOIN products ON products.id = popular_products.product_id " +
			" INNER JOIN product_translations pt on pt.product_id = products.id " +
			"LEFT JOIN entity_files ef on ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' and ef.entity_id = products.id" +
			" LEFT JOIN files f on f.id = ef.file_id INNER JOIN brands br ON br.id = products.brand_id " +
			"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
	).Limit(20).Find(&popularProducts).Error

	buildPopularProducts(popularProducts)

	blogs, err := getBlogsForHomePage()

	homePageModel := model.HomePageModel{Products: popularProducts, BlogModel: blogs}

	return homePageModel, err
}

func buildPopularProducts(popularProducts []model.RelatedProductsModel) {
	for index, product := range popularProducts {
		popularProducts[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		popularProducts[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		popularProducts[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}

func getBlogsForHomePage() ([]model.BlogModel, error) {

	blogs := []model.BlogModel{}

	err := database.Database.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err

}
