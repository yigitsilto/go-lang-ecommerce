package controllers

import (
	"ecommerce/exceptions"
	model "ecommerce/models"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomePageController struct {
	homePageService services.HomePageService
}

func NewHomePageController(homePageService services.HomePageService) *HomePageController {
	return &HomePageController{
		homePageService: homePageService,
	}
}

func (h *HomePageController) GetHomePage(c *gin.Context) {

	user, _ := c.Get("user")
	authUser := model.User{}
	if user != nil {
		authUser = user.(model.User)
	}

	homePageModel, err := h.homePageService.GetHomePage(&authUser)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"data": homePageModel})
}
