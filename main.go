package main

import (
	"gop-api/app/database"
	"gop-api/routes"
)

func main() {
	db := database.Init()
	router := routes.Init(db)

	router.Run(":8080")
}
