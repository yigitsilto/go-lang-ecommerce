package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type HomePageController struct {
	homePageService services.HomePageService
}

func NewHomePageController(homePageService services.HomePageService) *HomePageController {
	return &HomePageController{
		homePageService: homePageService,
	}
}

func (h *HomePageController) GetHomePage(c *fiber.Ctx) error {
	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	homePageModel, err := h.homePageService.GetHomePage(&authUser)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": homePageModel})
}
