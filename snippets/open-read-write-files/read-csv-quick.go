package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	const fileName = "elements.csv"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}
	const outputWidth = 85
	verticalSeparator := " || "
	horizontalSeparator := strings.Repeat("-", outputWidth)
	for i, record := range data {
		if i > 0 {
			fmt.Println(horizontalSeparator)
		}
		output := strings.Join(record, verticalSeparator)
		outputIndent := (outputWidth - len(output)) / 2
		if outputIndent > 0 {
			fmt.Print(strings.Repeat(" ", outputIndent))
		}
		fmt.Println(output)
	}
}
