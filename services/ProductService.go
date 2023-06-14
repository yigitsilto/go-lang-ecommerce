package services

import (
	"ecommerce/Repositories"
	model "ecommerce/models"
)

func GetProductsByBrand(slug string, page int) (model.Pagination, error) {

	products, err := Repositories.FindPageableProductsByBrandSlug(slug, page)

	return products, err

}

func FindProductById(id string) (model.Product, error) {

	product, err := Repositories.FindProductById(id)

	return product, err
}
