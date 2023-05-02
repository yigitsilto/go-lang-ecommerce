package controllers

import (
	"ecommerce/dto/brand"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBrands(c *gin.Context) {

	brands := services.GetAllBrands()

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func FindById(c *gin.Context) {

	b := services.FindBrandById(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": b})

}

func CreateBrand(c *gin.Context) {

	var input brand.CreateBrandDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := services.CreateBrand(input)

	c.JSON(http.StatusCreated, gin.H{"data": b})

}

func DeleteBrand(c *gin.Context) {

	services.DeleteBrand(c.Param("id"))

	c.JSON(http.StatusNoContent, gin.H{"data": ""})

}
