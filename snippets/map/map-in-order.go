package main

import (
	"fmt"
	"sort"
)

func main() {
	// Allocate memory for the map
	var myMap = make(map[int]string)
	myMap[0] = "test"
	myMap[2] = "sample"
	myMap[1] = "GoLang is Fun!"

	// Let's get a slice of all keys
	keySlice := make([]int, 0)
	for key, _ := range myMap {
		keySlice = append(keySlice, key)
	}

	// Now sort the slice
	sort.Ints(keySlice)

	// Iterate over all keys in a sorted order
	for _, key := range keySlice {
		fmt.Printf("Key: %d, Value: %s\n", key, myMap[key])
	}
}
