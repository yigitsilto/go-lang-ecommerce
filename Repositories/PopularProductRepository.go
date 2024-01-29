package Repositories

import (
	model "ecommerce/dto"
	"ecommerce/utils"
	"gorm.io/gorm"
)

type PopularProductRepository interface {
	GetAllRelatedProducts(companyGroupId float64) ([]model.Product, error)
	GetAllPopularCategories() ([]model.PopularCategoryModel, error)
	GetAllHiglightsProducts(companyGroupId float64) ([]model.Product, error)
	GetAllDailyPopularProducts(companyGroupId float64) (model.DailyProducts, error)
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
			"products.id, products.slug, products.tax, products.product_order,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, products.price2 as price2, products.price3 as price3, products.price4 as price4, products.price5 as price5,  "+
				" f.path AS path, products.is_active, popular_products.created_at, popular_products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN products ON products.id = popular_products.product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	err := query.Where(
		"products.is_active =?", true,
	).Order("popular_products.sort_order").Limit(20).Find(&popularProducts).Error

	pp.productUtil.BuildProducts(popularProducts, groupCompanyIdInt)

	return popularProducts, err

}

func (pp *PopularProductRepositoryImpl) GetAllHiglightsProducts(companyGroupId float64) (
	[]model.Product, error,
) {

	popularProducts := []model.Product{}

	groupCompanyIdInt := int(companyGroupId)

	query := pp.db.Table("highlights_products").
		Select(
			"products.id, products.slug, products.tax, products.product_order,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, highlights_products.created_at, highlights_products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN products ON products.id = highlights_products.product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	if groupCompanyIdInt != 0 {

		query = query.Select(
			"products.id, products.slug, products.tax,  products.product_order, products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, highlights_products.created_at, highlights_products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)
	}

	err := query.Where(
		"products.is_active =?", true,
	).Order("highlights_products.sort_order").Limit(20).Find(&popularProducts).Error

	pp.productUtil.BuildProducts(popularProducts, groupCompanyIdInt)

	return popularProducts, err

}

func (pp *PopularProductRepositoryImpl) GetAllDailyPopularProducts(companyGroupId float64) (
	model.DailyProducts, error,
) {

	popularProducts := []model.Product{}
	var detail model.DailyProductsInformation

	groupCompanyIdInt := int(companyGroupId)

	query := pp.db.Table("todays_popular_products").
		Select(
			"products.id, products.slug, products.tax, products.product_order,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock," +
				" brt.name AS brand_name, pt.name, todays_popular_products.video_url as video_url,  " +
				" f.path AS path, products.is_active, todays_popular_products.created_at, todays_popular_products.updated_at, " +
				" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN products ON products.id = todays_popular_products.product_id " +
				"INNER JOIN product_translations pt ON pt.product_id = products.id " +
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' " +
				"LEFT JOIN files f ON f.id = ef.file_id " +
				"INNER JOIN brands br ON br.id = products.brand_id " +
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id  ",
		)

	if groupCompanyIdInt != 0 {

		query = query.Select(
			"products.id, products.slug, products.tax,  products.product_order, products.short_desc as short_description, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, todays_popular_products.video_url as video_url, "+
				" f.path AS path, products.is_active, todays_popular_products.created_at, todays_popular_products.updated_at, "+
				" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)
	}

	err := query.Where(
		"products.is_active =?", true,
	).Order("todays_popular_products.sort_order").Limit(20).Find(&popularProducts).Error

	pp.productUtil.BuildProducts(popularProducts, groupCompanyIdInt)

	if len(popularProducts) > 0 {
		err = pp.db.Table("entity_files").
			Select("(select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_type = 'FleetCart\\\\TodaysPopularProduct' ORDER BY efs.created_at LIMIT 1) as image_path ").Limit(1).Find(&detail).Error

		detail.ImagePath = pp.productUtil.BuildImagePaths(detail.ImagePath)
		detail.VideoUrl = popularProducts[0].VideoUrl
	}

	returnModel := model.DailyProducts{Products: popularProducts, Detail: detail}

	return returnModel, err

}

func (pp *PopularProductRepositoryImpl) GetAllPopularCategories() ([]model.PopularCategoryModel, error) {

	popularCategories := []model.PopularCategoryModel{}

	err := pp.db.Select(" c.slug, ct.name, c.id, f.path").Table("popular_categories").Joins(
		"INNER JOIN categories c ON c.id = popular_categories.category_id " +
			"INNER JOIN category_translations ct ON ct.category_id = c.id " +
			"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Category\\\\Entities\\\\Category' AND ef.entity_id = c.id and ef.zone = 'logo' " +
			"LEFT JOIN files f ON f.id = ef.file_id",
	).Find(&popularCategories).Error

	pp.productUtil.BuildPopularCategory(popularCategories)

	return popularCategories, err
}
