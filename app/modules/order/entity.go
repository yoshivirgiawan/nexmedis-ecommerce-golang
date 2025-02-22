package order

import (
	"ecommerce/app/modules/product"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID           string `gorm:"primaryKey;type:varchar(36)"`
	UserID       string `gorm:"not null;type:varchar(36)"`
	Reference    string
	Status       string
	Total        float64
	PaidAt       *time.Time `gorm:"default:null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	OrderDetails []OrderDetail
}

func (order *Order) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.New().String()
	return
}

type OrderDetail struct {
	ID        string          `gorm:"primaryKey;type:varchar(36)"`
	OrderID   string          `gorm:"not null;type:varchar(36)"`
	Order     Order           `gorm:"foreignKey:order_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID string          `gorm:"not null;type:varchar(36)"`
	Product   product.Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity  int             `gorm:"not null"`
	Price     float64         `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (orderDetail *OrderDetail) BeforeCreate(tx *gorm.DB) (err error) {
	orderDetail.ID = uuid.New().String()
	return
}
