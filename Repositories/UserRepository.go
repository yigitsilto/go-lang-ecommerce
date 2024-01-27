package Repositories

import (
	model "ecommerce/dto"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByEmail(id float64) (model.UserMeModel, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindUserByEmail(id float64) (model.UserMeModel, error) {

	var user model.UserMeModel

	err := u.db.Table("users").Where("id =?", int(id)).Find(&user).Error

	return user, err

}
