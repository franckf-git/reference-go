package main

import (
	"fmt"
	"os"
)

func main() {
	const fileName = "test.txt"
	fileContents, err := os.ReadFile(fileName)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", fileContents)
}
