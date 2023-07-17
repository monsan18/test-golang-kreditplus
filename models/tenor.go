package models

type Tenor struct {
	TenorId	int `gorm:"primaryKey" json:"tenor_id"`
	Nik	string `gorm:"type:varchar(10)" json:"nik"`
	LimitId	int `json:"limit_id"`
	TenorMonth	int `json:"tenor_month"`
}