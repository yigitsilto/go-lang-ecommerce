package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (p *ProductController) GetProductsByBrand(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page parameter must be an integer"})
		return
	}
	user, _ := c.Get("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := p.productService.GetProductsByBrand(c.Param("slug"), page, c.Query("order"), &authUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (p *ProductController) FindProductById(c *gin.Context) {

	product, err := p.productService.FindProductById(c.Param("id"))

	if err != nil || product.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func (p *ProductController) FindByCategorySlug(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page parameter must be an integer"})
		return
	}
	user, _ := c.Get("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := p.productService.GetProductsByCategorySlug(
		c.Param("slug"), page, c.Query("filterBy"), c.Query("order"), &authUser,
	)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func (p *ProductController) FindFiltersForProducts(c *gin.Context) {

	filters, err := p.productService.FindFiltersForProduct()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"data": filters})
}
