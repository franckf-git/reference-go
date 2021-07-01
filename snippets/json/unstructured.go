package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	unstructuredJson := `{"os": {"Windows": "Windows OS","Mac": "OSX","Linux": "Ubuntu"},"compilers": "gcc"}`

	var result map[string]interface{}

	json.Unmarshal([]byte(unstructuredJson), &result)

	fmt.Println(result["os"]) // map[Linux:Ubuntu Mac:OSX Windows:Windows OS]
}
