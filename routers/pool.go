package routers

import (
	"secret-santa/controllers"

	"github.com/gin-gonic/gin"
)

func SetupPoolRouter(g *gin.Engine) {
	poolGroup := g.Group("/v1/pool")
	{
		poolGroup.GET("/all", controllers.GetPoolsFromJSONFile)
		poolGroup.GET("/:id", controllers.GetPoolById)
		poolGroup.POST("/", controllers.PostNewPool)
	}
}
