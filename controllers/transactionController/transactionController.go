package transactionController

import (
	"golang-dev-kreditplus/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-dev-kreditplus/helper"
)

func Index(c *gin.Context) {
	var transaction []models.TransactionResponse

	query := `select trx.transaction_id, trx.contract_no, trx.otr_no, trx.admin_fee, trx.installment_number,
	trx.interest_number, a.asset_name, trx.nik, c.fullname, trx.inserted_date from transactions trx
	join assets a on trx.asset_id = a.asset_id 
	join consumens c on c.nik = trx.nik `

	if err := models.DB.Raw(query).Scan(&transaction).Error; err != nil {
		response := helper.APIResponse("Get All Transaction Error", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Get All Transaction Success", http.StatusOK, "success", transaction)
		c.JSON(http.StatusOK, response)
		return
}

func NewTransaction(c *gin.Context){
	var trx models.Transaction
	
	if err := c.ShouldBindJSON(&trx); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}


	models.DB.Create(&trx)
	c.JSON(http.StatusOK, gin.H{"new_transaction_added": trx})
}