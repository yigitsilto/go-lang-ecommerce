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
	GetFiltersForProduct() ([]dto.FilterModel, error)
}

type ProductRepositoryImpl struct {
	db          *gorm.DB
	productUtil utils.ProductUtilInterface
}

func NewProductRepository(db *gorm.DB, productUtil utils.ProductUtilInterface) ProductRepository {
	return &ProductRepositoryImpl{db: db, productUtil: productUtil}
}

func (p *ProductRepositoryImpl) GetFiltersForProduct() ([]dto.FilterModel, error) {
	var filters []entities.Filters

	err := p.db.Table("filters").Select("*").
		Joins(
			" INNER JOIN filter_values fv on filters.id = fv.filter_id ",
		).
		Joins("INNER JOIN product_filter_values pfv ON fv.id = pfv.filter_value_id").
		Joins("INNER JOIN categories c ON c.slug = 'cilt-bakimi'").
		Joins("INNER JOIN product_categories pc ON pfv.product_id = pc.product_id").
		Where("filters.status =?", true).Find(&filters).Error

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

	perPage := 12
	offset := (page - 1) * perPage

	var products []dto.Product

	query := p.db.Table("products").
		Select(
			"products.id, products.slug, products.short_desc, products.tax,  products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
		Joins(
			"INNER JOIN product_translations pt ON pt.product_id = products.id "+
				"LEFT JOIN entity_files ef ON ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' AND ef.entity_id = products.id and ef.zone = 'base_image' "+
				"LEFT JOIN files f ON f.id = ef.file_id "+
				"INNER JOIN brands br ON br.id = products.brand_id "+
				"INNER JOIN brand_translations brt ON brt.brand_id = br.id",
		).
		Where("products.is_active = true AND br.slug = ?", slug)

	if groupCompanyIdInt != 0 {
		query = query.Select(
			"products.id, products.slug, products.short_desc, products.tax,  pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
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

	p.productUtil.BuildProducts(products)

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
			"products.id, products.slug, products.short_desc, products.tax, products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
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
			"products.id, products.slug, products.short_desc, products.tax,  pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
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

	err := query.
		Find(&products).
		Error

	p.productUtil.BuildProducts(products)

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
			"distinct products.id, products.slug, products.tax,   products.short_desc, products.price as price, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
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
		filterArray := strings.Split(filterBy, ",")
		query = query.Where(
			"EXISTS (SELECT 1 FROM product_filter_values AS pfv WHERE pfv.product_id = products.id AND pfv.filter_value_id IN (?))",
			filterArray,
		)
	}

	if groupCompanyIdInt != 0 {
		query = query.Select(
			"products.id, products.slug, products.short_desc, products.tax,   pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
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

	p.productUtil.BuildProducts(products)

	pagination := dto.Pagination{Data: products}

	return pagination, err

}
