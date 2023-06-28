package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
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
	blogRepository           Repositories.BlogRepository
	redis                    *config.RedisClient
}

func NewHomePageService(
	repository Repositories.SliderRepository, popularProductsRepository Repositories.PopularProductRepository,
	productRepository Repositories.ProductRepository,
	blogRepository Repositories.BlogRepository,
	redisClient *config.RedisClient,
) HomePageService {
	return &HomePageServiceImpl{
		sliderRepository: repository, popularProductRepository: popularProductsRepository,
		productRepository: productRepository,
		blogRepository:    blogRepository,
		redis:             redisClient,
	}
}

func (h *HomePageServiceImpl) GetHomePage(user *model.User) (model.HomePageModel, error) {

	var popularProducts []model.PopularProductsModel
	var blogs []model.BlogModel
	var sliders []model.Slider

	homePageFromCache, err := h.retrieveDataFromCache(blogs, sliders)

	userInformation, err := h.productRepository.GetUsersCompanyGroup(user)
	if err != nil || userInformation == 0 {
		popularProducts, _ = h.popularProductRepository.GetAllRelatedProducts()
	} else {
		popularProducts, _ = h.popularProductRepository.GetAllRelatedProductsWithUserSpecialPrices(userInformation)
	}

	if err == nil {
		return model.HomePageModel{
			Products: popularProducts, Slider: homePageFromCache.Slider, BlogModel: homePageFromCache.BlogModel,
		}, nil
	}

	homePageModel, err := h.retrieveDataFromDatabase(blogs, sliders, user)

	return model.HomePageModel{
		Products: popularProducts, Slider: homePageModel.Slider, BlogModel: homePageModel.BlogModel,
	}, err
}

func (h *HomePageServiceImpl) retrieveDataFromDatabase(
	blogs []model.BlogModel, sliders []model.Slider, user *model.User,
) (model.HomePageModel, error) {
	var wg sync.WaitGroup

	wg.Add(2)

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
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil

}

func (h *HomePageServiceImpl) getBlogsForHomePage() ([]model.BlogModel, error) {
	var blogs []model.BlogModel

	blogs, err := h.blogRepository.GetBlogsForHomePage()

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
	blogs []model.BlogModel, sliders []model.Slider,
) (model.HomePageModel, error) {
	blogsFromCache, err := h.redis.Get("blogs")
	slidersFromCache, err := h.redis.Get("sliders")

	if err != nil {
		return model.HomePageModel{}, err
	}
	err = json.Unmarshal([]byte(blogsFromCache), &blogs)
	err = json.Unmarshal([]byte(slidersFromCache), &sliders)

	if err != nil {
		return model.HomePageModel{}, err
	}

	homePageModel := model.HomePageModel{
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil
}
