package controllers

import (
	"ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type GalleryController struct {
	service services.GalleryService
}

func NewGalleryController(
	service services.GalleryService,
) *GalleryController {
	return &GalleryController{
		service: service,
	}
}

func (h *GalleryController) GetAllGalleries(c *fiber.Ctx) error {
	var galleries []dto.Gallery
	var err error
	galleries, err = h.service.GetGallery()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": galleries})
}
