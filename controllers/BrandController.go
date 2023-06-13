package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBrands(c *gin.Context) {

	brands, err := services.GetAllBrands()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func FindById(c *gin.Context) {

	b, err := services.FindBrandById(c.Param("id"))

	if err != nil || b.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})

}
