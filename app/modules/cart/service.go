package cart

import (
	"errors"
	"fmt"
)

type Service interface {
	GetCarts(userID string) ([]Cart, error)
	GetCartByID(inputDetail GetCartDetailInput) (Cart, error)
	AddToCart(input CreateCartInput) (Cart, error)
	UpdateCart(inputDetail GetCartDetailInput, input UpdateCartInput) (Cart, error)
	DeleteCart(inputDetail GetCartDetailInput) error
}

type service struct {
	repository Repository
}

func NewService() *service {
	repository := NewRepository()
	return &service{repository}
}

func (s *service) GetCarts(userID string) ([]Cart, error) {
	carts, err := s.repository.FindAll(userID)
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *service) GetCartByID(inputDetail GetCartDetailInput) (Cart, error) {
	cart, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return cart, err
	}

	if cart.UserID != inputDetail.User.ID {
		return cart, errors.New("cart not found")
	}

	return cart, nil
}

func (s *service) AddToCart(input CreateCartInput) (Cart, error) {
	cart, err := s.repository.FindByProductID(input.ProductID, input.User.ID)
	fmt.Println(err)
	if err != nil {
		return cart, err
	}

	if cart.ID != "" {
		cart.Quantity += input.Quantity
		return s.repository.Update(cart)
	}

	cart = Cart{
		UserID:    input.User.ID,
		ProductID: input.ProductID,
		Quantity:  input.Quantity,
	}
	return s.repository.Add(cart)
}

func (s *service) UpdateCart(inputDetail GetCartDetailInput, input UpdateCartInput) (Cart, error) {
	cart, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return cart, err
	}

	if cart.UserID != inputDetail.User.ID {
		return cart, errors.New("cart not found")
	}

	cart.Quantity = input.Quantity
	return s.repository.Update(cart)
}

func (s *service) DeleteCart(inputDetail GetCartDetailInput) error {
	cart, err := s.repository.FindByID(inputDetail.ID)
	if err != nil {
		return errors.New("cart not found")
	}

	if cart.UserID != inputDetail.User.ID {
		return errors.New("cart not found")
	}
	return s.repository.Delete(inputDetail.ID)
}
