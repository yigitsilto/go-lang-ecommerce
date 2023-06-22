package routes

import (
	"ecommerce/Repositories"
	"ecommerce/controllers"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// dependency injections for brands
	brandRepository := &Repositories.BrandRepositoryImpl{}
	brandService := services.NewBrandService(brandRepository)
	brandController := controllers.NewBrandController(brandService)

	// dependency injections for homePage
	sliderRepository := &Repositories.SliderRepositoryImpl{}
	productRepository := &Repositories.ProductRepositoryImpl{}
	popularProductsRepository := &Repositories.PopularProductRepositoryImpl{}
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
