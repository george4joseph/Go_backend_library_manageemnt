package config

import (
	"hello/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("/home/mary/Desktop/ma/Libm.db"), &gorm.Config{QueryFields: true})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Member{})
	db.AutoMigrate(&models.Book{})
	DB = db

}
