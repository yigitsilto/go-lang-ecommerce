package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
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

func (h *BrandController) GetAllBrands(c *fiber.Ctx) error {
	brands, err := h.brandService.GetAllBrands()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": brands})
}

func (h *BrandController) FindById(c *fiber.Ctx) error {
	b, err := h.brandService.FindBrandById(c.Params("id"))

	if err != nil || b.Slug == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.EntityNotFoundException.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": b})
}
