package migrations

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name       string
	Category   Category
	CategoryID uint
}

type Category struct {
	gorm.Model
	Name string
}
