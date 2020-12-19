package models

import "gorm.io/gorm"

// Product Model
type Product struct {
	gorm.Model
	Name  string  `gorm:"unique_index"`
	Desc  string  `sql:"type:text;"`
	Price float32 `sql:"type:decimal(10,2);"`
}
