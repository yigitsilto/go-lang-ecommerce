package database

var productSelect string = " products.id, products.slug, products.tax, products.product_order,  products.short_desc as short_description, products.price, products.special_price, products.qty, products.in_stock, brt.name AS brand_name, pt.name, products.price2 as price2, products.price3 as price3, products.price4 as price4, products.price5 as price5,   f.path AS path, products.is_active, popular_products.created_at, popular_products.updated_at, (select fs.path from entity_files efs INNER JOIN files fs ON fs.id = efs.file_id WHERE efs.entity_id = products.id and efs.zone != 'base_image' ORDER BY efs.created_at LIMIT 1) as second_image "

// GetProductSelectQuery returns the product select query
func GetProductSelectQuery() string {
	return productSelect
}
