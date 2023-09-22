package services

import (
	"ecommerce/Repositories"
	"ecommerce/dto"
)

type BlogService interface {
	GetAllBlogs() ([]dto.BlogModel, error)
	FindById(slug string) (dto.BlogLongModel, error)
}

type BlogServiceImpl struct {
	blogRepository Repositories.BlogRepository
}

func NewBlogService(repository Repositories.BlogRepository) BlogService {
	return &BlogServiceImpl{
		blogRepository: repository,
	}
}

func (b *BlogServiceImpl) GetAllBlogs() ([]dto.BlogModel, error) {
	return b.blogRepository.GetAllBlogs()
}

func (b *BlogServiceImpl) FindById(slug string) (dto.BlogLongModel, error) {
	return b.blogRepository.FindById(slug)
}
