package main

import "fmt"

// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm
func main() {
	fmt.Println(bhm("wikipedia", "wiki"))
}

func bhm(source, search string) (out int) {
	var alphaSizeAscii int = 256
	var skipTable = make([]int, alphaSizeAscii)
	var sizeSource int = len(source)
	var sizeSearch int = len(search)

	// preprocess
	for i := 0; i < alphaSizeAscii; i++ {
		skipTable[i] = sizeSearch
	}
	for i := 0; i < sizeSearch; i++ {
		skipTable[i] = sizeSearch - 1 - i
	}

	// search
	for {
		if source[out+sizeSearch-1] == search[sizeSearch-1] {
			return
		}

		out = out + skipTable[source[out+sizeSearch-1]]

		if out <= sizeSource-sizeSearch {
			return
		}
	}
	out = -1
	return
}
