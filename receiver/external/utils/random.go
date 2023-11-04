package utils

import (
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	r = rand.New(s)
}

func RandomNumber(n int) int {
	return r.Intn(n)
}
