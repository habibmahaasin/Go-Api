package main

import "github.com/gin-gonic/gin"

func main() {
	// conf, err := config.Init()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// db := database.Init(conf)
	// router := routes.Init(db)
	router := gin.Default()
	router.Run(":8080")
}
