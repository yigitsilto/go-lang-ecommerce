package services

import (
	"ecommerce/Repositories"
	model "ecommerce/models"
)

type ProductService interface {
	GetProductsByBrand(slug string, page int, orderBy string, user *model.User) (model.Pagination, error)
	FindProductById(id string) (model.Product, error)
}

type ProductServiceImpl struct {
	productRepository Repositories.ProductRepository
}

func NewProductService(productRepository Repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (p *ProductServiceImpl) GetProductsByBrand(
	slug string, page int, orderBy string, user *model.User,
) (model.Pagination, error) {

	userInformation, err := p.productRepository.GetUsersCompanyGroup(user)
	if err != nil || userInformation == 0 {

		products, err := p.productRepository.FindPageableProductsByBrandSlug(slug, page, orderBy)

		return products, err
	}

	products, err := p.productRepository.FindPageableProductsByBrandSlugWithUserPrices(
		slug, page, orderBy, userInformation,
	)

	return products, err

}

func (p *ProductServiceImpl) FindProductById(id string) (model.Product, error) {

	product, err := p.productRepository.FindProductById(id)

	return product, err
}
