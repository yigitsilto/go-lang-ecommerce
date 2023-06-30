package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
)

type SettingsServiceInterface interface {
	GetSettings() (dto.GeneralSettingsModel, error)
}

type SettingsServiceImpl struct {
	repository Repositories.SettingsRepositoryInterface
}

func NewSettingsService(repository Repositories.SettingsRepositoryInterface) SettingsServiceInterface {
	return &SettingsServiceImpl{repository: repository}
}

func (s *SettingsServiceImpl) GetSettings() (dto.GeneralSettingsModel, error) {
	settings, err := s.repository.GetSettings()

	return settings, err

}
