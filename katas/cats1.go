package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var filename string = "cats.json"

type Auto []struct {
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Children []Children `json:"children,omitempty"`
}
type Children struct {
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Children []Children `json:"children,omitempty"`
}

func main() {
	var cats []string
	data, _ := os.ReadFile(filename)
	dataFixStandard := bytes.ReplaceAll(data, []byte(`'`), []byte(`"`))
	catNames(dataFixStandard, &cats)
	fmt.Printf("%#v\n", cats)
}

func catNames(folder []byte, cats *[]string) {
	var payload Auto
	err := json.Unmarshal(folder, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Printf("%#v\n", cats)
	for _, v := range payload {
		if v.Type == "image" {
			println("add cats")
			*cats = append(*cats, v.Name)
		} else if v.Type == "folder" {
			println("recursive")
			catNames(v.Children, cats)
		}
	}

}
