package product

import (
	"ecommerce/config"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Product, error)
	FindAllWithFilters(take int, skip int, search string) ([]Product, error)
	Save(product Product) (Product, error)
	FindByID(ID string) (Product, error)
	Update(product Product) (Product, error)
	Delete(ID string) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db := config.DB
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) FindAllWithFilters(take int, skip int, search string) ([]Product, error) {
	var products []Product
	err := r.db.Limit(take).Offset(skip).Where("name LIKE ?", "%"+search+"%").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) Save(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) FindByID(ID string) (Product, error) {
	var product Product

	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) Delete(ID string) (Product, error) {
	var product Product
	err := r.db.Where("id = ?", ID).Delete(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
