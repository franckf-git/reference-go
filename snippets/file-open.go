package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening file!!!")
	}
	defer file.Close()

	// declare chunk size
	const maxSz = 4

	// create buffer
	b := make([]byte, maxSz)

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(string(b[:readTotal])) // print content from buffer
	}
}
