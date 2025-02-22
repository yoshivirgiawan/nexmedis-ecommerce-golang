package order

import (
	"ecommerce/app/modules/cart"
	"ecommerce/config"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(userID string) ([]Order, error)
	Create(order Order) (Order, error)
	FindByID(ID string) (Order, error)
	Update(order Order) (Order, error)
	FindCartsByIDs(cartIDs []string) ([]cart.Cart, error)
	DeleteCarts(cartIDs []string) error
	MarkAsPaid(orderID string) error
	MarkAsUnpaid(orderID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db := config.DB
	return &repository{db}
}

func (r *repository) FindAll(userID string) ([]Order, error) {
	var orders []Order
	if err := r.db.Preload("OrderDetails.Product").Where("user_id = ?", userID).Order("created_at desc").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *repository) Create(order Order) (Order, error) {
	tx := r.db.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return Order{}, err
	}

	var newOrder Order
	if err := tx.Preload("OrderDetails.Product").First(&newOrder, "id = ?", order.ID).Error; err != nil {
		tx.Rollback()
		return Order{}, err
	}

	tx.Commit()
	return newOrder, nil
}

func (r *repository) FindByID(ID string) (Order, error) {
	var order Order
	if err := r.db.Preload("OrderDetails.Product").First(&order, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Order{}, errors.New("order not found")
		}
		return Order{}, err
	}
	return order, nil
}

func (r *repository) Update(order Order) (Order, error) {
	if err := r.db.Save(&order).Error; err != nil {
		return Order{}, err
	}

	var updatedOrder Order
	if err := r.db.Preload("OrderDetails.Product").First(&updatedOrder, "id = ?", order.ID).Error; err != nil {
		return Order{}, err
	}

	return updatedOrder, nil
}

func (r *repository) FindCartsByIDs(cartIDs []string) ([]cart.Cart, error) {
	var carts []cart.Cart
	if err := r.db.Preload("Product").Where("id IN (?)", cartIDs).Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *repository) DeleteCarts(cartIDs []string) error {
	if err := r.db.Where("id IN (?)", cartIDs).Delete(&cart.Cart{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) MarkAsPaid(orderID string) error {
	now := time.Now()
	err := r.db.Model(&Order{}).Where("id = ?", orderID).Update("paid_at", now).Update("status", "paid").Error
	return err
}

func (r *repository) MarkAsUnpaid(orderID string) error {
	err := r.db.Model(&Order{}).Where("id = ?", orderID).Update("paid_at", nil).Update("status", "pending").Error
	return err
}
