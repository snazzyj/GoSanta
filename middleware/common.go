package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func AddCommonFieldsToRequest(g *gin.Context) {
	fmt.Println("time.Now()", time.Now())
	g.Set("createdAt", time.Now())
	g.Set("createdBy", "API")
	g.Next()
}
