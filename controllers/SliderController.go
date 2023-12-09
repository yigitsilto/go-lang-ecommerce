package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SliderController struct {
	sliderService services.SliderService
}

func NewSliderController(sliderService services.SliderService) *SliderController {
	return &SliderController{
		sliderService: sliderService,
	}
}

func (h *SliderController) GetSlider(c *gin.Context) {

	sliders, err := h.sliderService.GetSliders()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": sliders})
}
