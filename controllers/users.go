package controllers

import (
	"net/http"

	"secret-santa/models"

	"github.com/gin-gonic/gin"
)

func GetUsersFromJsonFile(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUserJSONFile())
}
func AddNewUser(c *gin.Context) {
	result, err := models.AddUser(models.UserModel{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
	} else {
		c.JSON(
			http.StatusOK,
			result,
		)
	}
}
