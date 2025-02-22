package user

import (
	"ecommerce/config"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]User, error)
	FindAllWithFilters(take int, skip int, search string) ([]User, error)
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByPhoneNumber(phoneNumber string) (User, error)
	FindByID(ID string) (User, error)
	Update(user User) (User, error)
	Delete(ID string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db := config.DB
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Where("role = ?", "user").Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) FindAllWithFilters(take int, skip int, search string) ([]User, error) {
	var users []User

	query := r.db.Model(&User{}).Where("role = ?", "user") // Filter default untuk role "user"

	// Tambahkan filter pencarian jika search tidak kosong
	if search != "" {
		query = query.Where("name LIKE ? OR email LIKE ? OR phone_number LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Tambahkan paginasi
	query = query.Offset(skip).Limit(take)

	// Jalankan query
	err := query.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Unscoped().Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByPhoneNumber(phone_number string) (User, error) {
	var user User
	err := r.db.Where("phone_number = ?", phone_number).Unscoped().Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID string) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Delete(ID string) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
