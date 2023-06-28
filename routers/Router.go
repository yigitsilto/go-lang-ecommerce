package routes

import (
	"ecommerce/Repositories"
	"ecommerce/config"
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	db := database.Database
	redisClient := config.NewRedisClient()

	// dependency injections for brands
	brandRepository := Repositories.NewBrandRepository(db)
	brandService := services.NewBrandService(brandRepository)
	brandController := controllers.NewBrandController(brandService)

	// dependency injections for homePage
	blogRepository := Repositories.NewBlogRepository(db)
	sliderRepository := Repositories.NewSliderRepository(db)
	productRepository := Repositories.NewProductRepository(db)
	popularProductsRepository := Repositories.NewPopularProductRepository(db)
	homePageService := services.NewHomePageService(
		sliderRepository, popularProductsRepository, productRepository, blogRepository, redisClient,
	)
	homePageController := controllers.NewHomePageController(homePageService)

	// dependency injections for products
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

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
}
