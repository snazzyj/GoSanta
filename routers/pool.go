package routers

import (
	"secret-santa/controllers"
	"secret-santa/middleware"

	"github.com/gin-gonic/gin"
)

func SetupPoolRouter(g *gin.Engine) {
	poolGroup := g.Group("/v1/pool")
	{
		poolGroup.GET("/all", controllers.GetPoolsFromJSONFile)
		poolGroup.GET("/:id", controllers.GetPoolById)
		poolGroup.POST("/", middleware.AddCommonFieldsToRequest, controllers.PostNewPool)
	}
}
