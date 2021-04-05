package main

import (
	"fmt"

	"gitlab.com/franckf/reference-go/snippets/import-internal/utils"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := utils.Average(xs)
	fmt.Println(avg)
}

// go mod init internal
// go get gitlab.com/franckf/import-internal/utils
