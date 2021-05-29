package main

import (
	"log"

	"Starter/routes"
	"Starter/src/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	middlewares.Start()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
