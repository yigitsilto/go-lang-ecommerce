package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PopularProductController struct {
	popularProductService services.PopularProductService
}

func NewPopularProductController(service services.PopularProductService) *PopularProductController {
	return &PopularProductController{
		popularProductService: service,
	}
}

func (h *PopularProductController) GetPopularProducts(c *gin.Context) {
	user, _ := c.Get("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := h.popularProductService.GetPopularProducts(&authUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
