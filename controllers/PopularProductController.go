package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
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

func (h *PopularProductController) GetPopularProducts(c *fiber.Ctx) error {
	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := h.popularProductService.GetPopularProducts(&authUser)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}

func (h *PopularProductController) GetHighlightsProducts(c *fiber.Ctx) error {
	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := h.popularProductService.GetHighlightsProducts(&authUser)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}

func (h *PopularProductController) GetDailyPopularProducts(c *fiber.Ctx) error {
	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := h.popularProductService.GetDailyPopularProducts(&authUser)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}
