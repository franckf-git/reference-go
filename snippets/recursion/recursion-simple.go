package main

import (
	"fmt"
)

func countDown(number int) int {
	newNumber := number - 1
	fmt.Println(newNumber)
	if newNumber > 0 {
		countDown(newNumber)
	}
	return newNumber
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	countDown(4)
	fmt.Println(factorial(4))
}
