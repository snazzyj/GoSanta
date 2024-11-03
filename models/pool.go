package models

import (
	"time"
)

type Pool struct {
	Common
	ID           int64     `json:"poolId"`
	Users        []int32   `json:"userIds"`
	Pairings     [][]int32 `json:"userPairings"`
	ExchangeDate time.Time `json:"exchangeDate"`
}
