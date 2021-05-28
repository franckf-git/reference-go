package main

import (
    "io/ioutil"
)

func main() {
    s := []byte("This is a string")                // convert string to byte slice
    ioutil.WriteFile("testfile.txt", s, 0644)      // the 0644 is octal representation of the filemode
}
