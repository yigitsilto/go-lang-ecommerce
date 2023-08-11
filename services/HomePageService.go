package services

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/dto"
	"encoding/json"
	"os"
	"sync"
)

type HomePageService interface {
	GetHomePage(user *dto.User) (dto.HomePageModel, error)
	getBlogsForHomePage() ([]dto.BlogModel, error)
	getSlidersForHomePage() ([]dto.Slider, error)
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

func (h *HomePageServiceImpl) GetHomePage(user *dto.User) (dto.HomePageModel, error) {

	var popularProducts []dto.Product
	var blogs []dto.BlogModel
	var sliders []dto.Slider
	var popularCategories []dto.PopularCategoryModel

	userInformation, err := h.productRepository.GetUsersCompanyGroup(user)

	popularProducts, _ = h.popularProductRepository.GetAllRelatedProducts(userInformation)

	popularCategories, _ = h.popularProductRepository.GetAllPopularCategories()

	homePageFromCache, err := h.retrieveDataFromCache(blogs, sliders)

	if err == nil {
		return dto.HomePageModel{
			Products: popularProducts, Slider: homePageFromCache.Slider, BlogModel: homePageFromCache.BlogModel,
			PopularCategories: popularCategories,
		}, nil
	}

	homePageModel, err := h.retrieveDataFromDatabase(blogs, sliders, user)

	return dto.HomePageModel{
		Products: popularProducts, Slider: homePageModel.Slider, BlogModel: homePageModel.BlogModel,
		PopularCategories: popularCategories,
	}, err
}

func (h *HomePageServiceImpl) retrieveDataFromDatabase(
	blogs []dto.BlogModel, sliders []dto.Slider, user *dto.User,
) (dto.HomePageModel, error) {
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

	homePageModel := dto.HomePageModel{
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil

}

func (h *HomePageServiceImpl) getBlogsForHomePage() ([]dto.BlogModel, error) {
	var blogs []dto.BlogModel

	blogs, err := h.blogRepository.GetBlogsForHomePage()

	return blogs, err
}

func (h *HomePageServiceImpl) getSlidersForHomePage() ([]dto.Slider, error) {
	sliders, err := h.sliderRepository.GetAllSliders()

	for index, slider := range sliders {
		sliders[index].Path = os.Getenv("IMAGE_APP_URL") + slider.Path
	}

	return sliders, err
}

func (h *HomePageServiceImpl) retrieveDataFromCache(
	blogs []dto.BlogModel, sliders []dto.Slider,
) (dto.HomePageModel, error) {
	blogsFromCache, err := h.redis.Get("blogs")
	slidersFromCache, err := h.redis.Get("sliders")

	if err != nil {
		return dto.HomePageModel{}, err
	}
	err = json.Unmarshal([]byte(blogsFromCache), &blogs)
	err = json.Unmarshal([]byte(slidersFromCache), &sliders)

	if err != nil {
		return dto.HomePageModel{}, err
	}

	homePageModel := dto.HomePageModel{
		BlogModel: blogs,
		Slider:    sliders,
	}

	return homePageModel, nil
}
