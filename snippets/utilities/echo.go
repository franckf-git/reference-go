package main

import (
	"os"
)

func echo(text ...string) {
	var textBytes []byte

	for i, t := range text {
		if i > 0 {
			textBytes = append(textBytes, ' ')
		}

		textBytes = append(textBytes, t...)
	}

	textBytes = append(textBytes, '\n')

	if _, err := os.Stdout.Write(textBytes); err != nil {
		panic(err)
	}
}

func main() {
	echo("Go", "is great!")
}
