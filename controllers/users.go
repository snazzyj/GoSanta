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
	c.JSON(
		http.StatusOK,
		models.AddUser(models.UserModel{
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
		}),
	)
}
