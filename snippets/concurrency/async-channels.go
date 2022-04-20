package main

import (
	"fmt"
	"math/rand"
	"time"
)

func funcwhotakesometime(value string, c chan string) {
	random := rand.Intn(7-1) + 1
	time.Sleep(time.Duration(random) * time.Second)

	fmt.Println("end at:", time.Now().Unix())
	c <- value
}

func main() {
	fmt.Println("Before the time is:", time.Now().Unix())
	listcmd := []string{"one", "two", "tree"}

	c := make(chan string, len(listcmd))

	for _, v := range listcmd {
		go funcwhotakesometime(v, c)

	}

	// obligé de faire une boucle pour avoir tous les messages (sinon on a que le dernier
	// c'est surtout parce qu'on n'utilise pas de waitGroup et de close
	result := make([]string, len(listcmd))
	for i, _ := range result {
		result[i] = <-c
		fmt.Println(result[i])
	}

	fmt.Println("After the time is:", time.Now().Unix())
}
