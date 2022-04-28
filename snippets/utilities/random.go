package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	randomUpperBound = 1 << 15 // two to the power of fifteen
)

func random() int {
	return int(rand.Int31n(randomUpperBound))
}

func init() {
	rand.Seed(time.Now().UnixNano()) // not for cryto
}

func main() {
	fmt.Printf("%d\n", random())
}
