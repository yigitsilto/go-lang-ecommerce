package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BannerController struct {
	bannerService services.BannerService
}

func NewBannerController(
	service services.BannerService,
) *BannerController {
	return &BannerController{
		bannerService: service,
	}
}

func (h *BannerController) GetAllBanners(c *gin.Context) {

	brands, err := h.bannerService.GetBanners()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brands})
}
