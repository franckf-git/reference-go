// dont work yet

package main

import (
	"fmt"
	//"math/rand"
	"time"
)

func funcwhotakesometime(value string, c chan string) {
	//random := rand.Intn(7-1) + 1
	//time.Sleep(time.Duration(random) * time.Second)

	time.Sleep(2 * time.Second)
	fmt.Println("end at:", time.Now().Unix())
	c <- value
}

func main() {
	fmt.Println("Before the time is:", time.Now().Unix())
	listcmd := []string{"one", "two", "tree"}

	c := make(chan string)

	for _, v := range listcmd {
		go funcwhotakesometime(v, c)

	}
	returnchannel := ""
	returnchannel <- c
	fmt.Println(returnchannel)
	fmt.Println("After the time is:", time.Now().Unix())
}

