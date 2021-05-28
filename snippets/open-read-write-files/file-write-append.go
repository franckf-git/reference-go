package main

import (
    "fmt"
    "os"
)

func main() {
    f, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    newLine := "This is a string which will be appended."
    _, err = fmt.Fprintln(f, newLine)
    if err != nil {
      fmt.Println(err)
    }
}
