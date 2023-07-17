package models

type Consumen struct {
	ConsumenId int32 `gorm:"primaryKey" json:"cons_id"`
	Nik string `gorm:"type:varchar(10)" json:"nik"`
	Fullname string `gorm:"type:varchar(100)" json:"fullname"`
	LegalName string `gorm:"type:varchar(120)" json:"legal_name"`
	BirthPlace string `gorm:"type:varchar(50)" json:"birth_place"`
	BirthDate string `gorm:"type:date" json:"birth_date"`
	Salary string `gorm:"type:decimal(18,2)" json:"salary"`
	IdCardPhoto	string `gorm:"type:varchar(500)" json:"id_card_photo_url"`
	SelphiePhoto	string	`gorm:"type:varchar(500)" json:"selphie_photo"`
}