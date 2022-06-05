package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

var (
	lineBreakRegExp = regexp.MustCompile(`\r?\n`)
)

func main() {
	const fileName = "test.txt"
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fileLines := lineBreakRegExp.Split(string(fileContents), -1)
	for i, line := range fileLines {
		fmt.Printf("%d) \"%s\"\n", i+1, line)
	}
}
