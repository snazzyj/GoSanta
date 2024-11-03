package models

import (
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
