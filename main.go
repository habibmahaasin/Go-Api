package main

import (
	"gop-api/app/config"
	"gop-api/app/database"
	"gop-api/routes"
	"log"
	"net/http"
)

func main() {
	conf, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	db := database.Init(conf)
	router := routes.Init(db)

	router.Run(":" + conf.App.Port)
	http.ListenAndServe(":"+conf.App.Port, nil)
}
