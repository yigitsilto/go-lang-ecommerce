package Repositories

import (
	"ecommerce/database"
	model "ecommerce/models"
)

func GetAllSliders() ([]model.Slider, error) {

	sliders := []model.Slider{}

	err := database.Database.Table("sliders").
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
