package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
)

type ProductService interface {
	GetProductsByBrand(slug string, page int, orderBy string, user *dto.User) (dto.Pagination, error)
	FindProductBySlug(slug string, user *dto.User) (dto.Product, error)
	GetProductsByCategorySlug(slug string, page int, filterBy string, order string, user *dto.User) (
		dto.Pagination, error,
	)
	FindFiltersForProduct(categorySlug string, filterId string) ([]dto.FilterModel, error)
}

type ProductServiceImpl struct {
	productRepository Repositories.ProductRepository
}

func NewProductService(productRepository Repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (p *ProductServiceImpl) GetProductsByBrand(
	slug string, page int, orderBy string, user *dto.User,
) (dto.Pagination, error) {

	userInformation, err := p.productRepository.GetUsersCompanyGroup(user)
	products, err := p.productRepository.FindPageableProductsByBrandSlug(slug, page, orderBy, userInformation)

	return products, err

}

func (p *ProductServiceImpl) FindProductBySlug(slug string, user *dto.User) (dto.Product, error) {

	userInformation, err := p.productRepository.GetUsersCompanyGroup(user)
	product, err := p.productRepository.FindProductBySlug(slug, userInformation)

	return product, err
}

func (p *ProductServiceImpl) FindFiltersForProduct(categorySlug string, filterId string) ([]dto.FilterModel, error) {

	filters, err := p.productRepository.GetFiltersForProduct(categorySlug, filterId)

	return filters, err
}

func (p *ProductServiceImpl) GetProductsByCategorySlug(
	slug string, page int, filterBy string, order string, user *dto.User,
) (dto.Pagination, error) {

	userInformation, err := p.productRepository.GetUsersCompanyGroup(user)

	products, err := p.productRepository.FindPageableProductsByCategorySlug(
		slug, page, filterBy, order, userInformation,
	)

	return products, err

}
