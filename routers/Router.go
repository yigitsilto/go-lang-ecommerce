package routes

import (
	"ecommerce/controllers"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	// dependency injections
	brandService := &services.BrandServiceImpl{}
	brandController := controllers.NewBrandController(brandService)

	productService := &services.ProductServiceImpl{}
	productController := controllers.NewProductController(productService)

	homePageService := &services.HomePageServiceImpl{}
	homePageController := controllers.NewHomePageController(homePageService)

	// Register the routers for brands.
	router.GET("/api/brands", brandController.GetAllBrands)
	router.GET("/api/brands/:id", brandController.FindById)

	// Register the routers for products.
	router.GET("/api/productsByBrand/:slug", productController.GetProductsByBrand)
	router.GET("/api/products/:id", productController.FindProductById)

	// Register the routers for homePage
	router.GET("/api/homePage", homePageController.GetHomePage)
}
