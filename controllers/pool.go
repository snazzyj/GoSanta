package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secret-santa/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPoolsFromJSONFile(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetPoolJsonFile())
}

func GetPoolById(c *gin.Context) {
	poolId, _ := strconv.Atoi(c.Param("id"))
	pool := models.GetPoolById(int32(poolId), models.DecodeData[models.PoolModel](models.OpenFile("pool.json")))
	c.JSON(http.StatusOK, pool)
}
func PostNewPool(c *gin.Context) {
	// fmt.Println(c.Params, "\n")
	for key, values := range c.Request.PostForm {
		fmt.Printf("Key: %s, Values: %v\n", key, values)
	}

	// Retrieve the string from the form
	pairingsStr := c.PostForm("pairings")

	fmt.Printf("The pairing string is...%s\n", pairingsStr)
	// Declare a variable to hold the result
	var pairings [][]int32

	// Unmarshal the JSON string into the pairings variable
	err := json.Unmarshal([]byte(pairingsStr), &pairings)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pairings format"})
		return
	}
	userIdsStr := c.PostForm("users")

	// Remove brackets and split by comma
	userIdsStr = strings.Trim(userIdsStr, "[]")
	ids := strings.Split(userIdsStr, ",")
	var userIds []int32
	for _, idStr := range ids {
		idStr = strings.TrimSpace(idStr) // Remove any whitespace
		if id, err := strconv.ParseInt(idStr, 10, 32); err == nil {
			userIds = append(userIds, int32(id))
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userId"})
			return
		}
	}
	layout := "01/03/2006" // e.g., "MM/dd/yyyy"
	fmt.Println("exchangeDate: ", c.PostForm("exchangeDate"))
	exchangeDate, _ := time.Parse(layout, c.PostForm("exchangeDate"))
	fmt.Println("exchangeDate after Parse: ", exchangeDate)
	result, err := models.AddPool(models.PoolModel{
		ID:           models.GenerateRandomNumber(),
		Users:        userIds,
		Pairings:     pairings,
		ExchangeDate: exchangeDate,
	})
	fmt.Print(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error creating new pool"})
		return
	}
}
