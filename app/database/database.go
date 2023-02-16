package database

import (
	"gop-api/app/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init(conf config.Conf) *gorm.DB {
	dsn := "host=" + conf.Database.Host + " user=" + conf.Database.User + " password=" + conf.Database.Pass + " dbname=" + conf.Database.Name + " port=" + conf.Database.Port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// dsn := conf.Database.User + ":" + conf.Database.Pass + "@tcp(" + conf.Database.Host + ":" + conf.Database.Port + ")/" + conf.Database.Name + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB Connection error")
	}

	// digunakan untuk auto migrate dari entity ke database
	// db.AutoMigrate(&models.User{})

	return db
}
