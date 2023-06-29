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
	productMap := make(map[int]dto.Product)
	var uniqueProducts []dto.Product

	// Sıralama işlevlerini depolamak için bir map oluştur
	sortFuncMap := map[string]func(i, j int) bool{
		"orderByNameAsc": func(i, j int) bool {
			return uniqueProducts[i].Name < uniqueProducts[j].Name
		},
		"orderByNameDesc": func(i, j int) bool {
			return uniqueProducts[i].Name > uniqueProducts[j].Name
		},
		"orderByPriceAsc": func(i, j int) bool {
			return uniqueProducts[i].Price < uniqueProducts[j].Price
		},
		"orderByPriceDesc": func(i, j int) bool {
			return uniqueProducts[i].Price > uniqueProducts[j].Price
		},
	}

	for _, product := range products {
		existingProduct, ok := productMap[product.ID]
		if !ok || product.CompanyPriceId > existingProduct.CompanyPriceId {
			productMap[product.ID] = product
		}
	}

	for _, product := range productMap {
		uniqueProducts = append(uniqueProducts, product)
	}

	// Sıralama fonksiyonunu uygula
	if sortFunc, ok := sortFuncMap[orderBy]; ok {
		sort.SliceStable(uniqueProducts, sortFunc)
	}

	return uniqueProducts
}
