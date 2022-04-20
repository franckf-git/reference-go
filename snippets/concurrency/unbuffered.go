package main

import (
	"fmt"
	"sync"
	"time"
)

func ConcurrentSum(l int) uint64 {
	var wg sync.WaitGroup
	c := make(chan uint64)
	for i := 0; i < l; i++ {
		wg.Add(1)
		go func(n int) {
			//Suppose here we do a heavy API call instead
			c <- uint64(n)
			time.Sleep(time.Millisecond)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return processResult(c)
}
func processResult(c <-chan uint64) uint64 {
	total := uint64(0)
	for n := range c {
		total += n
	}
	return total
}
func main() {
	sum := ConcurrentSum(10)
	//Output: Sum =  45
	fmt.Println("Sum = ", sum)
}
