package api

import (
	"net/http"

	"Starter/src/middlewares"
	"Starter/src/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var td *models.Todo
	if err := c.ShouldBindJSON(&td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	tokenAuth, err := middlewares.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	userId, err := middlewares.FetchAuth(tokenAuth)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	_ = userId
	// td.UserID = userId
	// fmt.Println(td.Title)
	//you can proceed to save the Todo to a database
	//but we will just return it to the caller here:
	c.JSON(http.StatusCreated, td)
}
