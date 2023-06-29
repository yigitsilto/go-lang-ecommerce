package Repositories

import (
	model "ecommerce/dto"
	"ecommerce/utils"
	"gorm.io/gorm"
)

type PopularProductRepository interface {
	GetAllRelatedProducts(companyGroupId float64) ([]model.Product, error)
}

type PopularProductRepositoryImpl struct {
	db          *gorm.DB
	productUtil utils.ProductUtilInterface
}

func NewPopularProductRepository(db *gorm.DB, productUtil utils.ProductUtilInterface) PopularProductRepository {
	return &PopularProductRepositoryImpl{db: db, productUtil: productUtil}
}

func (pp *PopularProductRepositoryImpl) GetAllRelatedProducts(companyGroupId float64) (
	[]model.Product, error,
) {

	popularProducts := []model.Product{}

	groupCompanyIdInt := int(companyGroupId)

	query := pp.db.Table("popular_products").
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
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	if groupCompanyIdInt != 0 {

		query = query.Select(
			"products.id, products.slug, products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, popular_products.created_at, popular_products.updated_at",
		).
			Joins(
				"INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id  <=  ? AND pp.price != 0 ",
				groupCompanyIdInt,
			)
	}

	err := query.Where(
		"products.is_active =?", true,
	).Order(" rand()").Limit(20).Find(&popularProducts).Error

	if groupCompanyIdInt != 0 {

		popularProducts = pp.productUtil.UniqueProductsWithPriceCalculation(popularProducts, "orderByIdAsc")

	}

	pp.productUtil.BuildProducts(popularProducts)

	return popularProducts, err

}
