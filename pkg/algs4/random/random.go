package random

import (
	"math/rand"
	"time"
)

func Uniform() float64 {
	return rand.Float64()
}

func UniformInt(lower, upper int) int {
	rand.Seed(time.Now().UnixNano())
	rng := upper - lower
	return rand.Intn(rng) + lower
}
