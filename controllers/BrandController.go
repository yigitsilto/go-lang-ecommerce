package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BrandController struct {
	brandService services.BrandService
}

func NewBrandController(
	brandService services.BrandService,
) *BrandController {
	return &BrandController{
		brandService: brandService,
	}
}

func (h *BrandController) GetAllBrands(c *gin.Context) {

	brands, err := h.brandService.GetAllBrands()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func (h *BrandController) FindById(c *gin.Context) {

	b, err := h.brandService.FindBrandById(c.Param("id"))

	if err != nil || b.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})

}
