package main

import (
    "image"
    "image/png"
    "os"
)
func main() {
    myImg := image.NewRGBA(image.Rect(0, 0, 12, 6))
        out, _ := os.Create("save.png")
        png.Encode(out, myImg)
        out.Close()
}
