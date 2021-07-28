package main

import (
	"log"

	"Starter/config"
	// "Starter/pkg/middlewares"
	"Starter/pkg/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// middlewares.Start()
}

func SetupRoutes() {
	main()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// migrating DB ["", "migrate", "drop"]
	db := config.MigrationDB("")
	_ = db

	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
