package main

import (
	"encoding/json"
	"fmt"
)

type Software struct {
	Name      string
	Developer string
}

func main() {

	softwaresJson := `[{"Name": "AutoCAD","Developer": "Autodesk"},{"Name": "Firefox","Developer": "Mozilla"},{"Name": "Chrome","Developer": "Google"}]`

	var softwares []Software

	err := json.Unmarshal([]byte(softwaresJson), &softwares)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v\n", softwares) // [{AutoCAD Autodesk} {Firefox Mozilla} {Chrome Google}]
}
