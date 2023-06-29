package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
)

type RelatedProductInterface interface {
	FindAllRelatedProducts(user *dto.User, productId string) ([]dto.Product, error)
}

type RelatedProductServiceImpl struct {
	repository        Repositories.RelatedProductRepositoryInterface
	productRepository Repositories.ProductRepository
}

func NewRelatedProductService(
	repositoryInterface Repositories.RelatedProductRepositoryInterface,
	productRepository Repositories.ProductRepository,
) RelatedProductInterface {
	return &RelatedProductServiceImpl{repository: repositoryInterface, productRepository: productRepository}
}

func (r *RelatedProductServiceImpl) FindAllRelatedProducts(user *dto.User, productId string) ([]dto.Product, error) {
	userInformation, err := r.productRepository.GetUsersCompanyGroup(user)

	products, err := r.repository.FindAllRelatedProducts(
		userInformation, productId,
	)

	if len(products) == 0 {
		products, err = r.repository.FindDummyRelatedProducts(
			userInformation,
		)
	}

	return products, err
}
