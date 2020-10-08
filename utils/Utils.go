package utils

import (
	"math"
	"math/rand"
)

func Rand() float64 {
	return rand.Float64()
}

func RandRange(a, b float64) float64 {
	if b < a {
		a, b = b, a
	}
	return a + (b-a)*rand.Float64()
}

func RandInt(a, b int) int {
	return int(RandRange(float64(a), float64(b)))
}

func Degrees2Rad(alpha float64) float64 {
	return alpha * math.Pi / 180.0
}
