package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GetMe(c *gin.Context) (dto.UserMeModel, error)
}

type AuthServiceImpl struct {
	userRepository Repositories.UserRepository
}

func NewAuthService(repository Repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: repository,
	}
}

func (h *AuthServiceImpl) GetMe(c *gin.Context) (dto.UserMeModel, error) {
	user, _ := c.Get("user")
	authUser := dto.User{}
	if user != nil {
		authUser = user.(dto.User)
	}

	return h.userRepository.FindUserByEmail(authUser.Email)
}
