package tenorController

import (
	"golang-dev-kreditplus/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-dev-kreditplus/helper"
	"strconv"
)

func Index(c *gin.Context) {
	var tenor []models.Tenor

	models.DB.Find(&tenor)
	c.JSON(http.StatusOK, gin.H{"tenor": tenor})
}

func NewTenor(c *gin.Context) {
	var tenor models.Tenor
	
	tenor.Nik = c.PostForm("nik")
	var value = c.PostForm("limit_value")
	var month = c.PostForm("tenor_month")


	limitId , _ := strconv.Atoi(value)
	tenor_month , _ := strconv.Atoi(month)

	tenor.LimitId = limitId
	tenor.TenorMonth = tenor_month


	//cek apakah Nik sudah terdaftar
	if tenor.Nik == "" {
		response := helper.APIResponse("Nik Can't be Empty", http.StatusBadRequest, "Error", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else if CheckNikConsumen(tenor.Nik, c) {
		return
	}

	//cek apakah limit yang dimasukkan sudah terdaftar
	if value == "" {
		response := helper.APIResponse("Limit Value Can't be Empty", http.StatusBadRequest, "Error", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	} else {
		//limitId := CheckLimit(strconv.Atoi(value), c)
	}

	if month == "" {
		response := helper.APIResponse("Tenor month can't be empty", http.StatusBadRequest, "Error", nil)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}


	models.DB.Create(&tenor)
	c.JSON(http.StatusOK, gin.H{"new_tenor_added": tenor})
}


//fungsi untuk cek nik yang dimasukkan sudah ada di table consumen
func CheckNikConsumen(NIK string, c *gin.Context) bool {
	if NIK != "" {
		checkNik := `SELECT nik FROM consumens WHERE nik = "` + NIK + `"`
		var nik string
		if models.DB.Raw(checkNik).Scan(&nik).RowsAffected == 0 {
			response := helper.APIResponse("Nik Doesn't Exist in Consumen data", http.StatusNotFound, "Error", nil)
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return true
		}
	}
	return false
}

//fungsi untuk cek limit yang dimasukkan sudah ada di table limit
// func CheckLimit(Limit int, c *gin.Context) int {
// 	if Limit != 0 {
// 		limitId := `SELECT limit_id from limits where limit_value = ` + Limit
// 		if models.DB.Raw(checkLimit).Scan(&Limit).RowsAffected == 0 {
// 			response := helper.APIResponse("Limit Doesn't Exist in Limits data", http.StatusNotFound, "Error", nil)
// 			c.AbortWithStatusJSON(http.StatusNotFound, response)
// 			return limitId
// 		}
// 	}
// 	return 0
//}