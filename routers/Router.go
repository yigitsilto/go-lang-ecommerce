package routes

import (
	"ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Register the routers for brands.
	router.GET("/api/brands", controllers.GetAllBrands)
	router.GET("/api/brands/:id", controllers.FindById)

	// Register the routers for products.
	router.GET("/api/productsByBrand/:slug", controllers.GetProductsByBrand)
	router.GET("/api/products/:id", controllers.FindProductById)
	router.GET("/api/homePage", controllers.GetHomePage)
}
