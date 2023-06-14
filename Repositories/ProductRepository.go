package Repositories

import (
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"os"
)

func FindPageableProductsByBrandSlug(slug string, page int, orderBy string) (model.Pagination, error) {

	if page < 1 {
		page = 1
	}
	perPage := 12

	// Sayfalama işlemi için offset hesapla
	offset := (page - 1) * perPage

	var products []model.Product

	err := database.Database.Table("products").
		Select(
			"products.id, products.slug, products.short_desc, products.price, products.special_price, products.qty, products.in_stock,"+
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
		Where("products.is_active = true AND br.slug = ?", slug).
		Offset(offset).
		Limit(perPage).
		Order(buildOrderByValues(orderBy)).
		Find(&products).Error

	buildProducts(products)

	pagination := model.Pagination{TotalRows: ProductsByBrandCount(slug), Data: products}

	return pagination, err

}

func ProductsByBrandCount(slug string) int64 {
	var count int64
	database.Database.Table("products").
		Select(
			"products.id, products.slug, products.short_desc, products.price, products.special_price, products.qty, products.in_stock,"+
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
		Where("products.is_active = true AND br.slug = ?", slug).
		Count(&count)

	return count
}

func FindProductById(id string) (model.Product, error) {

	product := model.Product{}

	err := database.Database.Table("products").Select("*").Joins(
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
		return "products.price asc"
	case "orderByPriceDesc":
		return " products.price desc"
	case "orderByNameAsc":
		return "pt.name asc"
	case "orderByNameDesc":
		return " pt.name desc"
	default:
		return " products.created_at"
	}
}
