package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
)

type PopularProductService interface {
	GetPopularProducts(user *dto.User) ([]dto.Product, error)
	GetHighlightsProducts(user *dto.User) ([]dto.Product, error)
	GetDailyPopularProducts(user *dto.User) (dto.DailyProducts, error)
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

func (h *PopularProductServiceImpl) GetHighlightsProducts(user *dto.User) ([]dto.Product, error) {
	var popularProducts []dto.Product

	userInformation, err := h.productRepository.GetUsersCompanyGroup(user)

	if err != nil {
		return nil, err
	}

	popularProducts, err = h.popularProductRepository.GetAllHiglightsProducts(userInformation)
	if err != nil {
		return nil, err
	}

	return popularProducts, nil
}

func (h *PopularProductServiceImpl) GetDailyPopularProducts(user *dto.User) (dto.DailyProducts, error) {
	var dailyProductModel dto.DailyProducts

	userInformation, err := h.productRepository.GetUsersCompanyGroup(user)

	if err != nil {
		return dailyProductModel, err
	}

	dailyProductModel, err = h.popularProductRepository.GetAllDailyPopularProducts(userInformation)
	if err != nil {
		return dailyProductModel, err
	}

	return dailyProductModel, nil
}
