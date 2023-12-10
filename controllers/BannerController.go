package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
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

func (h *BannerController) GetAllBanners(c *fiber.Ctx) error {
	banners, err := h.bannerService.GetBanners()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": banners})
}
