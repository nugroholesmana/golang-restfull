package models

// Product Model
type Product struct {
	Name  string  `gorm:"unique_index"`
	Desc  string  `sql:"type:text;"`
	Price float32 `sql:"type:decimal(10,2);"`
}
