package Repositories

import (
	model "ecommerce/models"
	"gorm.io/gorm"
)

type SliderRepository interface {
	GetAllSliders() ([]model.Slider, error)
}

type SliderRepositoryImpl struct {
	db *gorm.DB
}

func NewSliderRepository(db *gorm.DB) SliderRepository {
	return &SliderRepositoryImpl{
		db: db,
	}

}
func (s *SliderRepositoryImpl) GetAllSliders() ([]model.Slider, error) {

	sliders := []model.Slider{}

	err := s.db.Table("sliders").
		Select("sliders.id, f.path, sst.file_id").
		Joins(
			"inner join slider_slides ss on sliders.id = ss.slider_id " +
				"inner join slider_slide_translations sst on ss.id = sst.slider_slide_id " +
				"inner join files f on sst.file_id = f.id",
		).
		Limit(2).
		Find(&sliders).Error

	return sliders, err

}
