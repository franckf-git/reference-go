package main

import (
	"fmt"
)

func main() {
	var q int = 42 // Assign 42 to q
	var p *int
	p = &q // Make p reference whatever q is referencing

	t := *p
	q += 1

	fmt.Printf("q:  %v\n", q)  // prints 43 because we added 1 to 42
	fmt.Printf("&q: %v\n", &q) // prints 0xc000014118
	fmt.Printf("p:  %v\n", p)  // prints 0xc000014118
	fmt.Printf("&p: %v\n", &p) //prints 0xc00000e028 because pointer another pointer (unnecessary confusion)
	fmt.Printf("*p: %v\n", *p) // prints 43 because p is referencing q
	fmt.Printf("t:  %v\n", t)  // Copy whatever p is referencing to t's own spot in memory
	fmt.Printf("&t: %v\n", &t)
}

// + Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
