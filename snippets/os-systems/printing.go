package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Fprint(os.Stderr, "Error ", 23, "\n")

	// show type
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Printf("%T\n", timeZone)

	// value
	t := "tttt"
	fmt.Printf("%v\n", t)

	// error
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())
}
