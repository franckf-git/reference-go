package main

import (
	"fmt"
	"math"
)

// Fibo return fibonacci suite on number th
func Fibo(number int) []int {
	result := []int{1, 1}
	for len(result) < number {
		nMinus1 := result[len(result)-1]
		nMinus2 := result[len(result)-2]
		current := nMinus1 + nMinus2
		newEntry := []int{current}
		result = append(result, newEntry...)
	}
	return result
}

// Binet return fibonacci result at number position
func Binet(number int) int {
	phi := (1 + math.Sqrt(5)) / 2
	fiboValue := math.Pow(phi, float64(number)) / math.Sqrt(5)
	result := math.Floor(fiboValue)
	return int(result + 1) // math floor is not like in JS so we have to add +1
}
func main() {
	resultF := Fibo(17)
	resultB := Binet(17)
	fmt.Print(resultF, resultB)
}
