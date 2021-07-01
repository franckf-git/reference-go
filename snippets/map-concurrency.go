package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Let's have a global map variable so that all goroutines can easily access it
var myMap map[int]string

func myGoRoutine(id int, numKeys int, wg *sync.WaitGroup) {
	defer wg.Done()

	for key, _ := range myMap {
		myMap[key] = strconv.Itoa(id)
	}

	for key, value := range myMap {
		fmt.Printf("Goroutine #%d -> Key: %d, Value: %s\n", id, key, value)
	}

}

func main() {
	// Initially set some values
	myMap = make(map[int]string)
	myMap[0] = "test"
	myMap[2] = "sample"
	myMap[1] = "GoLang is Fun!"

	// Get the number of keys
	numKeys := len(myMap)

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go myGoRoutine(i, numKeys, &wg)
	}

	// Blocking wait
	wg.Wait()

	// Iterate over all keys
	for key, value := range myMap {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}
}
// potential data race condition
