package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
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

func (h *SliderController) GetSlider(c *fiber.Ctx) error {
	sliders, err := h.sliderService.GetSliders()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": sliders})
}
