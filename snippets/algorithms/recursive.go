package main

import "fmt"

func main() {
	countDown(100)
}

func countDown(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	next := n - 1
	countDown(next)

}
