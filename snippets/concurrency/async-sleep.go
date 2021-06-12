package main

import (
	"fmt"
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
	for _, v := range listcmd {
		go funcwhotakesometime(v)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("After the time is:", time.Now().Unix())
}
// weird use case
