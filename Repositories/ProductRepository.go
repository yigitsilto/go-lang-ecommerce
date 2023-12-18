package Repositories

import (
	"ecommerce/dto"
	"ecommerce/entities"
	"ecommerce/utils"
	"gorm.io/gorm"
	"strings"
)

type ProductRepository interface {
	FindPageableProductsByBrandSlug(slug string, page int, orderBy string, groupCompanyId float64) (
		dto.Pagination, error,
	)
	GetUsersCompanyGroup(user *dto.User) (float64, error)
	FindProductBySlug(slug string, groupCompanyId float64) (dto.Product, error)
	FindPageableProductsByCategorySlug(
		slug string, page int, filterBy string, order string, groupCompanyId float64,
	) (dto.Pagination, error)
	GetFiltersForProduct(categorySlug string, filterId string) ([]dto.FilterModel, error)
}

type ProductRepositoryImpl struct {
	db          *gorm.DB
	productUtil utils.ProductUtilInterface
}

func NewProductRepository(db *gorm.DB, productUtil utils.ProductUtilInterface) ProductRepository {
	return &ProductRepositoryImpl{db: db, productUtil: productUtil}
}

func (p *ProductRepositoryImpl) GetFiltersForProduct(categorySlug string, filterId string) ([]dto.FilterModel, error) {
	var filters []entities.Filters

	var filterValueIds []int
	var categoryId int
	err := p.db.Table("categories").Select("id").Where("slug=?", categorySlug).Pluck("id", &categoryId).Error

	subquery := p.db.Table("products p").
		Joins(
			"INNER JOIN product_categories pc ON p.id = pc.product_id AND p.is_active=? AND pc.category_id = ?", true,
			categoryId,
		)

	subquery = subquery.Select("p.id")

	err = p.db.Table("filter_values").
		Joins("INNER JOIN product_filter_values pfv ON filter_values.id = pfv.filter_value_id").
		Where("pfv.product_id IN (?)", subquery).
		Select("filter_values.id as filter_value_id").
		Group("filter_value_id").
		Find(&filterValueIds).Error

	err = p.db.Table("filters").Preload("Values", "id IN(?)", filterValueIds).Where(
		"status =?", true,
	).Find(&filters).Error

	return convertToFilterModel(filters), err

}
func convertToFilterModel(filters []entities.Filters) []dto.FilterModel {
	var filterListModel []dto.FilterModel

	for _, filter := range filters {
		filterModel := dto.FilterModel{
			Id:    filter.ID,
			Slug:  filter.Slug,
			Title: filter.Title,
		}

		for _, value := range filter.Values {
			filterValueModel := dto.FilterValuesModel{
				Id:    value.Id,
				Slug:  value.Slug,
				Title: value.Title,
			}
			filterModel.Values = append(filterModel.Values, filterValueModel)
		}
		filterListModel = append(filterListModel, filterModel)

	}

	return filterListModel
}

func (p *ProductRepositoryImpl) FindPageableProductsByBrandSlug(
	slug string, page int, orderBy string, groupCompanyId float64,
) (dto.Pagination, error) {
	groupCompanyIdInt := int(groupCompanyId)

	if page < 1 {
		page = 1
	}

	perPage := 50
	offset := (page - 1) * perPage

	var products []dto.Product

	query := p.db.Table("products").
		Select(
			"products.id, products.product_order, products.slug, products.short_desc as short_description, products.tax,  products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN product_translations pt ON pt.product_id = products.id "+
				"INNER JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' "+
				"INNER JOIN files f ON f.id = ef.file_id "+
				"INNER JOIN brands br ON br.id = products.brand_id "+
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id",
		).
		Where("products.is_active = true AND br.slug = ?", slug)

	if groupCompanyIdInt != 0 {
		query = query.Select(
			"products.id, products.product_order, products.product_order, products.slug, products.short_desc as short_description, products.tax,  pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)

		perPage = 60
		offset = (page - 1) * perPage
	}

	err := query.
		Offset(offset).
		Limit(perPage).
		Order(p.productUtil.BuildOrderByValues(&orderBy)).
		Find(&products).
		Error

	p.productUtil.BuildProducts(products, groupCompanyIdInt)

	pagination := dto.Pagination{Data: products}

	return pagination, err
}

func (p *ProductRepositoryImpl) GetUsersCompanyGroup(user *dto.User) (float64, error) {

	if user.Group == 0 {
		return 0, nil
	}
	userInformation := dto.UserInformation{}

	err := p.db.Table("users").Select("users.email, c.company_price_id as company_group_id ").
		Joins("INNER JOIN company c ON c.id = users.company_group_id ").
		Find(
			&userInformation, "email =?", user.Email,
		).Error

	return userInformation.CompanyGroupId, err

}

func (p *ProductRepositoryImpl) FindProductBySlug(slug string, groupCompanyId float64) (dto.Product, error) {

	groupCompanyIdInt := int(groupCompanyId)

	var products []dto.Product

	query := p.db.Table("products").
		Select(
			"products.id, products.slug, products.short_desc as short_description, products.tax, products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN product_translations pt ON pt.product_id = products.id "+
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' "+
				"LEFT JOIN files f ON f.id = ef.file_id "+
				"LEFT JOIN brands br ON br.id = products.brand_id "+
				"LEFT JOIN brand_translations brt ON brt.brand_id = br.id",
		).
		Where("products.is_active = true AND products.slug = ?", slug)

	if groupCompanyIdInt != 0 {
		query = query.Select(
			"products.id, products.slug, products.short_desc as short_description, products.tax,  pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)

	}

	err := query.
		Find(&products).
		Error

	p.productUtil.BuildProducts(products, groupCompanyIdInt)

	return products[0], err

}

func (p *ProductRepositoryImpl) FindPageableProductsByCategorySlug(
	slug string, page int, filterBy string, order string, groupCompanyId float64,
) (dto.Pagination, error) {

	groupCompanyIdInt := int(groupCompanyId)

	var id int
	err := p.db.Select("id").Table("categories").Where("slug = ?", slug).Scan(&id).Error

	if err != nil {
		return dto.Pagination{}, err
	}

	if page < 1 {
		page = 1
	}
	perPage := 12

	// Sayfalama işlemi için offset hesapla
	offset := (page - 1) * perPage

	var products []dto.Product

	query := p.db.Table("products").
		Select(
			"distinct products.id, products.slug, products.tax,  products.product_order,  products.short_desc as short_description, products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
		Joins(
			"INNER JOIN product_translations pt ON pt.product_id = products.id "+
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' "+
				"LEFT JOIN files f ON f.id = ef.file_id "+
				"INNER JOIN brands br ON br.id = products.brand_id "+
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id",
		).
		Where(
			"products.is_active = true AND EXISTS (SELECT 1 FROM product_categories AS pc WHERE pc.category_id = ? AND pc.product_id = products.id)",
			id,
		)

	if filterBy != "" {
		query.Where(
			"EXISTS (SELECT 1 FROM product_filter_values AS pfv WHERE pfv.product_id = products.id)",
		)
		filterArray := strings.Split(filterBy, ",")
		var filterValues []dto.FilterIdValues
		p.db.Select("filter_id, id").Table("filter_values").Where(
			"id IN (?)", filterArray,
		).Find(&filterValues)
		filterValueIDs := make(map[string][]string)

		for _, filterValue := range filterValues {
			filterValueIDs[filterValue.FilterId] = append(filterValueIDs[filterValue.FilterId], filterValue.Id)
		}

		for _, ids := range filterValueIDs {
			query.Where(
				" products.id IN (select pfv1.product_id from product_filter_values pfv1 WHERE pfv1.filter_value_id IN (?))",
				ids,
			)
		}

	}

	if groupCompanyIdInt != 0 {
		query = query.Select(
			"products.id, products.slug, products.product_order, products.short_desc as short_description, products.tax,   pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
			" (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image ",
		).
			Joins(
				"INNER JOIN "+
					"( SELECT product_id, MAX(company_price_id) AS max_company_price_id FROM product_prices WHERE company_price_id <= ? AND price != 0  GROUP BY product_id"+
					" ) max_pp ON max_pp.product_id = products.id  INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id = max_pp.max_company_price_id",
				groupCompanyIdInt,
			)

		offset = (page - 1) * perPage
	}

	err = query.
		Offset(offset).
		Limit(perPage).
		Order(p.productUtil.BuildOrderByValues(&order)).
		Find(&products).Error

	p.productUtil.BuildProducts(products, groupCompanyIdInt)

	pagination := dto.Pagination{Data: products}

	return pagination, err

}
