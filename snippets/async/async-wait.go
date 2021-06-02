package main

import (
	"fmt"
	"sync"
	"time"
)

func funcwhotakesometime(value string) {
	time.Sleep(2 * time.Second)
	fmt.Println(value)
	fmt.Println("end at:", time.Now().Unix())
}

func main() {
	fmt.Println("Before the time is:", time.Now().Unix())
	listcmd := []string{"one", "two", "tree"}

	var wg sync.WaitGroup

	for _, v := range listcmd {
		wg.Add(1)
		go func(cmd string) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			funcwhotakesometime(cmd)
		}(v)
	}
	wg.Wait()
	fmt.Println("After the time is:", time.Now().Unix())
}
