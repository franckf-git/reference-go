package main

import (
    "fmt"
    "image/color"
    "image/png"
    "log"
    "os"
)

func main() {
    // This example uses png.Decode which can only decode PNG images.
    pngFile, err := os.Open("png.png")
    if err != nil {
        log.Fatal(err)
    }
    defer pngFile.Close()

    // Consider using the general image.Decode as it can sniff and decode any registered image format.
    img, err := png.Decode(pngFile)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(img)

    levels := []string{" ", "░", "▒", "▓", "█"}

    for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
        for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
            c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
            level := c.Y / 51 // 51 * 5 = 255
            if level == 5 {
                level--
            }
            fmt.Print(levels[level])
        }
        fmt.Print("\n")
    }
}
