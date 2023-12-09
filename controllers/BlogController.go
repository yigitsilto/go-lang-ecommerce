package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
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

func (b *BlogController) GetAllBlogsByLimit(c *gin.Context) {
	brands, err := b.blogService.GetAllBlogsByLimit()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func (b *BlogController) GetAllBlogs(c *gin.Context) {
	brands, err := b.blogService.GetAllBlogs()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func (h *BlogController) FindById(c *gin.Context) {

	b, err := h.blogService.FindById(c.Param("slug"))

	if err != nil || b.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})

}
