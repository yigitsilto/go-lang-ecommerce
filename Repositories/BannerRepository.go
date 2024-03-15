package Repositories

import (
	model "ecommerce/dto"
	"ecommerce/utils"
	"gorm.io/gorm"
)

type BannerRepository interface {
	GetBanners() ([]model.Banner, error)
}

type BannerRepositoryImpl struct {
	db          *gorm.DB
	productUtil utils.ProductUtilInterface
}

func NewBannerRepository(db *gorm.DB, productUtil utils.ProductUtilInterface) BannerRepository {
	return &BannerRepositoryImpl{
		db:          db,
		productUtil: productUtil,
	}
}

func (b *BannerRepositoryImpl) GetBanners() ([]model.Banner, error) {
	var banners []model.Banner
	err := b.db.Table("banners").
		Select("distinct banners.id, banners.created_at, banners.updated_at, banners.title, banners.description, banners.link_url, f.path AS image ").
		Joins(
			"LEFT JOIN entity_files ef ON ef.entity_type = 'FleetCart\\\\Banner' AND ef.entity_id = banners.id and ef.zone = 'base_image'  " +
				"LEFT JOIN files f ON f.id = ef.file_id ",
		).
		Order("created_at desc").
		Where("banners.is_active = true").
		Limit(2).
		Find(&banners).Error

	return b.buildImagePaths(banners), err
}

func (b *BannerRepositoryImpl) buildImagePaths(banners []model.Banner) []model.Banner {
	for index, banner := range banners {
		banners[index].Image = b.productUtil.BuildImagePaths(banner.Image)
	}

	return banners
}
