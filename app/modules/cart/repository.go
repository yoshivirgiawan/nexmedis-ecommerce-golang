package cart

import (
	"ecommerce/config"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll(userID string) ([]Cart, error)
	Add(cart Cart) (Cart, error)
	FindByID(ID string) (Cart, error)
	FindByProductID(productID string, userID string) (Cart, error)
	Update(cart Cart) (Cart, error)
	Delete(ID string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db := config.DB
	return &repository{db}
}

func (r *repository) FindAll(userID string) ([]Cart, error) {
	var carts []Cart
	if err := r.db.Preload("Product").Where("user_id = ?", userID).Order("created_at desc").Find(&carts).Error; err != nil {
		return nil, err
	}
	return carts, nil
}

func (r *repository) Add(cart Cart) (Cart, error) {
	if err := r.db.Create(&cart).Error; err != nil {
		return Cart{}, err
	}

	var newCart Cart
	if err := r.db.Preload("Product").First(&newCart, "id = ?", cart.ID).Error; err != nil {
		return Cart{}, err
	}

	return newCart, nil
}

func (r *repository) FindByID(ID string) (Cart, error) {
	var cart Cart
	if err := r.db.Preload("Product").First(&cart, "id = ?", ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Cart{}, errors.New("cart not found")
		}
		return Cart{}, err
	}
	return cart, nil
}

func (r *repository) FindByProductID(productID string, userID string) (Cart, error) {
	var cart Cart
	if err := r.db.Find(&cart, "product_id = ? AND user_id = ?", productID, userID).Error; err != nil {
		return Cart{}, err
	}
	return cart, nil
}

func (r *repository) Update(cart Cart) (Cart, error) {
	if err := r.db.Save(&cart).Error; err != nil {
		return Cart{}, err
	}

	var updatedCart Cart
	if err := r.db.Preload("Product").First(&updatedCart, "id = ?", cart.ID).Error; err != nil {
		return Cart{}, err
	}

	return updatedCart, nil
}

func (r *repository) Delete(ID string) error {
	if err := r.db.Delete(&Cart{}, "id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}
