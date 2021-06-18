package main

import (
	"math"
)

// Binet return fibonacci result at number position
func Binet(number int) int {
	phi := (1 + math.Sqrt(5)) / 2
	fiboValue := math.Pow(phi, float64(number)) / math.Sqrt(5)
	result := math.Floor(fiboValue)
	return int(result + 1) // math floor is not like in JS so we have to add +1
}
