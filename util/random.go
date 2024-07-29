package util

import (
	"math/rand"
	"time"
)

var rng *rand.Rand

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)
}

func GetRandomNum(min, max int) int {
	return rng.Intn(max-min+1) + min
}

func GetRandomSeed() uint32 {
	return rng.Uint32()
}
