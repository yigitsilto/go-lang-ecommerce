package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
	"github.com/gofiber/fiber/v2"
)

type AuthService interface {
	GetMe(c *fiber.Ctx) (dto.UserMeModel, error)
}

type AuthServiceImpl struct {
	userRepository Repositories.UserRepository
}

func NewAuthService(repository Repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: repository,
	}
}

func (h *AuthServiceImpl) GetMe(c *fiber.Ctx) (dto.UserMeModel, error) {
	user := c.Locals("user")
	authUser := dto.User{}
	if user != nil {
		authUser = user.(dto.User)
	}

	return h.userRepository.FindUserByEmail(authUser.Id)
}
