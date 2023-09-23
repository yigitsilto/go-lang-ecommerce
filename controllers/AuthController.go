package controllers

import (
	"ecommerce/exceptions"
	"ecommerce/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (h *AuthController) GetMe(c *gin.Context) {
	user, err := h.authService.GetMe(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": exceptions.ServerError.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
