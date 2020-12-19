package main

import (
	"./config"
	"./routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	defer config.DB.Close()

	server := gin.Default()

	server.GET("/", routes.GetHome)
	go server.GET("/product/:id", routes.GetProduct)
	go server.POST("/article", routes.PostProduct)

	server.Run(":8000")
}
