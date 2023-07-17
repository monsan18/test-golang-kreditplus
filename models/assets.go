package models

import (
	"time"
)

type Assets struct {
	AssetId		int32 `gorm:"primaryKey" json:"asset_id"`
	AssetName	string `gorm:"type:varchar(150)" json:"asset_name"`
	AssetDesc	string `gorm:"type:varchar(250)" json:"asset_desc"`
	InsertedDate	time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"inserted_date"`
}