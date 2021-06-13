package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := &Person{"Jack", 22}

	v, _ := xml.MarshalIndent(p, "", " ")

	fmt.Println(string(v))
}
