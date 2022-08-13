package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var filename string = "cats.json"

type Children struct {
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	Children []Children `json:"children,omitempty"`
}

func main() {
	var cats []string
	var payload []Children

	data, _ := os.ReadFile(filename)
	dataFixStandard := bytes.ReplaceAll(data, []byte(`'`), []byte(`"`))

	err := json.Unmarshal(dataFixStandard, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	catNames(payload, &cats)
	fmt.Printf("%#v\n", cats)
}

func catNames(payload []Children, cats *[]string) {
	for _, v := range payload {
		if v.Type == "image" {
			*cats = append(*cats, v.Name)
		} else if v.Type == "folder" {
			catNames(v.Children, cats)
		}
	}
}
