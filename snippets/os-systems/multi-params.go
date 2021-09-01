package main

import (
	"fmt"
)

func variadicFunc(arg ...int) {
	for i := range arg {
		fmt.Println(arg[i])
	}
}

// pass a slice of interface{}
func printAll(arg ...interface{}) {
	fmt.Println(arg...) // forward it here
}

func doWork(a int, b string, num ...int) {
	fmt.Println(num)
}

func main() {
	variadicFunc(1, 2, 3) // 3 arguments
	// output:
	// 1
	// 2
	// 3
	variadicFunc(3, 4, 5, 6, 7, 8) // 6 arguments
	// output:
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8

	// pass all types of arguments using interface
	printAll(1, "Jane", 3.4, 4.7890, 'a') // prints 1 Jane 3.4 4.789 97

	doWork(1, "Jane", 3, 4) // prints [3 4]
}
