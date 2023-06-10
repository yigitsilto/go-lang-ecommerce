package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHomePage(c *gin.Context) {
	homePageModel, err := services.GetHomePage()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": homePageModel})
}
