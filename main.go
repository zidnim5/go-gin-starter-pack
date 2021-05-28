package main

import (
	"fmt"
	"log"
	"os"

	"Starter/src/api"
	"Starter/src/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	middlewares.Start()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("PORT"))

	router.POST("/login", api.Login)
	log.Fatal(router.Run(":8080"))
}
