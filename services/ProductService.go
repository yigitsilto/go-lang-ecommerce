package services

import (
	"ecommerce/Repositories"
	model "ecommerce/models"
)

type ProductService interface {
	GetProductsByBrand(slug string, page int, orderBy string, user *model.User) (model.Pagination, error)
	FindProductById(id string) (model.Product, error)
}

type ProductServiceImpl struct{}

func (p *ProductServiceImpl) GetProductsByBrand(
	slug string, page int, orderBy string, user *model.User,
) (model.Pagination, error) {

	userInformation, err := Repositories.GetUsersCompanyGroup(user)
	if err != nil || userInformation == 0 {

		products, err := Repositories.FindPageableProductsByBrandSlug(slug, page, orderBy)

		return products, err
	}

	products, err := Repositories.FindPageableProductsByBrandSlugWithUserPrices(slug, page, orderBy, userInformation)

	return products, err

}

func (p *ProductServiceImpl) FindProductById(id string) (model.Product, error) {

	product, err := Repositories.FindProductById(id)

	return product, err
}
