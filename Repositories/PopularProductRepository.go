package Repositories

import (
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"gorm.io/gorm"
	"os"
)

type PopularProductRepository interface {
	GetAllRelatedProducts() ([]model.PopularProductsModel, error)
	GetAllRelatedProductsWithUserSpecialPrices(companyGroupId float64) ([]model.PopularProductsModel, error)
}

type PopularProductRepositoryImpl struct {
	db *gorm.DB
}

func NewPopularProductRepository(db *gorm.DB) PopularProductRepository {
	return &PopularProductRepositoryImpl{db: db}
}

func (pp *PopularProductRepositoryImpl) GetAllRelatedProducts() ([]model.PopularProductsModel, error) {

	popularProducts := []model.PopularProductsModel{}

	err := database.Database.Table("popular_products").
		Select(
			"products.id, products.slug, products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock," +
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

func (pp *PopularProductRepositoryImpl) GetAllRelatedProductsWithUserSpecialPrices(companyGroupId float64) (
	[]model.PopularProductsModel, error,
) {

	groupCompanyIdInt := int(companyGroupId)

	popularProducts := []model.PopularProductsModel{}

	err := database.Database.Table("popular_products").
		Select(
			"products.id, products.slug, products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, popular_products.created_at, popular_products.updated_at",
		).
		Joins(
			"INNER JOIN products ON products.id = popular_products.product_id "+
				"INNER JOIN product_translations pt ON pt.product_id = products.id "+
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' "+
				"LEFT JOIN files f ON f.id = ef.file_id "+
				"INNER JOIN brands br ON br.id = products.brand_id "+
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id "+
				"INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id  <=  ? AND pp.price != 0 WHERE products.is_active = true   order by rand() ",
			groupCompanyIdInt,
		).
		Limit(100).
		Find(&popularProducts).Error

	productMap := make(map[int]model.PopularProductsModel)
	for _, product := range popularProducts {
		existingProduct, ok := productMap[product.ID]
		if !ok || product.CompanyPriceId > existingProduct.CompanyPriceId {
			productMap[product.ID] = product
		}
	}

	// Sonuçları al
	var uniqueProducts []model.PopularProductsModel
	for _, product := range productMap {
		uniqueProducts = append(uniqueProducts, product)
	}

	buildPopularProducts(uniqueProducts)

	return uniqueProducts, err

}

func buildPopularProducts(popularProducts []model.PopularProductsModel) {
	for index, product := range popularProducts {
		popularProducts[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		popularProducts[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		popularProducts[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}
