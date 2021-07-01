package main

import (
	"encoding/json"
	"fmt"
)

type App struct {
	Name string
}

func main() {
	apps := []App{
		{Name: "Google Play"},
		{Name: "Evernote"},
		{Name: "Buffer"},
	}

	appsJson, err := json.Marshal(apps)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(appsJson)) // [{"Name":"Google Play"},{"Name":"Evernote"},{"Name":"Buffer"}]

}
