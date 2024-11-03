package controllers

import (
	"net/http"
	"secret-santa/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsersFromJsonFile(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetUserJSONFile())
}
func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, _ := models.GetUserById(int32(userId), models.DecodeData[models.UserModel](models.OpenFile("users.json")))
	c.JSON(http.StatusOK, user)
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
