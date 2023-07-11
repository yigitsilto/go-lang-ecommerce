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
		productPrice := pu.calculateTax(product)
		products[index].PriceFormatted = fmt.Sprintf("%.2f TRY", productPrice)
		products[index].Price = productPrice
		products[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", product.SpecialPrice)
		products[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
	}
}

func (pu *ProductUtilImpl) calculateTax(product dto.Product) float64 {
	tax := product.Tax
	productPrice := product.Price
	if tax != 0 {
		productPrice += (product.Price * tax) / 100
	}
	return productPrice
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
