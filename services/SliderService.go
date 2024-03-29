package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
	"encoding/json"
	"os"
)

type SliderService interface {
	GetSliders() ([]dto.Slider, error)
}

type SliderServiceImpl struct {
	sliderRepository Repositories.SliderRepository
	redis            *config.RedisClient
}

func NewSliderService(
	repository Repositories.SliderRepository,
	redisClient *config.RedisClient,
) SliderService {
	return &SliderServiceImpl{
		sliderRepository: repository,
		redis:            redisClient,
	}
}

func (h *SliderServiceImpl) GetSliders() ([]dto.Slider, error) {
	var sliders []dto.Slider
	/*slidersFromCache, err := h.retrieveDataFromCache(sliders)

	if err == nil {
		return slidersFromCache, nil
	}*/

	sliderFromDatabase, err := h.retrieveDataFromDatabase(sliders)

	return sliderFromDatabase, err

}

func (h *SliderServiceImpl) retrieveDataFromDatabase(
	sliders []dto.Slider,
) ([]dto.Slider, error) {

	sliders, _ = h.getSlidersFromDatabase()
	slidersCache, _ := json.Marshal(sliders)
	h.redis.Set("sliders", string(slidersCache))

	return sliders, nil

}

func (h *SliderServiceImpl) getSlidersFromDatabase() ([]dto.Slider, error) {
	sliders, err := h.sliderRepository.GetAllSliders()

	for index, slider := range sliders {
		sliders[index].Path = os.Getenv("IMAGE_APP_URL") + slider.Path
	}

	return sliders, err
}

func (h *SliderServiceImpl) retrieveDataFromCache(
	sliders []dto.Slider,
) ([]dto.Slider, error) {
	slidersFromCache, err := h.redis.Get("sliders")

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(slidersFromCache), &sliders)

	if err != nil {
		return nil, err
	}

	return sliders, nil
}
