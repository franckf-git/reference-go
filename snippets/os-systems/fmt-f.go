package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		os.Exit(1)
	}
	filename := os.Args[1]
	destination, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create:", err)
		os.Exit(1)
	}
	defer destination.Close()
	fmt.Fprintf(destination, "[test]:")
	fmt.Fprintf(destination, "This can be used to store data output from any function to %s\n", filename)
}
