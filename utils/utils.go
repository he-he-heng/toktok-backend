package utils

import (
	"math/rand"
	"time"
)

func RandomDigit() int {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(900000) + 100000 // 100000 ~ 999999
}
