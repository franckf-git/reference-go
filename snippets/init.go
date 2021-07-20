package main

import "fmt"

func init() {
	fmt.Println("toujours premier en 2è")
}

func main() {
	fmt.Println("principal")
}

func init() {
	fmt.Println("toujours premier")
}
