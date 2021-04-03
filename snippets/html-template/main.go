package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("templates/hello.gohtml")

	if err != nil {
		panic(err)
	}

	data := struct {
		Name   string
		Skills []string
	}{
		Name: "John Doe",
		Skills: []string{
			"C++",
			"Java",
			"Python",
		},
	}

	err = t.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
