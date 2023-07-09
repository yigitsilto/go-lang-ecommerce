package utils

import (
	"ecommerce/dto"
	"fmt"
	"os"
)

type ProductUtilInterface interface {
	BuildProducts(products []dto.Product)
	BuildOrderByValues(orderBy *string) string
}

type ProductUtilImpl struct {
}

func (pu *ProductUtilImpl) BuildProducts(products []dto.Product) {
	for index, product := range products {
		products[index].PriceFormatted = fmt.Sprintf("%.2f TRY", product.Price)
		products[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		products[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}

func (pu *ProductUtilImpl) BuildOrderByValues(orderBy *string) string {
	switch *orderBy {
	case "orderByPriceAsc":
		return " price asc"
	case "orderByPriceDesc":
		return " price desc"
	case "orderByNameAsc":
		return "pt.name asc"
	case "orderByNameDesc":
		return " pt.name desc"
	default:
		return " products.created_at"
	}
}
