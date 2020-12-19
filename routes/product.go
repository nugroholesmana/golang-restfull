package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nugroholesmana/golang-restfull/config"
	"github.com/nugroholesmana/golang-restfull/models"
	"gorm.io/gorm"
)

// GetHome Fungsi halaman utama dari home
func GetHome(ctx *gin.Context) {
	items := []models.Product{}
	config.DB.Find(&items)
	ctx.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses ke home",
		"data":    items,
	})
}

// GetProduct untuk mengambil satu product
func GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	var item models.Product

	result := config.DB.First(&item, "id = ?", productID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(404, gin.H{
			"status": "error",
			"msg":    "Record not found",
		})
		ctx.Abort()
		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses ke product detail",
		"data":    item,
	})
}

// PostProduct untuk insert data product
func PostProduct(ctx *gin.Context) {
	title := ctx.PostForm("title")
	ctx.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses ke product detail",
		"title":   title,
	})
}
