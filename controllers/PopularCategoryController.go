package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
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

func (h *PopularCategoryController) GetPopularCategories(c *fiber.Ctx) error {
	categories, err := h.popularProductService.GetPopularCategories()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": categories})
}
