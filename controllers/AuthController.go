package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (h *AuthController) GetMe(c *fiber.Ctx) error {
	user, err := h.authService.GetMe(c)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}
