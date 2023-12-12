package Repositories

import (
	model "ecommerce/dto"
	"gorm.io/gorm"
)

type GalleryRepository interface {
	GetGallery() ([]model.Gallery, error)
}

type GalleryRepositoryImpl struct {
	db *gorm.DB
}

func NewGalleryRepository(db *gorm.DB) GalleryRepository {
	return &GalleryRepositoryImpl{
		db: db,
	}
}

func (b *GalleryRepositoryImpl) GetGallery() ([]model.Gallery, error) {
	var banners []model.Gallery
	err := b.db.Table("galleries").
		Select("galleries.image_path, galleries.link_url, galleries.id ").
		Order("created_at desc").
		Limit(8).
		Find(&banners).Error

	return banners, err
}
