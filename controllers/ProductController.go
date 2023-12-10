package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (p *ProductController) GetProductsByBrand(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Page parameter must be an integer"})
	}

	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := p.productService.GetProductsByBrand(c.Params("slug"), page, c.Query("order"), &authUser)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}

func (p *ProductController) FindProductBySlug(c *fiber.Ctx) error {
	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	product, err := p.productService.FindProductBySlug(c.Params("slug"), &authUser)

	if err != nil || product.Slug == "" {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.EntityNotFoundException.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": product})
}

func (p *ProductController) FindByCategorySlug(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Page parameter must be an integer"})
	}

	user := c.Locals("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := p.productService.GetProductsByCategorySlug(
		c.Params("slug"), page, c.Query("filterBy"), c.Query("order"), &authUser,
	)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": products})
}

func (p *ProductController) FindFiltersForProducts(c *fiber.Ctx) error {
	filters, err := p.productService.FindFiltersForProduct(c.Query("categorySlug"), c.Query("filterId"))

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"data": exceptions.ServerError.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"data": filters})
}
