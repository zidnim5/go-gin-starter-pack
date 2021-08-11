package routes

import (
	"github.com/gin-gonic/gin"

	"Starter/pkg/controllers"
	"Starter/pkg/middlewares"
)

func Testing(c *gin.Context) {
	c.JSON(200, gin.H{
		"text": "Hello World.",
	})
}

func Routes(route *gin.Engine) {

	/** TODO separete routes */
	routesHello(route)

	route.POST("/api/login", controllers.Login)
	// route.POST("/api/todo", middlewares.Authorize(), controllers.CreateTodo)
	route.POST("/api/todo", controllers.CreateTodo)
	route.GET("/api/todo", middlewares.Authorize(), controllers.GetTodo)
	route.GET("/api/todo/:id", controllers.GetTodoId)
	route.DELETE("/api/todo/:id", controllers.DeleteTodo)
	route.PATCH("/api/todo/:id", controllers.UpdateTodo)
}
