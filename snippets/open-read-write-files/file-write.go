package main

import (
    "fmt"
    "os"
)

func main() {

  // create the file
  f, err := os.Create("test.txt")
  if err != nil {
    fmt.Println(err)
  }
  // close the file with defer
  defer f.Close()

  // do operations

  //write directly into file
  f.Write([]byte("a string"))

  // write a string
  f.WriteString("\nThis is a pretty long string")

  // // write from a specific offset
  f.WriteAt([]byte("another string"), 42)     // 12 is the offset from start
}
