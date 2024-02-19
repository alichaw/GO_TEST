package main

import (
	"GO_TEST/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Admin API 路由
	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/ads", controllers.CreateAd)
	}

	// Public API 路由
	router.GET("/ads", controllers.GetAds)

	router.Run(":8080") // 啟動服務器
}
