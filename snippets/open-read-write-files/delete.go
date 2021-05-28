package main

import (
    "fmt"
    "os"
)

func main() {
  err := os.Remove("test.txt")  // remove a single file
  if err != nil {
    fmt.Println(err)
  }
}
