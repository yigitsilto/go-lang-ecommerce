package routes

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/services"
	"ecommerce/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
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

	// Register the routers for brands.
	router.GET("/api/brands", brandController.GetAllBrands)
	router.GET("/api/brands/:id", brandController.FindById)

	// Register the routers for products.
	router.GET("/api/productsByBrand/:slug", productController.GetProductsByBrand)
	router.GET("/api/productsByCategory/:slug", productController.FindByCategorySlug)
	router.GET("/api/products/:slug", productController.FindProductBySlug)
	router.GET("/api/getFiltersForProduct", productController.FindFiltersForProducts)

	// Register the routers for homePage
	router.GET("/api/homePage", homePageController.GetHomePage)

	// related products
	router.GET("/api/relatedProducts", relatedProductController.FindAllRelatedProducts)

	// settings
	router.GET("/api/settings", settingsController.GetSettings)

	// blogs
	router.GET("/api/blogs", blogController.GetAllBlogs)
	router.GET("/api/blogs/:slug", blogController.FindById)

}
