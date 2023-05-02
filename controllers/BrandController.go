package controllers

import (
	"ecommerce/dto/brand"
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBrands(c *gin.Context) {

	brands := services.GetAllBrands()

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func FindById(c *gin.Context) {

	b, err := services.FindBrandById(c.Param("id"))

	if err != nil || b.Title == "" {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.EntityNotFoundException.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})

}

func CreateBrand(c *gin.Context) {

	var input brand.CreateBrandDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := services.CreateBrand(input)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"data": exceptions.DuplicateValueException})
		return

	}

	c.JSON(http.StatusCreated, gin.H{"data": b})

}

func DeleteBrand(c *gin.Context) {

	services.DeleteBrand(c.Param("id"))

	c.JSON(http.StatusNoContent, gin.H{"data": ""})

}

func UpdateBrand(c *gin.Context) {

	var input brand.CreateBrandDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := services.UpdateBrand(c.Param("id"), input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": exceptions.ServerError.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": b})

}
