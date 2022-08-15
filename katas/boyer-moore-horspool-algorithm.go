package main

import (
	"fmt"
	"log"
)

// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore%E2%80%93Horspool_algorithm
func main() {
	fmt.Println(bhm("wikipedia", "edia"))
}

func bhm(source, search string) (out int) {
	var alphaSizeAscii [256]int
	var sizeSource int = len(source)
	var sizeSearch int = len(search)
	var iSource int
	var iSearch int

	// preprocess
	// - init
	for i := 0; i < len(alphaSizeAscii); i++ {
		alphaSizeAscii[i] = sizeSearch
	}
	log.Printf("alphaSizeAscii-init: %#+v\n", alphaSizeAscii)
	// - build
	for i := 0; i < (sizeSearch - 1); i++ {
		alphaSizeAscii[search[i]] = (sizeSearch - 1) - i
	}
	log.Printf("alphaSizeAscii-build: %#+v\n", alphaSizeAscii)

	log.Printf("iSource-before-search: %#+v\n", iSource)
	log.Printf("iSearch-before-search: %#+v\n", iSearch)
	// search
	for (iSource + sizeSearch) <= sizeSource {
		iSearch = sizeSearch - 1
		log.Printf("iSource-begin-loop: %#+v\n", iSource)
		log.Printf("iSearch-begin-loop: %#+v\n", iSearch)
		for source[iSource+iSearch] == search[iSearch] {
			if iSearch == 0 {
				out = iSource
				return
			}
			log.Printf("iSource-after-if: %#+v\n", iSource)
			log.Printf("iSearch-after-if: %#+v\n", iSearch)
			iSearch--
		}
		iSource = iSource + alphaSizeAscii[source[iSource+iSearch]]
		log.Printf("iSource-after-loop: %#+v\n", iSource)
		log.Printf("iSearch-after-loop: %#+v\n", iSearch)
	}
	out = -1
	return
}
