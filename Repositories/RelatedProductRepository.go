package Repositories

import (
	model "ecommerce/dto"
	"ecommerce/utils"
	"gorm.io/gorm"
)

type RelatedProductRepositoryInterface interface {
	FindAllRelatedProducts(groupCompanyId float64, productId string) ([]model.Product, error)
	FindDummyRelatedProducts(groupCompanyId float64) (
		[]model.Product, error,
	)
}

type RelatedProductRepositoryImpl struct {
	db          *gorm.DB
	productUtil utils.ProductUtilInterface
}

func NewRelatedProductRepository(
	db *gorm.DB, productUtil utils.ProductUtilInterface,
) RelatedProductRepositoryInterface {
	return &RelatedProductRepositoryImpl{
		db: db, productUtil: productUtil,
	}
}

func (r *RelatedProductRepositoryImpl) FindAllRelatedProducts(groupCompanyId float64, productId string) (
	[]model.Product, error,
) {
	var products []model.Product

	groupCompanyIdInt := int(groupCompanyId)
	limit := 12

	query := r.db.Table("related_products").
		Select(
			"products.id, products.slug, products.tax,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock," +
				" brt.name AS brand_name, pt.name, " +
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
		Joins(
			"INNER JOIN products products ON products.id = related_products.related_product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	if groupCompanyIdInt != 0 {

		query = query.Select(
			"products.id, products.slug, products.tax,  products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)

	}

	err := query.Where(
		"products.is_active =? AND related_products.product_id =?", true, productId,
	).Order(" rand()").Limit(limit).Find(&products).Error

	r.productUtil.BuildProducts(products)

	return products, err
}

func (r *RelatedProductRepositoryImpl) FindDummyRelatedProducts(groupCompanyId float64) (
	[]model.Product, error,
) {
	var products []model.Product

	groupCompanyIdInt := int(groupCompanyId)
	limit := 12

	query := r.db.Table("products").
		Select(
			"products.id, products.slug, products.tax,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock," +
				" brt.name AS brand_name, pt.name, " +
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
		Joins(
			"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	if groupCompanyIdInt != 0 {

		query = query.Select(
			"products.id, products.slug, products.tax,  products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)

	}

	err := query.Where(
		"products.is_active =? ", true,
	).Limit(limit).Find(&products).Error

	r.productUtil.BuildProducts(products)

	return products, err
}
