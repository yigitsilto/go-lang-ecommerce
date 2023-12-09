package services

import (
	"ecommerce/Repositories"
	model "ecommerce/dto"
)

type BannerService interface {
	GetBanners() ([]model.Banner, error)
}

type BannerServiceImpl struct {
	repository Repositories.BannerRepository
}

func NewBannerService(repository Repositories.BannerRepository) BannerService {
	return &BannerServiceImpl{
		repository: repository,
	}
}

func (s *BannerServiceImpl) GetBanners() ([]model.Banner, error) {
	return s.repository.GetBanners()
}
