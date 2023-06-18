package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
	"os"
	"sync"
)

func GetHomePage(user *model.User) (model.HomePageModel, error) {
	var wg sync.WaitGroup

	var popularProducts []model.PopularProductsModel
	var blogs []model.BlogModel
	var sliders []model.Slider

	wg.Add(3)

	go func() {
		defer wg.Done()
		userInformation, err := Repositories.GetUsersCompanyGroup(user)
		if err != nil || userInformation == 0 {
			popularProducts, _ = Repositories.GetAllRelatedProducts()
		} else {
			popularProducts, _ = Repositories.GetAllRelatedProductsWithUserSpecialPrices(userInformation)
		}

	}()

	go func() {
		defer wg.Done()
		blogs, _ = getBlogsForHomePage()
	}()

	go func() {
		defer wg.Done()
		sliders, _ = getSlidersForHomePage()
	}()

	wg.Wait()

	homePageModel := model.HomePageModel{
		Products:  popularProducts,
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil
}

func getBlogsForHomePage() ([]model.BlogModel, error) {
	var blogs []model.BlogModel

	err := database.Database.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err
}

func getSlidersForHomePage() ([]model.Slider, error) {
	sliders, err := Repositories.GetAllSliders()

	for index, slider := range sliders {
		sliders[index].Path = os.Getenv("IMAGE_APP_URL") + slider.Path
	}

	return sliders, err
}
