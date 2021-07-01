package main

import (
	"encoding/json"
	"fmt"
)

type Game struct {
	Name   string
	Rating float64
}

func main() {
	codString := `{"Name": "Call of Duty", "Rating": 8.4}`

	var cod Game

	err := json.Unmarshal([]byte(codString), &cod)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", cod) // {Name:Call of Duty Rating:8.4}
}
