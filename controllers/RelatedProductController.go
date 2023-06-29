package controllers

import (
	model "ecommerce/dto"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
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

func (r *RelatedProductController) FindAllRelatedProducts(c *gin.Context) {

	if c.Query("productId") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"data": exceptions.BadRequest.Error()})
		return
	}
	user, _ := c.Get("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	products, err := r.service.FindAllRelatedProducts(&authUser, c.Query("productId"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}
