package migrations

import (
	"time"
)

type Product struct {
	ID         uint       `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	Name       string     `gorm:"not null" json:"name"`
	Category   Category   `json:"category"gorm:"foreignkey:CategoryID"`
	CategoryID uint       `gorm:"not null" json:"category_id"`
}

type Category struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `gorm:"not null;unique" json:"name"`
	Products  []Product  `gorm:"foreignkey:CategoryID" json:"products"`
}

type CategoryWithCount struct {
	ID           uint       `json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Name         string     `json:"name"`
	ProductCount int64      `json:"product_count"`
}
