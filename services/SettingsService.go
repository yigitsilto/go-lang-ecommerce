package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
)

type SettingsServiceInterface interface {
	GetSettings() (dto.GeneralSettingsModel, error)
}

type SettingsServiceImpl struct {
	repository  Repositories.SettingsRepositoryInterface
	redisClient *config.RedisClient
}

func NewSettingsService(
	repository Repositories.SettingsRepositoryInterface, client *config.RedisClient,
) SettingsServiceInterface {
	return &SettingsServiceImpl{repository: repository, redisClient: client}
}

func (s *SettingsServiceImpl) GetSettings() (dto.GeneralSettingsModel, error) {

	//redis, err := s.redisClient.Get("settings")
	var settings dto.GeneralSettingsModel
	var err error
	//if err != nil {
	settings, err = s.repository.GetSettings()
	//settingsValue, _ := json.Marshal(settings)

	//s.redisClient.Set("settings", string(settingsValue))
	//}
	//else {
	//		err = json.Unmarshal([]byte(redis), &settings)
	//}

	return settings, err

}
