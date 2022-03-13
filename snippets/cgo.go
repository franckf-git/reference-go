package main

/*
#include <stdlib.h>
#include <time.h>
*/
import "C"

import (
	"fmt"
)

func main() {

	// Seed rand generator.
	seed(int(C.time(nil)))

	// Output 10 random numbers from 0-100.
	for i := 1; i <= 10; i++ {
		rand := random()
		fmt.Println(rand)
	}
}

func random() int {
	return int(C.rand()) % 100
}

func seed(i int) {
	C.srand(C.uint(i))
}
