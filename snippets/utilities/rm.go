package main

import (
	"os"
	"path/filepath"
)

func rm(path string) {
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, dirEntry := range dir {
		childPath := filepath.Join(
			path,
			dirEntry.Name(),
		)

		if err := os.RemoveAll(childPath); err != nil {
			panic(err)
		}
	}
}

func main() {
	rm("test")
}
