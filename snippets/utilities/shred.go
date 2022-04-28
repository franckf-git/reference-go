package main

import (
	"crypto/rand"
	"io/fs"
	"os"
)

func shred(fileName string, iterations int, roundUp bool) {
	file, err := os.OpenFile(fileName, 0666, fs.FileMode(os.O_RDWR))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	endLocation, err := file.Seek(0, 2)
	if err != nil {
		panic(err)
	}

	if roundUp {
		roundedEndLocation := endLocation / 1024
		if endLocation%1024 != 0 {
			roundedEndLocation++
		}
		roundedEndLocation *= 1024

		endLocation = roundedEndLocation
	}

	data := make([]byte, endLocation)

	for i := 0; i < iterations; i++ {
		if _, err := rand.Read(data); err != nil {
			panic(err)
		}

		if _, err := file.WriteAt(data, 0); err != nil {
			panic(err)
		}

		if err := file.Sync(); err != nil {
			panic(err)
		}
	}
}

func main() {
	shred("test.txt", 5, true)
}
