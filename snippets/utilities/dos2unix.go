package main

import (
	"os"
	"regexp"
)

var (
	windowsNewlineRegexp = regexp.MustCompile(`\r\n`)
)

func dos2unix(path string) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	fileData = windowsNewlineRegexp.ReplaceAllLiteral(
		fileData,
		[]byte{'\n'},
	)

	if err := os.WriteFile(path, fileData, 0666); err != nil {
		panic(err)
	}
}

func main() {
	dos2unix("test.txt")
}
