package controllers

import (
	"ecommerce/database"
	"ecommerce/dto/brand"
	"ecommerce/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllBrands(c *gin.Context) {

	var brands []model.Brand

	database.Database.Find(&brands)

	c.JSON(http.StatusOK, gin.H{"data": brands})
}

func FindById(c *gin.Context) {

	b := model.Brand{}

	database.Database.Where("id=?", c.Param("id")).Find(&b)

	c.JSON(http.StatusOK, gin.H{"data": b})

}

func CreateBrand(c *gin.Context) {

	var input brand.CreateBrandDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b := model.Brand{Title: input.Title}

	err := database.Database.Create(&b).Error

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"data": "Böyle bir kayıt zaten var!"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": b})

}

func DeleteBrand(c *gin.Context) {

	id := c.Param("id")

	b := model.Brand{}

	database.Database.Where("id=?", id).Find(&b)

	err := database.Database.Delete(&b).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Bir hata meydana geldi"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"data": ""})

}
