package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetProductsByBrand(c *gin.Context) {

	page, err := strconv.Atoi(c.Query("page"))

	products, err := services.GetProductsByBrand(c.Param("slug"), page)

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
