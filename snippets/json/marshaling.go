package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string
	Author string
}

func main() {
	book := Book{"C++ programming language", "Bjarne Stroutsrup"}
	res, err := json.Marshal(book)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res)) // {"Name":"C++ programming language","Author":"Bjarne Stroutsrup"}
}
