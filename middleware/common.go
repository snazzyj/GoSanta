package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func AddCommonFieldsToRequest(g *gin.Context) {
	g.Set("createdAt", time.Now())
	g.Set("createdBy", "API")
	g.Next()
}
