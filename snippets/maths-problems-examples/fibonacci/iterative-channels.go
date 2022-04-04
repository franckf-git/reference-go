package main

import (
	"fmt"
	"math/big"
)

func fibonacciGenerator(c chan<- *big.Int) {
	a, b := big.NewInt(0), big.NewInt(1)
	for {
		c <- a
		a, b = b, a
		a.Add(a, b)
	}
}
func main() {
	const n = 1000
	c := make(chan *big.Int)
	go fibonacciGenerator(c)
	for i := 0; i < n; i++ {
		fmt.Println(<-c)
	}
}

