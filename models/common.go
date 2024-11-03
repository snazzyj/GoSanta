package models

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"math/rand"
)

type Common struct {
	CreatedAt time.Time `json:"createdAt"`
	Createdby string    `json:"createdBy"`
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
