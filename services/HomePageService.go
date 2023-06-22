package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
	"os"
	"sync"
)

type HomePageService interface {
	GetHomePage(user *model.User) (model.HomePageModel, error)
	getBlogsForHomePage() ([]model.BlogModel, error)
	getSlidersForHomePage() ([]model.Slider, error)
}

type HomePageServiceImpl struct {
	sliderRepository         Repositories.SliderRepository
	popularProductRepository Repositories.PopularProductRepository
	productRepository        Repositories.ProductRepository
}

func NewHomePageService(
	repository Repositories.SliderRepository, popularProductsRepository Repositories.PopularProductRepository,
	productRepository Repositories.ProductRepository,
) HomePageService {
	return &HomePageServiceImpl{
		sliderRepository: repository, popularProductRepository: popularProductsRepository,
		productRepository: productRepository,
	}
}

func (h *HomePageServiceImpl) GetHomePage(user *model.User) (model.HomePageModel, error) {
	var wg sync.WaitGroup

	var popularProducts []model.PopularProductsModel
	var blogs []model.BlogModel
	var sliders []model.Slider

	wg.Add(3)

	go func() {
		defer wg.Done()
		userInformation, err := h.productRepository.GetUsersCompanyGroup(user)
		if err != nil || userInformation == 0 {
			popularProducts, _ = h.popularProductRepository.GetAllRelatedProducts()
		} else {
			popularProducts, _ = h.popularProductRepository.GetAllRelatedProductsWithUserSpecialPrices(userInformation)
		}

	}()

	go func() {
		defer wg.Done()
		blogs, _ = h.getBlogsForHomePage()
	}()

	go func() {
		defer wg.Done()
		sliders, _ = h.getSlidersForHomePage()
	}()

	wg.Wait()

	homePageModel := model.HomePageModel{
		Products:  popularProducts,
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil
}

func (h *HomePageServiceImpl) getBlogsForHomePage() ([]model.BlogModel, error) {
	var blogs []model.BlogModel

	err := database.Database.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err
}

func (h *HomePageServiceImpl) getSlidersForHomePage() ([]model.Slider, error) {
	sliders, err := h.sliderRepository.GetAllSliders()

	for index, slider := range sliders {
		sliders[index].Path = os.Getenv("IMAGE_APP_URL") + slider.Path
	}

	return sliders, err
}
