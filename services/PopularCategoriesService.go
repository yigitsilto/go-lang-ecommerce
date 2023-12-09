package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
	"encoding/json"
)

type PopularCategoriesService interface {
	GetPopularCategories() ([]dto.PopularCategoryModel, error)
}

type PopularCategoriesServiceImpl struct {
	popularProductRepository Repositories.PopularProductRepository
	redis                    *config.RedisClient
}

func NewPopularCategoriesService(
	repository Repositories.PopularProductRepository,
	redisClient *config.RedisClient,
) PopularCategoriesService {
	return &PopularCategoriesServiceImpl{
		popularProductRepository: repository,
		redis:                    redisClient,
	}
}

func (h *PopularCategoriesServiceImpl) GetPopularCategories() ([]dto.PopularCategoryModel, error) {

	var popularCategories []dto.PopularCategoryModel
	var err error

	popularCategories, err = h.retrieveDataFromCache(popularCategories)

	if err == nil {
		return popularCategories, nil
	}

	return h.popularProductRepository.GetAllPopularCategories()

}

func (h *PopularCategoriesServiceImpl) retrieveDataFromCache(
	categories []dto.PopularCategoryModel,
) ([]dto.PopularCategoryModel, error) {
	categoriesFromCache, err := h.redis.Get("popularCategories")

	if err != nil {
		categories, err = h.popularProductRepository.GetAllPopularCategories()
		categoriesCache, _ := json.Marshal(categories)
		h.redis.Set("popularCategories", string(categoriesCache))

		return categories, err
	}
	err = json.Unmarshal([]byte(categoriesFromCache), &categories)

	if err != nil {
		return nil, err
	}

	return categories, nil
}
