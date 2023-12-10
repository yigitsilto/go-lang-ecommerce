package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type SettingController struct {
	service services.SettingsServiceInterface
}

func NewSettingController(
	service services.SettingsServiceInterface,
) *SettingController {
	return &SettingController{
		service: service,
	}
}

func (h *SettingController) GetSettings(c *fiber.Ctx) error {

	settings, err := h.service.GetSettings()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": settings})
}
