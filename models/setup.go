package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang_dev_kreditplus?charset=utf8&parseTime=True"))

	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Limits{}, &Consumen{}, &Assets{}, &Tenor{}, &Transaction{})
	DB = database
}
