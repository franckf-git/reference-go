package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	// get file
	out, _ := os.Create("novel.txt") //give any name
	defer out.Close()

	resp, _ := http.Get("http://www.gutenberg.org/files/18581/18581.txt")
	defer resp.Body.Close()
	io.Copy(out, resp.Body)

	// read file
	file, err := os.Open("novel.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}
