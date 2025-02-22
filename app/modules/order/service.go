package order

import (
	"errors"
	"time"
)

type Service interface {
	GetOrders(userID string) ([]Order, error)
	GetOrderByID(inputDetail GetOrderInput) (Order, error)
	CreateOrder(input CreateOrderInput) (Order, error)
	PayOrder(inputDetail GetOrderInput) (Order, error)
}

type service struct {
	repository Repository
}

func NewService() *service {
	repository := NewRepository()
	return &service{repository}
}

func (s *service) GetOrders(userID string) ([]Order, error) {
	return s.repository.FindAll(userID)
}

func (s *service) GetOrderByID(inputDetail GetOrderInput) (Order, error) {
	return s.repository.FindByID(inputDetail.ID)
}

func (s *service) CreateOrder(input CreateOrderInput) (Order, error) {
	carts, err := s.repository.FindCartsByIDs(input.CartIDs)
	if err != nil {
		return Order{}, err
	}

	if len(carts) == 0 {
		return Order{}, errors.New("no cart items found")
	}

	order := Order{
		UserID:       input.User.ID,
		Reference:    "ORD-" + time.Now().Format("20060102150405"),
		Status:       "pending",
		Total:        0,
		PaidAt:       nil,
		OrderDetails: []OrderDetail{},
	}

	var total float64
	for _, cart := range carts {
		orderDetail := OrderDetail{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     cart.Product.Price,
		}
		total += float64(cart.Quantity) * cart.Product.Price
		order.OrderDetails = append(order.OrderDetails, orderDetail)
	}

	order.Total = total

	newOrder, err := s.repository.Create(order)
	if err != nil {
		return Order{}, err
	}

	err = s.repository.DeleteCarts(input.CartIDs)
	if err != nil {
		return Order{}, err
	}

	return newOrder, nil
}

func (s *service) PayOrder(inputDetail GetOrderInput) (Order, error) {
	s.repository.MarkAsPaid(inputDetail.ID)

	return s.repository.FindByID(inputDetail.ID)
}
