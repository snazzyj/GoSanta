package controllers

import (
	"fmt"
	"math/rand/v2"
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
func GenerateShuffledIndexes(arrayLength int) []int32 {
	indexes := make([]int32, arrayLength)
	for i := 0; i < arrayLength; i++ {
		indexes[i] = int32(i)
	}
	for {
		// Shuffle the indexes
		rand.Shuffle(arrayLength, func(i, j int) {
			indexes[i], indexes[j] = indexes[j], indexes[i]
		})
		// Check for derangement
		isDerangement := true
		for i := 0; i < arrayLength; i++ {
			if indexes[i] == int32(i) {
				isDerangement = false
				break
			}
		}
		if isDerangement {
			return indexes
		}
	}
}

func GeneratePoolPairs(c *gin.Context) {
	users := models.ParseUserIds(c.Query("users"))
	fmt.Println(users)
	// usersFromDb :=
	var userPairs [][]int32
	shuffledIndexes := GenerateShuffledIndexes(len(users))
	for i, idx := range shuffledIndexes {
		currentSlice := users[i]
		selectedSlice := users[idx]
		if currentSlice == selectedSlice {
			newIdx := (idx + 1) % int32(len(users))
			selectedSlice = users[newIdx]
		}
		pair := []int32{currentSlice, selectedSlice}
		userPairs = append(userPairs, pair)
	}
	c.JSON(http.StatusOK, userPairs)
}
