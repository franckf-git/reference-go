package main

import (
	"fmt"
	"os"
)

func cat(fileName string, printLineNumbers bool) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if printLineNumbers {
		outputData := make([]byte, 0, len(fileData))

		var lineCount int

		for i, d := range fileData {
			if d == '\n' || i == 0 {
				if i != 0 {
					outputData = append(outputData, d)
				}

				lineCount++

				startOfLineString := fmt.Sprintf("%6d  ", lineCount)

				outputData = append(outputData, startOfLineString...)

				if i == 0 {
					outputData = append(outputData, d)
				}
			} else {
				outputData = append(outputData, d)
			}
		}

		fmt.Printf("%s\n", outputData)
	} else {
		fmt.Println(string(fileData))
	}
}

func main() {
	cat("test.txt", true)
}
