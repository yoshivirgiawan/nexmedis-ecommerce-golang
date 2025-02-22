package cart

import (
	"ecommerce/app/modules/product"
	"ecommerce/app/modules/user"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Cart struct {
	ID        string          `gorm:"primaryKey;type:varchar(36)"`
	UserID    string          `gorm:"not null;type:varchar(36)"`
	User      user.User       `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID string          `gorm:"not null;type:varchar(36)"`
	Product   product.Product `gorm:"foreignKey:product_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity  int             `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (cart *Cart) BeforeCreate(tx *gorm.DB) (err error) {
	cart.ID = uuid.New().String()
	return
}
