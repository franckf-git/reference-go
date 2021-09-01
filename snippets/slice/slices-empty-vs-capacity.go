package main

import (
	"fmt"
)

func main() {
	numbers := []int{}
	for i := 0; i < 4; i++ {
		// reallocate memory at every iteration
		numbers = append(numbers, i)
		fmt.Println(len(numbers))
	}
	fmt.Println(numbers)

	fmt.Println("==========")

	numbers2 := make([]int, 4)
	for i := 0; i < cap(numbers2); i++ {
		// use same memory
		numbers2[i] = i
		fmt.Println(len(numbers2))
	}
	fmt.Println(numbers2)
}

