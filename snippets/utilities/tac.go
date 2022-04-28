package main

import (
	"os"
)

func printReversedLine(lineData []byte) {
	l := len(lineData)
	mid := l / 2

	for i := 0; i < mid; i++ {
		j := l - i - 1

		lineData[i], lineData[j] = lineData[j], lineData[i]
	}

	if _, err := os.Stdout.Write(lineData); err != nil {
		panic(err)
	}
}

func tac(fileName string, fixNewlineBug bool) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	endLocation, err := file.Seek(0, 2)
	if err != nil {
		panic(err)
	}

	fileData := make([]byte, endLocation)

	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	_, err = file.Read(fileData)
	if err != nil {
		panic(err)
	}

	if fileData[endLocation-1] == 0 { // EOF
		endLocation--
	}

	var line []byte

	for i := endLocation - 1; i >= 0; i-- {
		d := fileData[i]

		line = append(line, d)

		if d == '\n' {
			if len(line) > 0 {
				printReversedLine(line)

				line = []byte{}
			}
		}
	}

	if len(line) > 0 {
		if fixNewlineBug {
			line = append(line, '\n')
		}

		printReversedLine(line)
	}
}

func main() {
    // bug: the last two lines of the file to be printed are shown together
	tac("test.txt", false)
}
