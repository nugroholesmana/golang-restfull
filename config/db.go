package config

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB di deklarasikan variabel secara global
var DB *gorm.DB

// InitDB merupakan untuk koneksikan database
func InitDB() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/experiment?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&models.Product{})

	if err != nil {
		panic("FAILED CONNECT DB")
	}
}
