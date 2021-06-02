package main

import (
	"fmt"
	"time"
)

func funcwhotakesometime(value string) {
	time.Sleep(2 * time.Second)
	fmt.Println(value)
}

func main() {
	fmt.Println("Before the time is:", time.Now().Unix())
	listcmd := []string{"one", "two", "tree"}
	for _, v := range listcmd {
		funcwhotakesometime(v)
	}
	fmt.Println("After the time is:", time.Now().Unix())
}
