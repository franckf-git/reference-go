package main

import "os"

func main() {
	os.Mkdir("newfolder", 0700)
	os.Chdir("newfolder")
}
