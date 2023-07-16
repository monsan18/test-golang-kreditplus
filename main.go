package main

import (
	"golang-dev-kreditplus/controllers/limitController"
	"golang-dev-kreditplus/controllers/consumenController"
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
	// r.DELETE("/api/product", productController.Delete)

	r.Run()

}
