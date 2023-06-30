package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SettingController struct {
	service services.SettingsServiceInterface
}

func NewSettingController(
	service services.SettingsServiceInterface,
) *SettingController {
	return &SettingController{
		service: service,
	}
}

func (h *SettingController) GetSettings(c *gin.Context) {

	settings, err := h.service.GetSettings()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": settings})
}
