package main

import (
	"fmt"
)

func main() {
	resultF := Fibo(17)
	resultB := Binet(17)
	resultR := Recursive(17)
	resultE := Effective(17)
	fmt.Println(resultF, resultB, resultR, resultE)
}
