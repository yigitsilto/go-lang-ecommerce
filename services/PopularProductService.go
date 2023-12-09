package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
)

type PopularProductService interface {
	GetPopularProducts(user *dto.User) ([]dto.Product, error)
}

type PopularProductServiceImpl struct {
	popularProductRepository Repositories.PopularProductRepository
	productRepository        Repositories.ProductRepository
	redis                    *config.RedisClient
}

func NewPopularProductService(
	repository Repositories.PopularProductRepository,
	productRepository Repositories.ProductRepository,
	redisClient *config.RedisClient,
) PopularProductService {
	return &PopularProductServiceImpl{
		popularProductRepository: repository,
		productRepository:        productRepository,
		redis:                    redisClient,
	}
}

func (h *PopularProductServiceImpl) GetPopularProducts(user *dto.User) ([]dto.Product, error) {
	var popularProducts []dto.Product

	userInformation, err := h.productRepository.GetUsersCompanyGroup(user)

	if err != nil {
		return nil, err
	}

	popularProducts, err = h.popularProductRepository.GetAllRelatedProducts(userInformation)
	if err != nil {
		return nil, err
	}

	return popularProducts, nil
}
