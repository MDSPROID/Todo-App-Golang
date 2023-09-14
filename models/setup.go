package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase(){
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/todoAppGolang"))
	if err != nil {
		panic(err)
	}
	
	database.AutoMigrate(&Product{})

	DB = database
}