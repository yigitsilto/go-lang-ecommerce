package Repositories

import (
	model "ecommerce/dto"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetBlogsForHomePage() ([]model.BlogModel, error)
}

type BlogRepositoryImpl struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &BlogRepositoryImpl{
		db: db,
	}
}

func (b BlogRepositoryImpl) GetBlogsForHomePage() ([]model.BlogModel, error) {
	var blogs []model.BlogModel
	err := b.db.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err
}
