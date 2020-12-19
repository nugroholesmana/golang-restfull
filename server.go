package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nugroholesmana/golang-restfull/config"
	"github.com/nugroholesmana/golang-restfull/routes"
)

func main() {
	config.InitDB()

	server := gin.Default()

	server.GET("/", routes.GetHome)
	go server.GET("/product/:id", routes.GetProduct)
	go server.POST("/article", routes.PostProduct)

	server.Run(":8000")
}
