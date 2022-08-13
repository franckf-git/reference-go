package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var filename string = "cats.json"

func main() {
	var cats []string
	data, _ := os.ReadFile(filename)
	dataFixStandard := bytes.ReplaceAll(data, []byte(`'`), []byte(`"`))
	catNames(dataFixStandard, &cats)
	fmt.Printf("%#v\n", cats)
}

func catNames(folder []byte, cats *[]string) {
	var payload []map[string]interface{}
	err := json.Unmarshal(folder, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Printf("%#v\n", cats)
	for _, v := range payload {
		if v["type"] == "image" {
			println("add cats")
			*cats = append(*cats, v["name"].(string))
		} else if v["type"] == "folder" {
			println("recursive")
			catNames([]byte(v["children"].([]interface{})), cats)
		}
	}

}
