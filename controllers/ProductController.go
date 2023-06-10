package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllProducts(c *gin.Context) {

	products, err := services.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

func FindProductById(c *gin.Context) {

	product, err := services.FindProductById(c.Param("id"))

	if err != nil || product.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}
