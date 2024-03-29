package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	model "ecommerce/dto"
)

type BannerService interface {
	GetBanners() ([]model.Banner, error)
}

type BannerServiceImpl struct {
	repository  Repositories.BannerRepository
	redisClient *config.RedisClient
}

func NewBannerService(repository Repositories.BannerRepository, client *config.RedisClient) BannerService {
	return &BannerServiceImpl{
		repository:  repository,
		redisClient: client,
	}
}

func (s *BannerServiceImpl) GetBanners() ([]model.Banner, error) {

	var banners []model.Banner

	//redis, err := s.redisClient.Get("banners")
	var err error
	//	if err != nil {
	banners, err = s.repository.GetBanners()
	//bannersValue, _ := json.Marshal(banners)
	//s.redisClient.Set("banners", string(bannersValue))

	//	}
	//else {
	//	err = json.Unmarshal([]byte(redis), &banners)
	//}

	return banners, err
}
