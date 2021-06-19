package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {

    // open file for reading
    f, e := os.Open("buffertest.txt")
    CheckError(e)

    // create a buffered reader
    br := bufio.NewReader(f)

    // peek n bytes
    // bbuf is a byte buffer of size 10
    bbuf := make([]byte, 10)
    bbuf, e = br.Peek(6)
    CheckError(e)

    // bbuf contents
    fmt.Println(string(bbuf)) // Hello

    // num read
    nr, e := br.Read(bbuf)
    CheckError(e)

    fmt.Println("Num bytes read", nr) // 6

    // read single byte
    singleByte, e := br.ReadByte()
    CheckError(e)

    fmt.Println("Single byte is", string(singleByte))     // w

    // reset buffer
    br.Reset(f)
}

func CheckError(e error) {
    if e != nil {
        fmt.Println(e)
    }
}
