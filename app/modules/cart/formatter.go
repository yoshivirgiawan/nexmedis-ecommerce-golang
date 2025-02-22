package cart

import (
	"ecommerce/app/modules/product"
	"ecommerce/helper"
)

type CartFormatter struct {
	ID        string           `json:"id"`
	Product   ProductFormatter `json:"product"`
	Quantity  int              `json:"quantity"`
	CreatedAt string           `json:"created_at"`
	UpdatedAt string           `json:"updated_at"`
}

type ProductFormatter struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func FormatProduct(product product.Product) ProductFormatter {
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

func FormatCart(cart Cart) CartFormatter {
	formatter := CartFormatter{
		ID:        cart.ID,
		Product:   FormatProduct(cart.Product),
		Quantity:  cart.Quantity,
		CreatedAt: cart.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: cart.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatCarts(carts []Cart) []CartFormatter {
	var cartsFormatter []CartFormatter

	for _, cart := range carts {
		cartFormatter := FormatCart(cart)
		cartsFormatter = append(cartsFormatter, cartFormatter)
	}

	return cartsFormatter
}
