package routers

import (
	"secret-santa/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(g *gin.Engine) {
	userGroup := g.Group("/v1/users")

	{
		userGroup.GET("/all", controllers.GetUsersFromJsonFile)
		userGroup.GET("/:id", controllers.GetUserById)
		userGroup.POST("/", controllers.AddNewUser)
	}
}
