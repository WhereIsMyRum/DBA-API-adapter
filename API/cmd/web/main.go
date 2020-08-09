package main

import (
	"dba-scraper.com/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	ginInstance := gin.Default()
	ginInstance.Use(gin.Recovery())
	ginInstance.GET("/api/*type", controllers.Fetch)
	ginInstance.Run(":10000")
}
