package main

import (
    "fmt"
    "os"
)

func main() {
  err := os.RemoveAll("directoryname")  // delete an entire directory
  if err != nil {
    fmt.Println(err)
  }
}
