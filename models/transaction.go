package models

import (
	"time"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	TransactionId 	int32	`gorm:"primaryKey" json:"transaction_id"`
	ContractNo	string	`gorm:"type:varchar(100)" json:"contract_no"`
	OtrNo	int	`json:"otr_no"`
	AdminFee	decimal.Decimal `gorm:"type:decimal(18,2)" json:"admin_fee"`
	InstallmentNumber	int	`json:"installment_number"`
	InterestNumber	decimal.Decimal `gorm:"type:decimal(18,2)" json:"interest_number"`
	AssetId	int32	`json:"asset_id"`
	Nik string	`gorm:"type:varchar(10)" json:"nik"`
	InsertedDate	time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"inserted_date"`
}

type TransactionResponse struct {
	TransactionId	string	`json:"id"`
	Nik	string	`json:"nik"`
	Fullname string	`json:"name"`
	AssetName	string	`json:"asset_name"`
	ContractNo	string	`json:"contract_no"`
	OtrNo	string	`json:"otr_no"`
	AdminFee	string `json:"admin_fee"`
	InstallmentNumber	string	`json:"installment_number"`
	InterestNumber	string	`json:"interest_number"`
	InsertedDate	string	`json:"inserted_date"`
}