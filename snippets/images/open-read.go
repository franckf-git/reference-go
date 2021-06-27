package main

import (
    "fmt"
    "image"
    "image/png"
    "log"
    "os"
)

func main() {
    pngFile, err := os.Open("png.png")
    if err != nil {
        log.Fatal(err)
    }
    defer pngFile.Close()

    imData, imType, err := image.Decode(pngFile)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(imData)
    fmt.Println(imType)

    png, err := png.Decode(pngFile)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(png)
}
