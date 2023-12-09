package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PopularCategoryController struct {
	popularProductService services.PopularCategoriesService
}

func NewPopularCategoryController(service services.PopularCategoriesService) *PopularCategoryController {
	return &PopularCategoryController{
		popularProductService: service,
	}
}

func (h *PopularCategoryController) GetPopularCategories(c *gin.Context) {

	categories, err := h.popularProductService.GetPopularCategories()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
