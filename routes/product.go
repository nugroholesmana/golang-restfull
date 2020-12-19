package routes

import (
	"../config"
	"../models"
	"github.com/gin-gonic/gin"
)

func GetHome(ctx *gin.Context) {
	items := []models.Product{}
	config.DB.Find(&items)
	ctx.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses ke home",
		"data":    items,
	})
}
func GetProduct(ctx *gin.Context) {
	productID := ctx.Param("id")

	var item models.Product

	if config.DB.First(&item, "id = ?", productID).RecordNotFound() {
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

func PostProduct(ctx *gin.Context) {
	title := ctx.PostForm("title")
	ctx.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses ke product detail",
		"title":   title,
	})
}
