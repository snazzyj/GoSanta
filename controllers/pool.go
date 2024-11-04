package controllers

import (
	"fmt"
	"net/http"
	"secret-santa/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const STANDARD_FORMAT = "01/03/2006" // e.g., "MM/dd/yyyy"

func GetPoolsFromJSONFile(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetPoolJsonFile())
}

func GetPoolById(c *gin.Context) {
	poolId, _ := strconv.Atoi(c.Param("id"))
	pool := models.GetPoolById(int32(poolId), models.DecodeData[models.PoolModel](models.OpenFile("pool.json")))
	c.JSON(http.StatusOK, pool)
}
func PostNewPool(c *gin.Context) {
	exchangeDate, _ := time.Parse(STANDARD_FORMAT, c.PostForm("exchangeDate"))
	result, err := models.AddPool(models.PoolModel{
		ID:           models.GenerateRandomNumber(),
		Users:        models.ParseUserIds(c.PostForm("users")),
		Pairings:     models.ParsePairings(c.PostForm("pairings")),
		ExchangeDate: exchangeDate,
		Common: models.Common{
			CreatedAt: c.MustGet("createdAt").(time.Time),
			CreatedBy: c.MustGet("createdBy").(string),
		},
	})
	fmt.Print(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating new pool"})
		return
	}
}
