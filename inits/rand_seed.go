package inits

import (
	"math/rand"
	"time"
)

func RandSeed() {
	rand.Seed(time.Now().Unix())
}
