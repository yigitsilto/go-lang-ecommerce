package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/database"
	model "ecommerce/models"
	"encoding/json"
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
	redis                    *config.RedisClient
}

func NewHomePageService(
	repository Repositories.SliderRepository, popularProductsRepository Repositories.PopularProductRepository,
	productRepository Repositories.ProductRepository,
	redisClient *config.RedisClient,
) HomePageService {
	return &HomePageServiceImpl{
		sliderRepository: repository, popularProductRepository: popularProductsRepository,
		productRepository: productRepository,
		redis:             redisClient,
	}
}

func (h *HomePageServiceImpl) GetHomePage(user *model.User) (model.HomePageModel, error) {
	var wg sync.WaitGroup

	var popularProducts []model.PopularProductsModel
	var blogs []model.BlogModel
	var sliders []model.Slider

	homePageFromCache, err := h.retrieveDataFromCache(blogs, sliders, popularProducts)

	if err == nil {
		return homePageFromCache, nil
	}

	wg.Add(3)

	go func() {
		defer wg.Done()
		userInformation, err := h.productRepository.GetUsersCompanyGroup(user)
		if err != nil || userInformation == 0 {
			popularProducts, _ = h.popularProductRepository.GetAllRelatedProducts()
		} else {
			popularProducts, _ = h.popularProductRepository.GetAllRelatedProductsWithUserSpecialPrices(userInformation)
		}
		popularProductsCache, _ := json.Marshal(popularProducts)
		h.redis.Set("products", string(popularProductsCache))

	}()

	go func() {
		defer wg.Done()
		blogs, _ = h.getBlogsForHomePage()
		blogsCache, _ := json.Marshal(blogs)
		h.redis.Set("blogs", string(blogsCache))
	}()

	go func() {
		defer wg.Done()
		sliders, _ = h.getSlidersForHomePage()
		slidersCache, _ := json.Marshal(sliders)
		h.redis.Set("sliders", string(slidersCache))
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

func (h *HomePageServiceImpl) retrieveDataFromCache(
	blogs []model.BlogModel, sliders []model.Slider, popularProducts []model.PopularProductsModel,
) (model.HomePageModel, error) {
	popularProductsFromCache, err := h.redis.Get("products")
	blogsFromCache, err := h.redis.Get("blogs")
	slidersFromCache, err := h.redis.Get("sliders")

	if err != nil {
		return model.HomePageModel{}, err
	}
	err = json.Unmarshal([]byte(popularProductsFromCache), &popularProducts)
	err = json.Unmarshal([]byte(blogsFromCache), &blogs)
	err = json.Unmarshal([]byte(slidersFromCache), &sliders)

	if err != nil {
		return model.HomePageModel{}, err
	}

	homePageModel := model.HomePageModel{
		Products:  popularProducts,
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil
}
