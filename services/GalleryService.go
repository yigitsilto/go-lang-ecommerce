package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	model "ecommerce/dto"
	"encoding/json"
)

type GalleryService interface {
	GetGallery() ([]model.Gallery, error)
}

type GalleryServiceImpl struct {
	repository  Repositories.GalleryRepository
	redisClient *config.RedisClient
}

func NewGalleryService(repository Repositories.GalleryRepository, client *config.RedisClient) GalleryService {
	return &GalleryServiceImpl{
		repository:  repository,
		redisClient: client,
	}
}

func (s *GalleryServiceImpl) GetGallery() ([]model.Gallery, error) {

	var galleries []model.Gallery

	redis, err := s.redisClient.Get("galleries")

	if err != nil {
		galleries, err = s.repository.GetGallery()
		galleriesValue, _ := json.Marshal(galleries)
		s.redisClient.Set("galleries", string(galleriesValue))

	} else {
		err = json.Unmarshal([]byte(redis), &galleries)
	}

	return galleries, err
}
