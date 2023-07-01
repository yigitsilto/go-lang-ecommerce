package utils

import (
	"ecommerce/dto"
	"fmt"
	"os"
	"sort"
)

type ProductUtilInterface interface {
	BuildProducts(products []dto.Product)
	BuildOrderByValues(orderBy *string) string
	UniqueProductsWithPriceCalculation(products []dto.Product, orderBy string) []dto.Product
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

func (pu *ProductUtilImpl) UniqueProductsWithPriceCalculation(products []dto.Product, orderBy string) []dto.Product {

	// Sıralama işlevlerini depolamak için bir map oluştur
	sortFuncMap := map[string]func(i, j int) bool{
		"orderByNameAsc": func(i, j int) bool {
			return products[i].Name < products[j].Name
		},
		"orderByNameDesc": func(i, j int) bool {
			return products[i].Name > products[j].Name
		},
		"orderByPriceAsc": func(i, j int) bool {
			return products[i].Price < products[j].Price
		},
		"orderByPriceDesc": func(i, j int) bool {
			return products[i].Price > products[j].Price
		},
	}

	// Sıralama fonksiyonunu uygula
	if sortFunc, ok := sortFuncMap[orderBy]; ok {
		sort.SliceStable(products, sortFunc)
	}

	return products
}
