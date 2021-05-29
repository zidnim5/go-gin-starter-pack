package routes

import (
	"github.com/gin-gonic/gin"

	"Starter/src/api"
	"Starter/src/middlewares"
)

func Testing(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func Routes(route *gin.Engine) {
	route.GET("/helloo", Testing)
	route.POST("/login", api.Login)
	route.POST("/api/todo", middlewares.Authorize(), api.CreateTodo)
	route.POST("/todo", api.CreateTodo)
}
