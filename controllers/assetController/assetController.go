package assetController


import (
	"golang-dev-kreditplus/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var assets []models.Assets

	models.DB.Find(&assets)
	c.JSON(http.StatusOK, gin.H{"list_assets": assets})
}

func NewAsset(c *gin.Context){
	var asset models.Assets
	
	if err := c.ShouldBindJSON(&asset); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}


	models.DB.Create(&asset)
	c.JSON(http.StatusOK, gin.H{"new_asset_added": asset})
}