package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
	"os"
)

func GetHomePage() (model.HomePageModel, error) {

	popularProducts, err := Repositories.GetAllRelatedProducts()

	blogs, err := getBlogsForHomePage() // get blogs

	sliders, err := getSlidersForHomePage() // get sliders

	homePageModel := model.HomePageModel{Products: popularProducts, BlogModel: blogs, Slider: sliders}

	return homePageModel, err
}

func getSlidersForHomePage() ([]model.Slider, error) {
	sliders, err := Repositories.GetAllSliders()

	for index, slider := range sliders {
		sliders[index].Path = os.Getenv("IMAGE_APP_URL") + slider.Path
	}

	return sliders, err

}

func getBlogsForHomePage() ([]model.BlogModel, error) {

	var blogs []model.BlogModel

	err := database.Database.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err

}
