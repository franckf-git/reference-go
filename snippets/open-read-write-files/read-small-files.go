package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    content, err := ioutil.ReadFile("text.txt")     // the file is inside the local directory
    if err != nil {
        fmt.Println("Err")
    }
    fmt.Println(string(content))    // This is some content
}
