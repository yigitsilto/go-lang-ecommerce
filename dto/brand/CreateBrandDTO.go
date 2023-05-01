package brand

type CreateBrandDTO struct {
	Title string `json:"Title" binding:"required"`
}
