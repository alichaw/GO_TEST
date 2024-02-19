package controllers

import (
	"GO_TEST/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAd(c *gin.Context) {
	var newAd models.Ad
	if err := c.ShouldBindJSON(&newAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 這裡添加將廣告保存到數據庫的邏輯

	c.JSON(http.StatusCreated, newAd)
}
func GetAds(c *gin.Context) {
	var queryParams struct {
		Offset   int    `form:"offset,default=0"`
		Limit    int    `form:"limit,default=5"`
		Age      int    `form:"age"`
		Gender   string `form:"gender"` // 注意：目前的模型中沒有性別條件
		Country  string `form:"country"`
		Platform string `form:"platform"`
	}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 這裡添加根據查詢參數過濾廣告的邏輯

	c.JSON(http.StatusOK, gin.H{"items": []models.Ad{}})
}
