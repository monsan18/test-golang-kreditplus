package models

type Limits struct {
	LimitId int32 `gorm:"primaryKey" json:"id"`
	LimitValue string `gorm:"type:decimal(18,2)" json:"value"`
}

