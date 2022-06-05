package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines) // ?
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
		//        /lorsque le slice à une certaine taille, le traiter
	}

	readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

	fmt.Println(fileLines)
}
