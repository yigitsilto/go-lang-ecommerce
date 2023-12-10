package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type BlogController struct {
	blogService services.BlogService
}

func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{
		blogService: service,
	}
}

func (b *BlogController) GetAllBlogsByLimit(c *fiber.Ctx) error {
	brands, err := b.blogService.GetAllBlogsByLimit()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": brands})
}

func (b *BlogController) GetAllBlogs(c *fiber.Ctx) error {
	brands, err := b.blogService.GetAllBlogs()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": brands})
}

func (h *BlogController) FindById(c *fiber.Ctx) error {

	b, err := h.blogService.FindById(c.Params("slug"))

	if err != nil || b.Slug == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.EntityNotFoundException.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": b})
}
