package main

import "with-tests/addition"
import "with-tests/substraction"
import "fmt"

func main() {
	fmt.Println(addition.Add(1,3))
	fmt.Println(substraction.Sub(9,4))
}

// go test ./...
