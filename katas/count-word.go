package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var osStdin = os.Stdin
	log.Printf("cw: %#+v\n", countWord(osStdin))
}

func countWord(input io.Reader) map[string]int {
	var results = make(map[string]int)
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords) // feel like cheating
	for scanner.Scan() {
		currentWord := scanner.Text()
		currentWord = strings.ToLower(currentWord)
		results[currentWord] = results[currentWord] + 1
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return results
}

// TODO use regex for simple and complex punctuations ?
// TODO accents à ?
// TODO sort response
