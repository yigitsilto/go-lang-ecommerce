package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type RelatedProductController struct {
	service services.RelatedProductInterface
}

func NewRelatedProductController(
	service services.RelatedProductInterface,
) *RelatedProductController {
	return &RelatedProductController{
		service: service,
	}
}

func (r *RelatedProductController) FindAllRelatedProducts(c *fiber.Ctx) error {
	if c.Query("productId") == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"data": exceptions.BadRequest.Error()})
	}

	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := r.service.FindAllRelatedProducts(&authUser, c.Query("productId"))

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}
