package random

import (
	"math/rand"
	"time"
)

func UniformInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}
