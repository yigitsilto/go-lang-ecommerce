package utils

import (
	"ecommerce/dto"
	"fmt"
	"os"
)

type ProductUtilInterface interface {
	BuildProducts(products []dto.Product, groupCompanyIdInt int)
	BuildPopularCategory(categories []dto.PopularCategoryModel)
	BuildOrderByValues(orderBy *string) string
	BuildImagePaths(imagePath string) string
}

type ProductUtilImpl struct {
}

func (pu *ProductUtilImpl) BuildProducts(products []dto.Product, companyId int) {
	for index, product := range products {
		productPrice, specialPrice := pu.calculateTax(product)
		products[index].PriceFormatted = fmt.Sprintf("%.2f TRY", productPrice)
		products[index].Price = productPrice
		if companyId > 1 {
			products[index].SpecialPrice = 0
			products[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", 0)

			// arangement of price for companies
			priceForCompany := pu.arrangeCompaniesPrices(products[index], companyId)
			products[index].Price = priceForCompany
			products[index].PriceFormatted = fmt.Sprintf("%.2f TRY", priceForCompany)

		} else {
			products[index].SpecialPrice = specialPrice
			products[index].SpecialPriceFormatted = fmt.Sprintf("%.2f TRY", specialPrice)
		}

		products[index].Path = os.Getenv("IMAGE_APP_URL") + product.Path
		if product.SecondImage != "" {
			products[index].SecondImage = os.Getenv("IMAGE_APP_URL") + product.SecondImage
		}
	}
}
func (pu *ProductUtilImpl) arrangeCompaniesPrices(product dto.Product, companyId int) float64 {
	if companyId <= 1 {
		return product.Price
	}

	switch companyId {
	case 5:
		if product.Price5 == 0 {
			return pu.arrangeCompaniesPrices(product, companyId-1)
		}
		return product.Price5
	case 4:
		if product.Price4 == 0 {
			return pu.arrangeCompaniesPrices(product, companyId-1)
		}
		return product.Price4
	case 3:
		if product.Price3 == 0 {
			return pu.arrangeCompaniesPrices(product, companyId-1)
		}
		return product.Price3
	case 2:
		if product.Price2 == 0 {
			return pu.arrangeCompaniesPrices(product, companyId-1)
		}
		return product.Price2
	default:
		return product.Price
	}
}

func (pu *ProductUtilImpl) BuildPopularCategory(categories []dto.PopularCategoryModel) {
	for index, category := range categories {
		if category.Path != "" {
			categories[index].Path = os.Getenv("IMAGE_APP_URL") + category.Path
		}
	}
}

func (pu *ProductUtilImpl) calculateTax(product dto.Product) (float64, float64) {
	tax := product.Tax
	productPrice := product.Price
	specialPrice := product.SpecialPrice
	if tax != 0 {
		productPrice += (product.Price * tax) / 100
		specialPrice += (product.SpecialPrice * tax) / 100
	}

	return productPrice, specialPrice
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
		return " products.product_order asc"
	}
}

func (pu *ProductUtilImpl) BuildImagePaths(imagePath string) string {
	if imagePath != "" {
		return os.Getenv("IMAGE_APP_URL") + imagePath
	} else {
		return ""
	}
}
