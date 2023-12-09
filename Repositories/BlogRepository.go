package Repositories

import (
	model "ecommerce/dto"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetBlogsForHomePage() ([]model.BlogModel, error)
	GetAllBlogs() ([]model.BlogModel, error)
	GetAllBlogsByLimit(limit int) ([]model.BlogModel, error)
	FindById(slug string) (model.BlogLongModel, error)
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

func (b BlogRepositoryImpl) GetAllBlogs() ([]model.BlogModel, error) {
	var blogs []model.BlogModel
	err := b.db.Table("blogs").Order("created_at desc").Find(&blogs).Error

	return blogs, err
}

func (b BlogRepositoryImpl) GetAllBlogsByLimit(limit int) ([]model.BlogModel, error) {
	var blogs []model.BlogModel
	err := b.db.Table("blogs").Order("created_at desc").Limit(limit).Find(&blogs).Error

	return blogs, err
}

func (b BlogRepositoryImpl) FindById(slug string) (model.BlogLongModel, error) {
	var blog model.BlogLongModel
	err := b.db.Table("blogs").Where("slug =?", slug).Find(&blog).Error

	return blog, err
}
