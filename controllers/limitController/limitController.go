package limitController

import (
	"golang-dev-kreditplus/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var limits []models.Limits

	models.DB.Find(&limits)
	c.JSON(http.StatusOK, gin.H{"limits": limits})
}

func InsertLimit(c *gin.Context){
	var limits models.Limits
	
	if err := c.ShouldBindJSON(&limits); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}


	models.DB.Create(&limits)
	c.JSON(http.StatusOK, gin.H{"limits": limits})
}