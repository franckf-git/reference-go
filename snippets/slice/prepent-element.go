package main

import "fmt"

func main() {
	fibonacciNumbers := []int{1, 2, 3, 5, 8, 13, 21}

	fibonacciNumbers = append(fibonacciNumbers, 0)

	copy(fibonacciNumbers[1:], fibonacciNumbers)

	fibonacciNumbers[0] = 1

	fmt.Println(fibonacciNumbers)
}
