package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	const fileName = "elements.csv"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		fmt.Printf("%s\n", strings.Join(record, "****"))
	}
}
