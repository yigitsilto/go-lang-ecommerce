package Repositories

import (
	model "ecommerce/dto"
	"gorm.io/gorm"
	"os"
)

type BrandRepository interface {
	FindAllBrands() ([]model.Brand, error)
}

type BrandRepositoryImpl struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &BrandRepositoryImpl{
		db: db,
	}

}

func (b *BrandRepositoryImpl) FindAllBrands() ([]model.Brand, error) {

	brands := []model.Brand{}

	err := b.db.Table("brands").Select("brands.id, brands.created_at, brands.updated_at,f.path, bt.name, brands.is_active,brands.slug").Joins(
		" INNER JOIN brand_translations bt on bt.brand_id = brands.id " +
			"LEFT JOIN entity_files ef on ef.entity_type = 'Modules\\\\Brand\\\\Entities\\\\Brand' and ef.entity_id = brands.id" +
			" LEFT JOIN files f on f.id = ef.file_id WHERE brands.is_active = true",
	).Find(&brands).Error

	addImagePathToValues(brands)

	return brands, err
}

func addImagePathToValues(brands []model.Brand) {

	for index, brand := range brands {
		brands[index].Path = os.Getenv("IMAGE_APP_URL") + brand.Path
	}

}
