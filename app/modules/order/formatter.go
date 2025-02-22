package order

import (
	"ecommerce/app/modules/product"
	"ecommerce/helper"
)

type OrderFormatter struct {
	ID           string                 `json:"id"`
	Reference    string                 `json:"reference"`
	OrderDetails []OrderDetailFormatter `json:"order_details"`
	Status       string                 `json:"status"`
	Total        float64                `json:"total"`
	PaidAt       *string                `json:"paid_at"`
	CreatedAt    string                 `json:"created_at"`
	UpdatedAt    string                 `json:"updated_at"`
}

type OrderDetailFormatter struct {
	ID        string           `json:"id"`
	Product   ProductFormatter `json:"product"`
	Quantity  int              `json:"quantity"`
	Price     float64          `json:"price"`
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

func FormatOrder(order Order) OrderFormatter {
	var PaidAt *string = nil
	if order.PaidAt != nil {
		formattedPaidAt := order.PaidAt.Format("2006-01-02 15:04:05")
		PaidAt = &formattedPaidAt
	}
	formatter := OrderFormatter{
		ID:           order.ID,
		Reference:    order.Reference,
		Status:       order.Status,
		Total:        order.Total,
		PaidAt:       PaidAt,
		CreatedAt:    order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    order.UpdatedAt.Format("2006-01-02 15:04:05"),
		OrderDetails: FormatOrderDetails(order.OrderDetails),
	}
	return formatter
}

func FormatOrders(orders []Order) []OrderFormatter {
	var ordersFormatter []OrderFormatter

	for _, order := range orders {
		orderFormatter := FormatOrder(order)
		ordersFormatter = append(ordersFormatter, orderFormatter)
	}

	return ordersFormatter
}

func FormatOrderDetail(orderDetail OrderDetail) OrderDetailFormatter {
	formatter := OrderDetailFormatter{
		ID:        orderDetail.ID,
		Product:   FormatProduct(orderDetail.Product),
		Quantity:  orderDetail.Quantity,
		Price:     orderDetail.Price,
		CreatedAt: orderDetail.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: orderDetail.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatOrderDetails(orderDetails []OrderDetail) []OrderDetailFormatter {
	var orderDetailsFormatter []OrderDetailFormatter

	for _, orderDetail := range orderDetails {
		orderDetailFormatter := FormatOrderDetail(orderDetail)
		orderDetailsFormatter = append(orderDetailsFormatter, orderDetailFormatter)
	}

	return orderDetailsFormatter
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
