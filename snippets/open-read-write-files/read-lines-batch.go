package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	var (
		fileName   string = "lines.txt"
		file       *os.File
		err        error
		scanner    *bufio.Scanner
		batchSize  int      = 6
		batchLines []string = make([]string, batchSize)
		counter    int
		i          int
	)

	file, err = os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for i = 0; scanner.Scan(); i++ {
		log.Printf("%d) \"%s\"\n", i, scanner.Text())
		batchLines[counter] = scanner.Text()
		counter++
		if counter == batchSize {
			processBatch(batchLines)
			counter = 0
		}
	}

	log.Printf("counter: %#+v\n", counter)
	processBatch(batchLines[:counter])

	err = scanner.Err()
	if err != nil {
		log.Fatalln(err)
	}
}

func processBatch(batchLines []string) {
	var v string
	for _, v = range batchLines {
		log.Printf("-----processBatch: %#+v\n", v)
	}
}
