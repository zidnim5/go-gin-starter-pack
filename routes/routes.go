package routes

import (
	"github.com/gin-gonic/gin"

	"Starter/pkg/api"
	"Starter/pkg/middlewares"
)

func Testing(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func Routes(route *gin.Engine) {
	route.GET("/helloo", Testing)
	route.POST("/api/login", api.Login)
	// route.POST("/api/todo", middlewares.Authorize(), api.CreateTodo)
	route.POST("/api/todo", api.CreateTodo)
	route.GET("/api/todo", middlewares.Authorize(), api.GetTodo)
	route.GET("/api/todo/:id", api.GetTodoId)
	route.DELETE("/api/todo/:id", api.DeleteTodo)
	route.PATCH("/api/todo/:id", api.UpdateTodo)
}
