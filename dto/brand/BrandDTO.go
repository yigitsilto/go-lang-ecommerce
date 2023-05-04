package brand

import model "ecommerce/models"

type BrandDTO struct {
	ID          int    `json:"id"`
	Slug        string `json:"slug"`
	Translation model.BrandTranslation
}
