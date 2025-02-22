package product

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          string  `gorm:"primaryKey;type:varchar(36)"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Image       string
	DeletedAt   *time.Time `gorm:"index"` // Menandakan waktu saat data dihapus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (product *Product) BeforeCreate(tx *gorm.DB) (err error) {
	product.ID = uuid.New().String()
	return
}
