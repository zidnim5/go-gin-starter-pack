package api

import (
	"fmt"
	"net/http"

	"Starter/src/models"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var td *models.Todo = &models.Todo{}
	if err := c.ShouldBindJSON(td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	err := models.CreateTodo(td)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, td)
	}
}

func GetTodo(c *gin.Context) {
	var todo []models.Todo
	if err := models.GetAllTodo(&todo); err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetTodoId(c *gin.Context) {
	var todo *models.Todo = &models.Todo{}
	if err := models.GetTodoById(todo, c.Params.ByName("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.UpdateTodo(&todo, c.Params.ByName("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Data not found",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Updated",
	})

}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	if err := models.DeleteTodo(&todo, c.Params.ByName("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}
