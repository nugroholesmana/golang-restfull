package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name  string  `gorm:"unique_index"`
	Desc  string  `sql:"type:text;"`
	Price float32 `sql:"type:decimal(10,2);"`
}
