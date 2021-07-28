package routes

import "github.com/gin-gonic/gin"

func routesHello (route *gin.Engine) {
	route.GET("/helloo2", Testing)
}