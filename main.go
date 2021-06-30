package main

import (
	"fmt"
	"log"
	"os"

	"Starter/databases"
	"Starter/pkg/middlewares"
	"Starter/routes"

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

	// migrating DB ["", "migrate", "drop"]
	db := databases.MigrationDB("")
	_ = db

	fmt.Println(os.Getenv("DB_NAME"))
	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":8080"))
}
