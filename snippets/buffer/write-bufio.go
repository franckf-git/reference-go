package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // open file for write
    f, e := os.OpenFile("buffertest.txt", os.O_WRONLY, 0666)
    CheckError(e)

    // create a buffered writer
        // here we create a sized buffer of 4 bytes and the default is 4096 bytes
    bw := bufio.NewWriterSize(f, 4)

    // write to buffer
    bw.Write([]byte("H"))
    bw.Write([]byte("e"))
    bw.Write([]byte("l"))
    bw.Write([]byte("l"))
    bw.Write([]byte("o"))
    bw.Write([]byte(" "))
    bw.Write([]byte("w"))
    bw.Write([]byte("o"))
    bw.Write([]byte("r"))
    bw.Write([]byte("l"))
    bw.Write([]byte("d"))

    // check how much is inside waiting to be written
    fmt.Println(bw.Buffered())            // 3

    // check available space left
    fmt.Println(bw.Available())           // 1
}

func CheckError(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

// Hello wo          // since the buffer can contain 4 bytes it will write 4 bytes at a time
                     // "hello wo" contains 8 bytes multiple of four
                     // The "rld" portion is inside the buffer that's why
                     // buffered returns 3
                     // available returns 1

// write to disk and clear buffer
bw.Flush()
