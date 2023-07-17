package main

import (
	"golang-dev-kreditplus/controllers/limitController"
	"golang-dev-kreditplus/controllers/consumenController"
	"golang-dev-kreditplus/controllers/assetController"
	"golang-dev-kreditplus/controllers/tenorController"
	"golang-dev-kreditplus/controllers/transactionController"
	"golang-dev-kreditplus/models"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

var db *gorm.DB


func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/limits", limitController.Index)
	r.POST("/api/limits", limitController.InsertLimit)

	r.GET("/api/consumen", consumenController.Index)
	r.POST("/api/consumen", consumenController.InsertConsumen)
	
	r.GET("/api/assets", assetController.Index)
	r.POST("/api/asset", assetController.NewAsset)

	r.GET("/api/tenor", tenorController.Index)
	r.POST("/api/tenor", tenorController.NewTenor)
	
	r.GET("/api/transaction", transactionController.Index)
	r.POST("/api/transaction", transactionController.NewTransaction)

	r.Run()

}
