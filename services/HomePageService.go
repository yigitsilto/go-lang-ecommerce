package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"os"
)

func GetHomePage() (model.HomePageModel, error) {

	popularProducts, err := Repositories.GetAllRelatedProducts()

	buildPopularProducts(popularProducts) // build popular products array

	blogs, err := getBlogsForHomePage() // get blogs

	sliders, err := getSlidersForHomePage() // get sliders

	homePageModel := model.HomePageModel{Products: popularProducts, BlogModel: blogs, Slider: sliders}

	return homePageModel, err
}

func buildPopularProducts(popularProducts []model.RelatedProductsModel) {
	for index, product := range popularProducts {
		popularProducts[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		popularProducts[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		popularProducts[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
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
