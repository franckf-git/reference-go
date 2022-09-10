package main

/*
#include <time.h>
#include <stdio.h>
#include <stdlib.h>

void myprint(char* s) {
	printf("%s\n", s);
}
*/
import "C"

import (
	"unsafe"
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

	hello()
}

func random() int {
	return int(C.rand()) % 100
}

func seed(i int) {
	C.srand(C.uint(i))
}

func hello() {
	cs := C.CString("Hello from stdio\n")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
}
