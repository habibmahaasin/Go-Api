package main

import (
	"gop-api/app/config"
	"gop-api/app/database"
	"gop-api/routes"
	"log"
	"os"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)
	router := routes.Init(db)

	router.Run(":" + os.Getenv("PORT"))
}
