package routes

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/services"
	"ecommerce/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router *fiber.App) {
	db := database.Database
	redisClient := config.NewRedisClient()
	productUtil := utils.ProductUtilImpl{}

	// blog repository
	blogRepository := Repositories.NewBlogRepository(db)

	// dependency injections for brands
	brandRepository := Repositories.NewBrandRepository(db)
	brandService := services.NewBrandService(brandRepository)
	brandController := controllers.NewBrandController(brandService)

	// dependency injections for homePage
	sliderRepository := Repositories.NewSliderRepository(db)
	productRepository := Repositories.NewProductRepository(db, &productUtil)
	popularProductsRepository := Repositories.NewPopularProductRepository(db, &productUtil)
	homePageService := services.NewHomePageService(
		sliderRepository, popularProductsRepository, productRepository, blogRepository, redisClient,
	)
	homePageController := controllers.NewHomePageController(homePageService)

	// dependency injections for products
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	// dependency injections for related products
	relatedProductRepository := Repositories.NewRelatedProductRepository(db, &productUtil)
	relatedProductService := services.NewRelatedProductService(relatedProductRepository, productRepository)
	relatedProductController := controllers.NewRelatedProductController(relatedProductService)

	// settings dependency injections
	settingsRepository := Repositories.NewSettingsRepository(db)
	settingsService := services.NewSettingsService(settingsRepository, redisClient)
	settingsController := controllers.NewSettingController(settingsService)

	// blogs dependency injections
	blogService := services.NewBlogService(blogRepository)
	blogController := controllers.NewBlogController(blogService)

	// auth dependency injections
	userRepository := Repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	// slider DI
	sliderService := services.NewSliderService(sliderRepository, redisClient)
	sliderController := controllers.NewSliderController(sliderService)

	// popular products DI
	popularProductService := services.NewPopularProductService(
		popularProductsRepository, productRepository, redisClient,
	)
	popularProductController := controllers.NewPopularProductController(popularProductService)

	// popular categories DI
	popularCategoriesService := services.NewPopularCategoriesService(popularProductsRepository, redisClient)
	popularCategoriesController := controllers.NewPopularCategoryController(popularCategoriesService)

	// banners DI
	bannerRepository := Repositories.NewBannerRepository(db, &productUtil)
	bannerService := services.NewBannerService(bannerRepository, redisClient)
	bannerController := controllers.NewBannerController(bannerService)

	// TODO home page yeni temaya geçilince kaldırılacak
	// Register the routers for brands.
	router.Get("/api/brands", brandController.GetAllBrands)
	router.Get("/api/brands/:id", brandController.FindById)

	// Register the routers for banners.
	router.Get("/api/banners", bannerController.GetAllBanners)

	// Register the routers for sliders
	router.Get("/api/sliders", sliderController.GetSlider)

	// Register the routers for popular products
	router.Get("/api/popular-products", popularProductController.GetPopularProducts)
	router.Get("/api/highlights-products", popularProductController.GetHighlightsProducts)
	router.Get("/api/daily-products", popularProductController.GetDailyPopularProducts)

	// Register the routers for popular categories
	router.Get("/api/popular-categories", popularCategoriesController.GetPopularCategories)

	// Register the routers for products.
	router.Get("/api/productsByBrand/:slug", productController.GetProductsByBrand)
	router.Get("/api/productsByCategory/:slug", productController.FindByCategorySlug)
	router.Get("/api/products/:slug", productController.FindProductBySlug)
	router.Get("/api/getFiltersForProduct", productController.FindFiltersForProducts)

	// Register the routers for homePage
	router.Get("/api/homePage", homePageController.GetHomePage)

	// related products
	router.Get("/api/relatedProducts", relatedProductController.FindAllRelatedProducts)

	// settings
	router.Get("/api/settings", settingsController.GetSettings)

	// blogs
	router.Get("/api/blogs", blogController.GetAllBlogs)
	router.Get("/api/home/blogs", blogController.GetAllBlogsByLimit)
	router.Get("/api/blogs/:slug", blogController.FindById)

	// auth
	router.Get("/api/auth/me", authController.GetMe)

}
