package Repositories

import (
	model "ecommerce/models"
	"fmt"
	"gorm.io/gorm"
	"os"
	"sort"
	"strings"
)

type ProductRepository interface {
	FindPageableProductsByBrandSlug(slug string, page int, orderBy string, groupCompanyId float64) (
		model.Pagination, error,
	)
	GetUsersCompanyGroup(user *model.User) (float64, error)
	FindProductById(id string) (model.Product, error)
	FindPageableProductsByCategorySlug(
		slug string, page int, filterBy string, order string, groupCompanyId float64,
	) (model.Pagination, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (p *ProductRepositoryImpl) FindPageableProductsByBrandSlug(
	slug string, page int, orderBy string, groupCompanyId float64,
) (model.Pagination, error) {
	groupCompanyIdInt := int(groupCompanyId)

	if page < 1 {
		page = 1
	}

	perPage := 12
	offset := (page - 1) * perPage

	var products []model.Product

	query := p.db.Table("products").
		Select(
			"products.id, products.slug, products.short_desc, products.price as price, products.special_price, products.qty, products.in_stock,"+
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
			"products.id, products.slug, products.short_desc, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
			Joins(
				"INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id  <= ? AND pp.price != 0",
				groupCompanyIdInt,
			)

		perPage = 60
		offset = (page - 1) * perPage
	}

	err := query.
		Offset(offset).
		Limit(perPage).
		Order(buildOrderByValues(orderBy)).
		Find(&products).
		Error

	if groupCompanyIdInt != 0 {
		products = uniqueProductsWithPriceCalculation(products, orderBy)
	}

	buildProducts(products)

	pagination := model.Pagination{Data: products}

	return pagination, err
}

func uniqueProductsWithPriceCalculation(products []model.Product, orderBy string) []model.Product {
	productMap := make(map[int]model.Product)
	var uniqueProducts []model.Product

	// Sıralama işlevlerini depolamak için bir map oluştur
	sortFuncMap := map[string]func(i, j int) bool{
		"orderByNameAsc": func(i, j int) bool {
			return uniqueProducts[i].Name < uniqueProducts[j].Name
		},
		"orderByNameDesc": func(i, j int) bool {
			return uniqueProducts[i].Name > uniqueProducts[j].Name
		},
		"orderByPriceAsc": func(i, j int) bool {
			return uniqueProducts[i].Price < uniqueProducts[j].Price
		},
		"orderByPriceDesc": func(i, j int) bool {
			return uniqueProducts[i].Price > uniqueProducts[j].Price
		},
	}

	for _, product := range products {
		existingProduct, ok := productMap[product.ID]
		if !ok || product.CompanyPriceId > existingProduct.CompanyPriceId {
			productMap[product.ID] = product
		}
	}

	for _, product := range productMap {
		uniqueProducts = append(uniqueProducts, product)
	}

	// Sıralama fonksiyonunu uygula
	if sortFunc, ok := sortFuncMap[orderBy]; ok {
		sort.SliceStable(uniqueProducts, sortFunc)
	}

	return uniqueProducts
}

func (p *ProductRepositoryImpl) GetUsersCompanyGroup(user *model.User) (float64, error) {

	if user.Group == 0 {
		return 0, nil
	}
	userInformation := model.UserInformation{}

	err := p.db.Table("users").Select("users.email, c.company_price_id as company_group_id ").
		Joins("INNER JOIN company c ON c.id = users.company_group_id ").
		Find(
			&userInformation, "email =?", user.Email,
		).Error

	return userInformation.CompanyGroupId, err

}

func (p *ProductRepositoryImpl) FindProductById(id string) (model.Product, error) {

	product := model.Product{}

	err := p.db.Table("products").Select("*").Joins(
		" INNER JOIN product_translations pt on pt.product_id = products.id "+
			"INNER JOIN entity_files ef on ef.entity_type = 'Modules\\\\Product\\\\Entities\\\\Product' and ef.entity_id = products.id"+
			" INNER JOIN files f on f.id = ef.file_id",
	).Where("products.id=?", id).Find(&product).Error

	return product, err

}

func buildProducts(products []model.Product) {
	for index, product := range products {
		products[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		products[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		products[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}

func buildOrderByValues(orderBy string) string {
	switch orderBy {
	case "orderByPriceAsc":
		return " price asc"
	case "orderByPriceDesc":
		return " price desc"
	case "orderByNameAsc":
		return "pt.name asc"
	case "orderByNameDesc":
		return " pt.name desc"
	default:
		return " products.created_at"
	}
}

func (p *ProductRepositoryImpl) FindPageableProductsByCategorySlug(
	slug string, page int, filterBy string, order string, groupCompanyId float64,
) (model.Pagination, error) {

	groupCompanyIdInt := int(groupCompanyId)

	var id int
	err := p.db.Select("id").Table("categories").Where("slug = ?", slug).Scan(&id).Error

	if err != nil {
		return model.Pagination{}, err
	}

	if page < 1 {
		page = 1
	}
	perPage := 12

	// Sayfalama işlemi için offset hesapla
	offset := (page - 1) * perPage

	var products []model.Product

	query := p.db.Table("products").
		Select(
			"distinct products.id, products.slug, products.short_desc, products.price as price, products.special_price, products.qty, products.in_stock,"+
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
			"products.id, products.slug, products.short_desc, pp.price as price, pp.company_price_id, products.special_price, products.qty, products.in_stock,"+
				" brt.name AS brand_name, pt.name, "+
				" f.path AS path, products.is_active, products.created_at, products.updated_at",
		).
			Joins(
				"INNER JOIN product_prices pp ON pp.product_id = products.id AND pp.company_price_id  <= ? AND pp.price != 0",
				groupCompanyIdInt,
			)

		perPage = 60
		offset = (page - 1) * perPage
	}

	err = query.
		Offset(offset).
		Limit(perPage).
		Order(buildOrderByValues(order)).
		Find(&products).Error

	if groupCompanyId != 0 {
		products = uniqueProductsWithPriceCalculation(products, order)
	}

	buildProducts(products)

	pagination := model.Pagination{Data: products}

	return pagination, err

}
