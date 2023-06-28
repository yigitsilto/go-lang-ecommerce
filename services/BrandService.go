package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/dto"
)

type BrandService interface {
	GetAllBrands() ([]model.Brand, error)
	FindBrandById(id string) (model.Brand, error)
}

type BrandServiceImpl struct {
	repository Repositories.BrandRepository
}

func NewBrandService(repository Repositories.BrandRepository) BrandService {
	return &BrandServiceImpl{
		repository: repository,
	}
}

func (s *BrandServiceImpl) GetAllBrands() ([]model.Brand, error) {

	brands, err := s.repository.FindAllBrands()
	return brands, err

}

func (s *BrandServiceImpl) FindBrandById(id string) (model.Brand, error) {

	b := model.Brand{}

	err := database.Database.Where("id=?", id).Preload("Translation").Find(&b).Error

	return b, err

}
