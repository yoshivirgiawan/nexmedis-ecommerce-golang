package product

import (
	"ecommerce/helper"
)

type ProductFormatter struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func FormatProduct(product Product) ProductFormatter {
	imageURL := ""
	if product.Image != "" {
		imageURL = helper.GetAsset(product.Image)
	}

	formatter := ProductFormatter{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Image:       imageURL,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatProducts(products []Product) []ProductFormatter {
	var productsFormatter []ProductFormatter

	for _, product := range products {
		productFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, productFormatter)
	}

	return productsFormatter
}
