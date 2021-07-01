package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Name   string
	Author string
}

type Software struct {
	Name      string
	Developer string
}

func main() {
	// encoding
	book := Book{"C++ programming language", "Bjarne Stroutsrup"}

	resbook, errbook := json.Marshal(book)

	if errbook != nil {
		fmt.Println(errbook)
	}

	fmt.Println(string(resbook)) // {"Name":"C++ programming language","Author":"Bjarne Stroutsrup"}

	// encoding multi
	apps := []Software{{Name: "AutoCAD", Developer: "Autodesk"}, {Name: "Firefox", Developer: "Mozilla"}, {Name: "Chrome", Developer: "Google"}}
	appsJson, errapps := json.Marshal(apps)

	if errapps != nil {
		fmt.Println(errapps)
	}

	fmt.Println(string(appsJson)) // [{"Name":"AutoCAD","Developer":"Autodesk"},{"Name":"Firefox","Developer":"Mozilla"},{"Name":"Chrome","Developer":"Google"}]

	// decoding
	String := `{"Name": "Robots", "Author": "Asimov"}`

	var setbook Book

	errset := json.Unmarshal([]byte(String), &setbook)

	if errset != nil {
		fmt.Println(errset)
	}

	fmt.Printf("%+v\n", setbook.Name) // Robots

	// decoding multi
	softwaresJson := `[{"Name": "AutoCAD","Developer": "Autodesk"},{"Name": "Firefox","Developer": "Mozilla"},{"Name": "Chrome","Developer": "Google"}]`

	var softwares []Software

	errsoft := json.Unmarshal([]byte(softwaresJson), &softwares)

	if errsoft != nil {
		fmt.Println(errsoft)
	}

	fmt.Printf("%v\n", softwares[1].Developer) // Mozilla

	// unstructured to map
	unstructuredJson := `{"os": {"Windows": "Windows OS","Mac": "OSX","Linux": "Ubuntu"},"compilers": "gcc"}`

	var result map[string]interface{}

	json.Unmarshal([]byte(unstructuredJson), &result)

	fmt.Println(result["os"]) // map[Linux:Ubuntu Mac:OSX Windows:Windows OS]

}

