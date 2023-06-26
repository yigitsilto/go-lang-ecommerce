package routes

import (
	"ecommerce/Repositories"
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	db := database.Database

	// dependency injections for brands
	brandRepository := Repositories.NewBrandRepository(db)
	brandService := services.NewBrandService(brandRepository)
	brandController := controllers.NewBrandController(brandService)

	// dependency injections for homePage
	sliderRepository := Repositories.NewSliderRepository(db)
	productRepository := Repositories.NewProductRepository(db)
	popularProductsRepository := Repositories.NewPopularProductRepository(db)
	homePageService := services.NewHomePageService(sliderRepository, popularProductsRepository, productRepository)
	homePageController := controllers.NewHomePageController(homePageService)

	// dependency injections for products
	productService := services.NewProductService(productRepository)
	productController := controllers.NewProductController(productService)

	// Register the routers for brands.
	router.GET("/api/brands", brandController.GetAllBrands)
	router.GET("/api/brands/:id", brandController.FindById)

	// Register the routers for products.
	router.GET("/api/productsByBrand/:slug", productController.GetProductsByBrand)
	router.GET("/api/products/:id", productController.FindProductById)

	// Register the routers for homePage
	router.GET("/api/homePage", homePageController.GetHomePage)
}
