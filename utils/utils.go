package utils

import (
	"math"
	"math/rand"
)

func Rand() float64 {
	return rand.Float64()
}

func RandRange(a, b float64) float64 {
	var c float64
	if b < a {
		c = b
		b = a
		a = c
	}
	return a + (b-a)*rand.Float64()
}

func Degrees2Rad(alpha float64) float64 {
	return alpha * math.Pi / 180.0
}
