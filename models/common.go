package models

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"time"

	"math/rand"
)

type Common struct {
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string    `json:"createdBy"`
}

func GenerateRandomNumber() int32 {
	return rand.Int31()
}

func OpenFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	return file
}

func DecodeData[T any](file *os.File) []T {
	var result []T
	var err error
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&result)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}

func searchForElement[T any](array []T, key string, value interface{}) *T {
	for _, item := range array {
		// Use reflection to get the value of the specified key
		v := reflect.ValueOf(item).FieldByName(key)
		if v.IsValid() && v.Interface() == value {
			return &item // Return pointer to the found item
		}
	}
	return nil // Return nil if not found
}
