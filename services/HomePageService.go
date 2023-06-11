package services

import (
	"ecommerce/Repositories"
	"ecommerce/database"
	model "ecommerce/models"
	"fmt"
	"os"
)

func GetHomePage() (model.HomePageModel, error) {

	popularProducts, err := Repositories.GetAllRelatedProducts()

	buildPopularProducts(popularProducts)

	blogs, err := getBlogsForHomePage()

	homePageModel := model.HomePageModel{Products: popularProducts, BlogModel: blogs}

	return homePageModel, err
}

func buildPopularProducts(popularProducts []model.RelatedProductsModel) {
	for index, product := range popularProducts {
		popularProducts[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		popularProducts[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		popularProducts[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}

func getBlogsForHomePage() ([]model.BlogModel, error) {

	blogs := []model.BlogModel{}

	err := database.Database.Table("blogs").Limit(2).Find(&blogs).Error

	return blogs, err

}
