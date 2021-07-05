package api

import (
	"fmt"
	"net/http"

	"Starter/pkg/models"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var td *models.Product = &models.Product{}

	/**
	nama
	slug
	quantity
	description
	view
	*/

	if err := c.ShouldBindJSON(td); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	err := models.CreateProduct(td)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, td)
	}
}

func UpdateProduct(c *gin.Context) {
	var product models.Product = models.Product{}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	if err := product.UpdateProduct(product, c.Params.ByName("slug")); err != nil {
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
