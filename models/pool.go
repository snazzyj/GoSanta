package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type PoolModel struct {
	Common
	ID           int32     `json:"poolId"`
	Users        []int32   `json:"userIds"`
	Pairings     [][]int32 `json:"userPairings"`
	ExchangeDate time.Time `json:"exchangeDate"`
}

func GetPoolJsonFile() []PoolModel {
	return DecodeData[PoolModel](OpenFile("pools.json"))
}

func GetPoolById(poolId int32, pool []PoolModel) *PoolModel {
	return searchForElement(pool, "id", poolId)
}
func AddPool(pool PoolModel) (bool, error) {
	pools := GetPoolJsonFile()
	if pools != nil {
		pools = append(pools, pool)
	} else {
		pools = make([]PoolModel, 0)
		pools = append(pools, pool)
	}

	jsonData, err := json.MarshalIndent(pools, "", "    ")
	if err != nil {
		fmt.Println("err creating json: ", err)
		return false, errors.New("error creating json format")
	}

	file, err := os.Create("pools.json")
	if err != nil {
		fmt.Println("err creating json file: ", err)
		return false, errors.New("error creating json file to os")
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("err writing to json file", err)
		return false, errors.New("error writing to the json file")
	}
	fmt.Println("successfully wrote to pools.json")
	return true, nil
}
func ParseUserIds(userIdsStr string) []int32 {
	// Remove brackets and split by comma
	userIdsStr = strings.Trim(userIdsStr, "[]")
	ids := strings.Split(userIdsStr, ",")
	var userIds []int32
	for _, idStr := range ids {
		idStr = strings.TrimSpace(idStr) // Remove any whitespace
		if id, err := strconv.ParseInt(idStr, 10, 32); err == nil {
			userIds = append(userIds, int32(id))
		} else {
			// throw an error
		}
	}
	return userIds
}
func ParsePairings(pairingsStr string) [][]int32 {
	// Declare a variable to hold the result
	var pairings [][]int32

	// Unmarshal the JSON string into the pairings variable
	err := json.Unmarshal([]byte(pairingsStr), &pairings)
	if err != nil {
		// throw an error
		// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pairings format"})
		// return
	}
	return pairings
}
