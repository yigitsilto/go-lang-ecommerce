package dto

type DailyProducts struct {
	Products []Product                `json:"products"`
	Detail   DailyProductsInformation `json:"detail"`
}

type DailyProductsInformation struct {
	ImagePath string `json:"image_path"`
}
