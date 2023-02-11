package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gop-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error")
	}

	// digunakan untuk auto migrate dari entity ke database
	// db.AutoMigrate(&models.User{})

	return db
}
