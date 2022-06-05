package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func printLines(lineChannel <-chan string, doneChannel chan<- struct{}) {
	var i int
	for line := range lineChannel {
		i++
		fmt.Printf("%d) \"%s\"\n", i, line)
	}
	doneChannel <- struct{}{}
}
func main() {
	const fileName = "lines.txt"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	lineChannel := make(chan string)
	doneChannel := make(chan struct{})
	go printLines(lineChannel, doneChannel)
	reader := bufio.NewReader(file)
	buffer := make([]byte, 4)
	var currentLineBuilder strings.Builder
	var seenCarriageReturn bool
	for {
		bytesRead, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
		chunk := buffer[:bytesRead]
		for _, b := range chunk {
			switch b {
			case '\n':
				seenCarriageReturn = false
				lineChannel <- currentLineBuilder.String()
				currentLineBuilder.Reset()
			case '\r':
				seenCarriageReturn = true
			default:
				if seenCarriageReturn {
					currentLineBuilder.WriteByte('\r')
					seenCarriageReturn = false
				}
				currentLineBuilder.WriteByte(b)
			}
		}
	}
	if currentLineBuilder.Len() > 0 {
		lineChannel <- currentLineBuilder.String()
	}
	close(lineChannel)
	<-doneChannel
}
