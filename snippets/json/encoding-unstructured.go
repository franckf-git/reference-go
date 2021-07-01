package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string
	City   string
}

type Person struct {
	Name    string
	Address Address
}

func main() {
	p := Person{
		Name: "Sherlock Holmes",
		Address: Address{
			"22/b Baker street",
			"London",
		},
	}

	str, err := json.Marshal(p)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(str)) // {"Name":"Sherlock Holmes","Address":{"Street":"22/b Baker street","City":"London"}}
}
